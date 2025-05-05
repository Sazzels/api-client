<script lang="ts">
	import { frontend } from './wailsjs/go/models';
	import Loader from './Loader.svelte';
	import ClipboardButton from './ClipboardButton.svelte';

	let { response, loading }: { response: frontend.RequestResponseDTO; loading: boolean } = $props();
	let isJsonResponse: boolean = $derived.by(() => {
		try {
			JSON.parse(response.responseBody);
			// eslint-disable-next-line @typescript-eslint/no-unused-vars
		} catch (e) {
			return false;
		}

		return true;
	});
</script>

{#if loading}
	<Loader></Loader>
{:else if response.error !== ''}
	<p data-testid="response-error" class="text-red-700">{response.error}</p>
{:else if response.elapsedTime === undefined}{:else}
	<div class="flex h-full flex-col overflow-x-hidden overflow-y-auto border p-2">
		{#if response.tlsSkipped === true}
			<p class="text-red-600">TLS could not be verified, skipped!</p>
			<hr class="my-2 mr-2" />
		{/if}
		<div class="mr-2 flex justify-between">
			<h4 class="text-text-highlight">Status-Code:</h4>
			<ClipboardButton id="status-code-clipboard-btn" data={response.statusCode.toFixed(0)} />
		</div>
		<p data-testid="status-code">{response.statusCode}</p>
		<hr class="my-2 mr-2" />
		<div class="mr-2 flex justify-between">
			<h4 class="text-text-highlight">Elapsed Time:</h4>
			<ClipboardButton id="elapsed-time-clipboard-btn" data={response.elapsedTime} />
		</div>
		<p>{response.elapsedTime}</p>
		<hr class="my-2 mr-2" />
		<div class="mr-2 flex justify-between">
			<h4 class="text-text-highlight">Send Header:</h4>
			<ClipboardButton id="send-header-clipboard-btn" data={JSON.stringify(response.sendHeader)} />
		</div>
		<div data-testid="send-headers">
			{#each Object.entries(response.sendHeader) as [key, value] (key)}
				<p><span class="text-text-response-headers">{key}</span> {value}</p>
			{/each}
		</div>
		<hr class="my-2 mr-2" />
		<div class="mr-2 flex justify-between">
			<h4 class="text-text-highlight">Received Header:</h4>
			<ClipboardButton id="received-header-clipboard-btn" data={JSON.stringify(response.receivedHeader)} />
		</div>
		{#each Object.entries(response.receivedHeader) as [key, value] (key)}
			<p class="mr-2 break-words">
				<span class="text-text-response-headers">{key}</span>
				{value}
			</p>
		{/each}
		{#if response.error !== ''}
			<hr class="my-2 mr-2" />
			<h4 class="text-text-highlight">Error:</h4>
			<p>{response.error}</p>
		{/if}
		<hr class="my-2 mr-2" />
		<div class="mr-2 flex justify-between">
			<h4 class="text-text-highlight">Payload:</h4>
			<ClipboardButton id="payload-clipboard-btn" data={response.responseBody} />
		</div>
		{#if isJsonResponse}
			<div class="mr-2">
				<pre data-testid="payload" class="overflow-x-auto whitespace-pre-wrap">{response.responseBody}</pre>
			</div>
		{:else}
			<div class="mr-2 break-words">{response.responseBody}</div>
		{/if}
	</div>
{/if}
