import { Route, Routes } from "react-router";
import { TrainersPage } from "~/pages/trainers";
import { Topbar } from "~/components/topbar";

const Layout = ({ children }: { children: React.ReactNode }) => (
    <div className="flex flex-col h-screen">
        <Topbar />
        <div className="flex-1 min-h-0">{children}</div>
    </div>
);

const App = () => {
    return (
        <Routes>
            <Route
                element={
                    <Layout>
                        <div>wip</div>
                    </Layout>
                }
                index
            />
            <Route
                path="/trainers"
                element={
                    <Layout>
                        <TrainersPage />
                    </Layout>
                }
            />
        </Routes>
    );
};

export { App };
