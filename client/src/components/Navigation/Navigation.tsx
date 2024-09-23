import { Link } from "react-router-dom";
import Logo from "../Logo/Logo";
import "./Navigation.css";

const Navigation: React.FC = () => {


	return (
		<>
			<div id="Navigation">
				<div id="NavigationHeader">
					<Link to="/#" className="mainLogoContainer">
						<Logo />
					</Link>
				</div>
				<div className="main-navigation-links">
					<Link to="/">
						Strona główna
					</Link>
					<Link to="/rozwiazania">
						Rozwiązania
					</Link>
					<Link to="/kontakt">
						Kontakt
					</Link>
				</div>
			
			</div>
		</>
	);
};

export default Navigation;
