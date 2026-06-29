/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./projects/rent-manager-frontend/src/**/*.{html,ts,scss}'],
  theme: {
    extend: {
      colors: {
        rent: {
          50: '#f8fafc',
          100: '#f1f5f9',
          500: '#64748b',
          700: '#334155',
          900: '#0f172a',
        },
      },
    },
  },
  plugins: [],
};
