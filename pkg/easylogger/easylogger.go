package easylogger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

func LogError(err error, msg string) {
	if nil != err {
		//log.Println(err, msg)
		log.Fatal(err, msg)
	}
}

func Print(v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println(v)
		return
	}

	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		fmt.Println(v)
		return
	}

	fmt.Println(out.String())
}
