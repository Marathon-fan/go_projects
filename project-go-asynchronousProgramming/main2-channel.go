
// sending and receiving messages(construct a channel and send message to it)
//
// ranging over channels(for range)


package main

import "time"
import "runtime"

func main() {
  runtime.GOMAXPROCS(8)

  ch := make(chan string) // create a channel
  doneCh := make(chan bool)

  go abcGen(ch)
  go printer(ch, doneCh)


  go abcGen()

  println("This comes first!")

  <- doneCh
}

func abcGen(ch chan string) {
  for l := byte('a'); l <= byte('z'); l++ {
    ch <- string(l)
  }
  close(ch)
 }

 func printer(ch chan string, doneCh chan bool) {
   for 1: range ch {
     println(l)
   }
   doneCh <- true
 }
