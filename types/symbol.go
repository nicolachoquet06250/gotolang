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

func (s Symbol) Is(v string) bool {
	switch Symbol(v) {
	case s:
		return true
	default:
		return false
	}
}

func SymbolFromString(v string) *Symbol {
	if Symbol(v).IsValid() {
		switch Symbol(v) {
		case EQUAL:
			return &EQUAL
		case OEPN_BRACKET:
			return &OEPN_BRACKET
		case CLOSE_BRACKET:
			return &CLOSE_BRACKET
		case OEPN_HOOK:
			return &OEPN_HOOK
		case CLOSE_HOOK:
			return &CLOSE_HOOK
		case OEPN_PARENTHESIS:
			return &OEPN_PARENTHESIS
		case CLOSE_PARENTHESIS:
			return &CLOSE_PARENTHESIS
		case DOUBLE_QUOTE:
			return &DOUBLE_QUOTE
		case SIMPLE_QUOTE:
			return &SIMPLE_QUOTE
		case SEMICOLUMN:
			return &SEMICOLUMN
		default:
			return nil
		}
	}
	return nil
}
