import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Create, Update, Delete } from './wailsjs/go/frontend/Environment';
import { AddHeader, RemoveHeader, UpdateHeader } from '$lib/wailsjs/go/frontend/Environment';

export class EnvironmentStore {
	private _environments: frontend.EnvironmentDTO[] = $state([]);

	get environments(): frontend.EnvironmentDTO[] {
		return this._environments;
	}

	constructor(environments: frontend.EnvironmentDTO[]) {
		this._environments = environments;
	}

	public create(name: string): void {
		const dto = new frontend.EnvironmentDTO();
		dto.name = name;
		Create(dto).then((environment) => {
			const environmentsReversed = this._environments.toReversed();
			environmentsReversed.push(environment);
			this._environments = environmentsReversed.toReversed();
		});
	}

	public delete(environment: frontend.EnvironmentDTO): void {
		Delete(environment).then(() => {
			this._environments = this._environments.filter((_environment) => {
				return environment.id !== _environment.id;
			});
		});
	}

	public update(environment: frontend.EnvironmentDTO): void {
		Update(environment).then((newEnvironment) => {
			const index = this._environments.findIndex((_environment) => _environment.id === newEnvironment.id);
			this._environments[index] = newEnvironment;
		});
	}

	public addHeader(environmentHeader: frontend.EnvironmentHeaderDTO, environment: frontend.EnvironmentDTO): void {
		AddHeader(environmentHeader, environment).then((newHeader) => {
			environment.header ??= [];
			environment.header.push(newHeader);
		});
	}

	public deleteHeader(
		environmentHeader: frontend.EnvironmentHeaderDTO,
		environment: frontend.EnvironmentDTO,
		index: number,
	): void {
		RemoveHeader(environmentHeader).then(() => {
			environment.header.splice(index, 1);
		});
	}
	public updateHeader(
		environmentHeader: frontend.EnvironmentHeaderDTO,
		environment: frontend.EnvironmentDTO,
		index: number,
	): void {
		UpdateHeader(environmentHeader).then((newHeader) => {
			environment.header[index] = newHeader;
		});
	}

	public getById(id: number): frontend.EnvironmentDTO {
		return this._environments.find((environment) => environment.id === id) || new frontend.EnvironmentDTO();
	}
}

const environmentStoreContextKey = 'environmentStore';

export function initializeEnvironmentStore(environments: frontend.EnvironmentDTO[]): void {
	setContext<EnvironmentStore>(environmentStoreContextKey, new EnvironmentStore(environments));
}

export function getEnvironmentStore(): EnvironmentStore {
	return getContext<EnvironmentStore>(environmentStoreContextKey);
}
