package helper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"infraguard-agent/helpers/configHelper"
	model "infraguard-agent/models"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func PreCheck() {

	// Make it config based
	// managerUrl := "http://localhost:4200/api/instance-info"

	goos := runtime.GOOS
	var getName any
	var getUserName any
	var getMachineId any
	var getPublicIp any
	var getHostName any

	if goos == "darwin" {
		println("System is Windows")
		getName, _ = exec.Command("bash", "-c", "id -F").Output()
		getUserName, _ = exec.Command("bash", "-c", "id -un").Output()
		getPublicIp, _ = exec.Command("bash", "-c", "curl ifconfig.me && echo").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname").Output()
		getMachineId, _ = exec.Command("bash", "-c", "ioreg -rd1 -c IOPlatformExpertDevice | awk '/IOPlatformUUID/ { split($0, line, \"\\\"\"); printf(\"%s\\n\", line[4]); }'").Output()
	}
	if goos == "windows" {
		println("System is Windows")
		// fmt.Println("Hello World")
		getName, _ = exec.Command("cmd", "/C", "hostname").Output()
		getUserName, _ = exec.Command("cmd", "/C", "whoami").Output()

		//ipconfig | findstr /r /c:"IPv4"
		out, _ := exec.Command("cmd", "/C", "ipconfig | findstr /r /c:IPv4").Output()
		getPublicIp = Getdata(string(out), "IP")
		getHostName, _ = exec.Command("cmd", "/C", "hostname").Output()

		//wmic NICCONFIG WHERE IPEnabled=true GET MACAddres
		out1, _ := exec.Command("cmd", "/C", "wmic NICCONFIG WHERE IPEnabled=true GET MACAddress").Output()
		getMachineId = Getdata(string(out1), "MAC")

	}
	if goos == "linux" {
		println("System is linux")
		getName, _ = exec.Command("bash", "-c", "cat /proc/sys/kernel/hostname").Output()
		getUserName, _ = exec.Command("bash", "-c", "whoami").Output()
		getPublicIp, _ = exec.Command("bash", "-c", "curl ifconfig.me && echo").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname -A").Output()
		getMachineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()
	}

	// Bind all data into the object
	obj := model.InstanceInfo{strings.TrimSpace(fmt.Sprintf("%s", getName)), strings.TrimSpace(fmt.Sprintf("%s", getUserName)), strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
		strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)), strings.TrimSpace(fmt.Sprintf("%s", getHostName)), goos, time.Now(), "Active"}

	jsonReq, _ := json.Marshal(obj)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(configHelper.GetString("ManagerURL"), "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Print(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(bodyBytes))
}
func Getdata(str, flag string) string {
	var result string
	switch flag {
	case "IP":
		res := strings.Split(str, ":")
		result = res[1]
	case "MAC":
		res := strings.Split(str, "\n")
		result = res[1]
	default:
		fmt.Println("default")
	}
	return result
}
