import { Route, Routes, useLocation } from "react-router";
import { AccountPage } from "~/pages/account";
import { HomePage } from "~/pages/home";
import { Topbar } from "~/components/topbar";
import { cn } from "~/lib/utils";
import { AccountProvider } from "./contexts/account";

const App = () => {
  const location = useLocation();
  return (
    <div className={cn("h-screen w-screen")}>
      <Topbar />
      <div key={location.pathname} className={cn("h-full", "page-transition")}>
        <Routes>
          <Route index element={<HomePage />} />
          <Route
            path="/user/:userId"
            element={
              <AccountProvider>
                <AccountPage />
              </AccountProvider>
            }
          />
        </Routes>
      </div>
    </div>
  );
};

export { App };
