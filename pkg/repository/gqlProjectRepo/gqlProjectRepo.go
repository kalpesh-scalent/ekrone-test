package gqlProjectRepo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kalpesh-scalent/ekrone-test/pkg/httpClient"
	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
	"github.com/kalpesh-scalent/ekrone-test/pkg/models"
	"github.com/sirupsen/logrus"
)

type GqlProjectRepo struct {
	httpClient *httpClient.HttpClient
	logger     *logger.Logger
	gqlQueries map[string]string
}

func GetGqlProjectRepo(gqlQueries map[string]string, httpClient *httpClient.HttpClient, logger *logger.Logger) *GqlProjectRepo {

	return &GqlProjectRepo{
		httpClient: httpClient,
		logger:     logger,
		gqlQueries: gqlQueries,
	}
}

func (g GqlProjectRepo) postGqlQuery(ctx context.Context, queryPath string, query string, response interface{}) (err error) {

	graphRequest := models.GraphRequest{
		Query: query,
	}

	var graphResp models.GraphResponse

	g.logger.WithContext(ctx).WithFields(
		logrus.Fields{
			"queryPath": queryPath,
			"query":     graphRequest,
		},
	).Info("GraphQL Request - Begin")

	statusCode, err := g.httpClient.Post(ctx, queryPath, graphRequest, &graphResp)
	if err != nil {
		g.logger.WithContext(ctx).Error(err)
		return err
	}

	g.logger.WithContext(ctx).WithFields(
		logrus.Fields{
			"queryPath":  queryPath,
			"query":      graphRequest,
			"statusCode": statusCode,
			"response":   string(graphResp.Data),
		},
	).Info("GraphQL Request - Begin")

	if statusCode != http.StatusOK {
		g.logger.WithContext(ctx).Error(errWrongResponseCode, fmt.Sprintf("expected: %v, received: %v", http.StatusOK, statusCode))
		return errWrongResponseCode
	}

	if response == nil {
		return nil
	}

	if len(graphResp.Errors) != 0 {
		g.logger.WithContext(ctx).Error(graphResp.Errors)
		return errResponse
	}

	err = json.Unmarshal(graphResp.Data, response)
	if err != nil {
		g.logger.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
