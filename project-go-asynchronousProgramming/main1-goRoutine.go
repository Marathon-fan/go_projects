
package main

import "time"
import "runtime"

func main() {
  runtime.GOMAXPROCS(8)

  go abcGen()

  println("This comes first!")

  time.sleep(100 * time.Millisecond)
}

func abcGen() {
  for l := byte('a'); l <= byte('z'); l++ {
    go println(string(l))
  }
 }
