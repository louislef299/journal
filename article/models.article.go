package journal

import "errors"

type entry struct {
	Date    int    `json:"Date"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var entryList = []entry{
	{Date: 1, Title: "finish codebase", Content: "transform the codebase into a journal"},
	{Date: 2, Title: "add ci/cd capabilities", Content: "add gitactions to both test the codebase and deploy it to docker"},
	{Date: 3, Title: "add persistent storage", Content: "give the todo application persistent storage(have to decDatee which type of storage)"},
	{Date: 4, Title: "deploy on minikube", Content: "deploy the todo app on minikube"},
	{Date: 5, Title: "deploy on eks", Content: "deploy your todo app on eks"},
	{Date: 6, Title: "automate minikube to eks deployment", Content: "automate the deployment from making local changes, testing those local changes, using github actions to test and deploy the new image to the eks cluster"},
}

func getAllEntries() []entry {
	return entryList
}

func getEntryByDate(Date int) (*entry, error) {
	for _, a := range entryList {
		if a.Date == Date {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
