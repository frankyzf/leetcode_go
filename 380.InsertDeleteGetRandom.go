package main

import "math/rand"

type RandomizedSet struct {
	con  map[int]int
	size int
}

func Constructor() RandomizedSet {
	p := RandomizedSet{
		con:  map[int]int{},
		size: 0,
	}
	return p
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.con[val]; ok {
		return false
	} else {
		this.con[val] = 1
		this.size++
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.con[val]; !ok {
		return false
	}
	delete(this.con, val)
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if this.size == 0 {
		return 0
	}
	n := rand.Intn(this.size)
	keys := make([]int, this.size)
	i := 0
	for k := range this.con {
		keys[i] = k
		i++
	}
	return keys[n]
}
