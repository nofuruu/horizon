// See https://svelte.dev/docs/kit/types#app.d.ts

type User = {
	id: string;
	name: string;
	email: string;
	username: string;
	role: string;
};

declare global {
	namespace App {
		// interface Error {}
		interface Locals {
			user: User | null;
			session: any | null;
		}
		// interface PageData {}
		// interface PageState {}
		// interface Platform {}
	}
}

export {};