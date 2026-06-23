import { App } from "~/app";
import { AuthProvider } from "~/contexts/auth";
import { ThemeProvider } from "~/contexts/theme";
import { createRoot } from "react-dom/client";
import { BrowserRouter } from "react-router";

import "~/index.css";

createRoot(document.getElementById("root")!).render(
  <ThemeProvider>
    <AuthProvider>
      <BrowserRouter>
        <App />
      </BrowserRouter>
    </AuthProvider>
  </ThemeProvider>
);
