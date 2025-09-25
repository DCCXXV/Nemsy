<script lang="ts">
	import type { LayoutData } from './$types';

	let props = $props<{ data: LayoutData; children: () => unknown }>();
</script>

{#if props.data.me}
	<div class="absolute inset-x-0 top-0 h-full bg-slate-200 z-0"></div>
{:else}
	<div class="relative min-h-screen bg-zinc-100 flex justify-center items-center">
		<div class="absolute inset-x-0 top-0 h-[calc(3/7*100vh)] bg-slate-200 z-0"></div>

		<div
			class="relative z-10 flex flex-col lg:flex-row w-full max-w-6xl p-4 lg:p-0 mt-24 lg:mt-0 mb-24"
		>
			<div
				class="flex-1 h-full flex justify-center items-center lg:items-start py-12 border-2 border-zinc-900 rounded lg:px-10 bg-zinc-100 shadow-[-8px_8px_0px_#314158] transition-all hover:shadow-none hover:translate-x-[-8px] hover:translate-y-[8px]"
			>
				<div
					class="max-w-md flex flex-col items-center lg:items-start text-left mx-auto lg:mx-0 w-full p-4"
				>
					<h1 class="text-5xl font-bold text-slate-700">Nemsy</h1>
					<p class="py-6">
						Comparte y accede a apuntes universitarios con facilidad. Todo lo que necesitas para
						estudiar mejor, en un solo lugar.
					</p>

					<div class="grid gap-3 mb-6 w-full">
						<div class="flex items-center gap-2 justify-start">
							<span class="text-success text-xl">✔</span>
							<span><b>Open Source:</b> Transparente y colaborativo.</span>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span class="text-success text-xl">✔</span>
							<span><b>Sin anuncios:</b> Tu atención en lo que importa.</span>
						</div>
						<div class="flex items-center gap-2 justify-start">
							<span class="text-success text-xl">✔</span>
							<span><b>Rápido y ligero:</b> Acceso instantáneo a tus apuntes.</span>
						</div>
					</div>
					<!--
					<button
						onclick={() => (window.location.href = loginUrl)}
						class="flex bg-amber-200 px-4 py-2 rounded h-10"
					>
						<svg
							class="mr-2 ml-1 w-4 h-4"
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
							></path>
						</svg>
						Iniciar sesión con Google
					</button>
					 -->
				</div>
			</div>

			<div class="flex-1 h-full flex justify-center items-center p-4 rounded-xl">
				<div class="w-full max-w-md text-base-100">
					<div style="height: 400px;"></div>
				</div>
			</div>
		</div>
	</div>
{/if}
