package repositories

import (
	"context"

	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
)

type ProjectRepo interface {
	GetProjects(ctx context.Context, num int) ([]models.Node, error)
}
