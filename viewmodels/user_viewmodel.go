package viewmodels

import (
	"time"

	"go-beckend/models"
)

// UserViewModel adalah representasi data yang akan dikirim ke client
type UserViewModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ConvertToUserViewModel mengubah model User menjadi UserViewModel
func ConvertToUserViewModel(user models.User) UserViewModel {
	return UserViewModel{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// ConvertToUserViewModels mengubah array model User menjadi array UserViewModel
func ConvertToUserViewModels(users []models.User) []UserViewModel {
	var viewModels []UserViewModel
	for _, user := range users {
		viewModels = append(viewModels, ConvertToUserViewModel(user))
	}
	return viewModels
}
