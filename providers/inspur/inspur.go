package inspur

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	v1 "github.com/hezhaomin/hardwarem/api/v1"
	"github.com/hezhaomin/hardwarem/utils"
)

type Inspur struct {
	*v1.HardwareManagerData
}

func NewInspur(hm *v1.HardwareManagerData) (inspur *Inspur) {
	return &Inspur{hm}
}
func hasError(info string) bool {
	return strings.Contains(info, "Success:")
}

func (inspur *Inspur) GetRaid() (raids []string, err error) {
	raids = []string{}
	cmd := fmt.Sprintf("%s -H %s  -U %s -P %s", v1.InspurToolBin, inspur.MIP, inspur.UserName, inspur.Password)
	args := "getVirtualDrive"
	out, err := utils.ExecCmd(cmd, args)
	if err != nil {
		return raids, err
	}
	outStr := string(out)
	if !hasError(outStr) {
		return raids, errors.New(outStr)
	}
	controller := 0
	for _, raid := range strings.Split(outStr, "\n") {
		if strings.Contains(raid, "ControllerID") {
			controller, err = strconv.Atoi(strings.Split(raid, ":")[0])
			if err != nil {
				return raids, err
			}
			raids = append(raids, fmt.Sprintf("%d", controller))

		}
		if strings.HasPrefix(raid, "TargetID") {
			id := strings.Split(raid, ":")[0]
			raids[controller] = raids[controller] + ":" + id
		}
	}
	return raids, err

}

func (inspur *Inspur) CreateJob(job string) (err error) {
	return err

}
func (inspur *Inspur) InitRaid(raid string, speed string) (err error) {
	initType := "FI"
	if speed == "full" {
		initType = "SFI"
	}
	raidInfo := strings.Split(raid, ":")
	args := fmt.Sprintf("-H %s  -U %s -P %s setVirtualDrive -CID %s -VD %s -OP %s", inspur.MIP, inspur.UserName, inspur.Password, raidInfo[0], raidInfo[1], initType)
	out, err := utils.ExecCmd(inspur.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if !hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}
func (inspur *Inspur) DeleteRaid(raid string) (err error) {

	raidInfo := strings.Split(raid, ":")
	args := fmt.Sprintf("-H %s  -U %s -P %s setVirtualDrive -CID %s -VD %s -OP DEL", inspur.MIP, inspur.UserName, inspur.Password, raidInfo[0], raidInfo[1])
	out, err := utils.ExecCmd(inspur.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if !hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}

func (inspur *Inspur) Reboot() (err error) {
	args := fmt.Sprintf("-H %s  -U %s -P %s  setPowerStatus -T reset", inspur.MIP, inspur.UserName, inspur.Password)
	out, err := utils.ExecCmd(inspur.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if !hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}
