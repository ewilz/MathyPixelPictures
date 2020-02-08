package main

import (
    "fmt"
    "syscall/js"
    "strconv"
)


func add(_ js.Value, i []js.Value) interface {} {
  fmt.Println(i)
  value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
  value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

  int1, _ := strconv.Atoi(value1)
  int2, _ := strconv.Atoi(value2)

  js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 + int2)
  fmt.Println(int1 + int2)
  return js.ValueOf(int1 + int2)
}

func subtract(this js.Value, i []js.Value) interface{} {
  value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
  value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

  int1, _ := strconv.Atoi(value1)
  int2, _ := strconv.Atoi(value2)

  fmt.Println(int1 - int2)
  js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1 - int2)
  return js.ValueOf(int1 - int2)
}

func registerCallbacks() {
  js.Global().Set("add", js.FuncOf(add))
  js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
  c := make(chan struct{}, 0)
  fmt.Println("Go WebAssembly Initialized")
  registerCallbacks()
  <-c
}
