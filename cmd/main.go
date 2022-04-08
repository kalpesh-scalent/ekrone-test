package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
	"github.com/kalpesh-scalent/ekrone-test/pkg/httpClient"
	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
	"github.com/kalpesh-scalent/ekrone-test/pkg/middleware"
	"github.com/kalpesh-scalent/ekrone-test/pkg/repository/gqlProjectRepo"
	"github.com/kalpesh-scalent/ekrone-test/pkg/restAPI"
	"github.com/kalpesh-scalent/ekrone-test/pkg/service/projectService"
)

func main() {

	// pass array of context keys, which this logger should log
	logger := logger.NewLogger([]string{"requestID"})

	//load necessary environment variables. Set to default values if not found.
	baseURL := GetEnvWithDefault("GRAPH_BASE_URL", "https://gitlab.com/api/graphql")
	graphQLFolder := GetEnvWithDefault("GRAPH_CONFIG_LOCATION", "graphql")

	//load graphQL queries in map
	gqlQueries, err := loadGraphQLQueries(graphQLFolder)
	if err != nil {
		log.Panic(err)
	}

	if len(gqlQueries) == 0 {
		log.Panic("no graphql queries found")
	}

	// fetch instance to custom httpClient
	httpClient := httpClient.GetHttpClient(baseURL, logger)

	//dependency injection in project
	getProjectsRepo := gqlProjectRepo.GetGqlProjectRepo(gqlQueries, httpClient, logger)
	projectService := projectService.GetProjectServiceLayer(logger, getProjectsRepo)

	//initialize API routes
	r := mux.NewRouter()
	apiRouter := restAPI.GetAPIRouter(projectService, logger)
	apiRouter.InitAPIRoutes(r)

	// add necessary middlewares
	r.Use(middleware.ContextMiddleware)

	log.Println("starting the server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic("cannot initialize server with error: " + err.Error())
	}

}

func GetEnvWithDefault(name, defaulValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaulValue
	}
	return val
}

func loadGraphQLQueries(path string) (map[string]string, error) {
	files, err := filepath.Glob(fmt.Sprintf("%s/*.graphql", path))
	if err != nil {
		panic(err)
	}

	gqlQueries := make(map[string]string)

	for _, file := range files {
		fileByte, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err)
		}

		queryName := strings.Replace(filepath.Base(file), filepath.Ext(file), "", 1)
		gqlQueries[queryName] = string(fileByte)
	}
	return gqlQueries, nil
}
