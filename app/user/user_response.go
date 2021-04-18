package user

import "time"

type ResponseUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	AuthToken string
	UpdatedAt time.Time
}

func UserResponseFormatter(user User, auth_token string) ResponseUser {
	formatter := ResponseUser{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		AuthToken: auth_token,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	return formatter
}
