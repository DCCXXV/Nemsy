<script lang="ts">
	import type { LayoutData } from './$types';

	let { data }: { data: LayoutData } = $props();
</script>

{#if data.me}
	<div class="absolute inset-x-0 top-0 h-[calc(3/7*100vh)] bg-slate-200 z-0"></div>
	<div
		class="relative z-10 min-w-1/2 mx-auto my-6 bg-zinc-100 border-2 border-slate-900 rounded h-screen"
	></div>
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
