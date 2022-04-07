package providers

import (
	"errors"

	v1 "github.com/hezhaomin/hardwarem/api/v1"
	"github.com/hezhaomin/hardwarem/providers/dell"
	"github.com/hezhaomin/hardwarem/providers/inspur"
)

type HardwareManager interface {
	InitRaid(raid string, speed string) error
	DeleteRaid(controller string, raid string) error
	Reboot() error
	CreateJob(job string) error
	GetRaid() (raids []string, err error)
}

func NewHarwareManager(hm *v1.HardwareManagerData) (provider HardwareManager, err error) {
	model := hm.Model
	switch model {
	case "dell":
		return dell.NewDell(hm), nil
	case "inspur":
		return inspur.NewInspur(hm), nil
	default:
		return provider, errors.New("not spuort hardware")
	}

}
