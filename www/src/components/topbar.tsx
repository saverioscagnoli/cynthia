import { ActionIcon, Button, useMantineColorScheme } from "@mantine/core";
import { useColorScheme } from "@mantine/hooks";
import { MonitorIcon, MoonIcon, SunIcon } from "lucide-react";

const Topbar = () => {
  const { colorScheme, toggleColorScheme } = useMantineColorScheme();

  return (
    <div className="w-full h-16 sticky top-0 flex items-center justify-between px-8 ">
      <div></div>
      <ActionIcon
        size="xl"
        variant="subtle"
        color="grape"
        onClick={toggleColorScheme}
      >
        {colorScheme == "dark" ? <SunIcon /> : <MoonIcon />}
      </ActionIcon>
    </div>
  );
};

export { Topbar };
