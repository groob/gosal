package reports

import (
	"encoding/json"
	"os/exec"
)

func Get_win32_bios() (win32_bios, error) {
	cmd := exec.Command("powershell", "gwmi", "win32_bios", "|", "convertto-json")

	o, err := cmd.Output()
	if err != nil {
		return win32_bios{}, err
	}

	var j win32_bios

	if err := json.Unmarshal(o, &j); err != nil {
		return win32_bios{}, err
	}

	return j, nil
}

type win32_bios struct {
	PSComputerName string `json:"PSComputerName"`
	SerialNumber   string `json:"SerialNumber"`
}
