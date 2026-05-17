import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	// No auth needed - set mock user for compatibility
	event.locals.user = {
		id: "system",
		name: "System Admin",
		email: "admin@Horizon.local",
		username: "admin",
		role: "admin"
	};
	event.locals.session = null;

	return resolve(event);
};