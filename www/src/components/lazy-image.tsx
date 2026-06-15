import { useEffect, useRef, useState, type ComponentProps } from "react";

type LazyImageProps = ComponentProps<"img">;

const LazyImage: React.FC<LazyImageProps> = ({
  src,
  width,
  height,
  ...props
}) => {
  const [visible, setVisible] = useState<boolean>(false);
  const ref = useRef<HTMLDivElement | null>(null);

  useEffect(() => {
    const observer = new IntersectionObserver(
      ([entry]) => {
        if (entry.isIntersecting) {
          setVisible(true);
          observer.disconnect();
        }
      },
      { rootMargin: "200px" },
    );

    if (ref.current) observer.observe(ref.current);

    return () => observer.disconnect();
  }, []);

  return (
    <div
      ref={ref}
      style={{ width, height }}
      className="flex items-center justify-center"
    >
      {visible && <img src={src} width={width} height={height} {...props} />}
    </div>
  );
};

export { LazyImage };
