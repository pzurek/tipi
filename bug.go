package tp

import "fmt"

// Bug struct
type Bug struct {
	Entity
	// Priority     *Priority      `json:"Priority,omitempty"`
	// Severity     *Severity      `json:"Severity,omitempty"`
	// CustomFields *[]CustomField `json:"CustomFields,omitempty"`
}

// type Priority struct {
// 	Id   *float64 `json:"Id,omitempty"`
// 	Name *string  `json:"Name,omitempty"`
// }

// type Severity struct {
// 	Id   *float64 `json:"Id,omitempty"`
// 	Name *string  `json:"Name,omitempty"`
// }

// type CustomField struct {
// 	Name  *string `json:"Name,omitempty"`
// 	Value *string `json:"Value,omitempty"`
// }

// BugResponse struct
type BugResponse struct {
	NextURL     *string   `json:"Next,omitempty"`
	PreviousURL *string   `json:"Previous,omitempty"`
	Bugs        *[]Entity `json:"Items"`
}

// BugService struct
type BugService struct {
	client *Client
}

// Get function
func (s *BugService) Get(id string) (*Entity, *Response, error) {

	url := fmt.Sprintf("Bugs/%s", id)

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

// List function
func (s *BugService) List() ([]Entity, error) {

	var resource []Entity

	rp, next, _, err := s.getPage("")

	if err != nil {
		return nil, err
	}

	resource = append(resource, *rp...)

	for next != nil {
		rp, nx, _, err := s.getPage(*next)
		if err != nil {
			return nil, err
		}
		next = nx
		resource = append(resource, *rp...)
	}

	return resource, err
}

func (s *BugService) getPage(url string) (*[]Entity, *string, *Response, error) {

	if url == "" {
		url = "Bugs"
	}

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	response := new(BugResponse)
	resp, err := s.client.Do(req, response)
	if err != nil {
		return nil, nil, resp, err
	}

	next := response.NextURL
	resource := response.Bugs
	return resource, next, resp, err
}
