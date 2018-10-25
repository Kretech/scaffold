package config

type MapRepository struct {
	items map[string]interface{}
}

func NewMapRepository() *MapRepository {
	return &MapRepository{
		make(map[string]interface{}, 16),
	}
}

func (repo *MapRepository) All() map[string]interface{} {
	return repo.items
}

func (repo *MapRepository) Has(key string) bool {
	return repo.items[key] != nil
}

func (repo *MapRepository) Get(key string) interface{} {
	return repo.items[key]
}

func (repo *MapRepository) Set(key string, value interface{}) {
	repo.items[key] = value
}
