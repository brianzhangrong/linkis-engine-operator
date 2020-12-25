package main

import (
	"fmt"
	"test/goclient"
	node "test/node"
)

var _ *node.Node = &node.Node{Data: 2}

func main() {

	// node.LNode
	// var n *node.Node
	// n = &node.Node{}
	// n.Data = 2
	// node.Test()
	//(*n).data = 1
	// goclient.Test()
	fmt.Println("-------------------------------")
	goclient.Test1()
	// a := "12345"
	// for _, i := range []rune(a) {
	// 	fmt.Println(i)
	// }
}
