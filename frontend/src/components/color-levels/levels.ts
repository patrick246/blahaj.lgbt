export interface Level {
    className: string;
    lower: number;
    upper: number
}

export function calculateLevels(levels: Level[], value: number): string[] {
    const c = [];
    for (const l of levels) {
        if (value >= l.lower && value < l.upper) {
            c.push(l.className);
        }
    }

    return c;
}