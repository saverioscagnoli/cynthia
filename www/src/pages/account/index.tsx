import { useEffect, useState } from "react";
import { useParams } from "react-router";
import { cn } from "~/lib/utils";
import { publicApi } from "~/lib/wrapper";
import { useAccount } from "~/contexts/account";
import { FirstCardRow } from "./first-row";
import { UserCard } from "./user-card";

const Account = () => {
  return (
    <div className={cn("h-full w-250", "flex flex-col gap-4", "mx-auto my-12")}>
      <UserCard />
      <FirstCardRow />
    </div>
  );
};

const AccountPage = () => {
  const { userId } = useParams();
  const { maybeUser, setUser } = useAccount();
  const [err, setErr] = useState<string | null>(null);

  const fetchUser = async () => {
    if (!userId) {
      setErr("404");
      return;
    }

    let res = await publicApi.getUser(userId);

    if (!res.ok) {
      setErr(res.error.message);
      return;
    }

    setUser(res.data);
  };

  useEffect(() => {
    fetchUser();
  }, [userId]);

  if (err) {
    return <div>Error {err}</div>;
  } else if (!maybeUser) {
    return <div>Loading...</div>;
  } else {
    return <Account />;
  }
};

export { AccountPage };
