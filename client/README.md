# React + TypeScript + Vite

# Bez CSS

## Client:

/src/client
/components
/Cart - Cart.tsx // Komponent koszyka
/ProductCard - ProductCard.tsx // Komponent wyświetlający pojedynczy produkt
/ProductsList - ProductsList.tsx // Komponent wyświetlający listę produktów
/TopNavigation - TopNavigation.tsx // Komponent górnej nawigacji

/pages
/StorePage - StorePage.tsx // Strona sklepu
/SingleProductPage - SingleProductPage.tsx // Strona pojedynczego produktu
/OrderPage - OrderPage.tsx // Strona zamówienia
/PaymentPage - PaymentPage.tsx // Strona płatności
/AccountPage - AccountPage.tsx // Strona konta użytkownika

/services
/products - productsService.ts // Usługi związane z produktami
/account - accountService.ts // Usługi związane z kontem użytkownika
/order - orderService.ts // Usługi związane z zamówieniami
/auth - authService.ts // Usługi związane z autoryzacją

/hooks - useCart.ts // Hook do zarządzania stanem koszyka

/store - cartStore.ts // Zarządzanie stanem koszyka - userStore.ts // Zarządzanie stanem użytkownika

/utils - formatPrice.ts // Narzędzie do formatowania cen - calculateTotal.ts // Narzędzie do obliczania całkowitej kwoty zamówienia

/assets - logo.svg // Zasoby, takie jak obrazy i ikony

/types - productTypes.ts // Typy związane z produktami - orderTypes.ts // Typy związane z zamówieniami

## Admin:

Komponenty będą generyczne, będą się wyświetlać jak w wordpressie, na każdej stronie będzie ta sama lista z przyciskami zarządzaj/edytuj/usuń, tylko wartosci będą inne - w ordersPage będą zamówienia, w clientsPage będą klienci itp.

/src/admin
/components
/List - List.tsx // Generyczny komponent listy (zarządzanie, edycja, usuwanie)
/Sidenav - Sidenav.tsx // Komponent bocznej nawigacji
/Header - Header.tsx // Komponent nagłówka

/pages
/OrdersPage - OrdersPage.tsx // Strona zamówień
/ProductsPage - ProductsPage.tsx // Strona produktów
/ClientsPage - ClientsPage.tsx // Strona klientów
/ReportsPage - ReportsPage.tsx // Strona raportów
/SettingsPage - SettingsPage.tsx // Strona ustawień

/services
/httpService - httpService.ts // Konfiguracja HTTP (np. Axios)
/auth - authService.ts // Usługi związane z autoryzacją admina
/users - usersService.ts // Usługi związane z użytkownikami (klientami)
/products - productsService.ts // Usługi związane z produktami
/orders - ordersService.ts // Usługi związane z zamówieniami
/payments - paymentsService.ts // Usługi związane z płatnościami

/hooks - useOrders.ts // Hook do zarządzania stanem zamówień

/store - adminStore.ts // Centralne zarządzanie stanem admina

/utils - formatDate.ts // Narzędzie do formatowania dat - calculateRevenue.ts // Narzędzie do obliczania przychodu

/types - orderTypes.ts // Typy związane z zamówieniami - userTypes.ts // Typy związane z użytkownikami

## Common:

/src/common
/services
/apiClient.ts // Klient HTTP (np. Axios) z interceptorami

/utils - validateEmail.ts // Narzędzia wspólne, np. walidacja e-mail - currencyFormatter.ts // Formatowanie cen

/types - productTypes.ts // Typy produktów - userTypes.ts // Typy użytkowników
