import { registerThemeToggle } from "./theme-toggle.js";

const app = () => {
    registerThemeToggle();
}

document.addEventListener('DOMContentLoaded', app);