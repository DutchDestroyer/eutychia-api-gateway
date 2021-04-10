package database

import "github.com/google/uuid"

// ProjectDAO contains all informations of the different projects
type ProjectDAO struct {
	ID           string
	Name         string
	Researchers  []string // UUIDS
	Participants []string // UUIDS
	TestIDs      []string //UUIDS
}

var projects []ProjectDAO = []ProjectDAO{
	{"497aeeaf-0d41-46c4-a5a1-8a88c7b61807", "Project1", []string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72"},
		[]string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "683c5de1-5172-4a94-bd3b-2d4bf58b6b73"},
		[]string{"25553260-2ae4-465c-8a64-6a5c3ab355d0", "25553260-2ae4-465c-8a64-6a5c3ab355d1", "25553260-2ae4-465c-8a64-6a5c3ab355d2"}},
	{"497aeeaf-0d41-46c4-a5a1-8a88c7b61808", "Project2", []string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72"},
		[]string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72", "683c5de1-5172-4a94-bd3b-2d4bf58b6b73"},
		[]string{"25553260-2ae4-465c-8a64-6a5c3ab355d3", "25553260-2ae4-465c-8a64-6a5c3ab355d4", "25553260-2ae4-465c-8a64-6a5c3ab355d5"}},
	{"497aeeaf-0d41-46c4-a5a1-8a88c7b61809", "Project3", []string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72"},
		[]string{"683c5de1-5172-4a94-bd3b-2d4bf58b6b72"},
		[]string{"25553260-2ae4-465c-8a64-6a5c3ab355d6", "25553260-2ae4-465c-8a64-6a5c3ab355d7", "25553260-2ae4-465c-8a64-6a5c3ab355d8"}},
}

//AddNewProject adds a new project to the database
func AddNewProject(projectName string, tests []string, researchers []string, participants []string) error {

	projectID := uuid.New()

	projects = append(projects, ProjectDAO{
		projectID.String(), projectName, researchers, participants, tests,
	})

	return nil
}

//GetProjects gets all the projects with the specific IDs
func GetProjects(projectIDs []string) ([]ProjectDAO, error) {
	var projectsToInclude []ProjectDAO = []ProjectDAO{}

	for i := range projects {
		for j := range projectIDs {
			if projects[i].ID == projectIDs[j] {
				projectsToInclude = append(projectsToInclude, projects[i])
			}
		}
	}

	return projectsToInclude, nil
}
