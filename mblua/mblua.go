package mblua

import (
  "github.com/yamakiller/mgolua/mlua"
)

func push_globals_table(L *mlua.State){
  L.PushValue(int(mlua.LUA_REGISTRYINDEX))
  L.PushNumber(float64(mlua.LUA_RIDX_GLOBALS))
  L.RawGet(-2)
  L.Remove(-2)
}

func BBeginModule(L *mlua.State, name *string) {
  if (name != nil) {
    L.PushString(*name)
    L.RawGet(-2)
  } else {
    push_globals_table(L)
  }
}

func BEndModule(L *mlua.State) {
  L.Pop(1)
}

func BModule(L *mlua.State, name *string) {
  if name != nil {
      L.PushString(*name)
      L.RawGet(-2)
      if (!L.IsTable(-1)) {
           L.Pop(1)
           L.NewTable()
           L.PushString(*name)
           L.PushValue(-2)
           L.RawSet(-4)
      }
  } else {
    push_globals_table(L)
  }

  if (name != nil) {
    L.RawSet(-3)
  } else {
    L.Pop(1)
  }
}

func BFunction(L *mlua.State, name string, f mlua.LuaGoFunction) {
  L.PushString(name)
  L.PushGoFunction(f)
  L.RawSet(-3)
}
