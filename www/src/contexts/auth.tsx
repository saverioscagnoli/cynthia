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
  token: string | null;
  logout: () => void;
  logged: boolean;
  authFetch: (input: string, init?: RequestInit) => Promise<Response>;
};

const AuthContext = createContext<AuthContextT | null>(null);

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const [user, setUser] = useState<User | null>(null);
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token")
  );

  useEffect(() => {
    let params = new URLSearchParams(window.location.search);
    let t = params.get("token");

    if (t) {
      localStorage.setItem("token", t);
      setToken(t);
      window.history.replaceState({}, "", window.location.pathname);
    }
  }, []);

  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
    setUser(null);
  };

  const fetchUser = async () => {
    let result = await privateApi.getLoggedUser(token!);

    if (!result.ok) {
      if (result.error.status === 401) {
        logout();
        return;
      }

      console.error(result.error);
      return;
    }

    setUser(result.data);
  };

  useEffect(() => {
    if (!token) {
      setUser(null);
      return;
    }

    fetchUser();
  }, [token]);

  const authFetch = (input: string, init?: RequestInit) => {
    return fetch(input, {
      ...init,
      headers: {
        ...init?.headers,
        Authorization: `Bearer ${token}`
      }
    });
  };

  const updateUser = (u: Partial<User>) => {
    setUser(prev => (prev ? { ...prev, ...u } : null));
  };

  return (
    <AuthContext.Provider
      value={{
        loggedUser: user,
        updateLoggedUser: updateUser,
        token,
        logout,
        logged: user !== null,
        authFetch
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
