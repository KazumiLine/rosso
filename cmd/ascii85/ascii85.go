package main

import (
   "fmt"
   "github.com/elt/rosso/ascii85"
   "os"
)

func main() {
   if len(os.Args) == 2 {
      name := os.Args[1]
      src, err := os.Open(name)
      if err != nil {
         panic(err)
      }
      defer src.Close()
      dst, err := os.Create(name + ".txt")
      if err != nil {
         panic(err)
      }
      defer dst.Close()
      if err := ascii85.Encode(dst, src); err != nil {
         panic(err)
      }
   } else {
      fmt.Println("ascii85 [file]")
   }
}