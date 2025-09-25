<script lang="ts">
	import '../app.css';
	import { env } from '$env/dynamic/public';
	import type { LayoutData } from './$types';
	import { clickOutside } from '$lib/actions/clickOutside';

	let props = $props<{ data: LayoutData; children: () => unknown }>();
	const loginUrl = `${env.PUBLIC_API_BASE_URL}/auth/login`;

	let isMenuOpen = $state(false);
	let detailsElement: HTMLDetailsElement | null = null;

	function closeMenu() {
		isMenuOpen = false;
	}
	function closeDetails() {
		if (detailsElement) {
			detailsElement.open = false;
		}
	}
</script>

<svelte:head>
	<link rel="icon" href="/favicon.png" />
</svelte:head>

<div class="min-h-screen flex flex-col">
	<div class="bg-slate-200 bg-opacity-70 z-50 flex items-center justify-between px-4 py-2">
		<div class="flex items-center">
			<div class="relative" use:clickOutside onoutclick={closeMenu}>
				<button
					aria-label="menú-móvil"
					class="lg:hidden h-10 px-2 py-2 mr-2 rounded bg-zinc-100 border-2 border-zinc-900 transition-colors inline-flex items-center cursor-pointer hover:bg-zinc-200"
					onclick={() => (isMenuOpen = !isMenuOpen)}
				>
					<svg
						xmlns="http://www.w3.org/2000/svg"
						class="h-5 w-5"
						fill="none"
						viewBox="0 0 24 24"
						stroke="currentColor"
					>
						<path
							stroke-linecap="round"
							stroke-linejoin="round"
							stroke-width="2"
							d="M4 6h16M4 12h8m-8 6h16"
						/>
					</svg>
				</button>
				{#if isMenuOpen}
					<ul
						class="absolute bg-zinc-100 rounded z-10 mt-3 w-52 p-2 list-none border-zinc-900 border-2 transition-opacity"
					>
						<li><a class="block px-4 py-2 hover:bg-zinc-200 rounded" href="/">Inicio</a></li>
						<li>
							<div class="relative">
								<button
									class="block w-full text-left px-4 py-2 hover:bg-zinc-200 rounded"
									onclick={() => (isMenuOpen = !isMenuOpen)}
								>
									Explorar
								</button>
							</div>
						</li>
						<li><a class="block px-4 py-2 hover:bg-zinc-200 rounded" href="#">Guardado</a></li>
					</ul>
				{/if}
			</div>

			<a
				href="/"
				class="h-10 text-xl text-zinc-900 font-bold px-4 py-2 bg-zinc-100 hover:brightness-90 border-2 border-zinc-900 rounded transition-colors inline-flex items-center cursor-pointer"
				>Nemsy</a
			>
		</div>

		<div class="hidden lg:flex flex-1 justify-center">
			<ul class="flex list-none space-x-4">
				<li>
					<a
						class="h-10 flex items-center px-4 py-2 hover:bg-zinc-100 border-2 hover:border-zinc-900 border-slate-200 rounded transition-colors"
						href="/">Inicio</a
					>
				</li>
				<li>
					<details
						class="relative"
						bind:this={detailsElement}
						use:clickOutside
						onoutclick={closeDetails}
					>
						<summary
							class="h-10 flex items-center px-4 py-2 hover:bg-zinc-100 border-2 focus:bg-zinc-100 focus:border-zinc-900 hover:border-zinc-900 border-slate-200 rounded cursor-pointer transition-colors"
							>Explorar</summary
						>
						<ul
							class="absolute mt-2 p-2 bg-zinc-100 border-zinc-900 border-2 shadow rounded list-none min-w-max"
						>
							<li>
								<a class="block px-4 py-2 hover:bg-zinc-200 rounded" href="#">Apuntes</a>
							</li>
							<li>
								<a class="block px-4 py-2 hover:bg-zinc-200 rounded" href="#">Universidades</a>
							</li>
						</ul>
					</details>
				</li>
				<li>
					<a
						class="h-10 flex items-center px-4 py-2 hover:bg-zinc-100 border-2 hover:border-zinc-900 border-slate-200 rounded transition-colors"
						href="#">Guardado</a
					>
				</li>
			</ul>
		</div>

		<div class="flex-none flex flex-row items-center gap-4">
			{#if props.data.me}
				<div
					class="h-10 flex items-center px-4 py-2 bg-zinc-100 text-zinc-900 border-zinc-900 border-2 rounded"
				>
					{props.data.me?.email?.split('@')[0]}
				</div>
			{:else}
				<button
					onclick={() => (window.location.href = loginUrl)}
					class="h-10 px-4 py-2 bg-rose-300 hover:brightness-90 border-zinc-900 border-2 rounded transition-colors inline-flex items-center cursor-pointer"
				>
					<svg
						class="mr-2 -ml-1 w-4 h-4"
						aria-hidden="true"
						focusable="false"
						data-prefix="fab"
						data-icon="google"
						role="img"
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 488 512"
					>
						<path
							fill="currentColor"
							d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"
						></path></svg
					>
					Iniciar sesión
				</button>
			{/if}
		</div>
	</div>

	{@render props.children?.()}
</div>

<footer class="bg-slate-300 text-slate-800 p-10 flex flex-col sm:flex-row justify-between gap-8">
	<aside class="flex flex-col sm:flex-row items-start sm:items-center gap-4 shrink-0">
		<svg
			width="64"
			height="64"
			viewBox="0 0 120 120"
			version="1.1"
			id="svg1"
			xmlns="http://www.w3.org/2000/svg"
		>
			<defs id="defs1" />
			<g id="layer1">
				<path
					style="fill:#e1ded9;stroke:#000000;stroke-width:5;stroke-dasharray:none;stroke-opacity:1"
					d="M 60,60 V 14.134969 H 45.449816 L 19.582822,52.564417 35.369296,87.7546 60,87.78916 M 60,60 V 14.134969 H 74.550184 L 100.41718,52.564417 84.630704,87.7546 60,87.78916"
					id="path1"
					transform="rotate(-179.81551,59.996145,63.274092)"
				/>
				<circle
					style="fill:#000000;fill-opacity:1"
					id="path2"
					cx="-60.195713"
					cy="-66.354675"
					r="10.233129"
					transform="rotate(-179.81551)"
				/>
				<rect
					style="fill:#000000;fill-opacity:1;stroke:#000000;stroke-width:5;stroke-dasharray:none;stroke-opacity:1"
					id="rect4"
					width="86.533226"
					height="26.9133"
					x="-103.46233"
					y="-36.065521"
					transform="rotate(-179.81551)"
				/>
			</g>
		</svg>
		<p class="text-sm">
			Nesmy.org
			<br />
			Recursos académicos compartidos por y para estudiantes.
		</p>
	</aside>

	<nav class="flex flex-grow flex-col gap-2 flex-1 sm:flex-none">
		<h6 class="text-sm font-bold tracking-wide text-slate-600 uppercase">Información</h6>
		<a
			href="#"
			class="text-slate-800 hover:text-slate-900 hover:underline transition-colors no-underline"
		>
			Sobre nosotros
		</a>
		<a
			href="#"
			class="text-slate-800 hover:text-slate-900 hover:underline transition-colors no-underline"
		>
			Contacto
		</a>
	</nav>
	<nav class="flex flex-col gap-2 flex-1 sm:flex-none">
		<h6 class="text-sm font-bold tracking-wide text-slate-600 uppercase">Legal</h6>
		<a
			href="/tos"
			class="text-slate-800 hover:text-slate-900 hover:underline transition-colors no-underline"
		>
			Términos de uso
		</a>
		<a
			href="/privacy"
			class="text-slate-800 hover:text-slate-900 hover:underline transition-colors no-underline"
		>
			Política de privacidad
		</a>
		<a
			href="/cookies"
			class="text-slate-800 hover:text-slate-900 hover:underline transition-colors no-underline"
		>
			Política de Cookies
		</a>
	</nav>
</footer>
