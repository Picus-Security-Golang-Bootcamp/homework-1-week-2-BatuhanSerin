package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	data := map[string]interface{}{}
	contents, err := ioutil.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(contents, &data); err != nil {
		panic(err)
	}

	names := data["psychological thriller movies"].(map[string]interface{})

	if os.Args[1] == "list" {
		fmt.Println("**************Movie List**************")
		for _, movieNames := range names {

			movieNames := fmt.Sprintf("%v", movieNames)

			fmt.Println(movieNames)

		}
	} else if os.Args[1] == "search" {
		args := ""

		for i := 2; i < len(os.Args); i++ {
			args = args + " " + string(os.Args[i])
		}
		args = strings.Title(strings.ToLower(args))[1:]

		flag := []bool{true}
		for _, movieNames := range names {

			movieNames := fmt.Sprintf("%v", movieNames)

			if movieNames == args {
				fmt.Printf("The movie is found: %s", movieNames)
				flag[0] = false
				break
			}

		}
		if flag[0] {
			fmt.Println("The movie is not found!")
		}
	}

}
