package some

import (
	"strconv"
)

func GetSome() map[string]interface{} {
	result := make(map[string]interface{})

	type ls struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var list []ls
	for i := 0; i < 10; i++ {
		list = append(list, ls{
			Name: "haha" + strconv.Itoa(i),
			Age:  i,
		})
	}
	result["list"] = list

	return result
}
