package lexer

type Lexer struct {
	input        string
	position     int  // current position of the input
	readPosition int  // current reading position (after current char)
	ch           byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
