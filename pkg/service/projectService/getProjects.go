package projectService

import (
	"context"
	"errors"

	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
)

func (p ProjectService) GetProjectNamesWithForkCount(ctx context.Context, num int) (resp models.ProjectsWithRepoCount, err error) {

	nodes, err := p.projectRepo.GetProjects(ctx, num)
	if err != nil {
		return resp, err
	}

	if len(nodes) == 0 {
		return resp, errors.New("empty list of projects returned")
	}

	resp.Names = nodes[0].Name
	resp.Count = nodes[0].ForksCount

	for i := 1; i < len(nodes); i++ {
		resp.Names = resp.Names + ", " + nodes[i].Name
		resp.Count += nodes[i].ForksCount
	}

	return
}
