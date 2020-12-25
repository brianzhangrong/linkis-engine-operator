package node

import "fmt"

type Node struct {
	Data int
	next *Node
}

func (node *Node) Reverse(){
	if Header.next ==nil || Header.next.next==nil {
		return 
	}
	pre:= Header.next
	current :=pre.next
	var next *Node
	pre.next=nil
	
	for current.next!=nil{
		next = current.next
		current.next=pre
		pre=current
		current=next
	}
	Header.next=current
	current.next=pre
}
func (n *Node) PrintAll(){
	p := Header
	for{	
		if p.next== nil {
			fmt.Println("===",p.Data)
			break
		}
		if p != Header{
			fmt.Println("===",p.Data)
		}
		p=p.next
	}
}
var Header *Node=&Node{-1,nil}

func init() {
	nums := []int{}
	var cur *Node =Header
   	for _,i := range nums{
		var node *Node= &Node{i,nil}
		cur.next=node
		cur = node
	}
	cur.PrintAll()
	cur.Reverse()
	fmt.Println("-----")
	cur.PrintAll()
}
