package main

import (
	"testing"
	"github.com/jinzhu/configor"
	"fmt"
)

func TestJsonConfigor(t *testing.T){
	var curConfig Config


	err := configor.Load(&curConfig,".\\resources\\config.yml")


	if err!= nil{
		t.Error(err)
	}

	fmt.Println(curConfig)
}

type Config struct {
	ClientInfo []struct{
		ClientId string
		ClientSecret string
	}
	MgoDB struct{
		Conn string
	}
}