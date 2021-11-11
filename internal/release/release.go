package release

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
)


type Service struct{
	DB *gorm.DB
}

type Release struct{
	gorm.Model
	ReleaseName string
	ProjectID int
	StartDate time.Time
	EndDate time.Time
}

type ReleaseService interface{
	GetAllRelease([]Release,error)
	AddRelease(release Release)(Release,error)
	UpdateRelease(ID uint,newRelease Release)(Release,error)
	DeleteRelease(ID uint)(error)
	GetRelease(ID uint)(Release,error)
}

func NewService(db *gorm.DB) *Service{
	return &Service{
		DB:db,
	}
}


func(s *Service)GetRelease(ID uint) (Release,error){
	var release Release


	result :=s.DB.First(&release,ID)


	
	if result.Error !=nil{
		fmt.Println(ID)
		return Release{},result.Error
	}
	
	return release,nil
}


func(s *Service)UpdateRelease(ID uint,newRelease Release)(Release,error){
	release, err := s.GetRelease(ID)
	if err != nil {
		return Release{}, err
	}

	if result := s.DB.Model(&release).Updates(newRelease); result.Error != nil {
		return Release{}, result.Error
	}

	return release, nil


}

func(s *Service)AddRelease(release Release)(Release,error){
	if result :=s.DB.Save(&release);result.Error !=nil{
		return Release{},result.Error
	}
	return release,nil
}

func(s *Service)GetAllRelease()([]Release,error){
	var release []Release

	if result :=s.DB.Find(&release);result.Error !=nil{
		return release,result.Error
	}
	return release,nil
}