package main

import "fmt"

//func main() {
//	buffered_channel := make(chan string, 2)
//	buffered_channel <- "foo"
//	buffered_channel <- "bar"
//	//buffered_channel <- "deadlock: all goroutines are asleep"
//
//	fmt.Println("channel length after add: ", len(buffered_channel))
//	fmt.Println("channel capacity after add: ", cap(buffered_channel))
//
//	fmt.Println(<-buffered_channel)
//	fmt.Println(<-buffered_channel)
//
//	fmt.Println("channel length after Pop: ", len(buffered_channel))
//	fmt.Println("channel capacity after Pop: ", cap(buffered_channel))
//
//	buffered_channel <- "baz"
//	out := <-buffered_channel
//	fmt.Println(out)
//	close(buffered_channel)
//	//buffered_channel <- "send to closed channel"
//}

func main() {
	bufferedchannel := make(chan string, 3)
	bufferedchannel <- "1"
	bufferedchannel <- "2"
	bufferedchannel <- "3"
	close(bufferedchannel)
	for i := range bufferedchannel {
		fmt.Println(i)
	}
}
