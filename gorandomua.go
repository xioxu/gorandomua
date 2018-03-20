package gorandomua

import (
	"encoding/json"
	"fmt"
	"os"
	"math/rand"
	"time"
)

type userAgent struct {
	Description string `json:"description"`
	UA string `json:"userAgent"`
	BrowserName string `json:"browserName"`
}

var uaCache []userAgent

func getData()  *[]userAgent {
	if uaCache == nil{
		rand.Seed(time.Now().UnixNano())
		json.Unmarshal([]byte(uaData), &uaCache)
	}

	return &uaCache
}

func guard(err error)  {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func GetRandom () string {
	getData()
	rdmIndex := rand.Intn(len(uaCache))
	return uaCache[rdmIndex].UA
}

func GetAll()  *[]string{
	getData()
	uas := make([]string,len(uaCache))

	for idx,val := range uaCache{
		uas[idx] = val.UA
	}

	return &uas
}
