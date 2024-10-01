package parser

import (
	"log"
	"regexp"
	"strconv"
)

type Tokenizer struct {
	tokens []Token
	curr   int
	line   int
	src    string
}

type Token struct {
	Lexeme    string
	TokenType Lexeme
	Value     interface{}
	Line      int
}

func NewTokenizer(source string) *Tokenizer {
	t := &Tokenizer{curr: -1, line: 0}
	t.src = source
	t.tokenize()
	return t
}

func (t *Tokenizer) tokenize() {
	for t.nextChar() {
		char := t.advance()
		switch char {
		case "(":
			t.createToken(char, nil, LEFT_PAREN, t.line)
		case ")":
			t.createToken(char, nil, RIGHT_PAREN, t.line)
		case "{":
			t.createToken(char, nil, LEFT_BRACE, t.line)
		case "}":
			t.createToken(char, nil, RIGHT_BRACE, t.line)
		case ",":
			t.createToken(char, nil, COMMA, t.line)
		case ".":
			t.createToken(char, nil, DOT, t.line)
		case "-":
			t.createToken(char, nil, MINUS, t.line)
		case "+":
			t.createToken(char, nil, PLUS, t.line)
		case "/":
			t.createToken(char, nil, SLASH, t.line)
		case "*":
			t.createToken(char, nil, STAR, t.line)
		case "var":
			t.createToken(char, nil, VAR, t.line)
		case "o":
			if t.match("r") {
				t.createToken("or", nil, OR, t.line)
				t.advance()
			}
		case "=":
			if t.match("=") {
				t.createToken("==", nil, EQUAL_EQUAL, t.line)
				t.advance()
			} else {
				t.createToken("=", nil, EQUAL, t.line)
			}
		case "!":
			if t.match("=") {
				t.createToken("!=", nil, NOT_EQUAL, t.line)
				t.advance()
			} else {
				t.createToken("!", nil, BANG, t.line)
				t.advance()
			}
		case "<":
			if t.match("=") {
				t.createToken("<", nil, LESS_EQUAL, t.line)
				t.advance()
			} else {
				t.createToken("<", nil, LESS, t.line)
				t.advance()
			}
		case ">":
			if t.match("=") {
				t.createToken(">=", nil, GREATER_EQUAL, t.line)
				t.advance()
			} else {
				t.createToken(">", nil, GREATER, t.line)
			}
		case `"`:
			t.string(t.curr)
		case " ":
		case "\r":
		case `\0`:
		case `\`:
			// Ignore whitespace.
			continue
		case "\n":
			t.line++
			continue
		default:
			if isDigit(char) {
				t.number(t.curr)
			} else if isAlpha(char) {
				t.identifier(t.curr)
			} else {
				log.Fatalln("Lox Error : " + char + " is invalid")
			}
		}
	}
	t.createToken("EOF", nil, EOF, t.line)
}

func (t *Tokenizer) GetTokens() []Token {
	return t.tokens
}

func (t *Tokenizer) createToken(lexeme string, value interface{}, tokenType Lexeme, line int) {
	token := Token{
		Lexeme:    lexeme,
		TokenType: tokenType,
		Value:     value,
		Line:      line,
	}
	log.Println(t.curr, t.line, lexeme)
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) match(character string) bool {
	return string(t.src[t.curr+1]) == character
}

func (t *Tokenizer) peek() string {
	if !t.nextChar() {
		return `\0`
	}
	return string(t.src[t.curr])
}
func (t *Tokenizer) peekNext() string {
	if t.curr+1 >= len(t.src) {
		return `\0`
	}
	return string(t.src[t.curr+1])
}

func (t *Tokenizer) nextChar() bool {
	return t.curr < len(t.src)
}

func (t *Tokenizer) advance() string {
	t.curr++
	return t.peek()
}

func (t *Tokenizer) string(start int) {
	t.advance()
	for t.peek() != `"` && t.nextChar() {
		if t.peek() == `\n` && t.nextChar() {
			t.line++
		}
		t.advance()
	}
	if !t.nextChar() {
		log.Printf("[ERROR]: Unterminated String at position : %d:%d, line : %d", start+1, t.curr-1, t.line)
		return
	}
	t.advance()

	var value string = t.src[start+1 : t.curr-1]
	t.createToken(value, value, STRING, t.line)
}

func (t *Tokenizer) number(start int) {
	for isDigit(t.peek()) {
		t.advance()
	}

	if t.peek() == "." && isDigit(t.peekNext()) {
		t.advance()

		for isDigit(t.peek()) {
			t.advance()
		}
	}
	value, err := strconv.ParseFloat(t.src[start:t.curr], 64)
	if err != nil {
		log.Println(err.Error())
		return
	}
	t.createToken("number", value, NUMBER, t.line)
}

func (t *Tokenizer) identifier(start int) {
	for isAlphaNumeric(t.peek()) {
		t.advance()
	}
	var token string = t.src[start:t.curr]
	tokenType, exists := Keywords[token]
	if !exists {
		t.createToken("", token, IDENTIFIER, t.line)
	} else {
		t.createToken("", token, tokenType, t.line)
	}

}

func isDigit(c string) bool {
	return c == "0" || c == "1" || c == "2" || c == "3" || c == "4" || c == "5" || c == "6" || c == "7" || c == "8" || c == "9"
}

func isAlpha(c string) bool {
	expression := regexp.MustCompile(`^[a-zA-Z0-9_]*$`)
	return expression.MatchString(c)
}

func isAlphaNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}
