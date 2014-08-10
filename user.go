package tp

import "fmt"

// User struct
type User struct {
	Kind                      *string  `json:"Kind,omitempty"`
	ID                        *float64 `json:"Id,omitempty"`
	FirstName                 *string  `json:"FirstName,omitempty"`
	LastName                  *string  `json:"LastName,omitempty"`
	Email                     *string  `json:"Email,omitempty"`
	Login                     *string  `json:"Login,omitempty"`
	CreateDate                *string  `json:"CreateDate,omitempty"`
	ModifyDate                *string  `json:"ModifyDate,omitempty"`
	DeleteDate                *string  `json:"DeleteDate,omitempty"`
	IsActive                  *bool    `json:"IsActive,omitempty"`
	IsAdministrator           *bool    `json:"IsAdministrator,omitempty"`
	WeeklyAvailableHours      *float64 `json:"WeeklyAvailableHours,omitempty"`
	CurrentAllocation         *float64 `json:"CurrentAllocation,omitempty"`
	CurrentAvailableHours     *float64 `json:"CurrentAvailableHours,omitempty"`
	AvailableFrom             *string  `json:"AvailableFrom,omitempty"`
	AvailableFutureAllocation *float64 `json:"AvailableFutureAllocation,omitempty"`
	AvailableFutureHours      *float64 `json:"AvailableFutureHours,omitempty"`
	IsObserver                *bool    `json:"IsObserver,omitempty"`
	Role                      *Role    `json:"Role,omitempty"`
}

// UserService struct
type UserService struct {
	client *Client
}

type userListResponse struct {
	Items []User `json:"Items,omitempty"`
}

// Get function returns a single user
func (s *UserService) Get(id string) (*User, *Response, error) {

	url := fmt.Sprintf("Users/%s)", id)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	resource := new(User)
	resp, err := s.client.Do(req, &resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, err
}

// GetByEmail function returns the first user with given email
func (s *UserService) GetByEmail(email string) (*User, *Response, error) {

	url := fmt.Sprintf("Users?where=(Email+eq+'%s')", email)

	req, err := s.client.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	list := new(userListResponse)
	resp, err := s.client.Do(req, &list)
	if err != nil {
		return nil, resp, err
	}

	var resource *User
	if len(list.Items) > 0 {
		resource = &list.Items[0]
	}

	return resource, resp, err
}
