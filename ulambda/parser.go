package ulambda

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrUnknown = errors.New("ulambda: unknown error")
)

func replaceParenth(s string) string {
	pr := regexp.MustCompile(`([\(\)])`)
	sp := regexp.MustCompile(`  +`)
	ht := regexp.MustCompile(`^ +| +$`)
	return ht.ReplaceAllString(sp.ReplaceAllString(pr.ReplaceAllString(s, ` $1 `), ` `), ``)
}

func lexer(s string) []string {
	return strings.Split(replaceParenth(s), " ")
}

func ParseLex(lex []string) (Node, error) {
	n := lex[0]
	if n[0] == '\\' {
		c, err := ParseLex(lex[1:])
		if err != nil {
			return nil, err
		}
		return NewLambda(n[1:], c), nil
	} else if len(lex) == 1 {
		return NewVar(n), nil
	} else {
		f, err := ParseLex(lex[0:1])
		if err != nil {
			return nil, err
		}
		a, err := ParseLex(lex[1:])
		if err != nil {
			return nil, err
		}
		return NewApp(f, a), nil
	}
	return nil, ErrUnknown
}

func ParseExpr(expr string) (Node, error) {
	return ParseLex(lexer(expr))
}
