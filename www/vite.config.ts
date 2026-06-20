import { defineConfig } from "vite";
import react, { reactCompilerPreset } from "@vitejs/plugin-react";
import babel from "@rolldown/plugin-babel";
import tailwind from "@tailwindcss/vite";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), babel({ presets: [reactCompilerPreset()] }), tailwind()],
  resolve: {
    tsconfigPaths: true,
  },
  server: {
    proxy: {
      "/auth": "http://localhost:3247",
      "/user": "http://localhost:3247",
      "/sprites": "http://localhost:3247",
    },
  },
});
