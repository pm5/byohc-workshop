package parser

import "testing"

func TestParseNodeBooleanAnd(t *testing.T) {
	n := ParseNode(`(\true
			(\false
			(\and
				(and true) true
			)(\a \b (a b) false)
		)(\a \b b)
		)(\a \b a)`)
}
