package repository

import (
	"a21hc3NpZ25tZW50/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Store(task *model.Task) error
	Update(id int, task *model.Task) error
	Delete(id int) error
	GetByID(id int) (*model.Task, error)
	GetList() ([]model.Task, error)
	GetTaskCategory(id int) ([]model.TaskCategory, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) *taskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Store(task *model.Task) error {
	err := t.db.Create(task).Error
	if err != nil {
		return err
	}

	return nil
}

func (t *taskRepository) Update(id int, task *model.Task) error {
	err := t.db.Model(&model.Task{}).Where("id = ?", task.ID).Updates(model.Task{Title: task.Title, Deadline: task.Deadline, Priority: task.Priority, CategoryID: task.CategoryID, Status: task.Status}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) Delete(id int) error {
	err := t.db.Where("id = ?", id).Delete(&model.Task{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *taskRepository) GetByID(id int) (*model.Task, error) {
	var task model.Task
	err := t.db.First(&task, id).Error
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (t *taskRepository) GetList() ([]model.Task, error) {
	tasks := []model.Task{}

	err := t.db.Find(&tasks).Error
	if err != nil {
		return []model.Task{}, err
	}

	return tasks, nil
}

func (t *taskRepository) GetTaskCategory(id int) ([]model.TaskCategory, error) {
	taskCategory := []model.TaskCategory{}
	err := t.db.Table("tasks").Select("tasks.id as id, tasks.title as title, categories.name as category").Joins("left join categories on tasks.category_id = categories.id").Where("tasks.id = ?", id).First(&taskCategory).Error
	if err != nil {
		return []model.TaskCategory{}, err
	}

	return taskCategory, nil
}
