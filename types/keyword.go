package types

type Keyword string

var (
	CONST Keyword = "const"
)

func (k Keyword) IsValid() bool {
	switch k {
	case CONST:
		return true
	default:
		return false
	}
}

func (k Keyword) Is(v string) bool {
	switch Keyword(v) {
	case CONST:
		return true
	default:
		return false
	}
}
