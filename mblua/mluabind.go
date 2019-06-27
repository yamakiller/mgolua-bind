package mblua

import (
	"github.com/yamakiller/mgolua/mlua"
)

func MBLuaPushGlobalsTable(L *mlua.State) {
	L.PushValue(mlua.LUA_REGISTRYINDEX)
	L.PushNumber(mlua.LUA_RIDX_GLOBALS)
	L.RawGet(-2)
	L.Remove(-2)
}

func MBLuaOpen(L *mlua.State) {

}
