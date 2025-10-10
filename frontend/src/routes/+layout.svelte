<script lang="ts">
	import '../app.css';
	import type { LayoutData } from './$types';
	import { clickOutside } from '$lib/actions/clickOutside';
	import { page } from '$app/state';

	let props = $props<{ data: LayoutData; children: () => unknown }>();

	let isMenuOpen = $state(false);
	let detailsElement: HTMLDetailsElement | null = $state(null);

	function closeMenu() {
		isMenuOpen = false;
	}
	function closeDetails() {
		if (detailsElement) {
			detailsElement.open = false;
		}
	}

	let currentPath = $derived(page.url.pathname);
</script>

<svelte:head>
	<link rel="icon" href="/favicon.png" />
</svelte:head>

{#if currentPath.includes('/auth')}
	{@render props.children?.()}
{:else}
	<div class="min-h-screen flex flex-col bg-zinc-100">
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
				<ul class="flex list-none bg-zinc-100 border-2 border-zinc-900 rounded">
					<li>
						<a
							class="h-10 flex items-center px-6 py-2 hover:bg-zinc-200 rounded transition-colors
							{currentPath === '/' ? 'font-bold bg-zinc-200' : ''}"
							href="/"
							>Inicio
						</a>
					</li>
					<li>
						<details
							class="relative"
							bind:this={detailsElement}
							use:clickOutside
							onoutclick={closeDetails}
						>
							<summary
								class="h-10 flex items-center px-6 py-2 hover:bg-zinc-200 rounded cursor-pointer transition-colors"
								>Explorar
							</summary>
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
							class="h-10 flex items-center px-6 py-2 hover:bg-zinc-200 rounded transition-colors"
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
						{props.data.me?.Email?.split('@')[0]}
					</div>
				{:else}
					<button
						onclick={() => (window.location.href = '/auth')}
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

	<footer class="bg-slate-300 text-slate-800 p-10 flex flex-col md:flex-row justify-between gap-8">
		<aside class="flex flex-col md:flex-row items-start md:items-center gap-4 shrink-0">
			<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 279 279" width="64" height="64"
				><defs></defs><g transform="matrix(1,0,0,1,-110.080572441,-110.080572441)">
					<g>
						<g>
							<path
								d="M110.08057244055081,110.08057244055163 L389.91942755944854,110.08057244055163 L389.91942755944876,389.91942755944933 L110.08057244055098,389.91942755944933 L110.08057244055081,110.08057244055163 Z"
								transform="matrix(1,0,0,1,0,0)"
								fill="#e2e8f0"
								stroke="#62748e"
								stroke-width="50"
								stroke-linecap="round"
								stroke-linejoin="round"
							/>
						</g>
						<g>
							<path
								d="M181.34488238152414,183.1439145308349 C181.34488238152414,183.1439145308349 213.91808699217927,183.1439145308349 213.91808699217927,183.1439145308349 C213.91808699217927,183.1439145308349 298.7704510145223,254.40226741171577 298.7704510145223,254.40226741171577 C298.7704510145223,254.40226741171577 298.8657639760751,183.1439145308349 298.8657639760751,183.1439145308349 C298.8657639760751,183.1439145308349 318.6551176184665,183.1439145308349 318.6551176184665,183.1439145308349 C318.6551176184665,183.1439145308349 318.6551176184665,316.8560854691606 318.6551176184665,316.8560854691606 C318.6551176184665,316.8560854691606 286.08191300781135,316.8560854691606 286.08191300781135,316.8560854691606 C286.08191300781135,316.8560854691606 201.14615014410964,245.85984323254982 201.14615014410964,245.85984323254982 C201.14615014410964,245.85984323254982 201.14615014410964,316.8560854691606 201.14615014410964,316.8560854691606 C201.14615014410964,316.8560854691606 181.34488238152414,316.8560854691606 181.34488238152414,316.8560854691606 C181.34488238152414,316.8560854691606 181.34488238152414,183.1439145308349 181.34488238152414,183.1439145308349 Z"
								fill="#62748e"
							/>
						</g>
					</g></g
				></svg
			>
			<p class="text-sm">
				Nemsy.org
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
{/if}
