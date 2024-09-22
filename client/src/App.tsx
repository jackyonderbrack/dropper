import { Form, Outlet, RouterProvider, createBrowserRouter } from "react-router-dom";
import Sidenav from "./admin/components/Sidenav/Sidenav";
import Login from "./admin/pages/login/Login";
import PrivateRoute from "./admin/components/PrivateRoute/PrivateRoute";
import { AccessProvider } from "./admin/contexts/AccessContext";

// function Layout() {
// 	const location = useLocation();

// 	useEffect(() => {
// 		ReactGA.send({
// 			hitType: "pageview",
// 			page: location.pathname + location.search,
// 		});
// 	}, [location]);

// 	return (
// 		<div id="Application">
// 			<Navigation />
// 			<div className="client-outlet">
// 				<ScrollToTop />
// 				<Outlet />
// 				<Footer />
// 				<TrackPageView />
// 				<CustomCookieConsent />
// 			</div>
// 		</div>
// 	);
// }

function AdminLayout() {
	return (
		<main id="Admin">
			<Sidenav />
			<div className="admin-outlet">
				<Outlet />
			</div>
		</main>
	);
}

const router = createBrowserRouter([
	// {
	// 	path: "/",
	// 	element: <Layout />, // Główny layout aplikacji
	// 	children: [
	// 		{ index: true, element: <HomePage /> },
	// 		{ path: "rozwiazania", element: <SolutionsPage /> },
	// 		{ path: "kontakt", element: <ContactPage /> },
	// 		{ path: "wyslano", element: <ContactPageSent /> },
	// 		{ path: "polityka-prywatnosci", element: <PrivacyPolicyPage /> },
	// 		{ path: "nowosci", element: <NewsPage /> },
	// 	],
	// },
	{
		path: "/login",
		element: <Login />,
	},
	{
		path: "/admin",
		element: <PrivateRoute />,
		children: [
			{
				element: <AdminLayout />,
				children: [
					// { index: true, element: <Dashboard /> },
					// { path: "posts", element: <Posts /> },
					// { path: "users", element: <Users /> },
					// { path: "reports", element: <Reports /> },
					// { path: "settings", element: <Settings /> },
					// { path: "account", element: <Account /> },
					{ path: ":type/new", element: <Form /> },
				],
			},
		],
	},
	// { path: "*", element: <NotFoundPage /> }, // Dla nieznanych ścieżek
]);

function App() {
	return (
		<div id="App">
			<AccessProvider>
				<RouterProvider router={router} />
			</AccessProvider>
		</div>
	);
}

export default App;
