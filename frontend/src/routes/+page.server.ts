import { PUBLIC_API_BASE_URL } from '$env/static/public';
import type { PageServerLoad } from './$types';
import type { Subject } from '$lib/types';

export const load: PageServerLoad = async ({ fetch, parent }) => {
	const { me } = await parent();

	let subjects: Subject[] = [];

	if (!me) {
		return { subjects };
	}

	try {
		const res = await fetch(`${PUBLIC_API_BASE_URL}/api/me/subjects`, {
			credentials: 'include'
		});

		if (res.ok) {
			subjects = await res.json();
			return { subjects };
		}
	} catch (err) {
		console.error('Error fetching subjects:', err);
	}

	return { subjects };
};
