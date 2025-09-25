import type { ActionReturn } from 'svelte/action';

type Attributes = {
	onoutclick: (e: CustomEvent) => void;
};

export function clickOutside(node: HTMLElement): ActionReturn<undefined, Attributes> {
	const handleClick = (event: MouseEvent) => {
		if (event.target instanceof Node && !node.contains(event.target)) {
			node.dispatchEvent(new CustomEvent('outclick'));
		}
	};

	document.addEventListener('click', handleClick, true);

	return {
		destroy() {
			document.removeEventListener('click', handleClick, true);
		}
	};
}
