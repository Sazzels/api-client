<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import { getEnvironmentStore } from '$lib/environmentStore.svelte.ts';

	let { environment }: { environment: frontend.EnvironmentDTO } = $props();

	const environmentStore = getEnvironmentStore();
	let newHeaderKey = $state('');
	let newHeaderValue = $state('');

	let update = (header: frontend.EnvironmentHeaderDTO, index: number): void => {
		environmentStore.updateHeader(header, environment, index);
	};

	let deleteHeader = (header: frontend.EnvironmentHeaderDTO, index: number): void => {
		environmentStore.deleteHeader(header, environment, index);
	};

	let appendHeader = (): void => {
		let header = new frontend.EnvironmentHeaderDTO();
		header.key = newHeaderKey;
		header.value = newHeaderValue;
		environmentStore.addHeader(header, environment);
		newHeaderKey = '';
		newHeaderValue = '';
	};
</script>

<div class="flex h-full flex-col">
	<div class="mt-2 flex flex-row gap-1 pr-4">
		<div class="relative mx-auto w-full">
			<input
				bind:value={newHeaderKey}
				id="new-header-key"
				type="text"
				placeholder="New Header Key"
				class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
				text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
			/>
			<label
				for="new-header-key"
				class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
				  peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
				  peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
				  before:block before:h-full before:w-full before:transition-all"
				>New Header Key
			</label>
		</div>
		<div class="relative mx-auto w-full">
			<input
				bind:value={newHeaderValue}
				id="new-header-value"
				type="text"
				placeholder="New Header Value"
				class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
				text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden focus-visible:outline-hidden"
			/>
			<label
				for="new-header-value"
				class="before:bg-background absolute -top-2 left-2 z-1 cursor-text px-2 text-xs transition-all
				  peer-placeholder-shown:top-2.5 peer-placeholder-shown:text-sm peer-autofill:-top-2 peer-focus:-top-2
				  peer-focus:cursor-default peer-focus:text-xs before:absolute before:top-0 before:left-0 before:z-[-1]
				  before:block before:h-full before:w-full before:transition-all"
				>New Header Value
			</label>
		</div>
		<button
			aria-label="save"
			onclick={() => {
				appendHeader();
			}}
		>
			<svg
				xmlns="http://www.w3.org/2000/svg"
				width="24"
				height="24"
				viewBox="0 0 24 24"
				fill="none"
				stroke="currentColor"
				stroke-width="2"
				stroke-linecap="round"
				stroke-linejoin="round"
			>
				<path stroke="none" d="M0 0h24v24H0z" fill="none" />
				<path d="M6 4h10l4 4v10a2 2 0 0 1 -2 2h-12a2 2 0 0 1 -2 -2v-12a2 2 0 0 1 2 -2" />
				<path d="M12 14m-2 0a2 2 0 1 0 4 0a2 2 0 1 0 -4 0" />
				<path d="M14 4l0 4l-6 0l0 -4" />
			</svg>
		</button>
	</div>
	<div data-testid="environment-headers" class="mt-4 flex max-h-full flex-col overflow-y-auto pr-4">
		{#each environment.header as header, iter (iter)}
			<div data-testid="environment-header" class="flex flex-row gap-1">
				<input
					bind:value={header.key}
					oninput={() => {
						update(header, iter);
					}}
					type="text"
					class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
					text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden
					 focus-visible:outline-hidden"
				/>
				<input
					bind:value={header.value}
					oninput={() => {
						update(header, iter);
					}}
					type="text"
					class="peer border-background-accent focus:border-text-accent relative h-10 w-full rounded-sm border px-4
					text-sm placeholder-transparent outline-hidden transition-all focus:outline-hidden
					focus-visible:outline-hidden"
				/>
				<button aria-label="delete" onclick={() => deleteHeader(header, iter)} class="h-10">
					<svg
						class="h-5"
						xmlns="http://www.w3.org/2000/svg"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
					>
						<path stroke="none" d="M0 0h24v24H0z" fill="none" />
						<path d="M4 7l16 0" />
						<path d="M10 11l0 6" />
						<path d="M14 11l0 6" />
						<path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
						<path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
					</svg>
				</button>
			</div>
		{/each}
	</div>
</div>
