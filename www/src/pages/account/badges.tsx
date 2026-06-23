import type React from "react";
import { CalendarIcon } from "@radix-ui/react-icons";
import { Badge } from "~/components/badge";
import { cn } from "~/lib/utils";
import type { User } from "~/contexts/auth";

type AccountBadgesProps = {
  user: User;
};

const AccountBadges: React.FC<AccountBadgesProps> = ({ user }) => {
  let joinedAt = new Date(user.created_at).toLocaleString("en-US", {
    month: "long",
    year: "numeric"
  });

  return (
    <div className={cn("flex items-center gap-2")}>
      <Badge leftIcon={<CalendarIcon />} label={`Joined ${joinedAt}`} />
    </div>
  );
};

export { AccountBadges };
