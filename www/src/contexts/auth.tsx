import { createContext, useContext, useEffect, useState } from "react";
import { getLoggedUser } from "~/lib/backend";

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
  banner?: Blob;
  created_at: string;
};

type AuthContextT = {
  user: User | null;
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
    let res = await fetch("/user/me", {
      headers: { Authorization: `Bearer ${token}` }
    });

    if (res.status === 401) {
      logout();
      return;
    }

    try {
      let user = await getLoggedUser(token!);
      setUser(user);
    } catch (err) {
      console.error(err);
    }
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

  return (
    <AuthContext.Provider
      value={{
        user,
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
