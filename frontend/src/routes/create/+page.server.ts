import { PUBLIC_API_BASE_URL } from '$env/static/public';
import { redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';
import type { Subject } from '$lib/types';

export const load: PageServerLoad = async ({ fetch, parent }) => {
	const { me } = await parent();

	if (!me) {
		throw redirect(302, '/auth');
	}

	if (!me.studyId) {
		throw redirect(302, '/');
	}

	let subjects: Subject[] = [];

	try {
		const res = await fetch(`${PUBLIC_API_BASE_URL}/api/me/subjects`, {
			credentials: 'include'
		});

		if (res.ok) {
			subjects = await res.json();
		}
	} catch (err) {
		console.error('Error fetching subjects:', err);
	}

	return { subjects };
};
