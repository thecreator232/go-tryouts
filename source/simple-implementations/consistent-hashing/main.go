package main

import (
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Node struct {
	server string
	hash   int
}

type HashRing struct {
	nodes           []Node
	vnodesPerServer int
}

func NewHashRing(vnodesPerServer int) *HashRing {
	return &HashRing{
		nodes:           []Node{},
		vnodesPerServer: vnodesPerServer,
	}
}

// func (h HashRing) Len() int {
// 	return len(h.node)
// }

// func (h HashRing) Less(i, j int) bool {
// 	return h.node[i].hash < h.node[j].hash
// }

// func (h HashRing) Swap(i, j int) {
// 	h.node[i], h.node[j] = h.node[j], h.node[i]
// 	// g := *h
// 	// g[i], g[j] = g[j], g[i]
// }

func (h *HashRing) AddServer(server string) {

	for i := 0; i < h.vnodesPerServer; i++ {
		vnode := server + "#" + strconv.Itoa(i)
		hash := int(crc32.ChecksumIEEE([]byte(vnode)))
		h.nodes = append(h.nodes, Node{server: server, hash: hash})
	}
	sort.Slice(h.nodes, func(i, j int) bool {
		return h.nodes[i].hash < h.nodes[j].hash
	})
}

func (h *HashRing) RemoveServer(server string) {
	newNodes := []Node{}

	for _, node := range h.nodes {
		if node.server != server {
			newNodes = append(newNodes, node)
		}
	}
	h.nodes = newNodes
}

func (h HashRing) GetServer(key string) string {

	if len(h.nodes) == 0 {
		return ""
	}
	hash := int(crc32.ChecksumIEEE([]byte(key)))

	idx := sort.Search(len(h.nodes), func(i int) bool {
		return h.nodes[i].hash >= hash
	})

	if idx == len(h.nodes) {
		idx = 0
	}

	return h.nodes[idx].server
}

func main() {
	hh := NewHashRing(4)
	hh.AddServer("127.0.0.1")
	hh.AddServer("www.google.com")

	fmt.Println("testing123 := ", hh.GetServer("testing123"))
	fmt.Println("testing456 := ", hh.GetServer("testing456"))
	fmt.Println("testing789 := ", hh.GetServer("testing789"))
	fmt.Println("testing1011 := ", hh.GetServer("testing1011"))
	fmt.Println("testing1213 := ", hh.GetServer("testing1213"))

	hh.AddServer("1.1.1.1")

	fmt.Println("after adding a new server ------- \n\n")

	fmt.Println("testing123 := ", hh.GetServer("testing123"))
	fmt.Println("testing456 := ", hh.GetServer("testing456"))
	fmt.Println("testing789 := ", hh.GetServer("testing789"))
	fmt.Println("testing1011 := ", hh.GetServer("testing1011"))
	fmt.Println("testing1213 := ", hh.GetServer("testing1213"))

	hh.RemoveServer("127.0.0.1")

	fmt.Println("after removing a new server ------- \n\n")

	fmt.Println("testing123 := ", hh.GetServer("testing123"))
	fmt.Println("testing456 := ", hh.GetServer("testing456"))
	fmt.Println("testing789 := ", hh.GetServer("testing789"))
	fmt.Println("testing1011 := ", hh.GetServer("testing1011"))
	fmt.Println("testing1213 := ", hh.GetServer("testing1213"))

}
