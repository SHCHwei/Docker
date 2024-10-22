package main

import (
    "github.com/gin-gonic/gin"
    "fmt"
)

func main(){

    r := gin.Default()

    r.POST("/ping", test)
    r.GET("/ping2", runs)
    r.Run(":7777")
}

type Person struct {
	Name    string `form:"name" binding:"required"`
	Address string `form:"address" binding:"required"`
}


func test(c *gin.Context){

	var person Person
	if c.ShouldBind(&person) == nil {
		fmt.Println("====== Only Bind By Query String ======")
		fmt.Println(person.Name)
		fmt.Println(person.Address)
	}else{
	    fmt.Println("====== Sorry, it's nil ======")
	}
	c.String(200, "Success")

}

func runs(c *gin.Context){
    fmt.Println("hello world")
}