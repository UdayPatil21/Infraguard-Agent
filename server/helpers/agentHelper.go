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
		getName, _ = exec.Command("bash", "-c", "id -F").Output()
		getUserName, _ = exec.Command("bash", "-c", "id -un").Output()
		getPublicIp, _ = exec.Command("bash", "-c", "curl ifconfig.me && echo").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname").Output()
		getMachineId, _ = exec.Command("bash", "-c", "ioreg -rd1 -c IOPlatformExpertDevice | awk '/IOPlatformUUID/ { split($0, line, \"\\\"\"); printf(\"%s\\n\", line[4]); }'").Output()
	}
	if goos == "windows" {
		fmt.Println("Hello from Windows")
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
		strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)), strings.TrimSpace(fmt.Sprintf("%s", getHostName)), goos, time.Now()}

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
