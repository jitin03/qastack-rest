package project

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/jitin07/qastack/internal/release"
)


type Service struct{
	DB *gorm.DB
}


type Project struct{
	gorm.Model
	Name string
	Release []release.Release	`gorm:"-"`
}


type ProjectService interface{
	GetAllProjects([]Project,error)
	AddProject(project Project) (Project,error)
	UpdateProject(ID uint,newProject Project)(Project,error)
	DeleteProject(ID uint)(error)
	GetProject(ID uint)(Project,error)

}

// NewService - returns a new comments service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}


func(s *Service)GetAllProjects()([]Project,error){
	var projects []Project

	if result :=s.DB.Find(&projects);result.Error !=nil{
		return projects, result.Error
	}

	return projects,nil
}


func(s *Service)GetProject(ID uint) (Project,error){
	var project Project
	var releases []release.Release

	result :=s.DB.First(&project,ID)
	s.DB.Model(&project).Related(&releases)

	project.Release = releases
	if result.Error !=nil{
		fmt.Println(ID)
		return Project{},result.Error
	}
	
	return project,nil
}

func(s *Service)DeleteProject(ID uint)(error){
	if result := s.DB.Delete(&Project{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func(s *Service) UpdateProject(ID uint,newProject Project)(Project,error){

	project, err := s.GetProject(ID)
	if err != nil {
		return Project{}, err
	}

	if result := s.DB.Model(&project).Updates(newProject); result.Error != nil {
		return Project{}, result.Error
	}

	return project, nil
}


func(s *Service)AddProject(project Project) (Project,error){
	fmt.Println(project)
	if result := s.DB.Save(&project); result.Error != nil {
		return Project{}, result.Error
	}
	return project, nil
}