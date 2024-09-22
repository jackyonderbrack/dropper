export interface User {
	id: number;
	name: string;
	email: string;
	password: string;
	isAdmin: boolean;
	company: Company;
	nip: string;
	address: Address;
	createdAt: Date;
	updatedAt: Date;
}

export interface Company {
	name: string;
	nip: string;
}

export interface Address {
	street: string;
	city: string;
	zipCode: string;
	country: string;
}

export type NewUser = Pick<User, "name" | "email" | "password">;

export type LoginUser = Pick<User, "email" | "password">;
