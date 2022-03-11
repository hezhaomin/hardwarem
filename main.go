package main

import (
	"fmt"

	v1 "github.com/hezhaomin/hardwarem/api/v1"
	"github.com/hezhaomin/hardwarem/providers"
)

func main() {
	hardware := v1.HardwareManagerData{
		UserName:    "root",
		Password:    "4ZGkhuHn@1",
		Model:       "Inspur",
		MIP:         "10.254.4.136",
		ToolToolBin: v1.InspurToolBin,
	}
	provider, err := providers.NewHarwareManager(&hardware)
	if err != nil {

		fmt.Println(err)
		return
	}
	raids, err := provider.GetRaid()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(raids)
}
