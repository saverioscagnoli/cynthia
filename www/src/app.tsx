import { Route, Routes } from "react-router";
import { AccountPage } from "~/pages/account";
import { HomePage } from "~/pages/home";
import { Topbar } from "~/components/topbar";
import { cn } from "~/lib/utils";

const App = () => {
  return (
    <div className={cn("h-screen w-screen")}>
      <Topbar />
      <Routes>
        <Route index element={<HomePage />} />
        <Route path="/account" element={<AccountPage />} />
      </Routes>
    </div>
  );
};

export { App };
