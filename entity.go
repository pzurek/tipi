package tp

import "fmt"

// Entity struct
type Entity struct {
	ID               *float64     `json:"Id,omitempty"`
	Name             *string      `json:"Name,omitempty"`
	StartDate        *string      `json:"StartDate,omitempty"`
	EndDate          *string      `json:"EndDate,omitempty"`
	CreateDate       *string      `json:"CreateDate,omitempty"`
	ModifyDate       *string      `json:"ModifyDate,omitempty"`
	LastCommentDate  *string      `json:"LastCommentDate,omitempty"`
	PlannedStartDate *string      `json:"PlannedStartDate,omitempty"`
	PlannedEndDate   *string      `json:"PlannedStartDate,omitempty"`
	EntityType       *EntityType  `json:"EntityType,omitempty"`
	EntityState      *EntityState `json:"EntityState,omitempty"`
	Owner            *Owner       `json:"Owner,omitempty"`
	Project          *Project     `json:"Project,omitempty"`
	Team             *Team        `json:"Team,omitempty"`
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
		return nil, resp, err
	}

	return resource, resp, err
}
