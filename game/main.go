package main

import "fmt"

type Key byte

type Player struct {
	Name string
	Keys []Key
}

const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey
)

func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}
	return fmt.Sprintf("<Key %d", k)
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

func (p *Player) FoundKey(k Key) error {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("invalid key: %#v", k)
	}
	if !containsKey(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func main() {
	var p1 Player
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
}
