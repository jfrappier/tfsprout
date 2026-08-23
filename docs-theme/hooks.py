"""
Build-time documentation generation for the tfsprout docs site.

Two kinds of page in this site are not written by hand:

* ``checks/<ID>.md`` -- one page per check, built from the README that lives
  beside the analyzer in ``passes/<ID>/`` or ``xpasses/<ID>/``. Keeping the
  prose next to the code means a new analyzer gets a documentation page for
  free, and there is never a second copy to drift.
* ``changelog.md`` -- the repository CHANGELOG, surfaced in the site nav.

Check metadata (description, category, removed status) is read from
``docs/reference/checks.md``, which ``scripts/check-docs-sync.sh`` already
guarantees lists exactly the checks that exist on disk.
"""

from __future__ import annotations

import os
import re
import textwrap

from mkdocs.structure.files import File

# --------------------------------------------------------------------------
# Metadata

REPO_BLOB = "https://github.com/jfrappier/tfsprout/blob/main"

CATEGORIES = {
    "AT": ("Acceptance tests", "at"),
    "R": ("Resources", "r"),
    "S": ("Schemas", "s"),
    "V": ("Validation", "v"),
}

# Checks that can rewrite source under -fix. Kept here rather than inferred,
# because "has a SuggestedFix" is not something the index table records.
FIXABLE = {"R007", "XR007", "XR008"}

# Rows in docs/reference/checks.md, e.g.
#   | [AT001](../checks/AT001.md) | check for `TestCase` missing ... | AST |
ROW = re.compile(
    r"^\|\s*\[(?P<id>X?(?:AT|R|S|V)\d{3})\]\([^)]*\)\s*\|"
    r"\s*(?P<desc>.*?)\s*\|"
    r"\s*(?P<type>[^|]*?)\s*\|\s*$"
)

REMOVED_PREFIX = re.compile(r"^\*\*REMOVED\*\*\s*\([^)]*\)\s*")


def _split_id(check_id: str) -> tuple[str, bool]:
    """Return the category prefix and whether this is an extra (X) check."""
    extra = check_id.startswith("X")
    body = check_id[1:] if extra else check_id
    prefix = re.match(r"^[A-Z]+", body).group(0)
    return prefix, extra


def _read_registry(docs_dir: str) -> list[dict]:
    """Parse the check index into an ordered list of check metadata."""
    index = os.path.join(docs_dir, "reference", "checks.md")
    checks: list[dict] = []
    with open(index, encoding="utf-8") as fh:
        for line in fh:
            m = ROW.match(line)
            if not m:
                continue
            check_id = m.group("id")
            desc = m.group("desc")
            removed = bool(REMOVED_PREFIX.match(desc))
            desc = REMOVED_PREFIX.sub("", desc)
            prefix, extra = _split_id(check_id)
            label, slug = CATEGORIES[prefix]
            checks.append(
                {
                    "id": check_id,
                    "description": desc,
                    # The sidebar renders descriptions as plain text, so the
                    # index table's inline code markers have to come off.
                    "description_plain": desc.replace("`", ""),
                    "removed": removed,
                    "extra": extra,
                    "fixable": check_id in FIXABLE,
                    "category": label,
                    "category_slug": slug,
                    "source": ("xpasses" if extra else "passes") + "/" + check_id,
                    "url": "checks/" + check_id + "/",
                }
            )
    return checks


# --------------------------------------------------------------------------
# README -> page transformation

ALERT = re.compile(
    r"^> \[!(?P<kind>NOTE|TIP|IMPORTANT|WARNING|CAUTION)\]\n"
    r"(?P<body>(?:^>.*\n?)*)",
    re.MULTILINE,
)

ADMONITION = {
    "NOTE": "note",
    "TIP": "tip",
    "IMPORTANT": "info",
    "WARNING": "warning",
    "CAUTION": "danger",
}


def convert_alerts(text: str) -> str:
    """
    Rewrite GitHub alert blockquotes as MkDocs admonitions.

    The check READMEs are read on GitHub as well as on this site, so they use
    the ``> [!NOTE]`` syntax GitHub renders natively. Python-Markdown does not
    understand it, so it is translated here.
    """

    def repl(m: re.Match) -> str:
        body = "\n".join(
            line[2:] if line.startswith("> ") else line[1:]
            for line in m.group("body").rstrip("\n").split("\n")
        )
        return (
            f'!!! {ADMONITION[m.group("kind")]}\n\n'
            + textwrap.indent(body, "    ")
            + "\n"
        )

    return ALERT.sub(repl, text)


SECTION = re.compile(r"^## +(?P<title>.+?)\s*$", re.MULTILINE)


def _sections(body: str) -> tuple[str, list[tuple[str, str]]]:
    """Split markdown into (preamble, [(heading, content), ...])."""
    matches = list(SECTION.finditer(body))
    if not matches:
        return body, []
    preamble = body[: matches[0].start()]
    out = []
    for i, m in enumerate(matches):
        end = matches[i + 1].start() if i + 1 < len(matches) else len(body)
        out.append((m.group("title"), body[m.end() : end]))
    return preamble, out


def _badges(check: dict) -> str:
    items = [
        (
            f'<span class="badge badge--cat badge--{check["category_slug"]}">'
            f'{check["category"]}</span>'
        )
    ]
    if check["extra"]:
        items.append('<span class="badge badge--extra">Extra &middot; tfsproutx</span>')
    else:
        items.append('<span class="badge badge--standard">Standard</span>')
    if check["fixable"]:
        items.append('<span class="badge badge--fixable">Fixable with -fix</span>')
    if check["removed"]:
        items.append('<span class="badge badge--removed">Removed &middot; no longer reports</span>')
    return '<p class="badges">' + "".join(items) + "</p>"


REMOVED_BANNER = [
    '!!! warning "This check no longer reports"',
    "",
    "    It targeted a Terraform Plugin SDK v1 API that no longer exists. The ID is",
    "    retained permanently so existing `//lintignore:` comments and CI flags keep",
    "    working. See [Removed checks](../reference/removed-checks.md).",
    "",
]


def _title_block(check: dict) -> list[str]:
    """Front matter, heading, badge row and one-line summary."""
    return [
        "---",
        f'title: {check["id"]}',
        f'check_id: {check["id"]}',
        "---",
        "",
        f'# {check["id"]}',
        "",
        _badges(check),
        "",
        f'<p class="check-summary" markdown="1">{check["description"]}</p>',
        "",
    ]


def _examples_block(flagged: str | None, passing: str | None) -> list[str]:
    """
    Render the flagged/passing pair as a tab set.

    The two are a direct A/B comparison and each block runs to twenty-odd
    lines, so tabs keep them one click apart instead of pushing the passing
    example below the fold. Either side may be missing.
    """
    if flagged is None and passing is None:
        return []

    out = ["## Examples", ""]
    if flagged is not None:
        out += ['=== "Flagged"', "", textwrap.indent(flagged.strip(), "    "), ""]
    if passing is not None:
        label = "Passing" if flagged is not None else "Passing code"
        out += [f'=== "{label}"', "", textwrap.indent(passing.strip(), "    "), ""]
    return out


def _body_sections(sections: list[tuple[str, str]]) -> list[str]:
    """
    Emit the sections in the order the README wrote them.

    The flagged/passing pair collapses into a single tabbed Examples block,
    placed where the flagged section stood so that a check's Options table
    still precedes its examples.
    """
    renamed = {"ignoring reports": "Ignoring reports", "options": "Options"}
    flagged = next((c for t, c in sections if t.lower() == "flagged code"), None)
    passing = next((c for t, c in sections if t.lower() == "passing code"), None)
    anchor = "flagged code" if flagged is not None else "passing code"

    out: list[str] = []
    for title, content in sections:
        low = title.lower()
        if low == anchor:
            out += _examples_block(flagged, passing)
        elif low in ("flagged code", "passing code"):
            continue  # the other half of the pair, already tabbed above
        else:
            out += [f"## {renamed.get(low, title)}", "", content.strip(), ""]
    return out


def _source_footer(check: dict) -> list[str]:
    """Point the reader at the README this page was generated from."""
    source = check["source"]
    return [
        "---",
        "",
        (
            f"*This page is generated from [`{source}/README.md`]"
            f"({REPO_BLOB}/{source}/README.md), which lives beside the analyzer.*"
        ),
    ]


def render_check_page(check: dict, readme: str) -> str:
    """Build the markdown for one check page from its README."""
    # Alerts are converted before anything is indented into a tab: the pattern
    # is anchored to the start of a line and would stop matching once the block
    # carries a four-space tab indent.
    readme = convert_alerts(readme)

    # Drop the "# AT001" heading; the page supplies its own title block.
    body = re.sub(r"\A#\s+\S+\s*\n", "", readme).lstrip("\n")
    preamble, sections = _sections(body)

    parts = _title_block(check)
    if check["removed"]:
        parts += REMOVED_BANNER
    parts += [preamble.strip(), ""]
    parts += _body_sections(sections)
    parts += _source_footer(check)

    return "\n".join(parts)


# --------------------------------------------------------------------------
# MkDocs hooks


def on_files(files, config):
    checks = _read_registry(config.docs_dir)
    repo_root = os.path.dirname(os.path.abspath(config.config_file_path))

    for check in checks:
        readme = os.path.join(repo_root, check["source"], "README.md")
        with open(readme, encoding="utf-8") as fh:
            content = render_check_page(check, fh.read())
        files.append(
            File.generated(config, f'checks/{check["id"]}.md', content=content)
        )

    changelog = os.path.join(repo_root, "CHANGELOG.md")
    if os.path.exists(changelog):
        with open(changelog, encoding="utf-8") as fh:
            files.append(
                File.generated(
                    config, "changelog.md", content=convert_alerts(fh.read())
                )
            )

    # Templates render the category sidebar from this.
    config.extra["checks"] = checks
    return files
