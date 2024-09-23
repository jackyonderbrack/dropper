import "./footer.css";
import Logo from "../Logo/Logo";

import { Link } from "react-router-dom";
import Button from "../Button/Button";
import PrivacyPolicy from "../PrivacyPolicy/PrivacyPolicy";

const Footer = () => {
	return (
		<footer id="Footer" className="background-top px-4">
			<section className="footerHeader grid grid-cols-4 justify-items-center align-items-center gap-2 px-4 py-2">
				<div className="footerHeaderCol">
					<Link to="/#" className="footerLogoContainer">
						<Logo />
					</Link>
				</div>
				<div className="footerHeaderCol">
					<nav className="footerNav flex flex-col">
						<Link to="/">Strona główna</Link>
						<Link to="/rozwiazania">Rozwiązania</Link>
						<Link to="/kontakt">Kontakt</Link>
					</nav>
				</div>
				<div className="footerHeaderCol">
					<div className="social_icons">
						<div className="socialIconContainer">
							<a href="example.com">Social 1</a>
						</div>
						<div className="socialIconContainer">
							<a href="example.com" target="_blank">
								Social 2
							</a>
						</div>
					</div>
				</div>
				<div className="footerHeaderCol">
					<Button linkTo="/kontakt" buttonText="Skontaktuj się z nami" theme="btn-outline" />
				</div>
			</section>
			<section className="footerBottom px-4">
				<PrivacyPolicy />
			</section>
		</footer>
	);
};

export default Footer;
