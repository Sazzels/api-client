import { Get } from '$lib/wailsjs/go/frontend/Configuration';
import { GetAll } from '$lib/wailsjs/go/frontend/Projects';
import { GetAll as GetAllCollections } from '../lib/wailsjs/go/frontend/Collections';
import { GetAll as GetAllRequests } from '../lib/wailsjs/go/frontend/Requests';
import { GetAll as GetAllEnvironments } from '../lib/wailsjs/go/frontend/Environment';

export const prerender = false;
export const ssr = false;

export async function load() {
	return {
		configuration: (await Get()) || [],
		projects: (await GetAll()) || [],
		collections: (await GetAllCollections()) || [],
		requests: (await GetAllRequests()) || [],
		environments: (await GetAllEnvironments()) || [],
		// delay: await delay(1000),
	};
}
