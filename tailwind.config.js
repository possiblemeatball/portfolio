/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./index.md", "./.vitepress/theme/**/*.{js,vue}"],
  theme: {
    extend: {
      fontFamily: {
        'display': ['"Red Hat Display"', 'ui-sans-serif', 'system-ui', 'sans-serif', '"Apple Color Emoji"', '"Segoe UI Emoji"', '"Segoe UI Symbol"', '"Noto Color Emoji"'],
        'sans': ['"Red Hat Text"', 'ui-sans-serif', 'system-ui', 'sans-serif', '"Apple Color Emoji"', '"Segoe UI Emoji"', '"Segoe UI Symbol"', '"Noto Color Emoji"'],
        'mono': ['"Red Hat Mono"', 'ui-monospace', 'SFMono-Regular', 'Menlo', 'Monaco', 'Consolas', '"Liberation Mono"', '"Courier New"', 'monospace'],
        'serif': ['Mate', 'ui-serif', 'Georgia', 'Cambria', '"Times New Roman"', 'Times', 'serif'],
        'cursive': ['cursive'],
      },
      animation: {
        'slide-top': 'slide-top 1s cubic-bezier(0.250, 0.460, 0.450, 0.940) both',
        'slide-bottom': 'slide-bottom 1s cubic-bezier(0.250, 0.460, 0.450, 0.940) both'
      },
      keyframes: {
        'slide-top': {
          '0%': { transform: 'translateY(-100px)', opacity: '0' },
          '100%': { transform: 'translateY(0px)', opacity: '1' }
        },
        'slide-bottom': {
          '0%': { transform: 'translateY(100px)', opacity: '0' },
          '100%': { transform: 'translateY(0px)', opacity: '1' }
        }
      }
    },
  },
  plugins: [
      require('@tailwindcss/typography'),
  ],
}

