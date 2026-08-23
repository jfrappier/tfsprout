/*
  Three small behaviours the Cinder theme does not provide:
  a persisted light/dark toggle, copy buttons on code blocks, and a filter
  over the check catalogue in the sidebar.

  The initial theme is applied by an inline script in docs-theme/overrides/main.html
  so that it lands before first paint; this file only handles the toggle.
*/
(function () {
    'use strict';

    var STORAGE_KEY = 'tfsprout-theme';

    // ---------------------------------------------------------- theme toggle

    function initThemeToggle() {
        var toggle = document.getElementById('theme-toggle');
        if (!toggle) return;

        toggle.addEventListener('click', function (event) {
            event.preventDefault();
            var root = document.documentElement;
            var next = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
            root.setAttribute('data-theme', next);
            try {
                localStorage.setItem(STORAGE_KEY, next);
            } catch (e) {
                /* private browsing: the choice just will not persist */
            }
        });
    }

    // ----------------------------------------------------------- copy button

    function initCopyButtons() {
        if (!navigator.clipboard) return;

        var blocks = document.querySelectorAll('.highlight');
        Array.prototype.forEach.call(blocks, function (block) {
            var code = block.querySelector('pre');
            if (!code) return;

            var button = document.createElement('button');
            button.type = 'button';
            button.className = 'copy-button';
            button.textContent = 'Copy';
            button.setAttribute('aria-label', 'Copy this code to the clipboard');

            button.addEventListener('click', function () {
                navigator.clipboard.writeText(code.innerText).then(function () {
                    button.textContent = 'Copied';
                    button.classList.add('copied');
                    setTimeout(function () {
                        button.textContent = 'Copy';
                        button.classList.remove('copied');
                    }, 1600);
                });
            });

            block.appendChild(button);
        });
    }

    // ---------------------------------------------------------- check filter

    function initCheckFilter() {
        var input = document.getElementById('check-filter-input');
        if (!input) return;

        var items = document.querySelectorAll('.check-list .check-item');
        var groups = document.querySelectorAll('.check-list .check-group');
        var empty = document.querySelector('.check-list-empty');

        function apply() {
            var query = input.value.trim().toLowerCase();
            var matches = 0;

            Array.prototype.forEach.call(items, function (item) {
                var hit = !query || item.dataset.check.indexOf(query) !== -1;
                item.hidden = !hit;
                if (hit) matches++;
            });

            // A group heading is only meaningful while something under it shows.
            Array.prototype.forEach.call(groups, function (group) {
                var visible = false;
                var node = group.nextElementSibling;
                while (node && !node.classList.contains('check-group')) {
                    if (!node.hidden) { visible = true; break; }
                    node = node.nextElementSibling;
                }
                group.hidden = !visible;
            });

            if (empty) empty.hidden = matches !== 0;
        }

        input.addEventListener('input', apply);

        // Escape clears the filter rather than leaving the list narrowed.
        input.addEventListener('keydown', function (event) {
            if (event.key === 'Escape') {
                input.value = '';
                apply();
            }
        });
    }

    // ------------------------------------------------------------------ init

    function init() {
        initThemeToggle();
        initCopyButtons();
        initCheckFilter();
    }

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
