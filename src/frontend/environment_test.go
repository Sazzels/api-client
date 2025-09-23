package frontend

import (
	"api-client/src/database"
	"api-client/src/test"
	"testing"
)

func TestEnvironment(t *testing.T) {
	userDir := test.UserDir{Dir: "./tmp-environment_test/"}
	defer userDir.Cleanup()
	databaseClient := database.NewClient(&userDir)
	database.AutoMigrate(databaseClient)
	environmentRepository := database.NewEnvironmentRepository(databaseClient)
	environment := NewEnvironment(environmentRepository)

	environmentDtos, err := environment.GetAll()
	if err != nil {
		t.Fatal(err)
	}
	if len(environmentDtos) != 0 {
		t.Fatal("environments should be empty")
	}

	environmentDto := EnvironmentDTO{
		Name: "superenv",
	}

	environmentDto, err = environment.Create(environmentDto)
	if err != nil {
		t.Fatal("error creating new environment", err)
	}
	if environmentDto.ID != 1 {
		t.Fatal("new environment id should be 1")
	}
	environmentDto.Name = "megaenv"
	environmentDto, err = environment.Update(environmentDto)
	if err != nil {
		t.Fatal("error updating environment", err)
	}
	if environmentDto.Name != "megaenv" {
		t.Fatal("new environment name should be megaenv")
	}

	environmentHeaderDto := EnvironmentHeaderDTO{
		Key:   "foo",
		Value: "bar",
	}
	environmentHeaderDto, err = environment.AddHeader(environmentHeaderDto, environmentDto)
	if err != nil {
		t.Fatal("could not add header")
	}
	environmentHeaderDto.Key = "boo"
	environmentHeaderDto.Value = "far"
	environmentHeaderDto, err = environment.UpdateHeader(environmentHeaderDto)

	err = environment.RemoveHeader(environmentHeaderDto)
	if err != nil {
		t.Fatal("error removing header")
	}

	err = environment.Delete(environmentDto)
	if err != nil {
		t.Fatal("error deleting environment", err)
	}
}
