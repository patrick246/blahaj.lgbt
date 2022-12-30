import "./store.css";
import {ColorLevel} from "../../color-levels/ColorLevel";
import {Level} from "../../color-levels/levels";

export function Store({name, number}: {name: string, number: number}) {
    const levels: Level[] = [{
        lower: 0,
        upper: 1,
        className: 'level-critical'
    }, {
        lower: 1,
        upper: 5,
        className: 'level-bad'
    }, {
        lower: 5,
        upper: 20,
        className: 'level-okay'
    }, {
        lower: 20,
        upper: Number.MAX_VALUE,
        className: 'level-good'
    }]

    return (
        <div className="store-card">
            <h1>{ name }</h1>
            <p><ColorLevel levels={levels} value={number}>{ number }</ColorLevel></p>
        </div>
    )
}