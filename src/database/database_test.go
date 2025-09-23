package database

import (
	"api-client/src/test"
	"testing"
)

func TestDatabase(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-database_test/"}
	defer userDir.Cleanup()
	databaseClient := NewClient(&userDir)
	AutoMigrate(databaseClient)

	projectRepository := NewRepository[Project](databaseClient)
	collectionRepository := NewCollectionRepository(databaseClient)
	httpRequestRepository := NewHttpRequestRepository(databaseClient)
	websocketRequestRepository := NewWebsocketRequestRepository(databaseClient)
	environmentRepository := NewEnvironmentRepository(databaseClient)

	project := &Project{
		Name:        "NewProject1",
		Collections: nil,
	}
	project, err := projectRepository.Create(project)
	if err != nil {
		t.Fatal("project was not created")
	}
	project.Name = "NewProject2"
	project, err = projectRepository.Update(project)
	if err != nil || project.Name != "NewProject2" {
		t.Fatal("project was not updated")
	}
	projects, err := projectRepository.GetAll()
	if err != nil {
		t.Fatal("projects was not retrieved")
	}
	if len(projects) != 1 {
		t.Fatal("expected 1 project")
	}
	_, err = projectRepository.GetById(project.ID)
	if err != nil {
		t.Fatalf("project with id %d was", project.ID)
	}

	collection := &Collection{
		Name:              "TestCollection1",
		HttpRequests:      nil,
		GrpcRequests:      nil,
		WebsocketRequests: nil,
		ProjectID:         project.ID,
	}
	collection, err = collectionRepository.Create(collection)
	if err != nil {
		t.Fatal("collection was not created")
	}
	collection.Name = "NewCollection2"
	collection, err = collectionRepository.Update(collection)
	if err != nil {
		t.Fatal("collection was not updated")
	}
	collections, err := collectionRepository.GetAll()
	if err != nil {
		t.Fatal("collections was not retrieved")
	}
	if len(collections) != 1 {
		t.Fatal("expected 1 collection")
	}

	httpRequest := &HttpRequest{
		Name:         "httpRequest1",
		CollectionID: collection.ID,
		Url:          "superurl",
		HttpRequestBody: HttpRequestBody{
			Type:    "none",
			Payload: "",
		},
	}
	httpRequest, err = httpRequestRepository.Create(httpRequest)
	if err != nil {
		t.Fatal("http request was not created")
	}
	httpRequest.Url = "superurl2"
	httpRequest.HttpRequestBody.Payload = "payload"
	httpRequest, err = httpRequestRepository.Update(httpRequest)
	if err != nil {
		t.Fatal("http request was not updated")
	}

	httpRequestHeader := &HttpRequestHeader{
		HttpRequestID: httpRequest.ID,
		Key:           "tom",
		Value:         "riddle",
	}
	httpRequestHeader, err = httpRequestRepository.CreateHeader(httpRequestHeader)
	if err != nil {
		t.Fatal("http request header was not created")
	}
	httpRequestHeader.Key = "lord"
	httpRequestHeader.Value = "voldemort"
	_, err = httpRequestRepository.UpdateHeader(httpRequestHeader)
	if err != nil {
		t.Fatal("http request header could not be updated")
	}

	httpRequestParameter := &HttpRequestParameter{
		HttpRequestID: httpRequest.ID,
		Key:           "hero",
		Value:         "aax",
	}
	httpRequestParameter, err = httpRequestRepository.CreateParameter(httpRequestParameter)
	if err != nil {
		t.Fatal("http request parameter was not created")
	}
	httpRequestParameter.Key = "tery"
	httpRequestParameter.Value = "on"
	_, err = httpRequestRepository.UpdateParameter(httpRequestParameter)
	if err != nil {
		t.Fatal("http request parameter could not be updated")
	}

	httpRequests, err := httpRequestRepository.GetAll()
	if err != nil {
		t.Fatal("http request was not retrieved")
	}
	if len(httpRequests) != 1 {
		t.Fatal("expected 1 http request")
	}
	fullHttpRequest, err := httpRequestRepository.GetById(httpRequest.ID)
	if err != nil {
		t.Fatal("http request was not retrieved")
	}
	fullHttpRequest, _ = httpRequestRepository.GetById(httpRequest.ID)
	if len(fullHttpRequest.HttpRequestParameter) != 1 {
		t.Fatal("expected 2 http request parameters")
	}

	fullHttpRequest, _ = httpRequestRepository.GetById(httpRequest.ID)
	if len(fullHttpRequest.HttpRequestHeader) != 1 {
		t.Fatal("expected 2 http request headers")
	}

	err = httpRequestRepository.DeleteHeader(&fullHttpRequest.HttpRequestHeader[0])
	if err != nil {
		t.Fatal("http request header was not deleted")
	}

	err = httpRequestRepository.DeleteParameter(&fullHttpRequest.HttpRequestParameter[0])
	if err != nil {
		t.Fatal("http request parameter was not deleted")
	}

	// environment
	environment := &Environment{
		Name: "superenv",
	}

	environment, err = environmentRepository.Create(environment)
	if err != nil {
		t.Fatal("could not create environment")
	}

	environment.Name = "megaenv"
	environment, err = environmentRepository.Update(environment)
	if err != nil {
		t.Fatal("could not update environment")
	}

	environments, err := environmentRepository.GetAll()
	if err != nil {
		t.Fatal("could not get all environments")
	}
	if len(environments) != 1 {
		t.Fatal("there should only be one environment")
	}

	environmentHeader := &EnvironmentHeader{
		EnvironmentID: environment.ID,
		Key:           "yolo",
		Value:         "swag",
	}
	environmentHeader, err = environmentRepository.CreateHeader(environmentHeader)
	if err != nil {
		t.Fatal("could not create environment header")
	}
	environmentHeader.Key = "bilbo"
	environmentHeader.Value = "beutlin"
	environmentHeader, err = environmentRepository.UpdateHeader(environmentHeader)
	if err != nil {
		t.Fatal("could not update environment header")
	}

	httpRequestDisabledEnvironmentHeader := &HttpRequestDisabledEnvironmentHeader{
		HttpRequestID:       httpRequest.ID,
		EnvironmentHeaderID: environmentHeader.ID,
	}
	httpRequestDisabledEnvironmentHeader, err = httpRequestRepository.AddDisabledHeader(httpRequestDisabledEnvironmentHeader)
	if err != nil {
		t.Fatal("environment header could not be disabled")
	}
	err = httpRequestRepository.RemoveDisabledHeader(httpRequestDisabledEnvironmentHeader)
	if err != nil {
		t.Fatal("environment header could not be enabled")
	}

	// websocket
	websocketRequest := &WebsocketRequest{
		Name:         "superwebsocketrequest1",
		CollectionID: collection.ID,
		Url:          "ws://localhost:12345",
	}

	websocketRequest, err = websocketRequestRepository.Create(websocketRequest)
	if err != nil {
		t.Fatal("websocket request could not be created")
	}

	websocketRequest.Name = "superwebsocketrequest2"
	websocketRequest, err = websocketRequestRepository.Update(websocketRequest)
	if err != nil {
		t.Fatal("websocket request could not be updated")
	}

	websocketRequests, err := websocketRequestRepository.GetAll()
	if err != nil || len(websocketRequests) != 1 {
		t.Fatal("could not get websocket requests")
	}

	websocketRequest, err = websocketRequestRepository.GetById(websocketRequest.ID)
	if err != nil {
		t.Fatal("could not get websocket request by id")
	}

	err = projectRepository.Delete(project)
	if err != nil {
		t.Fatal("project was not deleted")
	}
	err = collectionRepository.Delete(collection)
	if err != nil {
		t.Fatal("collection was not deleted")
	}
	err = httpRequestRepository.Delete(httpRequest)
	if err != nil {
		t.Fatal("http request was not deleted")
	}
	err = websocketRequestRepository.Delete(websocketRequest)
	if err != nil {
		t.Fatal("websocket request was not deleted")
	}
	err = environmentRepository.DeleteHeader(environmentHeader)
	if err != nil {
		t.Fatal("environment header was not deleted")
	}
	err = environmentRepository.Delete(environment)
	if err != nil {
		t.Fatal("environment was not deleted")
	}
}
