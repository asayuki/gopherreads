import { setActiveMenuItem } from './site-nav.js';

document.addEventListener('DOMContentLoaded', () => {
    setActiveMenuItem();
    document.body.addEventListener('htmx:afterSettle', setActiveMenuItem);
});