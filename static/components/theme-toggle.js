class ThemeToggle extends HTMLElement {
    static register() {
        globalThis.customElements.define('theme-toggle', this);
    }

    connectedCallback() {
        if (this.shadowRoot) {
            return
        }

        this.attachShadow({ mode: 'open' });

        const theme = this.getTheme();
        this.setTheme(theme);

        const cssProperties = new CSSStyleSheet();
        cssProperties.replaceSync(`
            @property --celestial-radius {
                syntax: '<length>';
                initial-value: 0px;
                inherits: false;
            }

            @property --toggle-background-start {
                syntax: '<color>';
                initial-value: #5c9de3;
                inherits: false;
            }

            @property --toggle-background-middle {
                syntax: '<color>';
                initial-value: #5c9de3;
                inherits: false;
            }

            @property --toggle-background-end {
                syntax: '<color>';
                initial-value: #5c9de3;
                inherits: false;
            }
        `);

        document.adoptedStyleSheets = [cssProperties];

        this.shadowRoot.innerHTML = `
            <link rel="stylesheet" href="${new URL('theme-toggle.css', import.meta.url)}">
            <input type="checkbox" name="theme-toggle" id="theme-toggle" aria-label="Switch between dark and light mode" />
            <label for="theme-toggle">
                <celestial-body></celestial-body>
                <day-sky>
                    <fluffy-cloud></fluffy-cloud>
                    <fluffy-cloud></fluffy-cloud>
                    <fluffy-cloud></fluffy-cloud>
                </day-sky>
                <night-sky>
                    <distant-star></distant-star>
                    <distant-star></distant-star>
                    <distant-star></distant-star>
                    <distant-star></distant-star>
                </night-sky>
            </label>
        `;

        if (theme === 'dark') {
            this.shadowRoot.querySelector('input').checked = true;
        }

        this.shadowRoot.querySelector('input').addEventListener('change', (v) => this.setTheme(v.target.checked ? 'dark' : 'light'));
    }

    getTheme = () => {
        let theme = localStorage.getItem('theme');
        if (!theme) {
            theme = window.matchMedia('(prefers-color-scheme: dark)') ? 'dark' : 'light';
            this.setTheme(theme);
        }

        return theme;
    }

    setTheme = (theme) => {
        localStorage.setItem('theme', theme);
        document.querySelector('html').setAttribute('data-theme', theme);
    }
}

ThemeToggle.register();