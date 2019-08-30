package main

//获取docker仓库的所有镜像

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type repositories struct {
	Repositories []string `json:repositories`
}

type imageTags struct {
	Name string   `json:name`
	Tags []string `json:tags`
}

func getImages(ip, port string) {
	url := "http://" + ip + ":" + port + "/v2/_catalog"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))

	var v repositories
	json.Unmarshal(body, &v)

	for _, images := range v.Repositories {
		getTags(ip, port, images)
	}
}

//get tags
func getTags(ip, port, images string) {
	url := "http://" + ip + ":" + port + "/v2/" + images + "/tags/list"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	// fmt.Println(string(body))

	v := imageTags{}
	json.Unmarshal(body, &v)
	for _, tag := range v.Tags {
		fmt.Printf("%s:%s/%s:%s\n", ip, port, v.Name, tag)
	}
}

func main() {
	getImages("127.0.0.1", "5000")
}
