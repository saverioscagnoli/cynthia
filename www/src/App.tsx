import { Route } from "react-router";
import { Routes } from "react-router";
import { TrainersPage } from "~/pages/trainers";
import { Topbar } from "~/components/topbar";

const App = () => {
  return (
    <Routes>
      <Route index element={<div>wip</div>} />
      <Route path="/trainers" element={<TrainersPage />} />
    </Routes>
  );
};

export { App };
