import { Address } from "./adress.types";

export interface Customer {
	id: string;
	name: string;
	email: string;
	phone?: string;
	adress: Address;
}
