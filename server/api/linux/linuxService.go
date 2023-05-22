package linux

import (
	"errors"
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"os"
	"os/exec"
	"strings"
	"time"
)

func sendCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sendCommandService")
	var machineId []byte
	// Check machine ID
	machineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()

	if input.MachineID != strings.TrimSpace(string(machineId)) {
		logger.Error("Error: Machine Id mismatched")
		return nil, errors.New("machine id mismatched")
	}
	logger.Info("Matched")
	cmd, err := exec.Command("bash", "-c", input.Command).Output()
	if err != nil {
		logger.Error("Error executing the command", err)
		return nil, err
	}
	logger.Info("OUT:sendCommandService")
	return string(cmd), nil
}

func executeScriptService(input model.Executable) (any, error) {
	logger.Info("IN:executeScriptService")

	//Create file and write script data
	fileName := time.Now().Format("01-02-2006") + ".sh"
	err := os.WriteFile(fileName, input.Script, 0777)
	if err != nil {
		logger.Error("Error saving script file on instance", err)
		return nil, err
	}
	//change permissions of script file
	_, err = exec.Command("bash", "-c", "chmod "+input.Permission+" "+fileName).Output()
	if err != nil {
		logger.Error("Error in change permissions", err)
		return nil, err
	}
	//execute script file
	_, err = exec.Command("bash", "-c", "./"+fileName).Output()
	if err != nil {
		logger.Error("Error executing the script file", err)
		return nil, err
	}
	logger.Info("OUT:executeScriptService")
	return "Script file executed successfully", nil
}

func sudoCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sudoCommandService")
	// var machineId []byte
	// Check machine ID
	// machineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()

	// if input.MachineID != strings.TrimSpace(string(machineId)) {
	// 	logger.Error("Error: Machine Id mismatched")
	// 	return nil, errors.New("machine id mismatched")
	// }
	// logger.Info("Matched")
	//sudo command
	cmd, err := exec.Command("echo", "Digi@2023 ", " |", "sudo", "-S", "-k", input.Command).Output()
	if err != nil {
		return "", err
	}
	logger.Info("OUT:sudoCommandService")
	return string(cmd), nil
}
