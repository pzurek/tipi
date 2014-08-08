package tp

// Company struct
type Company struct {
	ID          *float64 `json:"Id,omitempty"`
	Name        *string  `json:"Name,omitempty"`
	URL         *string  `json:"Url,omitempty"`
	Description *string  `json:"Description,omitempty"`
}

// CustomField struct
type CustomField struct {
	Name  *string `json:"Name,omitempty"`
	Type  *string `json:"Type,omitempty"`
	Value *string `json:"Value,omitempty"`
}

// EntityState struct
type EntityState struct {
	ID   *float64 `json:"Id,omitempty"`
	Name *string  `json:"Name,omitempty"`
}

// EntityType struct
type EntityType struct {
	ID   *float64 `json:"Id,omitempty"`
	Name *string  `json:"Name,omitempty"`
}

// Iteration struct
type Iteration struct {
	ID *float64 `json:"Id,omitempty"`
}

// Owner struct
type Owner struct {
	ID        *float64 `json:"Id,omitempty"`
	FirstName *string  `json:"FirstName,omitempty"`
	LastName  *string  `json:"LastName,omitempty"`
}

// Priority struct
type Priority struct {
	ID   *float64 `json:"Id,omitempty"`
	Name *string  `json:"Name,omitempty"`
}

// Process struct
type Process struct {
	ID          *float64 `json:"Id,omitempty"`
	Name        *string  `json:"Name,omitempty"`
	IsDefault   *bool    `json:"IsDefault,omitempty"`
	Description *string  `json:"Description,omitempty"`
}

// Role struct
type Role struct {
	ID   *float64 `json:"Id,omitempty"`
	Name *string  `json:"Name,omitempty"`
}

// Severity struct
type Severity struct {
	ID   *float64 `json:"Id,omitempty"`
	Name *string  `json:"Name,omitempty"`
}
