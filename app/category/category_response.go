package category

type ResponseCategory struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func CategoryResponseFormatter(category Category) ResponseCategory {
	formatter := ResponseCategory{
		ID:   category.ID,
		Name: category.Name,
	}
	return formatter
}
