<script lang="ts">
	import '../app.css';
	import favicon from '$lib/assets/favicon.svg';
	import { onMount } from 'svelte';

	let me = $state(null);

	onMount(async () => {
		const res = await fetch(`${import.meta.env.VITE_API_BASE_URL}/api/me`, {
			credentials: 'include'
		});
		if (res.ok) {
			me = await res.json();
		}
	});

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>
<div class="min-h-screen flex flex-col">
	<div class="navbar sticky top-0 bg-base-100/70 backdrop-blur-md shadow-sm z-50">
		<div class="navbar-start">
			<div class="dropdown">
				<div tabindex="0" role="button" class="btn btn-ghost lg:hidden">
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
				</div>
				<ul
					tabindex="-1"
					class="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow"
				>
					<li><a href="/">Inicio</a></li>
					<li>
						<a>Explorar</a>
						<ul class="p-2">
							<li><a>Apuntes</a></li>
							<li><a>Universidades</a></li>
						</ul>
					</li>
					<li><a>Guardado</a></li>
				</ul>
			</div>
			<a href="/" class="btn btn-ghost text-xl">Nemsy</a>
		</div>
		<div class="navbar-center hidden lg:flex">
			<ul class="menu menu-horizontal px-1">
				<li><a href="/">Inicio</a></li>
				<li>
					<details>
						<summary>Explorar</summary>
						<ul class="p-2">
							<li><a>Apuntes</a></li>
							<li><a>Universidades</a></li>
						</ul>
					</details>
				</li>
				<li><a>Guardado</a></li>
			</ul>
		</div>
		<div class="navbar-end gap-4">
			<label class="swap swap-rotate">
				<!-- this hidden checkbox controls the state -->
				<input type="checkbox" class="theme-controller sr-only" value="business" />

				<!-- sun icon -->
				<svg
					class="swap-off h-5 w-5 fill-current"
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
				>
					<path
						d="M5.64,17l-.71.71a1,1,0,0,0,0,1.41,1,1,0,0,0,1.41,0l.71-.71A1,1,0,0,0,5.64,17ZM5,12a1,1,0,0,0-1-1H3a1,1,0,0,0,0,2H4A1,1,0,0,0,5,12Zm7-7a1,1,0,0,0,1-1V3a1,1,0,0,0-2,0V4A1,1,0,0,0,12,5ZM5.64,7.05a1,1,0,0,0,.7.29,1,1,0,0,0,.71-.29,1,1,0,0,0,0-1.41l-.71-.71A1,1,0,0,0,4.93,6.34Zm12,.29a1,1,0,0,0,.7-.29l.71-.71a1,1,0,1,0-1.41-1.41L17,5.64a1,1,0,0,0,0,1.41A1,1,0,0,0,17.66,7.34ZM21,11H20a1,1,0,0,0,0,2h1a1,1,0,0,0,0-2Zm-9,8a1,1,0,0,0-1,1v1a1,1,0,0,0,2,0V20A1,1,0,0,0,12,19ZM18.36,17A1,1,0,0,0,17,18.36l.71.71a1,1,0,0,0,1.41,0,1,1,0,0,0,0-1.41ZM12,6.5A5.5,5.5,0,1,0,17.5,12,5.51,5.51,0,0,0,12,6.5Zm0,9A3.5,3.5,0,1,1,15.5,12,3.5,3.5,0,0,1,12,15.5Z"
					/>
				</svg>

				<!-- moon icon -->
				<svg
					class="swap-on h-5 w-5 fill-current"
					xmlns="http://www.w3.org/2000/svg"
					viewBox="0 0 24 24"
				>
					<path
						d="M21.64,13a1,1,0,0,0-1.05-.14,8.05,8.05,0,0,1-3.37.73A8.15,8.15,0,0,1,9.08,5.49a8.59,8.59,0,0,1,.25-2A1,1,0,0,0,8,2.36,10.14,10.14,0,1,0,22,14.05,1,1,0,0,0,21.64,13Zm-9.5,6.69A8.14,8.14,0,0,1,7.08,5.22v.27A10.15,10.15,0,0,0,17.22,15.63a9.79,9.79,0,0,0,2.1-.22A8.11,8.11,0,0,1,12.14,19.73Z"
					/>
				</svg>
			</label>
			{#if me}
				<div>{me?.email?.split('@')[0] ?? ''}</div>
			{:else}
				<a href="/auth" class="btn btn-soft btn-primary"
					><svg
						class="mr-2 -ml-1 w-4 h-4"
						aria-hidden="true"
						focusable="false"
						data-prefix="fab"
						data-icon="google"
						role="img"
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 488 512"
						><path
							fill="currentColor"
							d="M488 261.8C488 403.3 391.1 504 248 504 110.8 504 0 393.2 0 256S110.8 8 248 8c66.8 0 123 24.5 166.3 64.9l-67.5 64.9C258.5 52.6 94.3 116.6 94.3 256c0 86.5 69.1 156.6 153.7 156.6 98.2 0 135-70.4 140.8-106.9H248v-85.3h236.1c2.3 12.7 3.9 24.9 3.9 41.4z"
						></path></svg
					>Iniciar sesión</a
				>
			{/if}
		</div>
	</div>

	{@render children?.()}
</div>

<footer class="footer sm:footer-horizontal bg-base-300 text-base-content p-10">
	<aside>
		<svg
			width="50"
			height="50"
			viewBox="0 0 24 24"
			xmlns="http://www.w3.org/2000/svg"
			fill-rule="evenodd"
			clip-rule="evenodd"
			class="fill-current"
		>
			<path
				d="M22.672 15.226l-2.432.811.841 2.515c.33 1.019-.209 2.127-1.23 2.456-1.15.325-2.148-.321-2.463-1.226l-.84-2.518-5.013 1.677.84 2.517c.391 1.203-.434 2.542-1.831 2.542-.88 0-1.601-.564-1.86-1.314l-.842-2.516-2.431.809c-1.135.328-2.145-.317-2.463-1.229-.329-1.018.211-2.127 1.231-2.456l2.432-.809-1.621-4.823-2.432.808c-1.355.384-2.558-.59-2.558-1.839 0-.817.509-1.582 1.327-1.846l2.433-.809-.842-2.515c-.33-1.02.211-2.129 1.232-2.458 1.02-.329 2.13.209 2.461 1.229l.842 2.515 5.011-1.677-.839-2.517c-.403-1.238.484-2.553 1.843-2.553.819 0 1.585.509 1.85 1.326l.841 2.517 2.431-.81c1.02-.33 2.131.211 2.461 1.229.332 1.018-.21 2.126-1.23 2.456l-2.433.809 1.622 4.823 2.433-.809c1.242-.401 2.557.484 2.557 1.838 0 .819-.51 1.583-1.328 1.847m-8.992-6.428l-5.01 1.675 1.619 4.828 5.011-1.674-1.62-4.829z"
			></path>
		</svg>
		<p>
			Nesmy.org
			<br />
			Recursos académicos compartidos por y para estudiantes.
		</p>
	</aside>
	<nav></nav>
	<nav>
		<h6 class="footer-title">Información</h6>
		<a class="link link-hover">Sobre nosotros</a>
		<a class="link link-hover">Contacto</a>
	</nav>
	<nav>
		<h6 class="footer-title">Legal</h6>
		<a href="/tos" class="link link-hover">Términos de uso</a>
		<a href="/privacy" class="link link-hover">Política de privacidad</a>
		<a href="/cookies" class="link link-hover">Política de Cookies</a>
	</nav>
</footer>
