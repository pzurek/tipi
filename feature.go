package tp

import "fmt"

// Feature struct
type Feature struct {
	Entity
}

// FeatureResponse struct
type FeatureResponse struct {
	NextURL     *string   `json:"Next,omitempty"`
	PreviousURL *string   `json:"Previous,omitempty"`
	Features    *[]Entity `json:"Items"`
}

// FeatureService struct
type FeatureService struct {
	client *Client
}

// Get funct
func (s *FeatureService) Get(id string) (*Entity, *Response, error) {
	url := fmt.Sprintf("Features/%s?format-json", id)
	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resource := new(Entity)
	resp, err := s.client.Do(req, &resource)
	if err != nil {
		return nil, nil, err
	}

	return resource, resp, err
}
