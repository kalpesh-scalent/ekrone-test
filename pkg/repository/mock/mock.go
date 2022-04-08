package mockProjectRepo

type MockProjectRepo struct{}

func GetMockProjectRepo() *MockProjectRepo {
	return &MockProjectRepo{}
}
