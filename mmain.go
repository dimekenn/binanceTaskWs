package main

import (
	"fmt"
	"log"
	"strconv"
)

func main()  {
	f := "54899.980000"
	i, err := strconv.ParseFloat(f, 64)
	if err !=nil{
		log.Fatal(err)
	}
	fmt.Println(i)
}
