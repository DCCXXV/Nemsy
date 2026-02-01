<script lang="ts">
	import { Combobox } from 'bits-ui';
	import { FileUpload } from 'melt/components';
	import CaretUpDown from 'phosphor-svelte/lib/CaretUpDown';
	import Check from 'phosphor-svelte/lib/Check';
	import Book from 'phosphor-svelte/lib/Book';
	import CaretDoubleUp from 'phosphor-svelte/lib/CaretDoubleUp';
	import CaretDoubleDown from 'phosphor-svelte/lib/CaretDoubleDown';
	import CloudArrowUp from 'phosphor-svelte/lib/CloudArrowUp';
	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let title = $state('');
	let selectedSubject = $state<string | undefined>(undefined);
	let description = $state('');
	let files = $state<File[]>([]);
	let searchValue = $state('');
	let isSubmitting = $state(false);
	let error = $state('');

	const subjects = $derived(
		data.subjects.map((s) => ({
			value: String(s.id),
			label: s.name
		}))
	);

	const filteredSubjects = $derived(
		searchValue === ''
			? subjects
			: subjects.filter((subject) =>
					subject.label.toLowerCase().includes(searchValue.toLowerCase())
				)
	);

	async function handleSubmit() {
		if (!title.trim() || !selectedSubject || files.length === 0) {
			error = 'Por favor, completa todos los campos obligatorios.';
			return;
		}

		isSubmitting = true;
		error = '';

		const formData = new FormData();
		formData.append('title', title);
		formData.append('subjectId', selectedSubject);
		formData.append('file', files[0]);
		if (description.trim()) {
			formData.append('description', description);
		}

		try {
			const res = await fetch(`${PUBLIC_API_BASE_URL}/api/resources`, {
				method: 'POST',
				credentials: 'include',
				body: formData
			});

			if (res.ok) {
				goto('/');
			} else {
				const text = await res.text();
				error = text || 'Error al subir el recurso.';
			}
		} catch (err) {
			error = 'Error de conexion.';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div class="bg-oatmeal-200 flex items-start justify-center pt-4 pb-6 min-h-screen">
	<div class="bg-zinc-100 border-2 border-zinc-900 rounded w-3/7 ml-4 p-4">
		{#if error}
			<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-2 rounded mb-4">
				{error}
			</div>
		{/if}

		<div class="flex flex-col mb-4">
			<label for="title">Titulo*</label>
			<input
				name="title"
				placeholder="Titulo del recurso"
				bind:value={title}
				class="bg-zinc-100 border-2 border-zinc-900 p-2 rounded focus:outline-none focus:ring-2 focus:ring-zinc-400"
			/>
		</div>

		<div class="flex flex-col mb-4">
			<label>Asignatura*</label>
			<Combobox.Root
				type="single"
				name="subject"
				bind:value={selectedSubject}
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
							{#each filteredSubjects as subject (subject.value)}
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
			<FileUpload
				onFilesChange={(newFiles) => {
					files = newFiles;
				}}
			>
				{#snippet children(fileUpload)}
					<input {...fileUpload.input} />
					<div
						{...fileUpload.dropzone}
						class="p-8 text-center border-2 border-dashed border-zinc-900 cursor-pointer"
					>
						{#if files.length > 0}
							<p class="font-medium">{files[0].name}</p>
							<p class="text-sm text-zinc-500">
								{(files[0].size / 1024 / 1024).toFixed(2)} MB
							</p>
						{:else if fileUpload.isDragging}
							Arrastra los archivos aqui
						{:else}
							<CloudArrowUp class="mx-auto size-8" />
							Clica o arrastra para subir tus archivos
						{/if}
					</div>
				{/snippet}
			</FileUpload>
		</div>

		<div class="flex flex-col">
			<label for="description">Descripcion (Opcional)</label>
			<textarea
				name="description"
				bind:value={description}
				placeholder="Describe el recurso que vas a subir para ayudar a que otros estudiantes entiendan de que se trata."
				class="bg-zinc-100 border-2 border-zinc-900 p-2 rounded focus:outline-none focus:ring-2 focus:ring-zinc-400"
			/>
		</div>

		<button
			onclick={handleSubmit}
			disabled={isSubmitting}
			class="bg-zinc-900 text-zinc-100 px-6 py-2 mt-4 rounded cursor-pointer hover:bg-zinc-700 disabled:opacity-50 disabled:cursor-not-allowed"
		>
			{isSubmitting ? 'Subiendo...' : 'Subir Recurso'}
		</button>
	</div>
</div>
