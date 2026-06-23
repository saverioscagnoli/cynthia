import { Route, Routes, useLocation } from "react-router";
import { AccountPage } from "~/pages/account";
import { HomePage } from "~/pages/home";
import { Topbar } from "~/components/topbar";
import { cn } from "~/lib/utils";
import { ProfileEditProvider } from "~/contexts/profile-edit";

const App = () => {
  const location = useLocation();
  return (
    <div className={cn("h-screen w-screen")}>
      <Topbar />
      <div key={location.pathname} className={cn("h-full", "page-transition")}>
        <Routes>
          <Route index element={<HomePage />} />
          <Route
            path="/account"
            element={
              <ProfileEditProvider>
                <AccountPage />
              </ProfileEditProvider>
            }
          />
        </Routes>
      </div>
    </div>
  );
};

export { App };
