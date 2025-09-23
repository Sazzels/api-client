package database

import (
	"api-client/src/configuration"
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	Filename      = "superdatabase.db"
	SqliteOptions = "?cache=shared&_fk=1"
)

func NewClient(userDir configuration.UserDir) *gorm.DB {
	err := os.MkdirAll(userDir.GetDataPath(), 0755)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create data directory")
	}
	database, err := gorm.Open(
		sqlite.Open("file:"+userDir.GetDataPath()+Filename+SqliteOptions),
		&gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		log.Fatal().Msgf("failed opening connection to sqlite: %v", err)
	}

	// this fixes database table lock error on sqlite3 during playwright tests
	db, _ := database.DB()
	db.SetMaxOpenConns(1)

	return database
}

func AutoMigrate(databaseClient *gorm.DB) {
	err := databaseClient.AutoMigrate(
		&EnvironmentHeader{},
		&Environment{},
		&HttpRequestDisabledEnvironmentHeader{},
		&HttpRequestHeader{},
		&HttpRequestParameter{},
		&HttpRequestBody{},
		&HttpRequest{},
		&GrpcRequest{},
		&WebsocketRequest{},
		&Collection{},
		&Project{},
	)
	if err != nil {
		log.Fatal().Msg("migration was not successful: " + err.Error())
	}

	CreateConstraintSafe(databaseClient, &Collection{}, "fk_collections_environment")
	CreateConstraintSafe(databaseClient, &HttpRequestBody{}, "fk_http_requests_http_request_body")
	CreateConstraintSafe(databaseClient, &HttpRequestHeader{}, "fk_http_requests_http_request_header")
	CreateConstraintSafe(databaseClient, &HttpRequestParameter{}, "fk_http_requests_http_request_parameter")
	CreateConstraintSafe(databaseClient, &HttpRequest{}, "fk_collections_http_requests")
	CreateConstraintSafe(databaseClient, &WebsocketRequest{}, "fk_collections_websocket_requests")
	CreateConstraintSafe(databaseClient, &GrpcRequest{}, "fk_collections_grpc_requests")
	CreateConstraintSafe(databaseClient, &Collection{}, "fk_projects_collections")
	CreateConstraintSafe(databaseClient, &EnvironmentHeader{}, "fk_environments_header")
	CreateConstraintSafe(databaseClient, &HttpRequestDisabledEnvironmentHeader{}, "fk_http_requests_http_request_disabled_environment_header")
	CreateConstraintSafe(databaseClient, &HttpRequestDisabledEnvironmentHeader{}, "fk_http_request_disabled_environment_headers_environment_header")

}

func CreateConstraintSafe[T any](databaseClient *gorm.DB, entity *T, constraint string) {
	if databaseClient.Migrator().HasConstraint(entity, constraint) == true {
		return
	}
	err := databaseClient.Migrator().CreateConstraint(entity, constraint)
	if err != nil {
		log.Fatal().Msg("error on constraint creation: " + err.Error())
	}
}

func DeleteConstraintSafe[T any](databaseClient *gorm.DB, entity *T, constraint string) error {
	if databaseClient.Migrator().HasConstraint(entity, constraint) == false {
		return nil
	}
	return databaseClient.Migrator().DropConstraint(entity, constraint)
}
