package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Collection struct {
	gorm.Model
	Name              string
	ProjectID         uint
	HttpRequests      []HttpRequest
	GrpcRequests      []GrpcRequest
	WebsocketRequests []WebsocketRequest
	EnvironmentID     *uint
	Environment       Environment
}

type CollectionRepository struct {
	database *gorm.DB
}

func NewCollectionRepository(database *gorm.DB) *CollectionRepository {
	return &CollectionRepository{database: database}
}

func (C *CollectionRepository) GetAll() ([]Collection, error) {
	var collections []Collection
	err := C.database.Find(&collections).Error
	if err != nil {
		return []Collection{}, err
	}

	return collections, nil
}

func (C *CollectionRepository) GetById(id uint) (*Collection, error) {
	collon := &Collection{}
	err := C.database.Model(&Collection{}).Where("id = ?", id).Preload("Environment.Header").Preload(clause.Associations).First(collon).Error
	if err != nil {
		return nil, err
	}

	return collon, nil
}

func (C *CollectionRepository) Create(collection *Collection) (*Collection, error) {
	err := C.database.Create(collection).Error
	if err != nil {
		return nil, err
	}

	return collection, nil
}

func (C *CollectionRepository) Update(collection *Collection) (*Collection, error) {
	err := C.database.Select("Name", "UpdatedAt", "ProjectID", "EnvironmentID").Updates(collection).Error
	if err != nil {
		return collection, err
	}

	return collection, nil
}

func (C *CollectionRepository) Delete(collection *Collection) error {
	return C.database.Delete(collection).Error
}
