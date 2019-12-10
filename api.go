package bignum

import (
	lua "github.com/yuin/gopher-lua"
)

var VERSION string = "1.0.1"
var LNAME string = "bignum"

var API = map[string]lua.LGFunction{
	"new":        LNum_New,
	"fromString": LNum_FromString,
}

var API_LNum = map[string]lua.LGFunction{
	"add":                LNum_Add,
	"sub":                LNum_Sub,
	"mul":                LNum_Mul,
	"div":                LNum_Div,
	"divRound":           LNum_DivRound,
	"mod":                LNum_Mod,
	"pow":                LNum_Pow,
	"round":              LNum_Round,
	"roundBank":          LNum_RoundBank,
	"roundCash":          LNum_RoundCash,
	"floor":              LNum_Floor,
	"ceil":               LNum_Ceil,
	"sin":                LNum_Sin,
	"cos":                LNum_Cos,
	"tan":                LNum_Tan,
	"equal":              LNum_Equal,
	"equals":             LNum_Equals,
	"greaterThan":        LNum_GreaterThan,
	"greaterThanOrEqual": LNum_GreaterThanOrEqual,
	"lessThan":           LNum_LessThan,
	"lessThanOrEqual":    LNum_LessThanOrEqual,
	"isNegative":         LNum_IsNegative,
	"isPositive":         LNum_IsPositive,
	"cmp":                LNum_Cmp,
	"exponent":           LNum_Exponent,
	"atan":               LNum_Atan,
	"float":              LNum_GetFloat,
	"intPart":            LNum_IntPart,
	"string":             LNum_String,
	"stringFixed":        LNum_StringFixed,
	"stringFixedBank":    LNum_StringFixedBank,
	"stringFixedCash":    LNum_StringFixedCash,
}

func Preload(L *lua.LState) {
	L.PreloadModule(LNAME, Loader)
}

func Loader(L *lua.LState) int {
	t := L.NewTable()
	lNum := L.NewTypeMetatable("num")
	L.SetField(lNum, "__index", L.SetFuncs(L.NewTable(), API_LNum))
	t.RawSetH(lua.LString("num"), lNum)
	t.RawSetH(lua.LString("__version__"), lua.LString(VERSION))
	L.SetFuncs(t, API)
	L.Push(t)
	return 1
}
