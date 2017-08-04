package main

import (
	"flag"
	"log"
	"once/once"
)

var filename string

func init() {
	flag.StringVar(&filename, "conf", "once.conf", "File path to configuration file")
	flag.Parse()
}

func main(){
	configuration, err := NewConfiguration(filename)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(configuration.OnceConfiguration.Domain)
	log.Println(*configuration.OnceConfiguration.RedisConf)

	err = once.InitOnce(configuration.OnceConfiguration)
	if err != nil {
		log.Fatal(err.Error())
	}
}
