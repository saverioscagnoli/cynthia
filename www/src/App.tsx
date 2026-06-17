import { Route, Routes } from "react-router";
import { TrainersPage } from "~/pages/trainers";
import { Topbar } from "~/components/topbar";
import { cn } from "./lib/utils";

const App = () => {
    return (
        <div className={cn("w-screen h-screen", "flex flex-col")}>
            <Topbar />
            <Routes>
                <Route index element={<div>wip</div>} />
                <Route path="/trainers" element={<TrainersPage />} />
            </Routes>
        </div>
    );
};

export { App };
