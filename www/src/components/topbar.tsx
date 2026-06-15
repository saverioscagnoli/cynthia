import {
    ActionIcon,
    Avatar,
    Button,
    Text,
    useMantineColorScheme,
} from "@mantine/core";
import { MoonIcon, SunIcon } from "lucide-react";
import { useAuth } from "~/contexts/auth";

const Topbar = () => {
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();
    const { user, logout } = useAuth();

    const discordAvatarUrl = user?.avatar
        ? `https://cdn.discordapp.com/avatars/${user.discordId}/${user.avatar}.png`
        : null;

    return (
        <div className="w-full h-16 sticky top-0 flex items-center justify-between px-8">
            <span className="font-semibold text-lg">Cynthia</span>

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
                    size="xl"
                    variant="subtle"
                    color="grape"
                    onClick={toggleColorScheme}
                >
                    {colorScheme === "dark" ? <SunIcon /> : <MoonIcon />}
                </ActionIcon>
            </div>
        </div>
    );
};

export { Topbar };
