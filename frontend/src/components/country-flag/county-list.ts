
export function getFlag(countryCode: string): string {
    countryCode = countryCode.toUpperCase()

    const length = [...countryCode].length;
    if (length != 2) {
        return "";
    }

    const firstCode = countryCode.codePointAt(0)
    if (firstCode === undefined) {
        return "";
    }

    const secondCode = countryCode.codePointAt(1)
    if (secondCode === undefined) {
        return "";
    }

    return String.fromCodePoint(firstCode - 65 + 0x1F1E6, secondCode - 65 + 0x1F1E6)
}