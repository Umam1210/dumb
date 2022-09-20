package categorysdto

type CategoryResponse struct {
	ID   int    `json:"id" gorm:"primary_key:auto_increment"`
	Name string `json:"name" form:"name" validate:"required"`
}
