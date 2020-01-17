// Code generated by goyacc -o promql/generated_parser.y.go promql/generated_parser.y. DO NOT EDIT.

//line promql/generated_parser.y:15
package promql

import __yyfmt__ "fmt"

//line promql/generated_parser.y:15

import (
	"sort"

	"github.com/prometheus/prometheus/pkg/labels"
)

//line promql/generated_parser.y:24
type yySymType struct {
	yys      int
	node     Node
	item     Item
	matchers []*labels.Matcher
	matcher  *labels.Matcher
	label    labels.Label
	labels   labels.Labels
}

const ERROR = 57346
const EOF = 57347
const COMMENT = 57348
const IDENTIFIER = 57349
const METRIC_IDENTIFIER = 57350
const LEFT_PAREN = 57351
const RIGHT_PAREN = 57352
const LEFT_BRACE = 57353
const RIGHT_BRACE = 57354
const LEFT_BRACKET = 57355
const RIGHT_BRACKET = 57356
const COMMA = 57357
const ASSIGN = 57358
const COLON = 57359
const SEMICOLON = 57360
const STRING = 57361
const NUMBER = 57362
const DURATION = 57363
const BLANK = 57364
const TIMES = 57365
const SPACE = 57366
const operatorsStart = 57367
const SUB = 57368
const ADD = 57369
const MUL = 57370
const MOD = 57371
const DIV = 57372
const LAND = 57373
const LOR = 57374
const LUNLESS = 57375
const EQL = 57376
const NEQ = 57377
const LTE = 57378
const LSS = 57379
const GTE = 57380
const GTR = 57381
const EQL_REGEX = 57382
const NEQ_REGEX = 57383
const POW = 57384
const operatorsEnd = 57385
const aggregatorsStart = 57386
const AVG = 57387
const COUNT = 57388
const SUM = 57389
const MIN = 57390
const MAX = 57391
const STDDEV = 57392
const STDVAR = 57393
const TOPK = 57394
const BOTTOMK = 57395
const COUNT_VALUES = 57396
const QUANTILE = 57397
const aggregatorsEnd = 57398
const keywordsStart = 57399
const OFFSET = 57400
const BY = 57401
const WITHOUT = 57402
const ON = 57403
const IGNORING = 57404
const GROUP_LEFT = 57405
const GROUP_RIGHT = 57406
const BOOL = 57407
const keywordsEnd = 57408
const startSymbolsStart = 57409
const START_LABELS = 57410
const START_LABEL_SET = 57411
const START_METRIC = 57412
const startSymbolsEnd = 57413

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"ERROR",
	"EOF",
	"COMMENT",
	"IDENTIFIER",
	"METRIC_IDENTIFIER",
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"LEFT_BRACKET",
	"RIGHT_BRACKET",
	"COMMA",
	"ASSIGN",
	"COLON",
	"SEMICOLON",
	"STRING",
	"NUMBER",
	"DURATION",
	"BLANK",
	"TIMES",
	"SPACE",
	"operatorsStart",
	"SUB",
	"ADD",
	"MUL",
	"MOD",
	"DIV",
	"LAND",
	"LOR",
	"LUNLESS",
	"EQL",
	"NEQ",
	"LTE",
	"LSS",
	"GTE",
	"GTR",
	"EQL_REGEX",
	"NEQ_REGEX",
	"POW",
	"operatorsEnd",
	"aggregatorsStart",
	"AVG",
	"COUNT",
	"SUM",
	"MIN",
	"MAX",
	"STDDEV",
	"STDVAR",
	"TOPK",
	"BOTTOMK",
	"COUNT_VALUES",
	"QUANTILE",
	"aggregatorsEnd",
	"keywordsStart",
	"OFFSET",
	"BY",
	"WITHOUT",
	"ON",
	"IGNORING",
	"GROUP_LEFT",
	"GROUP_RIGHT",
	"BOOL",
	"keywordsEnd",
	"startSymbolsStart",
	"START_LABELS",
	"START_LABEL_SET",
	"START_METRIC",
	"startSymbolsEnd",
}
var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line promql/generated_parser.y:223

//line yacctab:1
var yyExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 4,
	1, 27,
	-2, 0,
}

const yyPrivate = 57344

const yyLast = 69

var yyAct = [...]int{

	5, 31, 18, 23, 40, 48, 44, 38, 29, 25,
	1, 9, 7, 10, 24, 21, 11, 36, 27, 45,
	37, 28, 47, 43, 30, 20, 16, 25, 6, 0,
	19, 42, 24, 32, 33, 41, 39, 22, 20, 34,
	35, 46, 8, 19, 13, 0, 0, 12, 17, 15,
	14, 0, 0, 9, 26, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 2, 3, 4,
}
var yyPact = [...]int{

	-2, -1000, 1, 0, 42, -1000, -1000, 36, -1000, 25,
	-1000, 0, -1000, -1000, -1000, -1000, 6, -1000, -1000, -1,
	-1000, 5, -1000, -1000, 2, -1000, -1000, -1000, 23, -1000,
	4, -1000, -1000, -1000, -1000, -1000, -1000, 7, -1000, 3,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
}
var yyPgo = [...]int{

	0, 28, 26, 2, 24, 16, 15, 42, 13, 3,
	10,
}
var yyR1 = [...]int{

	0, 10, 10, 10, 10, 1, 1, 1, 2, 2,
	2, 3, 3, 3, 3, 4, 4, 4, 4, 8,
	8, 8, 5, 5, 7, 7, 7, 7, 6, 6,
	6, 9, 9, 9, 9,
}
var yyR2 = [...]int{

	0, 2, 2, 2, 1, 3, 4, 2, 3, 1,
	2, 3, 3, 2, 1, 1, 1, 1, 1, 2,
	1, 1, 1, 1, 3, 4, 2, 0, 3, 1,
	2, 3, 3, 2, 1,
}
var yyChk = [...]int{

	-1000, -10, 68, 69, 70, 2, -1, 11, -7, 11,
	-8, -5, -7, 2, 8, 7, -2, 12, -3, 7,
	2, -6, 12, -9, 7, 2, -7, 12, 15, 2,
	-4, 2, 34, 35, 40, 41, 12, 15, 2, 34,
	2, 12, -3, 19, 2, 12, -9, 19, 2,
}
var yyDef = [...]int{

	0, -2, 0, 27, -2, 4, 1, 0, 2, 0,
	3, 27, 20, 21, 22, 23, 0, 7, 9, 0,
	14, 0, 26, 29, 0, 34, 19, 5, 0, 10,
	0, 13, 15, 16, 17, 18, 24, 0, 30, 0,
	33, 6, 8, 11, 12, 25, 28, 31, 32,
}
var yyTok1 = [...]int{

	1,
}
var yyTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48, 49, 50, 51,
	52, 53, 54, 55, 56, 57, 58, 59, 60, 61,
	62, 63, 64, 65, 66, 67, 68, 69, 70, 71,
}
var yyTok3 = [...]int{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
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

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := yyPact[state]
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && yyChk[yyAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || yyExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := yyExca[i]
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		token = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = yyTok3[i+0]
		if token == char {
			token = yyTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
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
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
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
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
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
			if yyn < 0 || yyn == yytoken {
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
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
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
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
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
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
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
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:126
		{
			yylex.(*parser).generatedParserResult.(*VectorSelector).LabelMatchers = yyDollar[2].matchers
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:128
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].labels
		}
	case 3:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:130
		{
			yylex.(*parser).generatedParserResult = yyDollar[2].labels
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:132
		{
			yylex.(*parser).errorf("unexpected %v", yylex.(*parser).token.desc())
		}
	case 5:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:138
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 6:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:140
		{
			yyVAL.matchers = yyDollar[2].matchers
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:142
		{
			yyVAL.matchers = []*labels.Matcher{}
		}
	case 8:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:148
		{
			yyVAL.matchers = append(yyDollar[1].matchers, yyDollar[3].matcher)
		}
	case 9:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:150
		{
			yyVAL.matchers = []*labels.Matcher{yyDollar[1].matcher}
		}
	case 10:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:152
		{
			yylex.(*parser).errorf("unexpected %v in label matching, expected \",\" or \"}\"", yylex.(*parser).token.desc())
		}
	case 11:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:157
		{
			yyVAL.matcher = yylex.(*parser).newLabelMatcher(yyDollar[1].item, yyDollar[2].item, yyDollar[3].item)
		}
	case 12:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:159
		{
			yylex.(*parser).errorf("unexpected %v in label matching, expected string", yylex.(*parser).token.desc())
		}
	case 13:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:161
		{
			yylex.(*parser).errorf("unexpected %v in label matching, expected label matching operator", yylex.(*parser).token.Val)
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:163
		{
			yylex.(*parser).errorf("unexpected %v in label matching, expected identifier or \"}\"", yylex.(*parser).token.desc())
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:167
		{
			yyVAL.item = yyDollar[1].item
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:168
		{
			yyVAL.item = yyDollar[1].item
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:169
		{
			yyVAL.item = yyDollar[1].item
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:170
		{
			yyVAL.item = yyDollar[1].item
		}
	case 19:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:176
		{
			yyVAL.labels = append(yyDollar[2].labels, labels.Label{Name: labels.MetricName, Value: yyDollar[1].item.Val})
			sort.Sort(yyVAL.labels)
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:178
		{
			yyVAL.labels = yyDollar[1].labels
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:180
		{
			yylex.(*parser).errorf("missing metric name or metric selector")
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:185
		{
			yyVAL.item = yyDollar[1].item
		}
	case 23:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:186
		{
			yyVAL.item = yyDollar[1].item
		}
	case 24:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:190
		{
			yyVAL.labels = labels.New(yyDollar[2].labels...)
		}
	case 25:
		yyDollar = yyS[yypt-4 : yypt+1]
//line promql/generated_parser.y:192
		{
			yyVAL.labels = labels.New(yyDollar[2].labels...)
		}
	case 26:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:194
		{
			yyVAL.labels = labels.New()
		}
	case 27:
		yyDollar = yyS[yypt-0 : yypt+1]
//line promql/generated_parser.y:196
		{
			yyVAL.labels = labels.New()
		}
	case 28:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:201
		{
			yyVAL.labels = append(yyDollar[1].labels, yyDollar[3].label)
		}
	case 29:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:203
		{
			yyVAL.labels = []labels.Label{yyDollar[1].label}
		}
	case 30:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:205
		{
			yylex.(*parser).errorf("unexpected %v in label set, expected \",\" or \"}\"", yylex.(*parser).token.desc())
		}
	case 31:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:211
		{
			yyVAL.label = labels.Label{Name: yyDollar[1].item.Val, Value: yylex.(*parser).unquoteString(yyDollar[3].item.Val)}
		}
	case 32:
		yyDollar = yyS[yypt-3 : yypt+1]
//line promql/generated_parser.y:213
		{
			yylex.(*parser).errorf("unexpected %v in label set, expected string", yylex.(*parser).token.desc())
		}
	case 33:
		yyDollar = yyS[yypt-2 : yypt+1]
//line promql/generated_parser.y:215
		{
			yylex.(*parser).errorf("unexpected %v in label set, expected \"=\"", yylex.(*parser).token.desc())
		}
	case 34:
		yyDollar = yyS[yypt-1 : yypt+1]
//line promql/generated_parser.y:217
		{
			yylex.(*parser).errorf("unexpected %v in label set, expected identifier or \"}\"", yylex.(*parser).token.desc())
		}
	}
	goto yystack /* stack new state and value */
}
