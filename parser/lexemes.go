package parser

type Lexeme int

const (
	LEFT_PAREN  Lexeme = 0
	RIGHT_PAREN Lexeme = 1
	LEFT_BRACE  Lexeme = 2
	RIGHT_BRACE Lexeme = 3
	COMMA       Lexeme = 4
	DOT         Lexeme = 5
	MINUS       Lexeme = 6
	PLUS        Lexeme = 7
	SEMICOLON   Lexeme = 8
	SLASH       Lexeme = 9
	STAR        Lexeme = 10

	EQUAL         Lexeme = 11
	NOT_EQUAL     Lexeme = 12
	BANG          Lexeme = 13
	EQUAL_EQUAL   Lexeme = 14
	GREATER       Lexeme = 15
	GREATER_EQUAL Lexeme = 16
	LESS          Lexeme = 17
	LESS_EQUAL    Lexeme = 18

	IDENTIFIER Lexeme = 19
	STRING     Lexeme = 20
	NUMBER     Lexeme = 21

	AND   Lexeme = 22
	CLASS Lexeme = 23
	ELSE  Lexeme = 24
	FALSE Lexeme = 25
	FUN   Lexeme = 26
	FOR   Lexeme = 27
	IF    Lexeme = 28
	NIL   Lexeme = 29
	OR    Lexeme = 30

	PRINT  Lexeme = 31
	RETURN Lexeme = 32
	SUPER  Lexeme = 33
	THIS   Lexeme = 34
	TRUE   Lexeme = 35
	VAR    Lexeme = 36
	WHILE  Lexeme = 37

	EOF Lexeme = 38
)
