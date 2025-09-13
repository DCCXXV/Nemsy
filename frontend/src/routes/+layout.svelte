<script lang="ts">
	import '../app.css';
	import { env } from '$env/dynamic/public';
	import type { LayoutData } from './$types';

	let props = $props<{ data: LayoutData; children: () => unknown }>();
	const loginUrl = `${env.PUBLIC_API_BASE_URL}/auth/login`;
</script>

<svelte:head>
	<link rel="icon" href="/favicon.png" />
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
			{#if props.data.me}
				<div class="badge p-4">{props.data.me?.email?.split('@')[0]}</div>
			{:else}
				<button onclick={() => (window.location.href = loginUrl)} class="btn btn-accent">
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

<footer class="footer sm:footer-horizontal bg-base-300 text-base-content p-10">
	<aside>
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
