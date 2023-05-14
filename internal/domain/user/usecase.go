package user

import (
	"time"

	"github.com/morris-zheng/go-slim/internal/domain"
)

type UseCase struct {
	svc  *domain.ServiceContext
	repo *Repo
}

func NewUseCase(svc *domain.ServiceContext) *UseCase {
	return &UseCase{
		svc:  svc,
		repo: NewRepo(svc),
	}
}

type QueryParams struct {
	Limit int `json:"limit" form:"limit"`
	Page  int `json:"page" form:"page"`
	// other condition
}

func (s *UseCase) Query(qp QueryParams) ([]User, int64, error) {
	var ul []User
	var total int64
	db := s.repo.DB.Model(&User{})

	if qp.Limit != 0 {
		if err := db.Count(&total).Error; err != nil {
			return []User{}, 0, err
		}

		db = db.Limit(qp.Limit).Offset((qp.Page - 1) * qp.Limit)
	}

	if err := db.Find(&ul).Error; err != nil {
		return []User{}, 0, err
	}

	return ul, total, nil
}

func (s *UseCase) Get(id int) (User, error) {
	u := User{}
	if err := s.repo.DB.Where("id=?", id).First(&u).Error; err != nil {
		return u, err
	}

	return u, nil
}

func (s *UseCase) Create(u User) error {
	u.CreateTime = time.Now()
	u.UpdateTime = time.Now()

	if err := s.repo.DB.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func (s *UseCase) Update(u User) error {
	u.UpdateTime = time.Now()

	if err := s.repo.DB.Where("id=?", u.Id).Save(&u).Error; err != nil {
		return err
	}

	return nil
}

func (s *UseCase) Delete(id int) error {
	if err := s.repo.DB.Where("id=?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}
