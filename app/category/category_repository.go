package category

import "gorm.io/gorm"

type Repository interface {
	InsertCategory(cat Category) (Category, error)
	FindCategory(name string) *Category
	UpdateCategory(cat Category) (Category, error)
	GetAll() []Category
	DeleteCategory(cat Category) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertCategory(cat Category) (Category, error) {
	err := r.db.Create(&cat).Error
	if err != nil {
		return cat, err
	}
	return cat, nil
}

func (r *repository) FindCategory(name string) *Category {
	var cat Category
	err := r.db.First(&cat, "name = ?", name)
	if err == nil {
		return &cat
	}
	return nil
}

func (r *repository) UpdateCategory(cat Category) (Category, error) {
	err := r.db.First(&cat).Error
	if err != nil {
		return cat, err
	}
	return cat, nil
}

func (r *repository) GetAll() []Category {
	var cat []Category
	get := r.db.Find(&cat)
	if get != nil {
		return cat
	}
	return nil
}

func (r *repository) DeleteCategory(cat Category) error {
	err := r.db.First(&cat).Error
	if err != nil {
		return err
	}
	return nil
}
