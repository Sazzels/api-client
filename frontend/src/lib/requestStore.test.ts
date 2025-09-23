import { assert, expect, test, vi } from 'vitest';
import { frontend } from '$lib/wailsjs/go/models';
import { RequestStore } from './requestStore.svelte';
import * as models from '$lib/wailsjs/go/frontend/HttpRequests';
import { RequestTypes } from '$lib/enums/RequestTypes';
import HttpRequestDto = frontend.HttpRequestDto;
import HttpRequestHeaderDto = frontend.HttpRequestHeaderDto;
import HttpRequestParameterDto = frontend.HttpRequestParameterDto;
import EnvironmentHeaderDTO = frontend.EnvironmentHeaderDTO;

test('create request store', async () => {
	vi.spyOn(models, 'Update').mockImplementation((httpRequestDto: HttpRequestDto): Promise<HttpRequestDto> => {
		return Promise.resolve(httpRequestDto);
	});

	vi.spyOn(models, 'Delete').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	vi.spyOn(models, 'Create').mockImplementation((httpRequestDto: HttpRequestDto): Promise<HttpRequestDto> => {
		httpRequestDto.id = 3;
		httpRequestDto.type = RequestTypes.HTTP;

		return Promise.resolve(httpRequestDto);
	});

	vi.spyOn(models, 'AddHeader').mockImplementation(
		(headerDto: HttpRequestHeaderDto, httpRequestDto: HttpRequestDto): Promise<HttpRequestHeaderDto> => {
			headerDto.httpRequestID = httpRequestDto.id;
			headerDto.id = 1;

			return Promise.resolve(headerDto);
		},
	);

	vi.spyOn(models, 'UpdateHeader').mockImplementation(
		(headerDto: HttpRequestHeaderDto): Promise<HttpRequestHeaderDto> => {
			return Promise.resolve(headerDto);
		},
	);

	vi.spyOn(models, 'RemoveHeader').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	vi.spyOn(models, 'AddParameter').mockImplementation(
		(parameterDto: HttpRequestParameterDto, httpRequestDto: HttpRequestDto): Promise<HttpRequestParameterDto> => {
			parameterDto.httpRequestID = httpRequestDto.id;
			parameterDto.id = 1;

			return Promise.resolve(parameterDto);
		},
	);

	vi.spyOn(models, 'UpdateParameter').mockImplementation(
		(headerDto: HttpRequestParameterDto): Promise<HttpRequestParameterDto> => {
			return Promise.resolve(headerDto);
		},
	);

	vi.spyOn(models, 'RemoveParameter').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	vi.spyOn(models, 'AddDisabledHeader').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	vi.spyOn(models, 'RemoveDisabledHeader').mockImplementation((): Promise<void> => {
		return Promise.resolve();
	});

	const requestDtos = [];
	const requestDtoOne = new HttpRequestDto(
		'{"id":1,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest0","type":"http","collectionId":1,"url":"url0"}',
	);
	const requestDtoTwo = new HttpRequestDto(
		'{"id":2,"updatedAt":"2024-08-28T20:49:26.369713204+02:00","name":"httpRequest1","type":"websocket","collectionId":2,"url":"url1"}',
	);
	requestDtos.push(requestDtoOne, requestDtoTwo);
	const requestStore = new RequestStore(requestDtos);
	expect(requestStore.getByCollectionId(1)[0]).toBe(requestDtoOne);
	expect(requestStore.getByCollectionId(2)[0]).toBe(requestDtoTwo);

	requestStore.create(1, 'yolo', RequestTypes.HTTP);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1)[0].name).toBe('yolo');
	});

	requestStore.getByCollectionId(1)[0].name = 'zwolo';
	await requestStore.update(requestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1)[0].name).toBe('zwolo');
	});

	const headerDto = new HttpRequestHeaderDto('{"key":"keykey","value":"valuevalue"}');
	requestStore.addHeader(headerDto, requestDtoOne);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.header[0].key).toBe('keykey');
		expect(request.header[0].value).toBe('valuevalue');
	});

	let request = requestStore.getByCollectionId(1)[1];
	assert(request instanceof HttpRequestDto);
	request.header[0].key = 'newkeykey';
	requestStore.updateHeader(request.header[0], request, 0);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.header[0].key).toBe('newkeykey');
	});

	requestStore.deleteHeader(request.header[0], request, 0);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.header.length).toBe(0);
	});

	const parameterDto = new HttpRequestParameterDto('{"key":"foo","value":"bar"}');
	requestStore.addParameter(parameterDto, requestDtoOne);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.parameter[0].key).toBe('foo');
		expect(request.parameter[0].value).toBe('bar');
	});

	request = requestStore.getByCollectionId(1)[1];
	assert(request instanceof HttpRequestDto);
	request.parameter[0].key = 'boo';
	requestStore.updateParameter(request.parameter[0], request, 0);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.parameter[0].key).toBe('boo');
	});

	requestStore.deleteParameter(request.parameter[0], request, 0);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.parameter.length).toBe(0);
	});

	const environmentHeaderDto = new EnvironmentHeaderDTO('{"id":1,"key":"tom","value":"riddle"}');
	requestStore.disableEnvironmentHeader(environmentHeaderDto, requestDtoOne);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.disabledEnvironmentHeader).toContain(1);
		expect(request.disabledEnvironmentHeader.length).toBe(1);
	});

	requestStore.enableEnvironmentHeader(environmentHeaderDto, requestDtoOne);
	await vi.waitFor(() => {
		const request = requestStore.getByCollectionId(1)[1];
		assert(request instanceof HttpRequestDto);
		expect(request.disabledEnvironmentHeader.length).toBe(0);
	});

	requestStore.delete(requestStore.getByCollectionId(1)[0]);
	await vi.waitFor(() => {
		expect(requestStore.getByCollectionId(1).length).toBe(1);
	});
});
