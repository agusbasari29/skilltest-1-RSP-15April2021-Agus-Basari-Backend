package user

import "gorm.io/gorm"

type Repository interface {
	InsertUser(user User) (User, error)
	UpdateUser(user User) (User, error)
	GetAllUser() []User
	DeleteUser(user User) error
	FindUserByEmail(email string) (User, error)
	FindEmail(email string) *User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) InsertUser(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	err := r.db.First(&user)
	if err != nil {
		return user, err.Error
	}
	r.db.Save(&user)
	return user, nil
}

func (r *repository) GetAllUser() []User {
	var user []User
	err := r.db.Find(&user)
	if err == nil {
		return nil
	}
	return user
}

func (r *repository) DeleteUser(user User) error {
	err := r.db.Delete(&user)
	if err != nil {
		return err.Error
	}
	return nil
}

func (r *repository) FindUserByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}


func (r *repository) FindEmail(email string) *User {
	var user User

	err := r.db.First(&user, "email = ?", email).Error
	if err == nil {
		return &user
	}
	return nil
}
