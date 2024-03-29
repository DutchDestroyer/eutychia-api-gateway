package services

import (
	"errors"
	"strconv"

	"github.com/DutchDestroyer/eutychia-api-gateway/database"
	"github.com/DutchDestroyer/eutychia-api-gateway/models"
)

type IGenTestService interface {
	GetAllGenericTests() ([]models.GenericTestOverview, error)
	GetTestsOfProject(string) ([]models.GenericTestOverview, error)
	GetTestData(string, string) (models.GenericTestData, error)
	getGenTestDBService() database.IGenericTestDBService
	getProjectDBService() database.IProjectDBService
	getGenQuestionDBService() database.IGenQuestionDBService
}

type GenTestService struct{}

// GetAllGenericTests gets all generic tests that are in the database
func (g *GenTestService) GetAllGenericTests() ([]models.GenericTestOverview, error) {
	return g.getGenTestDBService().GetAllGenericTests()
}

// GetTestsOfProject gets all the tests of a project
func (g *GenTestService) GetTestsOfProject(projectID string) ([]models.GenericTestOverview, error) {
	projects, errGetProjects := g.getProjectDBService().GetProjects([]string{projectID})

	if errGetProjects != nil {
		return []models.GenericTestOverview{}, errGetProjects
	}

	var testIDs []string

	for i := range projects {
		testIDs = append(testIDs, projects[i].TestIDs...)
	}

	tests, errGetTests := g.getGenTestDBService().GetTestsOfIDs(testIDs)

	if errGetTests != nil {
		return []models.GenericTestOverview{}, errGetTests
	}

	var modTest []models.GenericTestOverview

	for i := range tests {
		modTest = append(modTest, models.GenericTestOverview{ID: tests[i].ID, Name: tests[i].Name, Type: tests[i].Type})
	}

	return modTest, nil
}

func (g *GenTestService) GetTestData(projectID string, testID string) (models.GenericTestData, error) {
	testData, err1 := g.getGenTestDBService().GetTestsOfIDs([]string{testID})

	if err1 != nil {
		return models.GenericTestData{}, err1
	}

	if len(testData) != 1 {
		return models.GenericTestData{}, errors.New("Size of array was unexpected, size is " + strconv.Itoa(len(testData)) + " instead of 1")
	}

	var questionsDAO, err2 = g.getGenQuestionDBService().GetQuestionsPerID(testData[0].QuestionIDs)

	if err2 != nil {
		return models.GenericTestData{}, err2
	}

	if len(questionsDAO) == 0 {
		return models.GenericTestData{}, errors.New("No questions found")
	}

	var questions []models.GenericQuestion

	for i := range questionsDAO {
		questions = append(questions,
			models.GenericQuestion{
				Question:     questionsDAO[i].Question,
				QuestionType: questionsDAO[i].QuestionType,
				Answers:      questionsDAO[i].Answers})
	}

	return models.GenericTestData{
		ID:             testData[0].ID,
		Name:           testData[0].Name,
		Title:          testData[0].Title,
		Type:           testData[0].Type,
		Description:    testData[0].Description,
		DisplayAnswers: testData[0].DisplayAnswers,
		FinalRemark:    testData[0].FinalRemark,
		Questions:      questions}, nil
}

func (g *GenTestService) getGenTestDBService() database.IGenericTestDBService {
	return &database.GenericTestDBService{}
}

func (g *GenTestService) getProjectDBService() database.IProjectDBService {
	return &database.ProjectDBService{}
}

func (g *GenTestService) getGenQuestionDBService() database.IGenQuestionDBService {
	return &database.GenQuestionDBService{}
}
