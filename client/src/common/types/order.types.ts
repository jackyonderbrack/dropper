import { Customer } from "./customer.types";
import { Product } from "./product.types";

// Główne dane zamówienia
export interface Order {
	id: string;
	products: Product[];
	customer: Customer;
	totalAmount: number;
	paymentStatus: "paid" | "pending" | "cancelled";
	orderStatus: "pending" | "shipped" | "delivered" | "cancelled";
	createdAt: string;
	updatedAt?: string;
}
