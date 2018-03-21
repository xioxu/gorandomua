package gorandomua

import (
	"encoding/json"
	"math/rand"
	"time"
)

type userAgent struct {
	Description string `json:"description"`
	UA string `json:"userAgent"`
	BrowserName string `json:"browserName"`
}

var uaCache []userAgent

func init() {
	if uaCache == nil{
		rand.Seed(time.Now().UnixNano())
		json.Unmarshal([]byte(uaData), &uaCache)
	}
}

func GetRandom () string {
	rdmIndex := rand.Intn(len(uaCache))
	return uaCache[rdmIndex].UA
}

func GetAll()  *[]string{
	uas := make([]string,len(uaCache))

	for idx,val := range uaCache{
		uas[idx] = val.UA
	}

	return &uas
}
