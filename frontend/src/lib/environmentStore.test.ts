import { expect, test, vi } from 'vitest';
import * as models from '$lib/wailsjs/go/frontend/Environment';
import { frontend } from '$lib/wailsjs/go/models.ts';
import { EnvironmentStore } from '$lib/environmentStore.svelte.ts';
import EnvironmentDTO = frontend.EnvironmentDTO;
import EnvironmentHeaderDTO = frontend.EnvironmentHeaderDTO;

test('create environment store', async () => {
	vi.spyOn(models, 'Update').mockImplementation((environmentDTO: EnvironmentDTO): Promise<EnvironmentDTO> => {
		return Promise.resolve(environmentDTO);
	});

	vi.spyOn(models, 'Delete').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	vi.spyOn(models, 'Create').mockImplementation((environmentDTO: EnvironmentDTO): Promise<EnvironmentDTO> => {
		environmentDTO.id = 3;

		return Promise.resolve(environmentDTO);
	});

	vi.spyOn(models, 'AddHeader').mockImplementation(
		(environmentHeaderDTO: EnvironmentHeaderDTO): Promise<EnvironmentHeaderDTO> => {
			environmentHeaderDTO.id = 1;

			return Promise.resolve(headerDto);
		},
	);

	vi.spyOn(models, 'UpdateHeader').mockImplementation(
		(environmentHeaderDTO: EnvironmentHeaderDTO): Promise<EnvironmentHeaderDTO> => {
			return Promise.resolve(environmentHeaderDTO);
		},
	);

	vi.spyOn(models, 'RemoveHeader').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	const environmentDtoOne = new EnvironmentDTO(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest0","type":"http","collectionId":1,"url":"url0"}',
	);
	const environmentDtoTwo = new EnvironmentDTO(
		'{"id":2,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest1","type":"websocket","collectionId":2,"url":"url1"}',
	);
	const environmentDtos = [environmentDtoOne, environmentDtoTwo];
	const environmentStore = new EnvironmentStore(environmentDtos);
	expect(environmentStore.getById(1)).toBe(environmentDtoOne);
	expect(environmentStore.getById(2)).toBe(environmentDtoTwo);

	environmentStore.create('superenv');
	await vi.waitFor(() => {
		expect(environmentStore.getById(3).name).toBe('superenv');
	});

	environmentStore.getById(1).name = 'updateenv';
	environmentStore.update(environmentStore.getById(1));
	await vi.waitFor(() => {
		expect(environmentStore.getById(1).name).toBe('updateenv');
	});

	const headerDto = new EnvironmentHeaderDTO('{"key":"keykey","value":"valuevalue"}');
	environmentStore.addHeader(headerDto, environmentDtoOne);
	await vi.waitFor(() => {
		const request = environmentStore.getById(1);
		expect(request.header[0].key).toBe('keykey');
		expect(request.header[0].value).toBe('valuevalue');
	});

	const environmentDTO = environmentStore.getById(1);
	environmentDTO.header[0].key = 'newkeykey';
	environmentStore.updateHeader(environmentDTO.header[0], environmentDTO, 0);
	await vi.waitFor(() => {
		const request = environmentStore.getById(1);
		expect(request.header[0].key).toBe('newkeykey');
	});

	environmentStore.deleteHeader(environmentDTO.header[0], environmentDTO, 0);
	await vi.waitFor(() => {
		const request = environmentStore.getById(1);
		expect(request.header.length).toBe(0);
	});

	environmentStore.delete(environmentStore.getById(1));
	await vi.waitFor(() => {
		expect(environmentStore.environments.length).toBe(2);
	});
});
