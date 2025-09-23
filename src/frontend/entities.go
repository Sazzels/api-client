package frontend

import (
	"api-client/src/database"
	"time"

	"gorm.io/gorm"
)

type ProjectDto struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
}

type CollectionDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Name          string    `json:"name"`
	ProjectID     uint      `json:"projectId"`
	EnvironmentID uint      `json:"environmentId"`
}

type WebsocketRequestDto struct {
	ID           uint      `json:"id"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Name         string    `json:"name"`
	Type         string    `json:"type"`
	CollectionID uint      `json:"collectionId"`
	Url          string    `json:"url"`
}

type HttpRequestDto struct {
	ID                        uint                      `json:"id"`
	UpdatedAt                 time.Time                 `json:"updatedAt"`
	Name                      string                    `json:"name"`
	Type                      string                    `json:"type"`
	CollectionID              uint                      `json:"collectionId"`
	Url                       string                    `json:"url"`
	Method                    string                    `json:"method"`
	Body                      HttpRequestBodyDto        `json:"body"`
	Parameter                 []HttpRequestParameterDto `json:"parameter"`
	Header                    []HttpRequestHeaderDto    `json:"header"`
	DisabledEnvironmentHeader []uint                    `json:"disabledEnvironmentHeader"`
}

type HttpRequestBodyDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Type          string    `json:"type"`
	Payload       string    `json:"payload"`
}

type HttpRequestParameterDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
}

type HttpRequestHeaderDto struct {
	ID            uint      `json:"id"`
	UpdatedAt     time.Time `json:"updatedAt"`
	HttpRequestID uint      `json:"httpRequestID"`
	Key           string    `json:"key"`
	Value         string    `json:"value"`
}

type Projects struct {
	projectRepository *database.Repository[database.Project]
}

type Collections struct {
	collectionRepository *database.CollectionRepository
}

type HttpRequests struct {
	httpRequestRepository *database.HttpRequestRepository
}

type WebsocketRequests struct {
	websocketRequestRepository *database.WebsocketRequestRepository
}

type Requests struct {
	httpRequests      *HttpRequests
	websocketRequests *WebsocketRequests
}

func NewProjects(projectRepository *database.Repository[database.Project]) *Projects {
	return &Projects{projectRepository}
}

func NewCollections(collectionRepository *database.CollectionRepository) *Collections {
	return &Collections{collectionRepository}
}

func NewHttpRequests(httpRequestRepository *database.HttpRequestRepository) *HttpRequests {
	return &HttpRequests{httpRequestRepository}
}

func NewWebsocketRequests(websocketRequestRepository *database.WebsocketRequestRepository) *WebsocketRequests {
	return &WebsocketRequests{websocketRequestRepository}
}

func NewRequests(httpRequests *HttpRequests, websocketRequests *WebsocketRequests) *Requests {
	return &Requests{httpRequests, websocketRequests}
}

func (P *Projects) Create(projectDto ProjectDto) (ProjectDto, error) {
	project := &database.Project{
		Name: projectDto.Name,
	}

	project, err := P.projectRepository.Create(project)
	if err != nil {
		return projectDto, err
	}

	projectDto = ProjectDto{
		ID:        project.ID,
		UpdatedAt: project.UpdatedAt,
		Name:      project.Name,
	}

	return projectDto, nil
}

func (P *Projects) Update(projectDto ProjectDto) (ProjectDto, error) {
	project := &database.Project{
		Model: gorm.Model{
			ID: projectDto.ID,
		},
		Name: projectDto.Name,
	}

	project, err := P.projectRepository.Update(project)
	if err != nil {
		return projectDto, err
	}

	projectDto = ProjectDto{
		ID:        project.ID,
		UpdatedAt: project.UpdatedAt,
		Name:      project.Name,
	}

	return projectDto, nil
}

func (P *Projects) Delete(projectDto ProjectDto) error {
	project := &database.Project{
		Model: gorm.Model{
			ID: projectDto.ID,
		},
	}

	return P.projectRepository.Delete(project)
}

func (P *Projects) GetAll() ([]ProjectDto, error) {
	projects, err := P.projectRepository.GetAll()
	if err != nil {
		return nil, err
	}
	projectDtos := make([]ProjectDto, len(projects))
	for iter, project := range projects {
		projectDto := ProjectDto{
			ID:        project.ID,
			UpdatedAt: project.UpdatedAt,
			Name:      project.Name,
		}

		projectDtos[iter] = projectDto
	}

	return projectDtos, nil
}

func (C *Collections) Create(collectionDto CollectionDto) (CollectionDto, error) {
	collection := &database.Collection{
		Name:      collectionDto.Name,
		ProjectID: collectionDto.ProjectID,
	}

	collection, err := C.collectionRepository.Create(collection)
	if err != nil {
		return collectionDto, err
	}
	collectionDto = CollectionDto{
		ID:        collection.ID,
		UpdatedAt: collection.UpdatedAt,
		Name:      collection.Name,
		ProjectID: collection.ProjectID,
	}

	return collectionDto, nil
}

func (C *Collections) Update(collectionDto CollectionDto) (CollectionDto, error) {
	collection := &database.Collection{
		Model: gorm.Model{
			ID: collectionDto.ID,
		},
		Name:      collectionDto.Name,
		ProjectID: collectionDto.ProjectID,
	}

	collection.EnvironmentID = nil
	if collectionDto.EnvironmentID > 0 {
		collection.EnvironmentID = &collectionDto.EnvironmentID
	}
	collection, err := C.collectionRepository.Update(collection)
	if err != nil {
		return collectionDto, err
	}
	collectionDto = CollectionDto{
		ID:        collection.ID,
		UpdatedAt: collection.UpdatedAt,
		Name:      collection.Name,
		ProjectID: collection.ProjectID,
	}

	return collectionDto, nil
}

func (C *Collections) Delete(collectionDto CollectionDto) error {
	collection := &database.Collection{
		Model: gorm.Model{
			ID: collectionDto.ID,
		},
	}

	return C.collectionRepository.Delete(collection)
}

func (C *Collections) GetAll() ([]CollectionDto, error) {
	collections, err := C.collectionRepository.GetAll()
	if err != nil {
		return nil, err
	}
	collectionDtos := make([]CollectionDto, len(collections))
	for iter, collection := range collections {
		var environementId uint = 0
		if collection.EnvironmentID != nil {
			environementId = *collection.EnvironmentID
		}
		collectionDtos[iter] = CollectionDto{
			ID:            collection.ID,
			UpdatedAt:     collection.UpdatedAt,
			Name:          collection.Name,
			ProjectID:     collection.ProjectID,
			EnvironmentID: environementId,
		}
	}

	return collectionDtos, nil
}

func (R *Requests) GetAll() ([]interface{}, error) {
	httpRequestDtos, err := R.httpRequests.GetAll()
	if err != nil {
		return nil, err
	}
	websocketRequestDtos, err := R.websocketRequests.GetAll()
	if err != nil {
		return nil, err
	}
	requestDtos := make([]interface{}, 0)
	for _, httpRequestDto := range httpRequestDtos {
		requestDtos = append(requestDtos, httpRequestDto)
	}
	for _, websocketRequestDto := range websocketRequestDtos {
		requestDtos = append(requestDtos, websocketRequestDto)
	}

	return requestDtos, err
}

func (H *HttpRequests) GetAll() ([]HttpRequestDto, error) {
	httpRequests, err := H.httpRequestRepository.GetAll()
	if err != nil {
		return nil, err
	}
	httpRequestDtos := make([]HttpRequestDto, len(httpRequests))
	for iter, request := range httpRequests {
		httpRequestDtos[iter] = H.buildDtoFromDatabase(request)
	}

	return httpRequestDtos, nil
}

func (W *WebsocketRequests) GetAll() ([]WebsocketRequestDto, error) {
	websocketRequests, err := W.websocketRequestRepository.GetAll()
	if err != nil {
		return nil, err
	}
	websocketRequestDtos := make([]WebsocketRequestDto, len(websocketRequests))
	for iter, websocketRequest := range websocketRequests {
		websocketRequestDtos[iter] = WebsocketRequestDto{
			ID:           websocketRequest.ID,
			UpdatedAt:    websocketRequest.UpdatedAt,
			Name:         websocketRequest.Name,
			Type:         "websocket",
			CollectionID: websocketRequest.CollectionID,
			Url:          websocketRequest.Url,
		}
	}

	return websocketRequestDtos, nil
}

func (W *WebsocketRequests) Create(websocketRequestDto WebsocketRequestDto) (WebsocketRequestDto, error) {
	websocketRequest := &database.WebsocketRequest{
		Name:         websocketRequestDto.Name,
		CollectionID: websocketRequestDto.CollectionID,
		Url:          websocketRequestDto.Url,
	}
	websocketRequest, err := W.websocketRequestRepository.Create(websocketRequest)
	if err != nil {
		return websocketRequestDto, err
	}
	websocketRequestDto.ID = websocketRequest.ID
	websocketRequestDto.UpdatedAt = websocketRequest.UpdatedAt
	websocketRequestDto.Type = "websocket"

	return websocketRequestDto, nil
}

func (W *WebsocketRequests) Update(websocketRequestDto WebsocketRequestDto) (WebsocketRequestDto, error) {
	websocketRequest, err := W.websocketRequestRepository.GetById(websocketRequestDto.ID)
	if err != nil {
		return websocketRequestDto, err
	}
	websocketRequest.Name = websocketRequestDto.Name
	websocketRequest.Url = websocketRequestDto.Url
	websocketRequest, err = W.websocketRequestRepository.Update(websocketRequest)
	if err != nil {
		return websocketRequestDto, err
	}

	return websocketRequestDto, nil
}

func (W *WebsocketRequests) Delete(websocketRequestDto WebsocketRequestDto) error {
	websocketRequest, err := W.websocketRequestRepository.GetById(websocketRequestDto.ID)
	if err != nil {
		return err
	}

	return W.websocketRequestRepository.Delete(websocketRequest)
}

func (H *HttpRequests) buildDtoFromDatabase(httpRequest database.HttpRequest) HttpRequestDto {
	parameterDtos := make([]HttpRequestParameterDto, len(httpRequest.HttpRequestParameter))
	for iter, parameter := range httpRequest.HttpRequestParameter {
		parameterDtos[iter] = HttpRequestParameterDto{
			ID:            parameter.ID,
			UpdatedAt:     parameter.UpdatedAt,
			HttpRequestID: parameter.HttpRequestID,
			Key:           parameter.Key,
			Value:         parameter.Value,
		}
	}

	headerDtos := make([]HttpRequestHeaderDto, len(httpRequest.HttpRequestHeader))
	for iter, header := range httpRequest.HttpRequestHeader {
		headerDtos[iter] = HttpRequestHeaderDto{
			ID:            header.ID,
			UpdatedAt:     header.UpdatedAt,
			HttpRequestID: header.HttpRequestID,
			Key:           header.Key,
			Value:         header.Value,
		}
	}

	disabledEnvironmentHeaderIds := make([]uint, len(httpRequest.HttpRequestDisabledEnvironmentHeader))
	for iter, environmentHeader := range httpRequest.HttpRequestDisabledEnvironmentHeader {
		disabledEnvironmentHeaderIds[iter] = environmentHeader.EnvironmentHeaderID
	}

	return HttpRequestDto{
		ID:           httpRequest.ID,
		UpdatedAt:    httpRequest.UpdatedAt,
		Name:         httpRequest.Name,
		CollectionID: httpRequest.CollectionID,
		Url:          httpRequest.Url,
		Type:         "http",
		Method:       httpRequest.Method,
		Body: HttpRequestBodyDto{
			ID:            httpRequest.HttpRequestBody.ID,
			UpdatedAt:     httpRequest.HttpRequestBody.UpdatedAt,
			HttpRequestID: httpRequest.HttpRequestBody.HttpRequestID,
			Type:          httpRequest.HttpRequestBody.Type,
			Payload:       httpRequest.HttpRequestBody.Payload,
		},
		Parameter:                 parameterDtos,
		Header:                    headerDtos,
		DisabledEnvironmentHeader: disabledEnvironmentHeaderIds,
	}

}

func (H *HttpRequests) Create(httpRequestDto HttpRequestDto) (HttpRequestDto, error) {
	httpRequest := &database.HttpRequest{
		Name:         httpRequestDto.Name,
		CollectionID: httpRequestDto.CollectionID,
		Url:          httpRequestDto.Url,
		HttpRequestBody: database.HttpRequestBody{
			Type:    "none",
			Payload: "",
		},
	}
	httpRequest, err := H.httpRequestRepository.Create(httpRequest)
	if err != nil {
		return httpRequestDto, err
	}
	httpRequestDto.ID = httpRequest.ID
	httpRequestDto.UpdatedAt = httpRequest.UpdatedAt
	httpRequestDto.Type = "http"
	httpRequestDto.Method = httpRequest.Method
	httpRequestDto.Body.HttpRequestID = httpRequest.ID
	httpRequestDto.Body.ID = httpRequest.HttpRequestBody.ID
	httpRequestDto.Body.Type = httpRequest.HttpRequestBody.Type
	httpRequestDto.Body.Payload = httpRequest.HttpRequestBody.Payload
	httpRequestDto.Body.UpdatedAt = httpRequest.HttpRequestBody.UpdatedAt

	return httpRequestDto, nil
}

func (H *HttpRequests) Update(httpRequestDto HttpRequestDto) (HttpRequestDto, error) {
	httpRequest := &database.HttpRequest{
		Model: gorm.Model{
			ID: httpRequestDto.ID,
		},
		Name:         httpRequestDto.Name,
		CollectionID: httpRequestDto.CollectionID,
		Url:          httpRequestDto.Url,
		Method:       httpRequestDto.Method,
		HttpRequestBody: database.HttpRequestBody{
			Model: gorm.Model{
				ID: httpRequestDto.Body.ID,
			},
			HttpRequestID: httpRequestDto.ID,
			Type:          httpRequestDto.Body.Type,
			Payload:       httpRequestDto.Body.Payload,
		},
	}
	newHttpRequest, err := H.httpRequestRepository.Update(httpRequest)
	if err != nil {
		return httpRequestDto, err
	}

	newHttpRequestDto := H.buildDtoFromDatabase(*newHttpRequest)
	newHttpRequestDto.Header = httpRequestDto.Header
	newHttpRequestDto.Parameter = httpRequestDto.Parameter
	newHttpRequestDto.DisabledEnvironmentHeader = httpRequestDto.DisabledEnvironmentHeader

	return newHttpRequestDto, nil
}

func (H *HttpRequests) Delete(httpRequestDto HttpRequestDto) error {
	httpRequest := &database.HttpRequest{
		Model: gorm.Model{
			ID: httpRequestDto.ID,
		},
	}

	return H.httpRequestRepository.Delete(httpRequest)
}

func (H *HttpRequests) AddHeader(httpRequestHeaderDto HttpRequestHeaderDto, httpRequestDto HttpRequestDto) (HttpRequestHeaderDto, error) {
	httpRequestHeader := &database.HttpRequestHeader{
		HttpRequestID: httpRequestDto.ID,
		Key:           httpRequestHeaderDto.Key,
		Value:         httpRequestHeaderDto.Value,
	}

	var err error
	httpRequestHeader, err = H.httpRequestRepository.CreateHeader(httpRequestHeader)
	if err != nil {
		return httpRequestHeaderDto, err
	}
	httpRequestHeaderDto.ID = httpRequestHeader.ID
	httpRequestHeaderDto.UpdatedAt = httpRequestHeader.UpdatedAt
	httpRequestHeaderDto.HttpRequestID = httpRequestHeader.HttpRequestID

	return httpRequestHeaderDto, nil
}

func (H *HttpRequests) UpdateHeader(httpRequestHeaderDto HttpRequestHeaderDto) (HttpRequestHeaderDto, error) {
	httpRequestHeader := &database.HttpRequestHeader{
		Model:         gorm.Model{ID: httpRequestHeaderDto.ID},
		HttpRequestID: httpRequestHeaderDto.HttpRequestID,
		Key:           httpRequestHeaderDto.Key,
		Value:         httpRequestHeaderDto.Value,
	}

	var err error
	httpRequestHeader, err = H.httpRequestRepository.UpdateHeader(httpRequestHeader)
	if err != nil {
		return httpRequestHeaderDto, err
	}
	httpRequestHeaderDto.UpdatedAt = httpRequestHeader.UpdatedAt

	return httpRequestHeaderDto, nil
}

func (H *HttpRequests) RemoveHeader(httpRequestHeaderDto HttpRequestHeaderDto) error {
	httpRequestHeader := &database.HttpRequestHeader{
		Model:         gorm.Model{ID: httpRequestHeaderDto.ID},
		HttpRequestID: httpRequestHeaderDto.HttpRequestID,
		Key:           httpRequestHeaderDto.Key,
		Value:         httpRequestHeaderDto.Value,
	}

	err := H.httpRequestRepository.DeleteHeader(httpRequestHeader)
	if err != nil {
		return err
	}

	return nil
}

func (H *HttpRequests) AddParameter(httpRequestParameterDto HttpRequestParameterDto, httpRequestDto HttpRequestDto) (HttpRequestParameterDto, error) {
	httpRequestHeader := &database.HttpRequestParameter{
		HttpRequestID: httpRequestDto.ID,
		Key:           httpRequestParameterDto.Key,
		Value:         httpRequestParameterDto.Value,
	}

	var err error
	httpRequestHeader, err = H.httpRequestRepository.CreateParameter(httpRequestHeader)
	if err != nil {
		return httpRequestParameterDto, err
	}
	httpRequestParameterDto.ID = httpRequestHeader.ID
	httpRequestParameterDto.UpdatedAt = httpRequestHeader.UpdatedAt
	httpRequestParameterDto.HttpRequestID = httpRequestHeader.HttpRequestID

	return httpRequestParameterDto, nil
}

func (H *HttpRequests) UpdateParameter(httpRequestParameterDto HttpRequestParameterDto) (HttpRequestParameterDto, error) {
	httpRequestHeader := &database.HttpRequestParameter{
		Model:         gorm.Model{ID: httpRequestParameterDto.ID},
		HttpRequestID: httpRequestParameterDto.HttpRequestID,
		Key:           httpRequestParameterDto.Key,
		Value:         httpRequestParameterDto.Value,
	}

	var err error
	httpRequestHeader, err = H.httpRequestRepository.UpdateParameter(httpRequestHeader)
	if err != nil {
		return httpRequestParameterDto, err
	}
	httpRequestParameterDto.UpdatedAt = httpRequestHeader.UpdatedAt

	return httpRequestParameterDto, nil
}

func (H *HttpRequests) RemoveParameter(httpRequestParameterDto HttpRequestParameterDto) error {
	httpRequestHeader := &database.HttpRequestParameter{
		Model:         gorm.Model{ID: httpRequestParameterDto.ID},
		HttpRequestID: httpRequestParameterDto.HttpRequestID,
		Key:           httpRequestParameterDto.Key,
		Value:         httpRequestParameterDto.Value,
	}

	err := H.httpRequestRepository.DeleteParameter(httpRequestHeader)
	if err != nil {
		return err
	}

	return nil
}

func (H *HttpRequests) AddDisabledHeader(environmentHeaderDto EnvironmentHeaderDTO, httpRequestDto HttpRequestDto) error {
	newEnvironmentHeader := &database.HttpRequestDisabledEnvironmentHeader{
		HttpRequestID:       httpRequestDto.ID,
		EnvironmentHeaderID: environmentHeaderDto.ID,
	}
	newEnvironmentHeader, err := H.httpRequestRepository.AddDisabledHeader(newEnvironmentHeader)
	if err != nil {
		return err
	}

	return nil
}

func (H *HttpRequests) RemoveDisabledHeader(environmentHeaderDto EnvironmentHeaderDTO, httpRequestDto HttpRequestDto) error {
	newEnvironmentHeader := &database.HttpRequestDisabledEnvironmentHeader{
		HttpRequestID:       httpRequestDto.ID,
		EnvironmentHeaderID: environmentHeaderDto.ID,
	}
	err := H.httpRequestRepository.RemoveDisabledHeader(newEnvironmentHeader)
	if err != nil {
		return err
	}

	return nil
}
