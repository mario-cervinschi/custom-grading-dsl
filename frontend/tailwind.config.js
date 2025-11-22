/** @type {import('tailwindcss').Config} */
module.exports = {
    content: [
        "./src/**/*.{html,ts}",
    ],
    theme: {
        extend: {
            fontFamily: {
                os: ['"JetBrains Mono"', 'monospace'],
            },
        },
    },
    plugins: [],
}