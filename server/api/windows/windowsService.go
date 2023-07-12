package windows

import (
	"errors"
	"fmt"
	"infraguard-agent/api/linux"
	helper "infraguard-agent/helpers"
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"os"
	"os/exec"
	"strings"
	"time"
)

func sendCommandService(input model.RunCommand) (any, error) {
	logger.Info("IN:sendCommandService")
	var getMachineId []byte
	// Check machine ID - MAC ID for windows
	getMachineId, _ = exec.Command("cmd", "/C", "wmic NICCONFIG WHERE IPEnabled=true GET MACAddress").Output()
	machineId := helper.Getdata(string(getMachineId), "MAC")

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

func ExecuteScriptService(input model.Executable) (model.CmdOutput, error) {
	logger.Info("IN:ExecuteScriptService")
	// var getMachineId []byte
	CmdOutput := model.CmdOutput{}
	// Check machine ID - MAC ID for windows
	// getMachineId, _ = exec.Command("cmd", "/C", "wmic NICCONFIG WHERE IPEnabled=true GET MACAddress").Output()
	// machineId := helper.Getdata(string(getMachineId), "MAC")

	// if input.MachineID != strings.TrimSpace(string(machineId)) {
	// 	logger.Error("Error: Machine Id mismatched")
	// 	return CmdOutput, errors.New("machine id cannot matched with server machine id")
	// }
	updatedString := linux.SanitizeScript(input.Script)

	//Create file and write script data
	fileName := time.Now().Format("01-02-2006") + ".ps1"
	file, _ := os.Create(fileName)
	_, err := file.WriteString(updatedString)
	// err = os.WriteFile(fileName, writeData, 0777)
	if err != nil {
		logger.Error("Error saving script file on instance", err)
		return CmdOutput, err
	}
	//change permissions of script file
	// _, err = exec.Command("cmd", "/C", "chmod "+model.Permissions+" "+fileName).Output()
	// if err != nil {
	// 	logger.Error("Error in change permissions", err)
	// 	return cmd, err
	// }

	// Sometimes windows give error
	//running script is disabled in this system
	//Perform following command

	// Set-ExecutionPolicy RemoteSigned

	//execute script file
	cmd := exec.Command("powershell", "./"+fileName)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		logger.Error("Error executing the script file", err)
		return CmdOutput, err
	}
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, "New-Item a.txt")
		fmt.Fprintln(stdin, "New-Item b.txt")
	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		logger.Error("Error getting command output", err)
		return CmdOutput, err
	}
	logger.Info("OUT:ExecuteScriptService")
	CmdOutput.Output = string(out)
	return CmdOutput, nil
}
