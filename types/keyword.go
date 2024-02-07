package types

type Keyword string

var (
	CONST                 Keyword = "const"
	Function                      = Keyword(FUNCTION)
	INLINE_COMMENT        Keyword = "//"
	MULTILINE_COMMENT     Keyword = "/*"
	MULTILINE_COMMENT_END Keyword = "*/"
)

func (k Keyword) IsValid() bool {
	switch k {
	case CONST, INLINE_COMMENT, MULTILINE_COMMENT, Function:
		return true
	default:
		return false
	}
}

func (k Keyword) Is(v string) bool {
	switch Keyword(v) {
	case k:
		return true
	default:
		return false
	}
}
