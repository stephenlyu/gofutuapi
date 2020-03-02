package futuapi

import "github.com/stephenlyu/gofutuapi"

func Init() {
	gofutuapi.FTAPIChannel_Init()
}

func UnInit() {
	gofutuapi.FTAPIChannel_UnInit()
}
