import franken from "franken-ui/shadcn-ui/preset-quick";

/** @type {import('tailwindcss').Config} */
export default {
  presets: [
    franken({
      theme: "zinc",
      except: [],
    }),
  ],
  content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
  safelist: [
    {
      pattern: /^uk-/,
    },
  ],
  theme: {
    extend: {},
  },
  plugins: [],
  darkMode: "selector",
};
