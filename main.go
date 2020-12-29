package main

import (
	"fmt"
	"linkis-engine-operator/engineinformers"
	"os"
	"os/signal"
	"syscall"
	// "google.golang.org/appengine/log"
	// node "linkis-engine-operator/node"
)

// var _ *node.Node = &node.Node{Data: 2}

func main() {

	// node.LNode
	// var n *node.Node
	// n = &node.Node{}
	// n.Data = 2
	// node.Test()
	//(*n).data = 1
	// goclient.Test()
	fmt.Println("-------------------------------")
	// shellintercept.CreateShellEnv()
	engineinformers.TestInformer()
	fmt.Println("-------------------------------")

	// goclient.Test1()
	// a := "12345"
	// for _, i := range []rune(a) {
	// 	fmt.Println(i)
	// }
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGUSR1)
	for {
		s := <-ch
		switch s {
		case syscall.SIGQUIT:
			fmt.Println("SIGSTOP")
			return
		case syscall.SIGSTOP:
			fmt.Println("SIGSTOP")
			return
		case syscall.SIGHUP:
			fmt.Println("SIGHUP")
			return
		case syscall.SIGKILL:
			fmt.Println("SIGKILL")
			return
		case syscall.SIGUSR1:
			fmt.Println("SIGUSR1")
			return
		default:
			fmt.Println("default")
			return
		}
	}
}
