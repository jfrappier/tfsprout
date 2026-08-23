/*
  Three small behaviours the Cinder theme does not provide:
  a persisted light/dark toggle, copy buttons on code blocks, and a filter
  over the check catalogue in the sidebar.

  The initial theme is applied by an inline script in
  docs-theme/overrides/main.html so that it lands before first paint; this file
  only handles the toggle.

  Wrapped in an IIFE so nothing here reaches the global scope.
*/
(() => {
    const STORAGE_KEY = 'tfsprout-theme';

    // ---------------------------------------------------------- theme toggle

    const initThemeToggle = () => {
        const toggle = document.getElementById('theme-toggle');
        if (!toggle) return;

        toggle.addEventListener('click', (event) => {
            event.preventDefault();
            const root = document.documentElement;
            const next = root.getAttribute('data-theme') === 'dark' ? 'light' : 'dark';
            root.setAttribute('data-theme', next);
            try {
                localStorage.setItem(STORAGE_KEY, next);
            } catch {
                /* private browsing: the choice just will not persist */
            }
        });
    };

    // ----------------------------------------------------------- copy button

    const initCopyButtons = () => {
        if (!navigator.clipboard) return;

        for (const block of document.querySelectorAll('.highlight')) {
            const code = block.querySelector('pre');
            if (!code) continue;

            const button = document.createElement('button');
            button.type = 'button';
            button.className = 'copy-button';
            button.textContent = 'Copy';
            button.setAttribute('aria-label', 'Copy this code to the clipboard');

            button.addEventListener('click', () => {
                navigator.clipboard.writeText(code.innerText).then(() => {
                    button.textContent = 'Copied';
                    button.classList.add('copied');
                    setTimeout(() => {
                        button.textContent = 'Copy';
                        button.classList.remove('copied');
                    }, 1600);
                });
            });

            block.appendChild(button);
        }
    };

    // ---------------------------------------------------------- check filter

    const initCheckFilter = () => {
        const input = document.getElementById('check-filter-input');
        if (!input) return;

        const items = document.querySelectorAll('.check-list .check-item');
        const groups = document.querySelectorAll('.check-list .check-group');
        const empty = document.querySelector('.check-list-empty');

        const apply = () => {
            const query = input.value.trim().toLowerCase();
            let matches = 0;

            for (const item of items) {
                const hit = !query || item.dataset.check.includes(query);
                item.hidden = !hit;
                if (hit) matches++;
            }

            // A group heading is only meaningful while something under it shows.
            for (const group of groups) {
                let visible = false;
                let node = group.nextElementSibling;
                while (node && !node.classList.contains('check-group')) {
                    if (!node.hidden) {
                        visible = true;
                        break;
                    }
                    node = node.nextElementSibling;
                }
                group.hidden = !visible;
            }

            if (empty) empty.hidden = matches !== 0;
        };

        input.addEventListener('input', apply);

        // Escape clears the filter rather than leaving the list narrowed.
        input.addEventListener('keydown', (event) => {
            if (event.key === 'Escape') {
                input.value = '';
                apply();
            }
        });
    };

    // ------------------------------------------------------------------ init

    const init = () => {
        initThemeToggle();
        initCopyButtons();
        initCheckFilter();
    };

    if (document.readyState === 'loading') {
        document.addEventListener('DOMContentLoaded', init);
    } else {
        init();
    }
})();
