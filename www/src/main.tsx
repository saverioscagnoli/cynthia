import { createRoot } from "react-dom/client";
import { App } from "./App.tsx";
import { BrowserRouter } from "react-router";
import { MantineProvider } from "@mantine/core";
import { AuthProvider } from "~/contexts/auth.tsx";

import "~/index.css";
import "@mantine/core/styles.css";

createRoot(document.getElementById("root")!).render(
    <AuthProvider>
        <MantineProvider defaultColorScheme="dark">
            <BrowserRouter>
                <App />
            </BrowserRouter>
        </MantineProvider>
        ,
    </AuthProvider>,
);
