import {Route, Routes} from "react-router-dom";
import {Header} from './components/header/header'
import {FrontPageNav} from "./components/frontpagenav/frontpagenav";
import {Stocklevels} from "./components/stocklevel-global/stocklevels";
import {CountryLevel} from "./components/stocklevel-country/countrylevel";
import {Map} from "./components/map/map";
import {Footer} from "./components/footer/footer";

export function App() {
    return (
        <>
            <Header />
            <Routes>
                <Route path="/" element={<FrontPageNav />} />
                <Route path="map" element={<Map />} />
                <Route path="stocklevels" element={<Stocklevels />} />
                <Route path="stocklevels/:country" element={<CountryLevel />} />
            </Routes>
            <Footer />
        </>
    )
}