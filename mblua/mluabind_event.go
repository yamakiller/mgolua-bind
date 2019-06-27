package mblua

import (
	"github.com/yamakiller/mgolua/mlua"
)

func module_index_event(L *mlua.State) int {
	L.PushString(".get")
	L.RawGet(-3)
	if (L.IsTable(-1)) {
		L.PushValue(2)
		L.RawGet(-2)
		if L.IsGFunction(-1) {
			L.Call(0, 1)
			return 1
		} else if (L.IsTable(-1)) {
			return 1
		}
	}

	if L.GetMetaTable(1) != 0 {
		L.PushString("__index")
		L.RawGet(-2)
		L.PushValue(1)
		L.PushValue(2)
		if (L.IsGFunction(-1)){
			L.Call(2, 1)
			return 1
		} else if (L.IsTable(-1)) {
			L.GetTable(-3)
			return 1
		}
	}
	L.PushNil()
	return 1
}

func MBLuaIsModuleMetatable(L *mlua.State) int {
	r := 0
	if L.GetMetaTable(-1) != 0 {
		L.PushString("__index")
		L.RawGet(-2)
		//r = (lua_tocfunction(L, -1) == module_index_event)
		L.Pop(2)
	}
	return r
}
