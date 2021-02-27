package database

// ProjectDAO contains all informations of the different projects
type ProjectDAO struct {
	ID           string
	Name         string
	researchers  []string
	participants []string
	tests        []string
}

var projects []ProjectDAO

//AddNewProject adds a new project to the database
func AddNewProject(projectName string, tests []string, researchers []string, participants []string) {

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
