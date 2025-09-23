<script lang="ts">
	import './styles.css';
	import Header from './Header.svelte';
	import EnvironmentModal from '$lib/EnvironmentModal.svelte';

	import { getConfigurationStore, initializeConfigurationStore } from '../lib/configurationStore.svelte.ts';
	import { initializeThemeStore } from '../lib/themeStore.svelte.ts';
	import { initializeProjectStore } from '../lib/projectStore.svelte.ts';
	import { initializeCollectionStore } from '../lib/collectionStore.svelte.ts';
	import { initializeRequestStore } from '../lib/requestStore.svelte.ts';
	import { initializeNavigationSystem } from '$lib/navigationSystem.svelte.ts';
	import { initializeWebsocketStore } from '$lib/websocketStore.svelte.ts';
	import { initializeSettingStore } from '$lib/settingStore.svelte.ts';
	import { initializeEnvironmentStore } from '$lib/environmentStore.svelte.ts';
	import SettingModal from '$lib/SettingModal.svelte';

	let { children, data } = $props();

	let showSettingModal: boolean = $state<boolean>(false);
	let showEnvironmentModal: boolean = $state<boolean>(false);

	initializeConfigurationStore(data.configuration);
	initializeProjectStore(data.projects);
	initializeCollectionStore(data.collections);
	initializeRequestStore(data.requests);
	initializeEnvironmentStore(data.environments);
	initializeThemeStore(getConfigurationStore());
	initializeSettingStore(getConfigurationStore());
	initializeNavigationSystem();
	initializeWebsocketStore();

	document.getElementById('loader')?.remove();

	function settingOnClick(): void {
		showSettingModal = true;
	}
	function environmentOnClick(): void {
		showEnvironmentModal = true;
	}
</script>

<div class="bg-background text-text flex h-[100vh] w-[100vw] flex-col overflow-hidden">
	<Header {settingOnClick} {environmentOnClick} />
	<main class="flex h-full w-full flex-1 flex-col overflow-hidden px-2 pb-2">
		{@render children()}
	</main>
</div>

<SettingModal bind:showSettingModal />
<EnvironmentModal bind:showEnvironmentModal />

<style>
</style>
