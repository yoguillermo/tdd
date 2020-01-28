package main

import "fmt"

func main()  {
  fmt.Printf(HolaMundo(""))
}

func HolaMundo(name string) string {
  if (name == ""){
    name = "Parrita mi hijo"
  }
  return "Hola, "+name
}