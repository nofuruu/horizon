import type { Handle } from "@sveltejs/kit";

export const handle: Handle = async ({ event, resolve }) => {
	// No auth needed - set mock user for compatibility
	event.locals.user = {
		id: "system",
		name: "root user",
		email: "root@localhost",
		username: "admin",
		role: "admin"
	};
	event.locals.session = null;

	return resolve(event);
};