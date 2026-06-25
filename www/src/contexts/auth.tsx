import { createContext, useContext, useEffect, useState } from "react";
import { privateApi } from "~/lib/wrapper";

type BagItem = {
  item_id: number;
  name: string;
  quantity: number;
  cost: number;
  fling_power?: number;
  fling_effect?: string;
};

type User = {
  id: string;
  username: string;
  discord_username: string;
  avatar_hash: string;
  money: number;
  sprite_id?: number;
  bag: BagItem[];
  banner: number[] | null;
  created_at: string;
};

type AuthContextT = {
  loggedUser: User | null;
  updateLoggedUser: (u: Partial<User>) => void;
  logout: () => void;
  logged: boolean;
};

const AuthContext = createContext<AuthContextT | null>(null);

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const [user, setUser] = useState<User | null>(null);

  const logout = async () => {
    await fetch("/api/auth/logout", { method: "POST", credentials: "include" });
    setUser(null);
  };

  const fetchUser = async () => {
    const result = await privateApi.getLoggedUser();
    if (!result.ok) {
      if (result.error.status === 401) {
        setUser(null);
        return;
      }
      console.error(result.error);
      return;
    }
    setUser(result.data);
  };

  useEffect(() => {
    fetchUser();
  }, []);

  const updateUser = (u: Partial<User>) => {
    setUser(prev => (prev ? { ...prev, ...u } : null));
  };

  return (
    <AuthContext.Provider
      value={{
        loggedUser: user,
        updateLoggedUser: updateUser,
        logout,
        logged: user !== null
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used inside AuthProvider");
  return ctx;
};

export { AuthContext, AuthProvider, useAuth };
export type { User, AuthContextT };
