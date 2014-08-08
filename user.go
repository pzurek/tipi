package tp

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
