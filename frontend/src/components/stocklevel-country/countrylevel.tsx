import {Link, useNavigate, useParams} from "react-router-dom";
import {useStoreAvailability} from "../../client/api";
import {Store} from "./store/store";
import {CountryFlag} from "../country-flag/country-flag";

import "./countrylevel.css";

export function CountryLevel() {
    const navigate = useNavigate();
    const {country} = useParams();

    if (!country) {
        navigate("../");
        return (<></>);
    }

    const [availability, isLoading] = useStoreAvailability(country);

    if (!availability || isLoading) {
        return <div>Loading...</div>;
    }

    return (
        <div className="container">
            <Link to="/stocklevels" className="country-backlink">← Back to countries</Link>
            <h1>Where can you get your Blahaj?</h1>
            <p>This page shows the stock levels at the stores in <CountryFlag code={country} />.</p>
            <div className="countrylevels">
                {
                    availability.map(a => <Store key={a.store_id} name={a.store_name} number={a.number} />)
                }
            </div>
        </div>
    );
}