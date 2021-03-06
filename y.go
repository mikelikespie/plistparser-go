//line gram.y:2
package plistparser

import __yyfmt__ "fmt"

//line gram.y:2
import (
	"encoding/hex"
	"strconv"
)

//line gram.y:10
type yySymType struct {
	yys   int
	tok   int
	val   interface{}
	str   string
	data  []byte
	array []interface{}
	pair  struct {
		key string
		val interface{}
	}
	pairs map[string]interface{}
}

const STRING = 57346
const IDENT = 57347

var yyToknames = []string{
	"STRING",
	"IDENT",
	"'{'",
	"'}'",
	"'('",
	"')'",
	"'<'",
	"'>'",
	"'='",
	"';'",
	"','",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line gram.y:143

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 24
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 44

var yyAct = []int{

	16, 2, 15, 22, 3, 19, 7, 8, 10, 25,
	9, 12, 11, 33, 24, 20, 7, 8, 10, 28,
	9, 23, 11, 20, 27, 30, 23, 29, 31, 32,
	7, 8, 1, 26, 7, 8, 21, 17, 6, 13,
	14, 4, 5, 18,
}
var yyPact = []int{

	12, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, 2,
	30, 21, -1000, 5, -5, -1000, -1000, -1000, 26, -1000,
	7, 16, -1000, -1000, -1000, 12, -1000, -1000, 12, -1000,
	-1000, -1000, 0, -1000,
}
var yyPgo = []int{

	0, 5, 43, 42, 4, 41, 40, 39, 0, 2,
	38, 36, 3, 32,
}
var yyR1 = []int{

	0, 13, 8, 8, 8, 8, 4, 4, 5, 5,
	7, 7, 6, 6, 9, 3, 3, 2, 2, 1,
	10, 12, 11, 11,
}
var yyR2 = []int{

	0, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 2, 1, 3, 1, 2, 3, 1, 2, 4,
	3, 1, 1, 2,
}
var yyChk = []int{

	-1000, -13, -8, -4, -5, -3, -10, 4, 5, 8,
	6, 10, 9, -7, -6, -9, -8, 7, -2, -1,
	-4, -11, -12, 5, 9, 14, 7, -1, 12, 11,
	-12, -9, -8, 13,
}
var yyDef = []int{

	0, -2, 1, 2, 3, 4, 5, 6, 7, 0,
	0, 0, 8, 0, 10, 12, 14, 15, 0, 17,
	0, 0, 22, 21, 9, 11, 16, 18, 0, 20,
	23, 13, 0, 19,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	8, 9, 3, 3, 14, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 13,
	10, 12, 11, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 6, 3, 7,
}
var yyTok2 = []int{

	2, 3, 4, 5,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line gram.y:40
		{
			yylex.(*lexer).result = yyS[yypt-0].val
		}
	case 2:
		//line gram.y:45
		{
			yyVAL.val = yyS[yypt-0].str
		}
	case 3:
		//line gram.y:48
		{
			yyVAL.val = yyS[yypt-0].array
		}
	case 4:
		//line gram.y:51
		{
			yyVAL.val = yyS[yypt-0].pairs
		}
	case 5:
		//line gram.y:54
		{
			yyVAL.val = yyS[yypt-0].data
		}
	case 6:
		//line gram.y:59
		{
			var e error
			if yyVAL.str, e = strconv.Unquote(yyS[yypt-0].str); e != nil {
				yylex.Error("invalid unquoting" + e.Error())
			}
		}
	case 7:
		//line gram.y:65
		{
			yyVAL.str, _ = strconv.Unquote("\"" + yyS[yypt-0].str + "\"")
		}
	case 8:
		//line gram.y:69
		{
			yyVAL.array = []interface{}{}
		}
	case 9:
		//line gram.y:72
		{
			yyVAL.array = yyS[yypt-1].array
		}
	case 10:
		//line gram.y:78
		{
			yyVAL.array = yyS[yypt-0].array
		}
	case 11:
		//line gram.y:81
		{
			yyVAL.array = yyS[yypt-1].array
		}
	case 12:
		//line gram.y:85
		{
			yyVAL.array = []interface{}{yyS[yypt-0].val}
		}
	case 13:
		//line gram.y:88
		{
			yyVAL.array = append(yyS[yypt-2].array, yyS[yypt-0].val)
		}
	case 14:
		//line gram.y:93
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 15:
		//line gram.y:98
		{
			yyVAL.pairs = map[string]interface{}{}
		}
	case 16:
		//line gram.y:101
		{
			yyVAL.pairs = yyS[yypt-1].pairs
		}
	case 17:
		//line gram.y:106
		{
			yyVAL.pairs = map[string]interface{}{yyS[yypt-0].pair.key: yyS[yypt-0].pair.val}
		}
	case 18:
		//line gram.y:109
		{
			yyVAL.pairs[yyS[yypt-0].pair.key] = yyS[yypt-0].pair.val
		}
	case 19:
		//line gram.y:114
		{
			yyVAL.pair.key, yyVAL.pair.val = yyS[yypt-3].str, yyS[yypt-1].val
		}
	case 20:
		//line gram.y:119
		{
			yyVAL.data = yyS[yypt-1].data
		}
	case 21:
		//line gram.y:123
		{
			var e error

			if len(yyS[yypt-0].str) != 4 {
				yylex.Error("Chunks inside < > brackets must be lenght of 4")
			}

			if yyVAL.data, e = hex.DecodeString(yyS[yypt-0].str); e != nil {
				yylex.Error("error decoding hex chunk" + e.Error())
			}
		}
	case 22:
		//line gram.y:134
		{
			yyVAL.data = yyS[yypt-0].data
		}
	case 23:
		//line gram.y:138
		{
			yyVAL.data = append(yyS[yypt-1].data, yyS[yypt-0].data...)
		}
	}
	goto yystack /* stack new state and value */
}
