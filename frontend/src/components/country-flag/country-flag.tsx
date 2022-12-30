import {getFlag} from "./county-list";

export function CountryFlag({code}: {code: string}) {
    return (<>{getFlag(code)}</>)
}