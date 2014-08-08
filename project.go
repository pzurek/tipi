package tp

// Project struct
type Project struct {
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
	IsActive          *bool         `json:"IsActive,omitempty"`
	IsProduct         *bool         `json:"IsProduct,omitempty"`
	Abbreviation      *string       `json:"Abbreviation,omitempty"`
	MailReplyAddress  *string       `json:"MailReplyAddress,omitempty"`
	Color             *string       `json:"Color,omitempty"`
	Progress          *float64      `json:"Progress,omitempty"`
	EntityType        *EntityType   `json:"EntityType,omitempty"`
	Owner             *Owner        `json:"Owner,omitempty"`
	LastCommentedUser *User         `json:"LastCommentedUser,omitempty"`
	Program           *Program      `json:"Program,omitempty"`
	Process           *Process      `json:"Process,omitempty"`
	Company           *Company      `json:"Company,omitempty"`
	CustomFields      []CustomField `json:"CustomFields,omitempty"`
}
