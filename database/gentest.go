package database

import "github.com/DutchDestroyer/eutychia-api-gateway/models"

//GenericTestDAO data model for the database
type GenericTestDAO struct {
	ID             string
	Name           string
	Type           string
	Title          string
	Description    string
	DisplayAnswers bool
	FinalRemark    string
	QuestionIDs    []string
}

var genericTests []GenericTestDAO = []GenericTestDAO{
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d0", "Test0", "generic", "Test 0", "This is test 0", false,
		"Thanks for completing this test0", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d1", "Test1", "generic", "Test1", "This is test 1", false,
		"Thanks for completing this test1", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d2", "Test2", "generic", "Test2", "This is test 2", false,
		"Thanks for completing this test2", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d3", "Test3", "generic", "Test3", "This is test 3", false,
		"Thanks for completing this test3", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d4", "Test4", "generic", "Test4", "This is test 4", false,
		"Thanks for completing this test4", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d5", "Test5", "generic", "Test5", "This is test 5", false,
		"Thanks for completing this test5", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d6", "Test6", "generic", "Test6", "This is test 6", false,
		"Thanks for completing this test6", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d7", "Test7", "generic", "Test7", "This is test 7", false,
		"Thanks for completing this test7", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d8", "Test8", "generic", "Test8", "This is test 8", false,
		"Thanks for completing this test8", []string{"dc919c8d-93fa-4f69-8086-e2395b9d01e0", "dc919c8d-93fa-4f69-8086-e2395b9d01e1", "dc919c8d-93fa-4f69-8086-e2395b9d01e2"},
	},
}

type IGenericTestDBService interface {
	GetTestsOfIDs([]string) ([]GenericTestDAO, error)
	GetAllGenericTests() ([]models.GenericTestOverview, error)
}

type GenericTestDBService struct{}

// GetTestsOfIDs get the tests with specific IDs
func (g *GenericTestDBService) GetTestsOfIDs(testIDs []string) ([]GenericTestDAO, error) {
	var testsToReturn []GenericTestDAO

	for i := range genericTests {
		for j := range testIDs {
			if genericTests[i].ID == testIDs[j] {
				testsToReturn = append(testsToReturn, genericTests[i])
			}
		}
	}

	return testsToReturn, nil
}

// GetAllGenericTests gets all generic tests that are in the database
func (g *GenericTestDBService) GetAllGenericTests() ([]models.GenericTestOverview, error) {
	var tests []models.GenericTestOverview

	for i := range genericTests {
		tests = append(tests, models.GenericTestOverview{
			ID:   genericTests[i].ID,
			Name: genericTests[i].Name,
			Type: genericTests[i].Type})
	}

	return tests, nil
}
