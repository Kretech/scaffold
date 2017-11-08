package config

type Repository struct {
	items map[string]interface{}
}

func NewRepository() *Repository {
	return &Repository{
		make(map[string]interface{}, 16),
	}
}

func (repo *Repository) All() map[string]interface{} {
	return repo.items
}

func (repo *Repository) Has(key string) bool {
	return repo.items[key] != nil
}

func (repo *Repository) Get(key string) interface{} {
	return repo.items[key]
}

func (repo *Repository) Set(key string, value interface{}) {
	repo.items[key] = value
}
