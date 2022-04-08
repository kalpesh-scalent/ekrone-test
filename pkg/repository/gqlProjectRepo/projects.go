package gqlProjectRepo

import (
	"context"
	"fmt"

	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
)

func (g GqlProjectRepo) GetProjects(ctx context.Context, num int) ([]models.Node, error) {

	queryGetProjects, found := g.gqlQueries["GetProjects"]
	if !found {
		g.logger.WithContext(ctx).Error(errQueryNotFound)
		return nil, errQueryNotFound
	}

	queryString := fmt.Sprintf(queryGetProjects, num)

	var responsesData models.GetProjects
	err := g.postGqlQuery(ctx, "", queryString, &responsesData)
	if err != nil {
		return nil, err
	}

	return responsesData.Projects.Nodes, nil
}
