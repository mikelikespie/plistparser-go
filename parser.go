package plistparser

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"text/scanner"
)

func isIdent(ch rune, i int) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	if ch >= 'A' && ch <= 'Z' {
		return true
	}
	if ch >= '0' && ch <= '9' {
		return true
	}

	switch ch {
	case '\\', '.', '_', '-', '$':
		return true

	}

	if i > 0 && ch == '/' {
		return true
	}

	return false
}

type lexer struct {
	s      scanner.Scanner
	inData bool
	result interface{}
}

func newLexer(reader io.Reader) (l *lexer) {
	l = &lexer{
		s: scanner.Scanner{
			Mode:        scanner.ScanIdents | scanner.ScanStrings | scanner.ScanComments | scanner.ScanIdents,
			IsIdentRune: isIdent,
		},
	}
	l.s.Init(reader)
	return
}

func (l *lexer) Lex(lval *yySymType) (tok int) {
	var accumulatedText string

	for stok := l.s.Scan(); ; stok = l.s.Scan() {
		switch stok {
		case scanner.EOF:
			return
		case scanner.Comment:
			continue
		case scanner.Ident:
			// We need to deal with individual '/' being counted as not an ident
			tok = IDENT
			lval.str = accumulatedText + l.s.TokenText()
			return
		case scanner.String:
			tok = STRING
			tt := l.s.TokenText()
			lval.str = tt
			return
		case '/':
			accumulatedText += "/"
			peeked := l.s.Peek()

			if isIdent(peeked, 1) {
				continue
			} else {
				tok = IDENT
				lval.str = accumulatedText
				println(lval.str, scanner.TokenString(l.s.Peek()))
				return
			}
		case '(', ')', '{', ';', '}', '=', ',', '<', '>':
			tok = int(stok)
			return
		default:
			log.Panicf("%d:%d: Unexpected token: %c", l.s.Line, l.s.Column, stok)

		}

	}

	return
}

func prettyPrint(v interface{}, indentLevel int) {
	printIndent := func (delta int) {
		for i := 0; i < indentLevel + delta; i++ {
			fmt.Printf("  ")
		}
	}
	switch v := v.(type) {
	case string:
		fmt.Print(strconv.Quote(v))
	case []interface{}:
		fmt.Printf("[\n")
		for _, val := range v {
			printIndent(1)

			prettyPrint(val, indentLevel + 1)
			fmt.Printf(",\n")
		}
		printIndent(0)
		fmt.Printf("]")

	case map[string]interface{}:
		fmt.Printf("{\n")
		for k, val := range v {
			printIndent(1)
			prettyPrint(k, indentLevel + 1)
			fmt.Printf(" = ")
			prettyPrint(val, indentLevel + 1)
			fmt.Printf(";\n")
		}
		printIndent(0)
		fmt.Printf("}")
	default:
		panic(v)

	}
}
func (l *lexer) Error(e string) {
	log.Panicf("%d:%d: %s last token text %s", l.s.Line, l.s.Column, e, l.s.TokenText())
}

func main() {
	l := newLexer(os.Stdin)
	yyParse(l)

	prettyPrint(l.result, 0)
}
