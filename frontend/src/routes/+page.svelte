<script lang="ts">
	import { env } from '$env/dynamic/public';
	import type { LayoutData } from './$types';

	let props = $props<{ data: LayoutData; children: () => unknown }>();

	const loginUrl = `${env.PUBLIC_API_BASE_URL}/auth/login`;
</script>

{#if props.data.me}{:else}
	<div class="hero min-h-screen bg-base-200">
		<div class="hero-content flex-col lg:flex-row w-full p-0 mb-24">
			<div
				class="flex-1 flex flex-col justify-center items-center lg:items-start p-10 text-center lg:text-left max-w-md"
			>
				<h1 class="text-5xl font-bold">Nemsy</h1>
				<p class="py-6 text-lg">
					Comparte y accede a apuntes universitarios con facilidad. Todo lo que necesitas para
					estudiar mejor, en un solo lugar.
				</p>

				<div class="grid gap-3 text-sm mb-6">
					<div class="flex items-center gap-2">
						<span class="text-success text-xl">✔</span>
						<span><b>Open Source:</b> Transparente y colaborativo.</span>
					</div>
					<div class="flex items-center gap-2">
						<span class="text-success text-xl">✔</span>
						<span><b>Sin anuncios:</b> Tu atención en lo que importa.</span>
					</div>
					<div class="flex items-center gap-2">
						<span class="text-success text-xl">✔</span>
						<span><b>Rápido y ligero:</b> Acceso instantáneo a tus apuntes.</span>
					</div>
				</div>

				<button onclick={() => (window.location.href = loginUrl)} class="btn btn-accent">
					<svg
						class="mr-2 ml-1 w-4 h-4"
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
					>
					Iniciar sesión con Google
				</button>
			</div>
			<div class="flex-1 bg-base-300 p-4 rounded-xl">
				<div class="w-full h-full flex items-center justify-center text-base-100">
					<div style="height: 400px;"></div>
				</div>
			</div>
		</div>
	</div>
{/if}
