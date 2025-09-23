package frontend

import (
	"api-client/src/database"
	"time"

	"gorm.io/gorm"
)

type EnvironmentDTO struct {
	ID        uint                   `json:"id"`
	UpdatedAt time.Time              `json:"updatedAt"`
	Name      string                 `json:"name"`
	Header    []EnvironmentHeaderDTO `json:"header"`
}

type EnvironmentHeaderDTO struct {
	ID        uint      `json:"id"`
	UpdatedAt time.Time `json:"updatedAt"`
	Key       string    `json:"key"`
	Value     string    `json:"value"`
}

type Environment struct {
	environmentRepository *database.EnvironmentRepository
}

func NewEnvironment(environmentRepository *database.EnvironmentRepository) *Environment {
	return &Environment{environmentRepository}
}

func (E *Environment) GetAll() ([]EnvironmentDTO, error) {
	environments, err := E.environmentRepository.GetAll()
	if err != nil {
		return []EnvironmentDTO{}, err
	}
	environmentsDTO := make([]EnvironmentDTO, len(environments))
	for iter, environment := range environments {
		environmentsDTO[iter] = EnvironmentDTO{
			ID:        environment.ID,
			UpdatedAt: environment.UpdatedAt,
			Name:      environment.Name,
			Header:    make([]EnvironmentHeaderDTO, len(environment.Header)),
		}
		for jiter, header := range environment.Header {
			environmentsDTO[iter].Header[jiter] = EnvironmentHeaderDTO{
				ID:        header.ID,
				UpdatedAt: header.UpdatedAt,
				Key:       header.Key,
				Value:     header.Value,
			}
		}
	}

	return environmentsDTO, nil
}

func (E *Environment) Create(environmentDTO EnvironmentDTO) (EnvironmentDTO, error) {
	environment := &database.Environment{
		Name:   environmentDTO.Name,
		Header: nil,
	}
	var err error
	environment, err = E.environmentRepository.Create(environment)
	if err != nil {
		return environmentDTO, err
	}
	environmentDTO.ID = environment.ID
	environmentDTO.UpdatedAt = environment.UpdatedAt

	return environmentDTO, nil
}

func (E *Environment) Update(environmentDTO EnvironmentDTO) (EnvironmentDTO, error) {
	environment := &database.Environment{
		Model: gorm.Model{
			ID: environmentDTO.ID,
		},
		Name: environmentDTO.Name,
	}

	environment, err := E.environmentRepository.Update(environment)
	if err != nil {
		return environmentDTO, err
	}

	environmentDTO = EnvironmentDTO{
		ID:        environment.ID,
		UpdatedAt: environment.UpdatedAt,
		Name:      environment.Name,
		Header:    environmentDTO.Header,
	}

	return environmentDTO, nil
}

func (E *Environment) Delete(environmentDTO EnvironmentDTO) error {
	environment := &database.Environment{
		Model: gorm.Model{ID: environmentDTO.ID},
	}

	return E.environmentRepository.Delete(environment)
}

func (E *Environment) AddHeader(environmentHeaderDTO EnvironmentHeaderDTO, environmentDTO EnvironmentDTO) (EnvironmentHeaderDTO, error) {
	environmentHeader := &database.EnvironmentHeader{
		EnvironmentID: environmentDTO.ID,
		Key:           environmentHeaderDTO.Key,
		Value:         environmentHeaderDTO.Value,
	}

	var err error
	environmentHeader, err = E.environmentRepository.CreateHeader(environmentHeader)
	if err != nil {
		return environmentHeaderDTO, err
	}

	environmentHeaderDTO.ID = environmentHeader.ID
	environmentHeaderDTO.UpdatedAt = environmentHeader.UpdatedAt

	return environmentHeaderDTO, nil
}

func (E *Environment) UpdateHeader(environmentHeaderDTO EnvironmentHeaderDTO) (EnvironmentHeaderDTO, error) {
	environmentHeader := &database.EnvironmentHeader{
		Model: gorm.Model{
			ID: environmentHeaderDTO.ID,
		},
		Key:   environmentHeaderDTO.Key,
		Value: environmentHeaderDTO.Value,
	}

	var err error
	environmentHeader, err = E.environmentRepository.UpdateHeader(environmentHeader)
	if err != nil {
		return environmentHeaderDTO, err
	}

	environmentHeaderDTO.UpdatedAt = environmentHeader.UpdatedAt

	return environmentHeaderDTO, nil
}

func (E *Environment) RemoveHeader(environmentHeaderDTO EnvironmentHeaderDTO) error {
	environmentHeader := &database.EnvironmentHeader{
		Model: gorm.Model{
			ID: environmentHeaderDTO.ID,
		},
		Key:   environmentHeaderDTO.Key,
		Value: environmentHeaderDTO.Value,
	}

	err := E.environmentRepository.DeleteHeader(environmentHeader)
	if err != nil {
		return err
	}

	return nil
}
