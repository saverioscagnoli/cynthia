import { useState, useRef, useEffect } from "react";

const LazyImage = ({ src, width = 160, height = 160 }) => {
    const [visible, setVisible] = useState(false);
    const ref = useRef(null);

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
            {visible && (
                <img
                    className="[image-rendering:pixelated]"
                    src={src}
                    width={width}
                    height={height}
                />
            )}
        </div>
    );
};

const N_TRAINERS = 1454;
const TrainersPage = () => {
    return (
        <div className="flex justify-center p-4">
            <div className="grid grid-cols-[repeat(10,160px)] gap-4">
                {Array.from({ length: N_TRAINERS }, (_, i) => (
                    <LazyImage
                        key={i}
                        src={`http://localhost:9000/sprites/trainer/${i}`}
                    />
                ))}
            </div>
        </div>
    );
};

export { TrainersPage };
