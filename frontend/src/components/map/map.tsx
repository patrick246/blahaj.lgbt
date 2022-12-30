import {CircleMarker, MapContainer, Popup, TileLayer} from "react-leaflet";

import "leaflet/dist/leaflet.css";
import "./map.css";
import {useStoreAvailability} from "../../client/api";
import {CountryFlag} from "../country-flag/country-flag";
import {ColorLevel} from "../color-levels/ColorLevel";
import {calculateLevels, Level} from "../color-levels/levels";

export function Map() {
    let [stores, loading] = useStoreAvailability();

    if (stores) {
        stores = stores.filter(store => store.location.longitude !== "n/a" && store.location.latitude !== "n/a")
    }

    const classLevels: Level[] = [{
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

    const colorLevels: Level[] = [{
        lower: 0,
        upper: 1,
        className: '#be0b0b'
    }, {
        lower: 1,
        upper: 5,
        className: '#ee971b'
    }, {
        lower: 5,
        upper: 20,
        className: '#ffd816'
    }, {
        lower: 20,
        upper: Number.MAX_VALUE,
        className: '#4bbd28'
    }]

    return (
        <div className="map-container">
            <MapContainer center={[49.336055, 19.19516]} zoom={4} className="map">
                <TileLayer
                    attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors'
                    url="https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png"
                />
                {
                    !loading && stores && stores.map(store =>
                        <>
                            <CircleMarker
                                center={[
                                    Number.parseFloat(store.location.longitude),
                                    Number.parseFloat(store.location.latitude),
                                ]}
                                radius={Math.max(2, Math.min((store.number / 20)+2, 25))}
                                color={calculateLevels(colorLevels, store.number).pop()}
                            >
                                <Popup>
                                    <h2>{store.store_name} <CountryFlag code={store.store_country}/></h2>
                                    <p className="popup-number"><ColorLevel levels={classLevels} value={store.number}>{store.number}</ColorLevel> 🦈</p>
                                </Popup>
                            </CircleMarker>
                        </>
                    )
                }
            </MapContainer>
        </div>
    );
}