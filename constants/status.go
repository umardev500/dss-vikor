package constants

type Status string

const (
	Active   Status = "active"
	Inactive Status = "inactive"
	Suspend  Status = "suspend"
	Pending  Status = "pending"
)
