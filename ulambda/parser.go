package ulambda

import (
	"regexp"
	"strings"
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

func ParseLex(lex []string) Node {
	n := lex[0]
	if n[0] == '\\' {
		return NewLambda(n[1:], ParseLex(lex[1:]))
	} else if len(lex) == 1 {
		return NewVar(n)
	} else {
		return NewApp(ParseLex(lex[0:1]), ParseLex(lex[1:]))
	}
	// raise err
}

func ParseExpr(expr string) Node {
	return ParseLex(lexer(expr))
}
