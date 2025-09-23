import { getContext, setContext } from 'svelte';
import { frontend } from './wailsjs/go/models';
import { Create, Update, Delete } from './wailsjs/go/frontend/HttpRequests';
import {
	Create as CreateWebsocket,
	Update as UpdateWebsocket,
	Delete as DeleteWebsocket,
} from './wailsjs/go/frontend/WebsocketRequests';
import HttpRequestDto = frontend.HttpRequestDto;
import WebsocketRequestDto = frontend.WebsocketRequestDto;
import { RequestTypes } from '$lib/enums/RequestTypes.ts';
import {
	AddDisabledHeader,
	AddHeader,
	AddParameter,
	RemoveDisabledHeader,
	RemoveHeader,
	RemoveParameter,
	UpdateHeader,
	UpdateParameter,
} from '$lib/wailsjs/go/frontend/HttpRequests';
import HttpRequestHeaderDto = frontend.HttpRequestHeaderDto;
import HttpRequestParameterDto = frontend.HttpRequestParameterDto;
import EnvironmentHeaderDTO = frontend.EnvironmentHeaderDTO;

export class RequestStore {
	private _requests: (HttpRequestDto | WebsocketRequestDto)[] = $state([]);

	constructor(requests: (HttpRequestDto | WebsocketRequestDto)[]) {
		this._requests = requests;
	}

	public getByCollectionId(collectionId: number): (HttpRequestDto | WebsocketRequestDto)[] {
		return this._requests.filter(
			(request: HttpRequestDto | WebsocketRequestDto) => request.collectionId === collectionId,
		);
	}

	public update(request: HttpRequestDto | WebsocketRequestDto): Promise<void> {
		return new Promise<void>((resolve) => {
			switch (request.type) {
				case RequestTypes.HTTP:
					Update(request as HttpRequestDto).then((newRequest: HttpRequestDto) => {
						const index = this._requests.findIndex((request) => request.id === newRequest.id);
						this._requests[index] = newRequest;
						resolve();
					});
					break;
				case RequestTypes.WEBSOCKET:
					UpdateWebsocket(request as WebsocketRequestDto).then((newRequest: WebsocketRequestDto) => {
						const index = this._requests.findIndex((request) => request.id === newRequest.id);
						this._requests[index] = newRequest;
						resolve();
					});
					break;
			}
		});
	}

	public addHeader(httpRequestHeaderDto: HttpRequestHeaderDto, httpRequestDto: HttpRequestDto): void {
		AddHeader(httpRequestHeaderDto, httpRequestDto).then((newHeader) => {
			httpRequestDto.header ??= [];
			httpRequestDto.header.push(newHeader);
		});
	}

	public deleteHeader(httpRequestHeaderDto: HttpRequestHeaderDto, httpRequestDto: HttpRequestDto, index: number): void {
		RemoveHeader(httpRequestHeaderDto).then(() => {
			httpRequestDto.header.splice(index, 1);
		});
	}
	public updateHeader(httpRequestHeaderDto: HttpRequestHeaderDto, httpRequestDto: HttpRequestDto, index: number): void {
		UpdateHeader(httpRequestHeaderDto).then((newHeader) => {
			httpRequestDto.header[index] = newHeader;
		});
	}

	public addParameter(httpRequestParameterDto: HttpRequestParameterDto, httpRequestDto: HttpRequestDto): void {
		AddParameter(httpRequestParameterDto, httpRequestDto).then((newParameter) => {
			httpRequestDto.parameter ??= [];
			httpRequestDto.parameter.push(newParameter);
		});
	}

	public deleteParameter(
		httpRequestParameterDto: HttpRequestParameterDto,
		httpRequestDto: HttpRequestDto,
		index: number,
	): void {
		RemoveParameter(httpRequestParameterDto).then(() => {
			httpRequestDto.parameter.splice(index, 1);
		});
	}

	public updateParameter(
		httpRequestParameterDto: HttpRequestParameterDto,
		httpRequestDto: HttpRequestDto,
		index: number,
	): void {
		UpdateParameter(httpRequestParameterDto).then((newParameter) => {
			httpRequestDto.parameter[index] = newParameter;
		});
	}

	public disableEnvironmentHeader(environmentHeaderDTO: EnvironmentHeaderDTO, httpRequestDto: HttpRequestDto): void {
		AddDisabledHeader(environmentHeaderDTO, httpRequestDto).then(() => {
			httpRequestDto.disabledEnvironmentHeader ??= [];
			httpRequestDto.disabledEnvironmentHeader.push(environmentHeaderDTO.id);
		});
	}

	public enableEnvironmentHeader(environmentHeaderDTO: EnvironmentHeaderDTO, httpRequestDto: HttpRequestDto): void {
		RemoveDisabledHeader(environmentHeaderDTO, httpRequestDto).then(() => {
			httpRequestDto.disabledEnvironmentHeader = httpRequestDto.disabledEnvironmentHeader.filter(
				(value) => value !== environmentHeaderDTO.id,
			);
		});
	}

	public create(collectionId: number, newRequestName: string, requestType: RequestTypes): void {
		switch (requestType) {
			case RequestTypes.HTTP: {
				const httpRequestDto = new HttpRequestDto();
				httpRequestDto.collectionId = collectionId;
				httpRequestDto.name = newRequestName;
				Create(httpRequestDto).then((newRequest: HttpRequestDto) => {
					const collections = this._requests.toReversed();
					collections.push(newRequest);
					this._requests = collections.toReversed();
				});
				break;
			}
			case RequestTypes.WEBSOCKET: {
				const websocketRequestDto = new WebsocketRequestDto();
				websocketRequestDto.collectionId = collectionId;
				websocketRequestDto.name = newRequestName;
				CreateWebsocket(websocketRequestDto).then((newRequest: WebsocketRequestDto) => {
					this._requests.push(newRequest);
				});
				break;
			}
		}
	}

	public delete(requestDto: HttpRequestDto | WebsocketRequestDto): void {
		switch (requestDto.type) {
			case RequestTypes.HTTP:
				Delete(requestDto as HttpRequestDto).then(() => {
					this._requests = this._requests.filter((request) => request.id !== requestDto.id);
				});
				break;
			case RequestTypes.WEBSOCKET:
				DeleteWebsocket(requestDto as WebsocketRequestDto).then(() => {
					this._requests = this._requests.filter((request) => request.id !== requestDto.id);
				});
				break;
		}
	}
}

const requestStoreContextKey = 'requestStore';

export function initializeRequestStore(requests: (HttpRequestDto | WebsocketRequestDto)[]): void {
	setContext<RequestStore>(requestStoreContextKey, new RequestStore(requests));
}

export function getRequestStore(): RequestStore {
	return getContext<RequestStore>(requestStoreContextKey);
}
