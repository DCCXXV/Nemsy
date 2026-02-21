<script lang="ts">
	import { Combobox } from 'bits-ui';
	import { FileUpload } from 'melt/components';
	import CaretUpDownIcon from 'phosphor-svelte/lib/CaretUpDownIcon';
	import CheckIcon from 'phosphor-svelte/lib/CheckIcon';
	import BookIcon from 'phosphor-svelte/lib/BookIcon';

	import CaretDoubleUpIcon from 'phosphor-svelte/lib/CaretDoubleUpIcon';
	import CaretDoubleDownIcon from 'phosphor-svelte/lib/CaretDoubleDownIcon';
	import CloudArrowUpIcon from 'phosphor-svelte/lib/CloudArrowUpIcon';
	import XIcon from 'phosphor-svelte/lib/XIcon';
	import PaperPlaneTiltIcon from 'phosphor-svelte/lib/PaperPlaneTiltIcon';

	import { PUBLIC_API_BASE_URL } from '$env/static/public';
	import { goto } from '$app/navigation';

	let { data } = $props();

	let title = $state('');
	let selectedSubject = $state<string | undefined>(undefined);
	let description = $state('');
	let selectedFiles = $state<Set<File>>(new Set());
	let searchValue = $state('');

	const filesArray = $derived(Array.from(selectedFiles));

	function formatFileSize(bytes: number): string {
		return (bytes / 1024 / 1024).toFixed(2);
	}

	function getFileExt(name: string): string {
		const dot = name.lastIndexOf('.');
		return dot >= 0 ? name.slice(dot + 1).toUpperCase() : '';
	}

	function removeFile(file: File) {
		selectedFiles = new Set([...selectedFiles].filter((f) => f !== file));
	}

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
		if (!title.trim() || !selectedSubject || selectedFiles.size === 0) {
			error = 'Por favor, completa todos los campos obligatorios.';
			return;
		}

		isSubmitting = true;
		error = '';

		const formData = new FormData();
		formData.append('title', title);
		formData.append('subjectId', selectedSubject);
		for (const file of selectedFiles) {
			formData.append('files', file);
		}
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
			error = 'Error de conexion: (' + err + ')';
		} finally {
			isSubmitting = false;
		}
	}
</script>

<div
	class="bg-zinc-100 flex items-start justify-center pt-4 pb-6 min-h-screen relative overflow-hidden"
>
	<img src="/img/tree.svg" alt="" class="absolute -bottom-46 right-12 w-250 pointer-events-none" />

	<div class="relative z-10 bg-zinc-50 border border-zinc-300 rounded-none w-3/7 ml-4 p-4">
		{#if error}
			<div class="bg-red-100 border border-red-400 text-red-700 px-4 py-2 rounded-none mb-4">
				{error}
			</div>
		{/if}

		<div class="flex flex-col mb-4">
			<label for="title">Titulo*</label>
			<input
				name="title"
				placeholder="Titulo del recurso"
				bind:value={title}
				class="bg-zinc-100 border border-zinc-300 p-2 text-zinc-700 rounded-none focus:ring-0 focus:border-blue-600"
			/>
		</div>

		<div class="flex flex-col mb-4">
			<p>Asignatura*</p>
			<Combobox.Root
				type="single"
				name="subject"
				bind:value={selectedSubject}
				onOpenChangeComplete={(o) => {
					if (!o) searchValue = '';
				}}
			>
				<div class="relative">
					<BookIcon class="absolute left-3 top-1/2 size-5 -translate-y-1/2 text-zinc-900" />
					<Combobox.Input
						oninput={(e) => (searchValue = e.currentTarget.value)}
						class="w-full h-10 pl-10 pr-10 bg-zinc-100 border border-zinc-300 rounded-none text-zinc-700 text placeholder:text-zinc-400 focus:ring-0 focus:border-blue-600"
						placeholder="Buscar asignatura"
						aria-label="Buscar asignatura"
					/>
					<Combobox.Trigger class="absolute right-3 top-1/2 -translate-y-1/2 cursor-pointer">
						<CaretUpDownIcon class="size-5 text-zinc-500" />
					</Combobox.Trigger>
				</div>
				<Combobox.Portal>
					<Combobox.Content
						class="z-50 bg-zinc-100 border border-zinc-300 rounded-none max-h-64 overflow-hidden"
						sideOffset={8}
					>
						<Combobox.ScrollUpButton
							class="flex w-full items-center justify-center py-1 text-zinc-500 hover:text-zinc-700"
						>
							<CaretDoubleUpIcon class="size-4" />
						</Combobox.ScrollUpButton>
						<Combobox.Viewport class="p-1">
							{#each filteredSubjects as subject (subject.value)}
								<Combobox.Item
									class="flex items-center px-3 py-2 text-sm rounded-none cursor-pointer select-none data-highlighted:bg-zinc-200 data-[state=checked]:bg-zinc-200"
									value={subject.value}
									label={subject.label}
								>
									{#snippet children({ selected })}
										{subject.label}
										{#if selected}
											<CheckIcon class="ml-auto size-4 text-zinc-700" />
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
							<CaretDoubleDownIcon class="size-4" />
						</Combobox.ScrollDownButton>
					</Combobox.Content>
				</Combobox.Portal>
			</Combobox.Root>
		</div>

		<div class="flex flex-col mb-4">
			<span>Archivos*</span>
			{#if filesArray.length > 0}
				<div class="flex flex-col gap-2 mt-2">
					{#each filesArray as file (file.name + file.size)}
						<div
							class="flex items-center gap-3 p-3 bg-zinc-100 border border-zinc-300 rounded-none"
						>
							<span class="flex-1 truncate text-sm font-medium">{file.name}</span>
							<span
								class="px-2 py-0.5 bg-lime-200 text-lime-700 text-xs font-semibold rounded-none uppercase"
							>
								{getFileExt(file.name)}
							</span>
							<span
								class="px-2 py-0.5 bg-amber-200 text-amber-700 text-xs font-semibold rounded-none uppercase"
							>
								{formatFileSize(file.size)} MB
							</span>
							<button
								type="button"
								onclick={() => removeFile(file)}
								class="text-zinc-500 hover:text-rose-500 cursor-pointer"
								aria-label="Eliminar archivo"
							>
								<XIcon />
							</button>
						</div>
					{/each}
				</div>
			{/if}
			<FileUpload multiple={true} bind:selected={selectedFiles}>
				{#snippet children(fileUpload)}
					<input {...fileUpload.input} />
					<div
						{...fileUpload.dropzone}
						class="p-8 text-center border border-dashed border-zinc-300 rounded-none cursor-pointer hover:bg-zinc-50 text-zinc-500 transition-colors {filesArray.length >
						0
							? 'mt-2'
							: ''}"
					>
						{#if fileUpload.isDragging}
							<CloudArrowUpIcon class="mx-auto size-8" />
							Suelta los archivos aqui
						{:else}
							<CloudArrowUpIcon class="mx-auto size-8" />
							Clica o arrastra para subir archivos
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
				class="bg-zinc-100 border border-zinc-300 p-2 text-zinc-700 rounded-none focus:ring-0 focus:border-blue-600"
			></textarea>
		</div>
		<div class="flex justify-end">
			<button
				onclick={handleSubmit}
				disabled={isSubmitting}
				class="flex items-center bg-lime-200 text-lime-700 px-6 py-2 mt-4 rounded-none cursor-pointer hover:bg-lime-100 disabled:opacity-50 disabled:cursor-not-allowed"
			>
				<PaperPlaneTiltIcon class="size-5 mr-2" />{isSubmitting ? 'Subiendo...' : 'Subir Recurso'}
			</button>
		</div>
	</div>
</div>
