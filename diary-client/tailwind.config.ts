import type { Config } from 'tailwindcss'

export default {
  content: ['./app/**/*.{js,ts,jsx,tsx}'],
  theme: {
    extend: {
      fontFamily: {
        "hand": ["Norda"]
      },
    },
  },
  plugins: [],
} satisfies Config

