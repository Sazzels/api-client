package database

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type HttpRequest struct {
	gorm.Model
	Name                                 string
	CollectionID                         uint
	Url                                  string
	Method                               string `gorm:"default:GET"`
	HttpRequestBody                      HttpRequestBody
	HttpRequestParameter                 []HttpRequestParameter
	HttpRequestHeader                    []HttpRequestHeader
	HttpRequestDisabledEnvironmentHeader []HttpRequestDisabledEnvironmentHeader
}

type HttpRequestDisabledEnvironmentHeader struct {
	gorm.Model
	HttpRequestID       uint
	EnvironmentHeaderID uint
	EnvironmentHeader   EnvironmentHeader
}

type HttpRequestBody struct {
	gorm.Model
	HttpRequestID uint
	Type          string `gorm:"default:none"`
	Payload       string
}

type HttpRequestParameter struct {
	gorm.Model
	HttpRequestID uint
	Key           string
	Value         string
}

type HttpRequestHeader struct {
	gorm.Model
	HttpRequestID uint
	Key           string
	Value         string
}

type HttpRequestRepository struct {
	database *gorm.DB
}

func NewHttpRequestRepository(database *gorm.DB) *HttpRequestRepository {
	return &HttpRequestRepository{database: database}
}

func (H *HttpRequestRepository) GetAll() ([]HttpRequest, error) {
	var httpRequests []HttpRequest
	err := H.database.Preload(clause.Associations).Find(&httpRequests).Error
	if err != nil {
		return []HttpRequest{}, err
	}

	return httpRequests, nil
}

func (H *HttpRequestRepository) GetById(id uint) (*HttpRequest, error) {
	httpRequest := &HttpRequest{}
	err := H.database.Model(&HttpRequest{}).Where("id = ?", id).Preload(clause.Associations).First(httpRequest).Error
	if err != nil {
		return nil, err
	}

	return httpRequest, nil
}

func (H *HttpRequestRepository) Create(request *HttpRequest) (*HttpRequest, error) {
	err := H.database.Create(request).Error
	if err != nil {
		return nil, err
	}

	return request, nil
}

func (H *HttpRequestRepository) Update(request *HttpRequest) (*HttpRequest, error) {
	err := H.database.Session(&gorm.Session{FullSaveAssociations: true}).Updates(request).Error
	if err != nil {
		return request, err
	}

	return request, nil
}
func (H *HttpRequestRepository) Delete(request *HttpRequest) error {
	return H.database.Delete(request).Error
}

func (H *HttpRequestRepository) CreateParameter(httpRequestParameter *HttpRequestParameter) (*HttpRequestParameter, error) {
	err := H.database.Create(httpRequestParameter).Error
	if err != nil {
		return nil, err
	}

	return httpRequestParameter, nil
}

func (H *HttpRequestRepository) UpdateParameter(httpRequestParameter *HttpRequestParameter) (*HttpRequestParameter, error) {
	err := H.database.Updates(httpRequestParameter).Error
	if err != nil {
		return nil, err
	}

	return httpRequestParameter, nil
}

func (H *HttpRequestRepository) DeleteParameter(httpRequestParameter *HttpRequestParameter) error {
	err := H.database.Delete(httpRequestParameter).Error
	if err != nil {
		return err
	}

	return nil
}

func (H *HttpRequestRepository) CreateHeader(httpRequestHeader *HttpRequestHeader) (*HttpRequestHeader, error) {
	err := H.database.Create(httpRequestHeader).Error
	if err != nil {
		return nil, err
	}

	return httpRequestHeader, nil
}

func (H *HttpRequestRepository) UpdateHeader(httpRequestHeader *HttpRequestHeader) (*HttpRequestHeader, error) {
	err := H.database.Updates(httpRequestHeader).Error
	if err != nil {
		return nil, err
	}

	return httpRequestHeader, nil
}

func (H *HttpRequestRepository) DeleteHeader(httpRequestHeader *HttpRequestHeader) error {
	err := H.database.Delete(httpRequestHeader).Error
	if err != nil {
		return err
	}

	return nil
}

func (H *HttpRequestRepository) AddDisabledHeader(httpRequestDisabledEnvironmentHeader *HttpRequestDisabledEnvironmentHeader) (*HttpRequestDisabledEnvironmentHeader, error) {
	err := H.database.Create(httpRequestDisabledEnvironmentHeader).Error
	if err != nil {
		return nil, err
	}

	return httpRequestDisabledEnvironmentHeader, nil
}

func (H *HttpRequestRepository) RemoveDisabledHeader(httpRequestDisabledEnvironmentHeader *HttpRequestDisabledEnvironmentHeader) error {
	err := H.database.Where("environment_header_id = ? and http_request_id = ?", httpRequestDisabledEnvironmentHeader.EnvironmentHeaderID, httpRequestDisabledEnvironmentHeader.HttpRequestID).Delete(&HttpRequestDisabledEnvironmentHeader{}).Error
	if err != nil {
		panic(err)
		return err
	}

	return nil
}
