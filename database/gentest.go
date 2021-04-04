package database

//GenericTestDAO data model for the database
type GenericTestDAO struct {
	ID             string
	Name           string
	Type           string
	Description    string
	DisplayAnswers bool
	FinalRemark    string
	Questions      []string
}

var tests []GenericTestDAO = []GenericTestDAO{
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d0", "Test0", "generic", "This is test 0", false,
		"Thanks for completing this test0", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d1", "Test1", "generic", "This is test 1", false,
		"Thanks for completing this test1", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d2", "Test2", "generic", "This is test 2", false,
		"Thanks for completing this test2", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d3", "Test3", "generic", "This is test 3", false,
		"Thanks for completing this test3", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d4", "Test4", "generic", "This is test 4", false,
		"Thanks for completing this test4", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d5", "Test5", "generic", "This is test 5", false,
		"Thanks for completing this test5", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d6", "Test6", "generic", "This is test 6", false,
		"Thanks for completing this test6", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d7", "Test7", "generic", "This is test 7", false,
		"Thanks for completing this test7", []string{},
	},
	{
		"25553260-2ae4-465c-8a64-6a5c3ab355d8", "Test8", "generic", "This is test 8", false,
		"Thanks for completing this test8", []string{},
	},
}

// GetTestsOfIDs get the tests with specific IDs
func GetTestsOfIDs(testIDs []string) ([]GenericTestDAO, error) {
	var testsToReturn []GenericTestDAO

	for i := range tests {
		for j := range testIDs {
			if tests[i].ID == testIDs[j] {
				testsToReturn = append(testsToReturn, tests[i])
			}
		}
	}

	return testsToReturn, nil
}
