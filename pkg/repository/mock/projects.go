package mockProjectRepo

import (
	"context"
	"errors"

	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
)

func (MockProjectRepo) GetProjects(ctx context.Context, num int) (n []models.Node, err error) {

	if len(Nodes) < num {
		return n, errors.New("data not found")
	}

	for i := 0; i < num; i++ {
		n = append(n, Nodes[i])
	}

	return n, nil
}

var Nodes = []models.Node{
	{
		Name:       "abc1",
		ForksCount: 1,
	},
	{
		Name:       "abc2",
		ForksCount: 2,
	},
	{
		Name:       "abc3",
		ForksCount: 3,
	},
	{
		Name:       "abc4",
		ForksCount: 4,
	},
	{
		Name:       "abc5",
		ForksCount: 5,
	},
	{
		Name:       "abc6",
		ForksCount: 6,
	},
	{
		Name:       "abc7",
		ForksCount: 7,
	},
	{
		Name:       "abc8",
		ForksCount: 8,
	},
	{
		Name:       "abc9",
		ForksCount: 9,
	},
	{
		Name:       "abc10",
		ForksCount: 10,
	},
}
