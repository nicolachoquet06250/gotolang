package types

type Symbol string

var (
	EQUAL             Symbol = "="
	OEPN_BRACKET      Symbol = "{"
	CLOSE_BRACKET     Symbol = "}"
	OEPN_HOOK         Symbol = "["
	CLOSE_HOOK        Symbol = "]"
	OEPN_PARENTHESIS  Symbol = "("
	CLOSE_PARENTHESIS Symbol = ")"
	DOUBLE_QUOTE      Symbol = `"`
	SIMPLE_QUOTE      Symbol = "'"
	SEMICOLUMN        Symbol = ";"
)

func (s Symbol) IsValid() bool {
	switch s {
	case EQUAL, OEPN_BRACKET, CLOSE_BRACKET, OEPN_HOOK, CLOSE_HOOK, OEPN_PARENTHESIS, CLOSE_PARENTHESIS, DOUBLE_QUOTE, SIMPLE_QUOTE, SEMICOLUMN:
		return true
	default:
		return false
	}
}
