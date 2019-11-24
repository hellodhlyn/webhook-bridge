package bridge

import (
	"errors"
	"testing"
)

const dockerHubInput = `
{
  "callback_url": "https://registry.hub.docker.com/u/svendowideit/testhook/hook/2141b5bi5i5b02bec211i4eeih0242eg11000a/",
  "push_data": {
    "images": [
      "27d47432a69bca5f2700e4dff7de0388ed65f9d3fb1ec645e2bc24c223dc1cc3",
      "51a9c7c1f8bb2fa19bcd09789a34e63f35abb80044bc10196e304f6634cc582c",
      "..."
    ],
    "pushed_at": 1.417566161e+09,
    "pusher": "trustedbuilder",
    "tag": "latest"
  },
  "repository": {
    "comment_count": 0,
    "date_created": 1.417494799e+09,
    "description": "",
    "dockerfile": "FROM scratch\nRUN sleep 10",
    "full_description": "Docker Hub based automated build from a GitHub repo",
    "is_official": false,
    "is_private": true,
    "is_trusted": true,
    "name": "testhook",
    "namespace": "svendowideit",
    "owner": "svendowideit",
    "repo_name": "svendowideit/testhook",
    "repo_url": "https://registry.hub.docker.com/u/svendowideit/testhook/",
    "star_count": 0,
    "status": "Active"
  }
}`

const template = `{"text": "Docker image '{{ .repository.repo_name }}' has built successfully."}`

func TestJSONConverter(t *testing.T) {
	converter := NewJSONConverter([]byte(template))
	output, err := converter.Convert(&Input{Payload: []byte(dockerHubInput)})
	if err != nil {
		t.Error(err)
	}
	if string(output.ConvertedPayload) != `{"text": "Docker image 'svendowideit/testhook' has built successfully."}` {
		t.Error(errors.New("invalid converted payload"))
	}
}
