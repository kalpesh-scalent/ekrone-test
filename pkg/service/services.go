package service

import (
	"context"

	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
)

type ProjectService interface {
	GetProjectNamesWithForkCount(context.Context, int) (models.ProjectsWithRepoCount, error)
}
