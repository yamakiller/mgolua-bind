package mblua

func module_index_event(L *State) int {
	L.PushString(".get")
	L.RawSet(-3)
	if !L.IsTable(-1) {
    L.PushValue(2);  /* key */
		L.RawGet(-2);
		if (lua_iscfunction(L,-1))
		{
      L.Call(0, 1)
			return 1;
		}
		else if (L.IsTable(-1)) {
      
    }
	}
}

func MBLuaIsModuleMetatable(L *State) int {
	r := 0
	if L.GetMetaTable(-1) != 0 {
		L.PushString("__index")
		L.RawGet(-2)
		//r = (lua_tocfunction(L, -1) == module_index_event)
		L.Pop(2)
	}
	return r
}
