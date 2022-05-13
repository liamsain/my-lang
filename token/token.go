package token

// is this really needed..
// type TokenType string

type Token struct {
	Type    string
	Literal string
}

const (
	Illegal = "Illegal"
	Eof     = "Eof"
)
