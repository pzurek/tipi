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

// EntityType struct
type EntityType struct {
	ID   *float64 `json:"Id"`
	Name *string  `json:"Name"`
}

// EntityState struct
type EntityState struct {
	ID   *float64 `json:"Id"`
	Name *string  `json:"Name"`
}

// Owner struct
type Owner struct {
	ID        *float64 `json:"Id"`
	FirstName *string  `json:"FirstName"`
	LastName  *string  `json:"LastName"`
}

// Project struct
type Project struct {
	ID   *float64 `json:"Id"`
	Name *string  `json:"Name"`
}

// Team struct
type Team struct {
	ID   *float64 `json:"Id"`
	Name *string  `json:"Name"`
}

// type EntityResponse struct {
// 	NextUrl     *string   `json:"Next,omitempty"`
// 	PreviousUrl *string   `json:"Previous,omitempty"`
// 	Bugs        *[]Entity `json:"Items"`
// }

// EntityService struct
type EntityService struct {
	client *Client
}

// Get func
func (s *EntityService) Get(t, id string) (*Entity, *Response, error) {

	url := fmt.Sprintf("%s/%s?format=json", t, id)

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
