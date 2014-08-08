package tp

// Comment struct
type Comment struct {
	ID          *float64 `json:"Id,omitempty"`
	Description *string  `json:"Description,omitempty"`
	ParentID    *float64 `json:"ParentId,omitempty"`
	CreateDate  *string  `json:"CreateDate,omitempty"`
	General     *General `json:"General,omitempty"`
	Owner       *Owner   `json:"Owner,omitempty"`
}

// CommentService struct
type CommentService struct {
	client *Client
}

// Create func
func (s *CommentService) Create(comment *Comment) (*Comment, *Response, error) {
	url := "Comments"

	req, err := s.client.NewRequest("POST", url, comment)
	if err != nil {
		return nil, nil, err
	}

	resource := new(Comment)
	resp, err := s.client.Do(req, &resource)
	if err != nil {
		return nil, resp, err
	}

	return resource, resp, err
}
