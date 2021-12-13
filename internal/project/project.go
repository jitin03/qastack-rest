package project

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
)


type Service struct{
	DB *gorm.DB
}

type Users struct {
	User_Id int		`gorm:"primary_key, AUTO_INCREMENT"`
	Username string
	Password string
	Email string
	Role string
	Project []Project `gorm:"ForeignKey:User_Id"`
}

type Project struct{
	Id string
	Name string
	User_Id int
}


type ProjectService interface{
	GetAllProjects(userId int)([]Project,error)
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


func(s *Service)GetAllProjects(userId int)([]Project,error){
	var projects []Project
	fmt.Printf("%s",userId)
	if result :=s.DB.Find(&projects,Project{User_Id: userId});result.Error !=nil{
		return projects, result.Error
	}

	return projects,nil
}


func(s *Service)GetProject(ID uint) (Project,error){
	var project Project
	var users []Users

	result :=s.DB.First(&project,ID)
	s.DB.Model(&project).Related(&users)

	//project.Release = users
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
	log.Info(project.Name)
	log.Info(project.User_Id)
	if result := s.DB.Save(&project); result.Error != nil {
		return Project{}, result.Error
	}else{
		log.Info(result)
	}

	return project, nil
}