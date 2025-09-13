import { PUBLIC_API_BASE_URL } from '$env/static/public';
import type { Handle } from '@sveltejs/kit';

export const handle: Handle = async ({ event, resolve }) => {
	const session = event.cookies.get('session_token');

	if (session) {
		try {
			const res = await event.fetch(`${PUBLIC_API_BASE_URL}/api/me`, {
				headers: {
					cookie: `session_token=${session}`
				}
			});

			if (res.ok) {
				event.locals.user = await res.json();
			} else {
				event.locals.user = null;
			}
		} catch (err) {
			console.error('Error fetching /api/me', err);
			event.locals.user = null;
		}
	} else {
		event.locals.user = null;
	}

	return resolve(event);
};
