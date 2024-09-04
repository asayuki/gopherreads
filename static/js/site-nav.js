export const setActiveMenuItem = () => {
    const path = window.location.pathname;
    document.querySelectorAll('nav.main li').forEach((item) => {
        const itemLink = item.querySelector('a');
        const itemPath = itemLink.getAttribute('hx-get');

        itemLink.setAttribute('data-active', itemPath === path ? true : false);
    });
}