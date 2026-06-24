import {
  createContext,
  useContext,
  useRef,
  useState,
  type Dispatch,
  type SetStateAction
} from "react";
import type { User } from "./auth";

type AccountContextT = {
  maybeUser: User | null;
  user: User;
  setUser: Dispatch<SetStateAction<User | null>>;
  updateUser: (user: Partial<User>) => void;
  isEditing: boolean;
  startEdit: () => void;
  stopEdit: () => void;
  toggleEdit: () => void;
  registerOnEditConfirm: (fn: () => void) => () => void;
  onEditConfirm: () => void;
};

const AccountContext = createContext<AccountContextT | null>(null);

const AccountProvider: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const [user, setUser] = useState<User | null>(null);
  const [editing, setEditing] = useState<boolean>(false);
  const confirmHandlers = useRef<Set<() => void>>(new Set());

  const updateUser = (us: Partial<User>) => {
    setUser(u => (u ? { ...u, ...us } : u));
  };

  const startEdit = () => setEditing(true);
  const stopEdit = () => setEditing(false);
  const toggleEdit = () => setEditing(e => !e);

  const registerOnEditConfirm = (fn: () => void) => {
    confirmHandlers.current.add(fn);
    return () => confirmHandlers.current.delete(fn); // returns unregister
  };

  const onEditConfirm = () => {
    confirmHandlers.current.forEach(fn => fn());
    stopEdit();
  };

  return (
    <AccountContext.Provider
      value={{
        maybeUser: user,
        get user() {
          return user!;
        },
        setUser,
        updateUser,
        isEditing: editing,
        startEdit,
        stopEdit,
        toggleEdit,
        registerOnEditConfirm,
        onEditConfirm
      }}
    >
      {children}
    </AccountContext.Provider>
  );
};

const useAccount = () => {
  const ctx = useContext(AccountContext);

  if (!ctx) throw new Error("useAccount must be used within AccountProvider");

  return ctx;
};

export { AccountProvider, useAccount };
