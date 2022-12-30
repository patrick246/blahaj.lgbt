import {ReactNode, useMemo} from "react";
import {calculateLevels, Level} from "./levels";

export function ColorLevel({levels, value, children}: {levels: Level[], value: number, children: ReactNode}) {
    const classes = useMemo<string[]>(() => {
        return calculateLevels(levels, value);
    }, [levels, value]);

    return (
        <span className={classes.join(" ")}>{children}</span>
    )

}