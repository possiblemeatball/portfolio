/** @type {import('tailwindcss').Config} */
export default {
  content: [
      "./index.html",
      "./src/**/*.{js,vue}"
  ],
  theme: {
    extend: {
        animation: {
            'slide-top': 'slide-top 1.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both',
            'slide-bottom': 'slide-bottom 1.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both'
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
  plugins: [],
}

