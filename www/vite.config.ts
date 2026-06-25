import { defineConfig } from "vite";
import react, { reactCompilerPreset } from "@vitejs/plugin-react";
import babel from "@rolldown/plugin-babel";
import tailwind from "@tailwindcss/vite";
import dotenv from "dotenv";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

dotenv.config({ path: path.resolve(__dirname, "../.env") });

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), babel({ presets: [reactCompilerPreset()] }), tailwind()],
  resolve: {
    tsconfigPaths: true,
  },
  server: {
    proxy: {
      "/api": "http://localhost:3247",
      //"/api": process.env.BASE_URL,
    },
  },
  define: {
    __API_URL__: JSON.stringify(process.env.BASE_URL),
  },
});
