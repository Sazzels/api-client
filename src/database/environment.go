package database

import (
	"gorm.io/gorm"
)

type Environment struct {
	gorm.Model
	Name   string `gorm:"unique"`
	Header []EnvironmentHeader
}

type EnvironmentHeader struct {
	gorm.Model
	EnvironmentID uint
	Key           string
	Value         string
}

type EnvironmentRepository struct {
	database *gorm.DB
}

func NewEnvironmentRepository(database *gorm.DB) *EnvironmentRepository {
	return &EnvironmentRepository{database}
}

func (E *EnvironmentRepository) GetAll() ([]Environment, error) {
	var environments []Environment
	err := E.database.Preload("Header").Find(&environments).Error
	if err != nil {
		return []Environment{}, err
	}

	return environments, nil
}

func (E *EnvironmentRepository) Create(environment *Environment) (*Environment, error) {
	err := E.database.Create(environment).Error
	if err != nil {
		return nil, err
	}

	return environment, nil
}

func (E *EnvironmentRepository) Update(environment *Environment) (*Environment, error) {
	err := E.database.Updates(environment).Error
	if err != nil {
		return environment, err
	}

	return environment, nil
}
func (E *EnvironmentRepository) Delete(environment *Environment) error {
	return E.database.Delete(environment).Error
}

func (E *EnvironmentRepository) DeleteHeader(environmentHeader *EnvironmentHeader) error {
	return E.database.Delete(environmentHeader).Error
}

func (E *EnvironmentRepository) CreateHeader(environmentHeader *EnvironmentHeader) (*EnvironmentHeader, error) {
	err := E.database.Create(environmentHeader).Error
	if err != nil {
		return nil, err
	}

	return environmentHeader, nil
}

func (E *EnvironmentRepository) UpdateHeader(environmentHeader *EnvironmentHeader) (*EnvironmentHeader, error) {
	err := E.database.Updates(environmentHeader).Error
	if err != nil {
		return nil, err
	}

	return environmentHeader, nil
}
