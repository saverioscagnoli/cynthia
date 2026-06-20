import type { Dispatch } from "react";
import { createContext, useContext, useEffect, useState } from "react";

type Theme = "auto" | "light" | "dark";

type ThemeContextT = {
  theme: Theme;
  setTheme: Dispatch<Theme>;
  toggle: () => void;
};

const ThemeContext = createContext<ThemeContextT | null>(null);

const ThemeProvider: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const [theme, setTheme] = useState<Theme>("auto");

  useEffect(() => {
    if (
      theme === "auto" &&
      window.matchMedia("(prefers-color-scheme: dark)").matches
    ) {
      document.documentElement.className = "dark";
      setTheme("dark");
    } else {
      document.documentElement.className = theme;
    }
  }, [theme]);

  const toggleTheme = () => setTheme(t => (t === "dark" ? "light" : "dark"));

  return (
    <ThemeContext.Provider value={{ theme, setTheme, toggle: toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

const useTheme = () => {
  const ctx = useContext(ThemeContext);

  if (!ctx) throw new Error("useTheme must be used inside ThemeProvider");

  return ctx;
};

export { ThemeProvider, useTheme };
