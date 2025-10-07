import { PUBLIC_API_BASE_URL } from '$env/static/public';
import type { LayoutServerLoad } from './$types';
import type { User } from '$lib/types';

export const load: LayoutServerLoad = async ({ fetch }) => {
	let me: User | null = null;

	try {
		const res = await fetch(`${PUBLIC_API_BASE_URL}/api/me`, {
			credentials: 'include'
		});

		if (res.ok) {
			me = await res.json();
		}
	} catch (err) {
		console.error('Error fetching /me:', err);
	}

	return { me };
};
