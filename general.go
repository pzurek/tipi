package tp

// General struct
type General struct {
	ID                *float64      `json:"Id,omitempty"`
	Name              *string       `json:"Name,omitempty"`
	Description       *string       `json:"Description,omitempty"`
	StartDate         *string       `json:"StartDate,omitempty"`
	EndDate           *string       `json:"EndDate,omitempty"`
	CreateDate        *string       `json:"CreateDate,omitempty"`
	ModifyDate        *string       `json:"ModifyDate,omitempty"`
	LastCommentDate   *string       `json:"LastCommentDate,omitempty"`
	Tags              *string       `json:"Tags,omitempty"`
	NumericPriority   *float64      `json:"NumericPriority,omitempty"`
	EntityType        *EntityType   `json:"EntityType,omitempty"`
	Owner             *Owner        `json:"Owner,omitempty"`
	LastCommentedUser *User         `json:"LastCommentedUser,omitempty"`
	Project           *Project      `json:"Project,omitempty"`
	CustomFields      []CustomField `json:"CustomFields,omitempty"`
}
