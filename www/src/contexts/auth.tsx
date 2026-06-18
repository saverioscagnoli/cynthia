import { createContext, useContext, useEffect, useState } from "react";
import { getTrainerMe } from "~/lib/backend";
import type { Trainer } from "~/types";

interface User {
  discordId: string;
  username: string;
  avatar: string;
  trainerId: number;
}

interface AuthContext {
  user: User | null;
  trainer: Trainer | null;
  token: string | null;
  logout: () => void;
  authFetch: (input: string, init?: RequestInit) => Promise<Response>;
}

const AuthContext = createContext<AuthContext | null>(null);

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const [token, setToken] = useState<string | null>(
    localStorage.getItem("token"),
  );

  const [user, setUser] = useState<User | null>(null);
  const [trainer, setTrainer] = useState<Trainer | null>(null);

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const t = params.get("token");
    if (t) {
      localStorage.setItem("token", t);
      setToken(t);
      window.history.replaceState({}, "", window.location.pathname);
    }
  }, []);

  // Make sure user fetch depends on token being set
  useEffect(() => {
    if (!token) {
      setUser(null);
      return;
    }

    fetch("/user/me", {
      headers: { Authorization: `Bearer ${token}` },
    })
      .then((res) => {
        if (res.status === 401) {
          logout();
          return null;
        }

        return res.json();
      })
      .then(async (data) => {
        if (data) {
          setUser(data);
          let t = getTrainerMe(data.id);
          setTrainer;
        }
      });
  }, [token]);

  const logout = () => {
    localStorage.removeItem("token");
    setToken(null);
    setUser(null);
    setTrainer(null);
  };

  const authFetch = (input: string, init?: RequestInit) => {
    return fetch(input, {
      ...init,
      headers: {
        ...init?.headers,
        Authorization: `Bearer ${token}`,
      },
    });
  };

  return (
    <AuthContext.Provider value={{ user, trainer, token, logout, authFetch }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const ctx = useContext(AuthContext);
  if (!ctx) throw new Error("useAuth must be used inside AuthProvider");
  return ctx;
};
