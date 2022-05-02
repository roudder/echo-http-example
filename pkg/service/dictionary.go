package service

import "awesomeProject/pkg/repository"

type DictionaryService struct {
	rep repository.DictRepository
}

type DictService interface {
	Get(key string) string
	Set(key string, value string)
}

func (d *DictionaryService) Get(key string) string {
	return d.rep.Get(key)
}

func (d *DictionaryService) Set(key string, value string) {
	d.rep.Set(key, value)
}

func NewDictionaryService(rep repository.DictRepository) DictService {
	return &DictionaryService{
		rep: rep,
	}
}
