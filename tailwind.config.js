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
            'slide-bottom': 'slide-bottom 1.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both',
            'fade-in': 'fade-in 1.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both',
            'fade-out': 'fade-out 1.5s cubic-bezier(0.250, 0.460, 0.450, 0.940) both'
        },
        keyframes: {
            'slide-top': {
                '0%': { transform: 'translateY(-100px)' },
                '100%': { transform: 'translateY(0px)' }
            },
            'slide-bottom': {
                '0%': { transform: 'translateY(100px)' },
                '100%': { transform: 'translateY(0px)' }
            },
            'fade-in': {
                '0%': { opacity: '0' },
                '100%': { opacity: '1' }
            },
            'fade-out': {
                '0%': { opacity: '1' },
                '100%': { opacity: '0' }
            }
        }
    },
  },
  plugins: [],
}

