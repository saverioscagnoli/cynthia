import { createContext, useContext, useState } from "react";

type ProfileEditContextT = {
  editing: boolean;
  startEditing: () => void;
  stopEditing: () => void;
  toggleEditing: () => void;
};

const ProfileEditContext = createContext<ProfileEditContextT | null>(null);

const ProfileEditProvider: React.FC<{ children: React.ReactNode }> = ({
  children
}) => {
  const [editing, setEditing] = useState(false);
  const startEditing = () => setEditing(true);
  const stopEditing = () => setEditing(false);
  const toggleEditing = () => setEditing(e => !e);

  return (
    <ProfileEditContext.Provider
      value={{ editing, startEditing, stopEditing, toggleEditing }}
    >
      {children}
    </ProfileEditContext.Provider>
  );
};

const useProfileEdit = () => {
  const ctx = useContext(ProfileEditContext);

  if (!ctx)
    throw new Error("useProfileEdit must be used inside ProfileEditProvider");

  return ctx;
};

export { ProfileEditProvider, useProfileEdit };
