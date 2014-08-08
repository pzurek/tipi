package tp

import "fmt"

// Entity struct
type Entity struct {
	ID               *float64     `json:"Id,omitempty"`
	Name             *string      `json:"Name"`
	StartDate        *string      `json:"StartDate"`
	EndDate          *string      `json:"EndDate"`
	CreateDate       *string      `json:"CreateDate"`
	ModifyDate       *string      `json:"ModifyDate"`
	LastCommentDate  *string      `json:"LastCommentDate"`
	PlannedStartDate *string      `json:"PlannedStartDate"`
	PlannedEndDate   *string      `json:"PlannedStartDate"`
	EntityType       *EntityType  `json:"EntityType"`
	EntityState      *EntityState `json:"EntityState"`
	Owner            *Owner       `json:"Owner"`
	Project          *Project     `json:"Project"`
	Team             *Team        `json:"Team"`
}

// EntityService struct
type EntityService struct {
	client *Client
}

// Get func
func (s *EntityService) Get(t, id string) (*Entity, *Response, error) {

	url := fmt.Sprintf("%s/%s", t, id)

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
