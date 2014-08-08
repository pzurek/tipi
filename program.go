package tp

// Program struct
type Program struct {
	ID                *float64      `json:"Id"`
	Name              *string       `json:"Name"`
	Description       *string       `json:"Description"`
	StartDate         *string       `json:"StartDate"`
	EndDate           *string       `json:"EndDate"`
	CreateDate        *string       `json:"CreateDate"`
	ModifyDate        *string       `json:"ModifyDate"`
	LastCommentDate   *string       `json:"LastCommentDate"`
	Tags              *string       `json:"Tags"`
	NumericPriority   *float64      `json:"NumericPriority"`
	EntityType        *EntityType   `json:"EntityType"`
	Owner             *Owner        `json:"Owner"`
	LastCommentedUser *User         `json:"LastCommentedUser"`
	Project           *Project      `json:"Project"`
	CustomFields      []CustomField `json:"CustomFields"`
}
