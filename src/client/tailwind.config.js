/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'dark-grey': '#202123',
        'grey-chat':'#343541',
        'light-grey': '#434654',
      },
    },
  },
  plugins: [],
}