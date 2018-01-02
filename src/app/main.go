package main

import (
	"flag"
	"qiniupkg.com/x/log.v7"
	"io/ioutil"
	"encoding/json"
	"base"
	"router"
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		log.Println("need config file")
		return
	}
	f := args[0]
	c, err := load(f)
	if err != nil {
		log.Println("load config failed", err)
		return
	}
	router.RunNewRouter(c)
}

func load(fileName string) (*base.Config, error) {
	var config base.Config
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
