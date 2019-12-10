package bignum

import (
	"github.com/shopspring/decimal"
	lua "github.com/yuin/gopher-lua"
)

func LNum_NewUD(L *lua.LState, gs decimal.Decimal) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("num"))
	return ud
}

func check_LNum(L *lua.LState, index int) decimal.Decimal {

	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(decimal.Decimal); ok {
		return v
	}
	L.ArgError(1, "bignum.num object expected")
	return decimal.Decimal{}
}

func LNum_New(L *lua.LState) int {
	n := L.CheckNumber(1)
	L.Push(LNum_NewUD(L, decimal.NewFromFloat(float64(n))))
	return 1
}

func LNum_FromString(L *lua.LState) int {

	n := L.CheckString(1)
	d, err := decimal.NewFromString(n)
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}
	L.Push(LNum_NewUD(L, d))
	L.Push(lua.LNil)
	return 2
}

func LNum_Add(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Add(n2)))
	return 1
}

func LNum_Sub(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Sub(n2)))
	return 1
}

func LNum_Mul(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Mul(n2)))
	return 1
}

func LNum_Div(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Div(n2)))
	return 1
}

func LNum_DivRound(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	p := L.CheckInt(3)
	L.Push(LNum_NewUD(L, n.DivRound(n2, int32(p))))
	return 1
}

func LNum_Equal(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.Equal(n2)))
	return 1
}

func LNum_Equals(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.Equals(n2)))
	return 1
}

func LNum_Ceil(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Ceil()))
	return 1
}

func LNum_Abs(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Abs()))
	return 1
}

func LNum_Floor(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Floor()))
	return 1
}

func LNum_Round(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(LNum_NewUD(L, n.Round(int32(p))))
	return 1
}

func LNum_RoundBank(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(LNum_NewUD(L, n.RoundBank(int32(p))))
	return 1
}

func LNum_RoundCash(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(LNum_NewUD(L, n.RoundCash(uint8(p))))
	return 1
}

func LNum_Tan(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Tan()))
	return 1
}

func LNum_Sin(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Sin()))
	return 1
}

func LNum_Cos(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Cos()))
	return 1
}

func LNum_Cmp(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LNumber(n.Cmp(n2)))
	return 1
}

func LNum_Exponent(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(lua.LNumber(float64(n.Exponent())))
	return 1
}

func LNum_GreaterThan(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.GreaterThan(n2)))
	return 1
}

func LNum_GreaterThanOrEqual(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.GreaterThanOrEqual(n2)))
	return 1
}

func LNum_LessThan(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.LessThan(n2)))
	return 1
}

func LNum_LessThanOrEqual(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(lua.LBool(n.LessThanOrEqual(n2)))
	return 1
}

func LNum_Atan(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(LNum_NewUD(L, n.Atan()))
	return 1
}

func LNum_String(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(lua.LString(n.String()))
	return 1
}

func LNum_StringFixed(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(lua.LString(n.StringFixed(int32(p))))
	return 1
}

func LNum_StringFixedBank(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(lua.LString(n.StringFixedBank(int32(p))))
	return 1
}

func LNum_StringFixedCash(L *lua.LState) int {
	n := check_LNum(L, 1)
	p := L.CheckInt(2)
	L.Push(lua.LString(n.StringFixedCash(uint8(p))))
	return 1
}

func LNum_GetFloat(L *lua.LState) int {
	n := check_LNum(L, 1)
	d, e := n.Float64()
	L.Push(lua.LNumber(d))
	L.Push(lua.LBool(e))
	return 2
}

func LNum_Mod(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Mod(n2)))
	return 2
}

func LNum_Pow(L *lua.LState) int {
	n := check_LNum(L, 1)
	n2 := check_LNum(L, 2)
	L.Push(LNum_NewUD(L, n.Pow(n2)))
	return 2
}

func LNum_IsNegative(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(lua.LBool(n.IsNegative()))
	return 1
}

func LNum_IsPositive(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(lua.LBool(n.IsPositive()))
	return 1
}

func LNum_IntPart(L *lua.LState) int {
	n := check_LNum(L, 1)
	L.Push(lua.LNumber(n.IntPart()))
	return 1
}
