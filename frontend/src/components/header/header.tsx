import BlahajRing from './blahaj-ring.svg';

import './header.css';

export function Header() {
    return (
        <div className="header">
            <h1><a href="/">blahaj.lgbt</a></h1>
            <div className="blahaj">
                <img src={BlahajRing} alt="Blahaj" width="720" />
            </div>
        </div>
    );
}