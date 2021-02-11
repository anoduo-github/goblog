package main

import (
	"fmt"
	"goblog/module/redis"
)

type user struct {
	Id   int
	Name string
}

func main() {
	redis.Init()
	u := new(user)
	u.Id = 1
	u.Name = "jack"
	err := redis.SetStruct("test", *u, 60)
	if err != nil {
		fmt.Println(err)
		return
	}
	id, err := redis.HGetInt("test", "Id")
	if err != nil {
		fmt.Println(err)
	}
	name, err := redis.HGetString("test", "Name")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)
	fmt.Println(name)
}
