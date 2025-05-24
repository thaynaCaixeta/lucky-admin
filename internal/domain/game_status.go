package domain

type GameStatus int

const (
	OnGoing GameStatus = iota
	Completed
	Cancelled
)

func (s GameStatus) String() string {
	switch s {
	case OnGoing:
		return "ON_GOING"
	case Completed:
		return "COMPLETED"
	case Cancelled:
		return "CANCELLED"
	default:
		return "UNKNOW"
	}
}
