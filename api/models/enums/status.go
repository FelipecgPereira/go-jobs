package enums

type Status string

const (
	StatusPending Status = "pending"
	StatusActive  Status = "active"
	StatusClosed  Status = "closed"
)

func (s Status) IsValid() bool {
	switch s {
	case StatusPending, StatusActive, StatusClosed:
		return true
	default:
		return false
	}
}
