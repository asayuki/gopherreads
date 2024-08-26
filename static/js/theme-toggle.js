class ThemeToggle extends HTMLElement {
    connectedCallback() {
        const theme = this.getTheme();
        this.setTheme(theme);

        this.innerHTML = `
            <input type="checkbox" name="theme-toggle" id="theme-toggle" />
            <label for="theme-toggle">
                <div class="celestial-body"></div>
            </label>
        `;

        if (theme === 'dark') {
            this.querySelector('input').checked = true;
        }

        this.querySelector('input').addEventListener('change', (v) => {
            console.log(v);
            this.setTheme(v.target.checked ? 'dark' : 'light');
        })
    }

    getTheme = () => {
        let theme = localStorage.getItem('theme');
        if (!theme) {
            theme = window.matchMedia("(prefers-color-scheme: dark)") ? 'dark' : 'light';
            this.setTheme(theme);
        }

        return theme;
    }

    setTheme = (theme) => {
        localStorage.setItem('theme', theme);
        document.querySelector('html').setAttribute('data-theme', theme);
    }
}

export const registerThemeToggle = () => customElements.define('theme-toggle', ThemeToggle);