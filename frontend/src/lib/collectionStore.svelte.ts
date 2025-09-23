import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Create, Delete, Update } from './wailsjs/go/frontend/Collections';

export class CollectionStore {
	private _collections: frontend.CollectionDto[] = $state([]);

	constructor(collections: frontend.CollectionDto[]) {
		this._collections = collections;
	}

	public getByProjectId(projectId: number): frontend.CollectionDto[] {
		return this._collections.filter((collection) => collection.projectId === projectId);
	}

	public create(projectId: number, newCollectionName: string): void {
		const collectionDTO = new frontend.CollectionDto();
		collectionDTO.projectId = projectId;
		collectionDTO.name = newCollectionName;
		Create(collectionDTO).then((newCollection: frontend.CollectionDto) => {
			const collections = this._collections.toReversed();
			collections.push(newCollection);
			this._collections = collections.toReversed();
		});
	}

	public delete(collectionDto: frontend.CollectionDto): void {
		Delete(collectionDto).then(() => {
			this._collections = this._collections.filter((collection) => collection.id !== collectionDto.id);
		});
	}

	public update(collectionDto: frontend.CollectionDto): void {
		Update(collectionDto).then((collection) => {
			const index = this._collections.findIndex((collection) => collection.id === collectionDto.id);
			this._collections[index] = collection;
		});
	}

	public getById(id: number): frontend.CollectionDto {
		return this._collections.find((collection) => collection.id === id) || new frontend.CollectionDto();
	}
}

const collectionStoreContextKey = 'collectionStore';

export function initializeCollectionStore(collections: frontend.CollectionDto[]): void {
	setContext<CollectionStore>(collectionStoreContextKey, new CollectionStore(collections));
}

export function getCollectionStore(): CollectionStore {
	return getContext<CollectionStore>(collectionStoreContextKey);
}
