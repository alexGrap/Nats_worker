package usecase

import (
	"L0/config"
	"L0/db/postgres"
	models "L0/internal"
	"errors"
	"log"
)

type UseCase struct {
	Repository postgres.PsqlRepo
	InMemory   map[string][]byte
}

func InitUseCase(conf *config.Config) models.UseCase {
	useCase := UseCase{}
	useCase.InMemory = make(map[string][]byte)
	useCase.Repository = postgres.InitPsqlDB(conf)
	err := useCase.Restore()
	if err != nil {
		log.Println(err)
	}
	return &useCase
}

func (useCase *UseCase) Save(id string, jsonModel []byte) error {
	_, value := useCase.InMemory[id]
	if value {
		log.Println(errors.New("already exist"))
	}
	useCase.InMemory[id] = jsonModel
	err := useCase.Repository.Create(id, jsonModel)
	if err != nil {
		return err
	}

	return nil
}

func (useCase *UseCase) Get(id string) ([]byte, error) {
	el, value := useCase.InMemory[id]
	if !value {
		return nil, errors.New("non stable model")
	}

	return el, nil
}

func (useCase *UseCase) GetAll() ([]models.StoreStruct, error) {
	var result []models.StoreStruct
	tmp := models.StoreStruct{}
	for key, value := range useCase.InMemory {
		tmp.Id = key
		tmp.Body = value
		result = append(result, tmp)
	}
	return result, nil
}

func (useCase *UseCase) Restore() error {
	rows, err := useCase.Repository.CheckIfExist()
	strut := models.StoreStruct{}
	if err != nil {
		err := useCase.Repository.InitTable()
		if err != nil {
			return errors.New("troubles with storage:" + err.Error())
		}
		return errors.New("empty storage. Table was created now")

	}
	for rows.Next() {
		if err := rows.Scan(&strut.Id, &strut.Body); err != nil {
			log.Println(err)
		}
		useCase.InMemory[strut.Id] = strut.Body
	}
	return nil
}
