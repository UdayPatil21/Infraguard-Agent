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
	logger.Log.Info("IN:sendCommandService")
	var machineId []byte
	// Check machine ID
	machineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()

	if input.MachineID != strings.TrimSpace(string(machineId)) {
		logger.Log.Sugar().Errorf("Error: Machine Id mismatched")
		return nil, errors.New("machine id mismatched")
	}
	logger.Log.Info("Matched")
	cmd, err := exec.Command("bash", "-c", input.Command).Output()
	if err != nil {
		logger.Log.Sugar().Errorf("Error executing the command", err)
		return nil, err
	}
	logger.Log.Info("OUT:sendCommandService")
	return string(cmd), nil
}
func SanitizeScript(script string) string {
	s2 := strings.Replace(script, `\n`, "\n", -1)
	// s := strings.ReplaceAll(s2, "\\", "")
	return s2
}
func ExecuteScriptService(input model.Executable) (model.CmdOutput, error) {
	logger.Log.Info("IN:ExecuteScriptService")
	// var script string
	// json.Unmarshal(input.Script, &script)
	cmd := model.CmdOutput{}
	updatedString := SanitizeScript(input.Script)
	// convert updatedString into bytes
	// writeData, err := json.Marshal(updatedString)
	// if err != nil {
	// 	logger.Log.Sugar().Errorf("Error unmarshaling updated string to bytes", err)
	// 	return "", err
	// }
	//Create file and write script data
	fileName := time.Now().Format("01-02-2006") + ".sh"
	file, _ := os.Create(fileName)
	_, err := file.WriteString(updatedString)
	// err = os.WriteFile(fileName, writeData, 0777)
	if err != nil {
		logger.Log.Sugar().Errorf("Error saving script file on instance", err)
		return cmd, err
	}
	//change permissions of script file
	_, err = exec.Command("bash", "-c", "chmod "+model.Permissions+" "+fileName).Output()
	if err != nil {
		logger.Log.Sugar().Errorf("Error in change permissions", err)
		return cmd, err
	}
	//execute script file
	out, err := exec.Command("bash", "./"+fileName).Output()
	if err != nil {
		logger.Log.Sugar().Errorf("Error executing the script file", err)
		return cmd, err
	}
	logger.Log.Info("OUT:ExecuteScriptService")
	cmd.Output = string(out)
	return cmd, nil
}
