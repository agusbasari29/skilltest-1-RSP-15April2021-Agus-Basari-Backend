package category

type Services interface {
	AddCategory(req RequestCategory) (Category, error)
	CheckExistCategory(req RequestCategory) *Category
	GetAllCategories() []Category
}

type services struct {
	repository Repository
}

func NewServices(repository Repository) *services {
	return &services{repository}
}

func (s *services) AddCategory(req RequestCategory) (Category, error) {
	cat := Category{}
	cat.Name = req.Name
	newCat, err := s.repository.InsertCategory(cat)
	if err != nil {
		return cat, err
	}
	return newCat, nil
}

func (s *services) CheckExistCategory(req RequestCategory) *Category {
	name := req.Name
	if cat := s.repository.FindCategory(name); cat != nil {
		return cat
	}
	return nil
}

func (s *services) GetAllCategories() []Category {
	get := s.repository.GetAll()
	if get != nil {
		return get
	}
	return nil
}
