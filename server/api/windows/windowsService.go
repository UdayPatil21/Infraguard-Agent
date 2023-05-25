package windows

import (
	"errors"
	helper "infraguard-agent/helpers"
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"os/exec"
	"strings"
)

func sendCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sendCommandService")
	var getMachineId []byte
	// Check machine ID - MAC ID for windows
	getMachineId, _ = exec.Command("cmd", "/C", "wmic NICCONFIG WHERE IPEnabled=true GET MACAddress").Output()
	machineId := helper.Getdata(string(getMachineId), "MAC")

	// machineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()
	if input.MachineID != strings.TrimSpace(string(machineId)) {
		logger.Error("Error: Machine Id mismatched")
		return nil, errors.New("machine id mismatched")
	}
	logger.Info("Matched")
	cmd, err := exec.Command("cmd", "/C", input.Command).Output()
	if err != nil {
		logger.Error("Error executing the command", err)
		return nil, err
	}
	logger.Info("OUT:sendCommandService")
	return string(cmd), nil
}
