package github

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "test",
		Description: "test",
		Homepage:    "test",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	// fmt.Println(string(bytes))

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, &target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
	assert.EqualValues(t, target.Description, request.Description)

	// assert.EqualValues(t, `{"name":"test","description":"test","homepage":"test","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`, string(bytes))

	// json, err := request.AsJson()
	// if err != nil {
	// 	t.Error(err)
	// }

	// expected := `{"name":"test","description":"test","homepage":"test","private":true,"has_issues":true,"has_projects":true,"has_wiki":true}`
	// if json != expected {
	// 	t.Errorf("Expected %s, got %s", expected, json)
	// }
}
