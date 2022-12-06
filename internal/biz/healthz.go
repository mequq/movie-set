package biz

import (
	"app/config"

	"github.com/rs/zerolog"
)

// HealthzRepo is the repository interface for healthz.
type HealthzRepo interface {
	// Readyness checks the readyness of the service.
	Readiness() error
	// Liveness checks the liveness of the service.
	Liveness() error
}

// HealthzUsecase is usecase
type HealthzUsecase struct {
	repo HealthzRepo
	conf *config.Config
}

// NewHealthzUsecase creates a new healthz usecase.
func NewHealthzUsecase(repo HealthzRepo, conf *config.Config, logger zerolog.Logger) *HealthzUsecase {
	return &HealthzUsecase{
		repo: repo,
		conf: conf,
	}
}

// Readyness checks the readyness of the service.
func (u *HealthzUsecase) Readiness() error {
	return u.repo.Readiness()
}

// Liveness checks the liveness of the service.
func (u *HealthzUsecase) Liveness() error {
	return u.repo.Liveness()
}
