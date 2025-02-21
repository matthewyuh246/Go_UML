package main

import "fmt"

type User struct{}

type IRepository interface {
	create(user User) User
	findById(id string) User
}

type Repository struct{}

func NewRepository() IRepository {
	return &Repository{}
}

func (r *Repository) create(user User) User {
	fmt.Println("RDBにUserを登録")
	return user
}

func (r *Repository) findById(id string) User {
	fmt.Printf("ID: %sのユーザーを検索", id)
	user := User{}
	return user
}

type IService interface {
	create(user User) User
	findById(id string) User
}

type Service struct {
	repository IRepository
}

func NewService(repo IRepository) IService {
	return &Service{
		repo,
	}
}

func (s *Service) create(user User) User {
	return s.repository.create(user)
}

func (s *Service) findById(id string) User {
	return s.repository.findById(id)
}

type IController interface {
	create(user User) User
	findById(id string) User
}

type Controller struct {
	service IService
}

func NewController(sv IService) IController {
	return &Service{
		sv,
	}
}

func (c *Controller) create(user User) User {
	return c.service.create(user)
}

func (c *Controller) findById(id string) User {
	return c.service.findById(id)
}

func main() {
	repository := NewRepository()
	service := NewService(repository)
	controller := NewController(service)
	controller.findById("123")
}
