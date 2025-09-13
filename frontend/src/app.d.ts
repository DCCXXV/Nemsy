import type { User } from '$lib/types';
declare global {
	namespace App {
		interface Locals {
			user: User | null;
		}

		interface LayoutData {
			me: User | null;
		}
	}
}

export {};
