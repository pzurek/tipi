package tp

// Assignable struct
type Assignable struct {
	ID                *float64      `json:"Id,omitempty"`
	Name              *string       `json:"Name,omitempty"`
	Description       *string       `json:"Description,omitempty"`
	StartDate         *string       `json:"StartDate,omitempty"`
	EndDate           *string       `json:"EndDate,omitempty"`
	CreateDate        *string       `json:"CreateDate,omitempty"`
	ModifyDate        *string       `json:"ModifyDate,omitempty"`
	LastCommentDate   *string       `json:"LastCommentDate,omitempty"`
	PlannedStartDate  *string       `json:"PlannedStartDate,omitempty"`
	PlannedEndDate    *string       `json:"PlannedEndDate,omitempty"`
	Tags              *string       `json:"Tags,omitempty"`
	NumericPriority   *float64      `json:"NumericPriority,omitempty"`
	Effort            *float64      `json:"Effort,omitempty"`
	EffortCompleted   *float64      `json:"EffortCompleted,omitempty"`
	EffortToDo        *float64      `json:"EffortToDo,omitempty"`
	Progress          *float64      `json:"Progress,omitempty"`
	TimeSpent         *float64      `json:"TimeSpent,omitempty"`
	TimeRemain        *float64      `json:"TimeRemain,omitempty"`
	EntityType        *EntityType   `json:"EntityType,omitempty"`
	Owner             *Owner        `json:"Owner,omitempty"`
	LastCommentedUser *User         `json:"LastCommentedUser,omitempty"`
	Project           *Project      `json:"Project,omitempty"`
	Release           *Release      `json:"Release,omitempty"`
	Iteration         *Iteration    `json:"Iteration,omitempty"`
	TeamIteration     *Iteration    `json:"TeamIteration,omitempty"`
	Team              *Team         `json:"Team,omitempty"`
	Priority          *Priority     `json:"Priority,omitempty"`
	EntityState       *EntityState  `json:"EntityState,omitempty"`
	CustomFields      []CustomField `json:"CustomFields,omitempty"`
}
