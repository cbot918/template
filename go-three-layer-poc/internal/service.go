package internal

import "github.com/jmoiron/sqlx"

type Service struct {
	R *Repository
}

func NewService(db *sqlx.DB) *Service {
	return &Service{
		R: NewRepository(db),
	}
}

func (s *Service) SignupService(user *SignupRequest) error {
	data := user
	// deal with jwt in the future

	return s.R.SignupRepository(data)
}

func (s *Service) InsertService() error {
	return s.R.InsertRepository()
}
