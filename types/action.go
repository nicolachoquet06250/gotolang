package types

type Action string

var (
	AFFECTATION Action = "="
)

func (a Action) IsValid() bool {
	switch a {
	case AFFECTATION:
		return true
	default:
		return false
	}
}
