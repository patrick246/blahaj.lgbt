import {Country} from "./country/country";
import {useCountryAvailability} from "../../client/api";

import "./stocklevels.css";

export function Stocklevels() {
    const [availability, isLoading] = useCountryAvailability();

    if (availability == undefined || isLoading) {
        return (
            <div className="container">Loading...</div>
        )
    }

    return (
        <div className="container">
            <h1>Where can you get your Blahaj?</h1>
            <p>This page shows you the stock levels in different countries. Click on a country to get individual store stock levels.</p>
            <div className="stocklevels">
                {
                    availability.map(entry => <Country key={entry.country} code={entry.country} number={entry.number} />)
                }
            </div>
        </div>
    )
}