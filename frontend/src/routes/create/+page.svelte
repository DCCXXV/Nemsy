<script lang="ts">
	import { Combobox } from 'bits-ui';
	import { FileUpload } from 'melt/components';
	import CaretUpDown from 'phosphor-svelte/lib/CaretUpDown';
	import Check from 'phosphor-svelte/lib/Check';
	import Book from 'phosphor-svelte/lib/Book';
	import CaretDoubleUp from 'phosphor-svelte/lib/CaretDoubleUp';
	import CaretDoubleDown from 'phosphor-svelte/lib/CaretDoubleDown';
	import CloudArrowUp from 'phosphor-svelte/lib/CloudArrowUp';

	const subjects = [
		{ value: 'modelado-software', label: 'Modelado de Software' },
		{ value: 'fundamentos-algoritmia', label: 'Fundamentos de la Algoritmia' },
		{ value: 'aplicaciones-web', label: 'Aplicaciones Web' },
		{ value: 'programacion-evolutiva', label: 'Programación evolutiva' }
	];

	let searchValue = $state('');

	const filteredSubjects = $derived(
		searchValue === ''
			? subjects
			: subjects.filter((subject) =>
					subject.label.toLowerCase().includes(searchValue.toLowerCase())
				)
	);
</script>

<div class="bg-pastel-200 flex items-start justify-center pt-4 pb-6 h-screen">
	<div class="bg-zinc-100 border-2 border-zinc-900 rounded w-3/7 ml-4 p-4">
		<div class="flex flex-col mb-4">
			<label for="title">Título*</label>
			<input
				name="title"
				placeholder="Título del recurso"
				class="bg-zinc-100 border-2 border-zinc-900 p-2 rounded focus:outline-none focus:ring-2 focus:ring-zinc-400"
			/>
		</div>
		<div class="flex flex-col mb-4">
			<label>Asignatura*</label>
			<Combobox.Root
				type="single"
				name="subject"
				onOpenChangeComplete={(o) => {
					if (!o) searchValue = '';
				}}
			>
				<div class="relative">
					<Book class="absolute left-3 top-1/2 size-5 -translate-y-1/2 text-zinc-900" />
					<Combobox.Input
						oninput={(e) => (searchValue = e.currentTarget.value)}
						class="w-full h-10 pl-10 pr-10 bg-zinc-100 border-2 border-zinc-900 rounded text-sm placeholder:text-zinc-400 focus:outline-none focus:ring-2 focus:ring-zinc-400"
						placeholder="Buscar asignatura"
						aria-label="Buscar asignatura"
					/>
					<Combobox.Trigger class="absolute right-3 top-1/2 -translate-y-1/2 cursor-pointer">
						<CaretUpDown class="size-5 text-zinc-500" />
					</Combobox.Trigger>
				</div>
				<Combobox.Portal>
					<Combobox.Content
						class="z-50 bg-zinc-100 border-2 border-zinc-900 rounded max-h-64 overflow-hidden"
						sideOffset={8}
					>
						<Combobox.ScrollUpButton
							class="flex w-full items-center justify-center py-1 text-zinc-500 hover:text-zinc-700"
						>
							<CaretDoubleUp class="size-4" />
						</Combobox.ScrollUpButton>
						<Combobox.Viewport class="p-1">
							{#each filteredSubjects as subject, i (i + subject.value)}
								<Combobox.Item
									class="flex items-center px-3 py-2 text-sm rounded cursor-pointer select-none data-[highlighted]:bg-zinc-200 data-[state=checked]:bg-zinc-300"
									value={subject.value}
									label={subject.label}
								>
									{#snippet children({ selected })}
										{subject.label}
										{#if selected}
											<Check class="ml-auto size-4 text-zinc-700" />
										{/if}
									{/snippet}
								</Combobox.Item>
							{:else}
								<span class="block px-3 py-2 text-sm text-zinc-500">
									No se encontraron resultados.
								</span>
							{/each}
						</Combobox.Viewport>
						<Combobox.ScrollDownButton
							class="flex w-full items-center justify-center py-1 text-zinc-500 hover:text-zinc-700"
						>
							<CaretDoubleDown class="size-4" />
						</Combobox.ScrollDownButton>
					</Combobox.Content>
				</Combobox.Portal>
			</Combobox.Root>
		</div>

		<div class="flex flex-col mb-4">
			<span>Archivo*</span>
			<FileUpload>
				{#snippet children(fileUpload)}
					<input {...fileUpload.input} />
					<div
						{...fileUpload.dropzone}
						class="p-8 text-center border-2 border-dashed border-zinc-900 cursor-pointer"
					>
						{#if fileUpload.isDragging}
							Arrastra los archivos aquí
						{:else}
							<CloudArrowUp class="mx-auto size-8" />
							Clica o arrastra para subir tus archivos
						{/if}
					</div>
				{/snippet}
			</FileUpload>
		</div>

		<div class="flex flex-col">
			<label for="title">Descripción (Opcional)</label>
			<textarea
				name="title"
				placeholder="Describe el recurso que vas a subir para ayudar a que otros estudiantes entiendan de que se trata."
				class="bg-zinc-100 border-2 border-zinc-900 p-2 rounded focus:outline-none focus:ring-2 focus:ring-zinc-400"
			/>
		</div>
		<button class="bg-zinc-900 text-zinc-100 px-6 py-2 mt-4 rounded cursor-pointer">
			Subir Recurso
		</button>
	</div>
</div>
