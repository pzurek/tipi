package tp

// Release struct
type Release struct {
	ID                *float64      `json:"Id"`
	Name              *string       `json:"Name"`
	Description       *string       `json:"Description"`
	StartDate         *string       `json:"StartDate"`
	EndDate           *string       `json:"EndDate"`
	CreateDate        *string       `json:"CreateDate"`
	ModifyDate        *string       `json:"ModifyDate"`
	Tags              *string       `json:"Tags"`
	NumericPriority   *float64      `json:"NumericPriority"`
	IsCurrent         *bool         `json:"IsCurrent"`
	Progress          *float64      `json:"Progress"`
	EntityType        *EntityType   `json:"EntityType"`
	Owner             *Owner        `json:"Owner"`
	LastCommentedUser *User         `json:"LastCommentedUser"`
	Project           *Project      `json:"Project"`
	CustomFields      []CustomField `json:"CustomFields"`
}
