package tp

import "fmt"

// UserStory struct
type UserStory struct {
	Entity
}

// UserStoryResponse struct
type UserStoryResponse struct {
	NextURL     *string  `json:"Next,omitempty"`
	PreviousURL *string  `json:"Previous,omitempty"`
	UserStories []Entity `json:"Items,omitempty"`
}

// UserStoryService struct
type UserStoryService struct {
	client *Client
}

// Get func
func (s *UserStoryService) Get(id string) (*Entity, *Response, error) {
	url := fmt.Sprintf("UserStories/%s?format-json", id)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resource := new(Entity)
	resp, err := s.client.Do(req, &resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, err
}
