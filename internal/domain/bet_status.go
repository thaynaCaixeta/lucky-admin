package domain

type BetStatus int

const (
	NotPaid BetStatus = iota
	Paid
)

func (s BetStatus) String() string {
	switch s {
	case NotPaid:
		return "NOT_PAID"
	case Paid:
		return "PAID"
	default:
		return "UNKNOW"
	}
}
