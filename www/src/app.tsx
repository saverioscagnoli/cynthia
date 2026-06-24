import { Route, Routes, useLocation } from "react-router";
import { AccountPage } from "~/pages/account";
import { HomePage } from "~/pages/home";
import { Topbar } from "~/components/topbar";
import { cn } from "~/lib/utils";
import { AccountProvider } from "./contexts/account";

const App = () => {
  const location = useLocation();

  return (
    <div className={cn("flex h-screen w-screen flex-col")}>
      <Topbar />
      <div
        key={location.pathname}
        className={cn("page-transition flex-1 overflow-y-auto")}
      >
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
