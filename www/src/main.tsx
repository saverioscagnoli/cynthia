import { createRoot } from "react-dom/client";
import { App } from "./App.tsx";

import "~/index.css";
import { BrowserRouter } from "react-router";
import { ThemeContextProvider } from "~/contexts/theme";

createRoot(document.getElementById("root")!).render(
    <ThemeContextProvider>
        <BrowserRouter>
            <App />
        </BrowserRouter>
    </ThemeContextProvider>,
);
