package v1

const (
	DellToolBin   = "/usr/sbin/racadm"
	InspurToolBin = "/usr/sbin/isrest"
)

type HardwareManagerData struct {
	UserName    string
	Password    string
	Model       string
	MIP         string
	ToolToolBin string
}
