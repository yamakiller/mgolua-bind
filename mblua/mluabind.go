package mblua

import (
	"github.com/yamakiller/mgolua/mlua"
)

func goclass_gc_event(L *mlua.State) int {
	return 0
}

func MBLuaModuleEnd(L *mlua.State) {
	L.Pop(1)
}

func MBLuaModule(L *mlua.State, name *string, hasva int) {
	if name != nil {
		L.PushString(*name) //TODO 逻辑有点问题
		L.RawGet(-2)
		if !L.IsTable(-1) {
			L.Pop(1)
			L.NewTable()
			L.PushString(*name)
			L.PushValue(-2)
			L.RawSet(-4)
		}
	} else {
		L.PushValue(mlua.LUA_RIDX_GLOBALS)
	}
}

func MBLuaOpen(L *mlua.State) {
	top := L.GetTop()
	L.PushString("mblua_opened")
	L.RawGet(mlua.LUA_REGISTRYINDEX)
	if !L.IsBoolean(-1) {
		//lua_pushstring(L,"tolua_opened"); lua_pushboolean(L,1); lua_rawset(L,LUA_REGISTRYINDEX);
		L.PushString("mblua_opened")
		L.PushBoolean(true)
		L.RawSet(mlua.LUA_REGISTRYINDEX)

		L.PushString("mblua_box")
		L.NewTable()

		L.NewTable()
		L.PushLiteral("__mode")
		L.PushLiteral("v")
		L.RawSet(-3)
		L.SetMetaTable(-2)
		L.RawSet(mlua.LUA_REGISTRYINDEX)

		L.PushString("mblua_super")
		L.NewTable()
		L.RawSet(mlua.LUA_REGISTRYINDEX)
		L.PushString("mblua_gc")
		L.NewTable()
		L.RawSet(mlua.LUA_REGISTRYINDEX)

		/* create gc_event closure */
		L.PushString("mblua_gc_event")
		L.PushString("mblua_gc")
		L.RawGet(mlua.LUA_REGISTRYINDEX)
		L.PushString("mblua_super")
		L.RawGet(mlua.LUA_REGISTRYINDEX)
		L.PushGoClosure(goclass_gc_event, 2)
		L.RawSet(mlua.LUA_REGISTRYINDEX)
		L.NewMetaTable("mblua_common_class")

		/*tolua_module(L, NULL, 0)
		tolua_beginmodule(L, NULL)
		tolua_module(L, "tolua", 0)
		tolua_beginmodule(L, "tolua")
		tolua_function(L, "type", tolua_bnd_type)
		tolua_function(L, "takeownership", tolua_bnd_takeownership)
		tolua_function(L, "releaseownership", tolua_bnd_releaseownership)
		tolua_function(L, "cast", tolua_bnd_cast)
		tolua_function(L, "inherit", tolua_bnd_inherit)*/
	}
	L.SetTop(top)
}
