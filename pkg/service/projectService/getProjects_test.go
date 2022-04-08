package projectService

import (
	"context"
	"strings"
	"testing"

	"github.com/kalpesh-scalent/ekrone-test/pkg/logger"
	mockProjectRepo "github.com/kalpesh-scalent/ekrone-test/pkg/repository/mock"
)

func TestGetProject(t *testing.T) {

	mockRepo := mockProjectRepo.GetMockProjectRepo()

	logger := logger.NewLogger([]string{})

	svc := GetProjectServiceLayer(logger, mockRepo)

	ctx := context.Background()

	testTable := []struct {
		testName         string
		num              int
		errorExpected    bool
		expectedCount    int
		expectedProjects string
	}{
		{
			testName:         "positive test 1",
			num:              5,
			errorExpected:    false,
			expectedCount:    countFirstNForks(5),
			expectedProjects: getFirstNprojects(5),
		},
		{
			testName:         "out of bound",
			num:              15,
			errorExpected:    true,
			expectedCount:    countFirstNForks(15),
			expectedProjects: getFirstNprojects(15),
		},
		{
			testName:         "positive test 2",
			num:              7,
			errorExpected:    false,
			expectedCount:    countFirstNForks(7),
			expectedProjects: getFirstNprojects(7),
		},
		{
			testName:         "zero elements",
			num:              0,
			errorExpected:    true,
			expectedCount:    countFirstNForks(0),
			expectedProjects: getFirstNprojects(0),
		},
		{
			testName:         "negative number",
			num:              -3,
			errorExpected:    true,
			expectedCount:    countFirstNForks(-3),
			expectedProjects: getFirstNprojects(-3),
		},
	}

	for _, test := range testTable {

		projects, err := svc.GetProjectNamesWithForkCount(ctx, test.num)
		if err != nil && !test.errorExpected {
			t.Errorf("%s : expected data, got error", test.testName)
			return
		}
		if err == nil && test.errorExpected {
			t.Errorf("%s : expected error, didn't get", test.testName)
			return
		}

		if test.expectedCount != projects.Count {
			t.Errorf("%s : expected: %v, got: %v", test.testName, test.expectedCount, projects.Count)
			return
		}

		if test.expectedProjects != projects.Names {
			t.Errorf("%s : expected: %v, got %v", test.testName, test.expectedProjects, projects.Names)
			return
		}
	}
}

func countFirstNForks(num int) int {

	if num > len(mockProjectRepo.Nodes) {
		return 0
	}

	count := 0
	for i := 0; i < num; i++ {
		count += mockProjectRepo.Nodes[i].ForksCount
	}

	return count
}

func getFirstNprojects(num int) string {

	if num > len(mockProjectRepo.Nodes) {
		return ""
	}

	stringSlice := []string{}
	for i := 0; i < num; i++ {
		stringSlice = append(stringSlice, mockProjectRepo.Nodes[i].Name)
	}

	return strings.Join(stringSlice, ", ")
}
