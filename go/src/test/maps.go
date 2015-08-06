package main

import "fmt"

//test to see if delete actualy deletes the value of the map key
func main(){
  num := 5
  m := make(map[string]int)
  numptr := &num
  fmt.Println(*numptr)
  m["test"] = num
  fmt.Println(m["test"])
  delete(m, "test")
  fmt.Println(*numptr)
  //It doesnt!! :)
}
