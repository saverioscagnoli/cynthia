import { BsDiscord } from "react-icons/bs";
import { Button } from "~/components/ui";
import { cn } from "~/lib/utils";

const HomePage = () => {
  return (
    <div
      className={cn(
        "h-[calc(100vh-5rem)] w-full",
        "flex flex-col items-center justify-center gap-8"
      )}
    >
      <img
        width={80 * 4}
        height={80 * 4}
        src="https://play.pokemonshowdown.com/sprites/trainers/cynthia-gen4.png"
        style={{ imageRendering: "pixelated" }}
      />
      <Button size="lg" colorScheme="blue" leftIcon={<BsDiscord />}>
        Invite to your Discord server!
      </Button>
    </div>
  );
};

export { HomePage };
