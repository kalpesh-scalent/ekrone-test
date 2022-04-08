package projectService

import (
	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
	repositories "github.com/kalpesh-scalent/ekrone-test/pkg/repository"
)

type ProjectService struct {
	logger      *logger.Logger
	projectRepo repositories.ProjectRepo
}

func GetProjectServiceLayer(l *logger.Logger, p repositories.ProjectRepo) *ProjectService {
	return &ProjectService{
		logger:      l,
		projectRepo: p,
	}
}
