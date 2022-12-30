import "./frontpagenav.css";
import {Link} from "react-router-dom";

export function FrontPageNav() {
    return (
        <div className="frontpage-nav">
            <div className="nav-item">
                <Link to="/map">
                    <div className="emoji">🗺️</div>
                    Map
                </Link>
            </div>

            <div className="nav-item">
                <Link to="/stocklevels">
                    <div className="emoji">📈</div>
                    Stock levels
                </Link>
            </div>
        </div>
    );
}