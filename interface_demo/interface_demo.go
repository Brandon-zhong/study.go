package main

import "sort"

func main() {
	var l = []User{User{id:1, name: "fkasd"}, User{id: 2, name: "dfasd"}}
	sort.Sort(list(l))

}

type User struct {
	id   int32
	name string
}

type list []User

func (u list) Len() int {
	return len(u)
}

func (u list) Less(i, j int) bool {
	return u[i].id > u[j].id
}
func (u list) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

type array [] User


