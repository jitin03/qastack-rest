package release

import (
	"fmt"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	DB *gorm.DB
}

type Release struct {
	// gorm.Model
	Id          string
	ReleaseName string
	ProjectID   string
	StartDate   string
	EndDate     string
}

type ReleaseService interface {
	GetAllRelease(projectId string) ([]Release, error)
	AddRelease(release Release) (Release, error)
	UpdateRelease(ID string, newRelease Release) (Release, error)
	DeleteRelease(ID string) error
	GetRelease(ID string) (Release, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) DeleteRelease(ID string) error {
	if result := s.DB.Delete(&Release{}, Release{Id: ID}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s *Service) GetRelease(ID string) (Release, error) {
	var release Release

	result := s.DB.First(&release, Release{Id: ID})

	if result.Error != nil {
		fmt.Println(ID)
		return Release{}, result.Error
	}

	return release, nil
}

func (s *Service) UpdateRelease(ID string, newRelease Release) (Release, error) {
	log.Info(ID)
	release, err := s.GetRelease(ID)
	if err != nil {
		return Release{}, err
	}

	if result := s.DB.Model(&release).Updates(newRelease); result.Error != nil {
		return Release{}, result.Error
	}

	return release, nil

}

func (s *Service) AddRelease(release Release) (Release, error) {
	log.Info(release.ReleaseName)
	log.Info(release.ProjectID)
	if result := s.DB.Save(&release); result.Error != nil {
		return Release{}, result.Error
	}
	return release, nil
}

func (s *Service) GetAllRelease(projectId string) ([]Release, error) {
	var release []Release

	if result := s.DB.Where("project_id = ?", projectId).Find(&release); result.Error != nil {
		return release, result.Error
	}
	return release, nil
}
