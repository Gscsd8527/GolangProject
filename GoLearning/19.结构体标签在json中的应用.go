package main

import (
	"encoding/json"
	"fmt"
)

// json指的是在转换json的时候的key
type Movie struct {
	Title string `json:"title"`
	Year int `json:"year"`
	Price int `json:"price"`
	Axtors []string `json:"axtors"`
}


func main() {
	movie := Movie{"战狼", 2015, 5, []string{"吴京", "倪大红"}}

	// 编码的过程  结构体  ---》 json
	jsonStr, err := json.Marshal(movie)
	if err != nil{
		fmt.Println("json marshal error", err)
		return
	}
	fmt.Printf("jsonStr = %s\n", jsonStr)

	// 解码的过程 jsonstr  ---》  结构体
	//jsonStr = {"title":"战狼","year":2015,"price":5,"axtors":["吴京","倪大红"]}
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil{
		fmt.Println("json numarshl error", err)
		return
	}
	fmt.Printf("%v \n", my_movie)


}
