package dell

import (
	"errors"
	"fmt"
	"strings"

	v1 "github.com/hezhaomin/hardwarem/api/v1"
	"github.com/hezhaomin/hardwarem/utils"
)

type Dell struct {
	*v1.HardwareManagerData
}

func NewDell(hm *v1.HardwareManagerData) (dell *Dell) {
	hm.ToolToolBin = v1.DellToolBin
	return &Dell{hm}
}
func hasError(info string) bool {
	return strings.Contains(info, "ERROR:")
}

func (dell *Dell) GetRaid() (raids []string, err error) {
	raids = []string{}
	args := fmt.Sprintf("-r %s  -u %s -p %s storage get vdisks", dell.MIP, dell.UserName, dell.Password)
	out, err := utils.ExecCmd(dell.ToolToolBin, args)
	if err != nil {
		return raids, err
	}
	outStr := string(out)
	if hasError(outStr) {
		return raids, errors.New(outStr)
	}
	result := strings.Split(string(out), "\n")

	for _, raid := range result {
		if strings.HasPrefix(raid, "Disk.Virtual") {
			raids = append(raids, raid)
		}
	}
	return raids, nil

}

func (dell *Dell) CreateJob(job string) (err error) {
	args := fmt.Sprintf("-r %s  -u %s -p %s jobqueue create %s", dell.MIP, dell.UserName, dell.Password, job)
	out, err := utils.ExecCmd(dell.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}
func (dell *Dell) InitRaid(raid string, speed string) (err error) {
	args := fmt.Sprintf("-r %s  -u %s -p %s init:%s -speed %s", dell.MIP, dell.UserName, dell.Password, raid, speed)
	out, err := utils.ExecCmd(dell.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}
func (dell *Dell) DeleteRaid(raid string) (err error) {
	args := fmt.Sprintf("-r %s  -u %s -p %s deletevd:%s", dell.MIP, dell.UserName, dell.Password, raid)
	out, err := utils.ExecCmd(dell.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}

func (dell *Dell) Reboot() (err error) {
	args := fmt.Sprintf("-r %s  -u %s -p %s serveraction hardreset", dell.MIP, dell.UserName, dell.Password)
	out, err := utils.ExecCmd(dell.ToolToolBin, args)
	if err != nil {
		return err
	}
	outStr := string(out)
	if hasError(outStr) {
		return errors.New(outStr)
	}

	return err

}
