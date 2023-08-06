package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (model.User, error)
	CreateUser(user model.User) (model.User, error)
	GetUserTaskCategory() ([]model.UserTaskCategory, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(email string) (model.User, error) {
	var User model.User
	err := r.db.Where("email = ?", email).First(&User).Error
	if err != nil {
		return model.User{}, err
	}

	return User, nil
}

func (r *userRepository) CreateUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetUserTaskCategory() ([]model.UserTaskCategory, error) {
	var Tasks []model.Task
	var UserTaskCategories []model.UserTaskCategory
	res := r.db.Find(&Tasks)
	if res.Error != nil {
		return []model.UserTaskCategory{}, res.Error
	}

	var cats []model.Category
	_ = r.db.Find(&cats)
	catsMap := make(map[int]model.Category)
	for _, cat := range cats {
		catsMap[cat.ID] = cat
	}

	var users []model.User
	_ = r.db.Find(&users)
	usersMap := make(map[int]model.User)
	for _, user := range users {
		usersMap[user.ID] = user
	}

	for _, task := range Tasks {
		cat := catsMap[task.CategoryID]
		user := usersMap[task.UserID]
		var temp model.UserTaskCategory
		if (user.Fullname != "") {
			temp.Category = cat.Name
			temp.Deadline = task.Deadline
			temp.Email = user.Email
			temp.Fullname = user.Fullname
			temp.ID = user.ID
			temp.Priority = task.Priority
			temp.Task = task.Title
			temp.Status = task.Status
			
			UserTaskCategories = append(UserTaskCategories, temp)
		}
	}

	return UserTaskCategories, nil // TODO: replace this	
}
