import {CountryFlag} from "../../country-flag/country-flag";

import "./country.css";
import {Link} from "react-router-dom";

export function Country({code, number}: {code: string, number: number}) {
    return (
        <Link to={code} className="country-card-link">
            <div className="country-card">
                <h1 title={code.toUpperCase()}><CountryFlag code={code} /></h1>
                <p>{ number }</p>
            </div>
        </Link>
    )
}