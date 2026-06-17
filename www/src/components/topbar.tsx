import {
    ActionIcon,
    Avatar,
    Button,
    Text,
    useMantineColorScheme,
} from "@mantine/core";
import { MoonIcon, SunIcon } from "lucide-react";
import { useRef } from "react";
import { useAuth } from "~/contexts/auth";
import { cn } from "~/lib/utils";

const Topbar = () => {
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();
    const { user, logout } = useAuth();

    const discordAvatarUrl = user?.avatar
        ? `https://cdn.discordapp.com/avatars/${user.discordId}/${user.avatar}.png`
        : null;

    return (
        <div className="w-full h-18 sticky top-0 flex items-center justify-between px-8">
            <img
                src="/cynthia-overworld.png"
                width={64}
                height={64}
                className={cn("mb-3")}
                style={{ imageRendering: "pixelated" }}
            />

            <div className="flex items-center gap-3">
                {user ? (
                    <>
                        {discordAvatarUrl && (
                            <Avatar
                                src={discordAvatarUrl}
                                radius="xl"
                                size="sm"
                            />
                        )}
                        <Text size="sm">{user.username}</Text>
                        <Button
                            variant="subtle"
                            color="red"
                            size="xs"
                            onClick={logout}
                        >
                            Logout
                        </Button>
                    </>
                ) : (
                    <Button
                        component="a"
                        href="/auth/login"
                        variant="filled"
                        color="indigo"
                        size="sm"
                    >
                        Login with Discord
                    </Button>
                )}

                <ActionIcon
                    size="lg"
                    variant="subtle"
                    color="grape"
                    onClick={toggleColorScheme}
                >
                    {colorScheme === "dark" ? (
                        <SunIcon size={18} />
                    ) : (
                        <MoonIcon size={18} />
                    )}
                </ActionIcon>
            </div>
        </div>
    );
};

export { Topbar };
