package tp

import "fmt"

// Task struct
type Task struct {
	Entity
}

// TaskResponse struct
type TaskResponse struct {
	NextURL     *string  `json:"Next,omitempty"`
	PreviousURL *string  `json:"Previous,omitempty"`
	Tasks       []Entity `json:"Items,omitempty"`
}

// TaskService struct
type TaskService struct {
	client *Client
}

// Get func
func (s *TaskService) Get(id string) (*Entity, *Response, error) {
	url := fmt.Sprintf("Tasks/%s?format-json", id)
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
