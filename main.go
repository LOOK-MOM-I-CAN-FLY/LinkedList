package main

import (
	"fmt"
	linkedList "list/package"
)

func main() {
	ll := &linkedList.List[int]{Value: 100}
	ll.Append(110)
	ll.Append(140)
	ll.Append(120)
	ll.Append(109)
	fmt.Println(ll)
}

