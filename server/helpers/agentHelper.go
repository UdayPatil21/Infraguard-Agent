package helper

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"infraguard-agent/helpers/configHelper"
	"infraguard-agent/helpers/logger"
	model "infraguard-agent/models"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

func GetNetworkIP() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	// Defer resp. Body.close )
	content, _ := ioutil.ReadAll(resp.Body)
	//buf: = new (bytes. Buffer)
	//buf. Readfrom (resp. Body)
	//s: = buf. String ()
	return string(content)
}

// type Response struct {
// 	Data   model.Clusters
// 	Status string
// }
type Response struct {
	Data   model.Clusters
	Status bool
	Error  any
}

func GetActivation() (model.Clusters, error) {

	activation := model.Clusters{}
	res := Response{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(configHelper.GetString("ManagerURL") + "getAgentActivation" + model.Activation_Id)
	if err != nil {
		logger.Error("Error getting activation data by name", err)
		return activation, err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(bodyBytes, &res)
	return res.Data, nil
}
func PreCheck() error {

	// Make it config based
	// managerUrl := "http://localhost:4200/api/instance-info"

	//Check your network ip address
	netIP := GetNetworkIP()
	//Get activation details from manager
	// activationDetais, err := GetActivation()
	// if err != nil {
	// 	logger.Error("error getting activation info", err)
	// 	return err
	// }
	// activationId, _ := uuid.Parse(model.Activation_Id)
	//Validate activation details
	// if activationDetais.ActivationID != model.Activation_Id && activationDetais.ActivationCode != model.Activation_Code {
	// 	logger.Error("NO ACTIVATION AVAILABLE FOR PROVIDED DETAILS! PLEASE PROVIDE CORRECT ACTIVATION DETAILS")
	// 	logger.Error(activationDetais.ActivationID, "model"+model.Activation_Id)
	// 	logger.Error(activationDetais.ActivationCode, "model"+model.Activation_Code)
	// 	//panic server because no need of further execution
	// 	panic("NO ACTIVATION AVAILABLE FOR PROVIDED DETAILS! PLEASE PROVIDE CORRECT ACTIVATION DETAILS")
	// }
	//get activation ID from details and used as AgentActivationID
	// activationNumber := activationDetais.ID
	goos := runtime.GOOS
	// var getName any
	var getUserName any
	var getMachineId any
	var getPublicIp any
	var getHostName any
	var getPrivateIp any
	// var getTimeZone any
	// var getDisk any
	var getImageName any

	if goos == "darwin" {
		println("System is Windows")
		// getName, _ = exec.Command("bash", "-c", "id -F").Output()
		getUserName, _ = exec.Command("bash", "-c", "id -un").Output()
		getPublicIp, _ = exec.Command("bash", "-c", "curl ifconfig.me && echo").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname").Output()
		getMachineId, _ = exec.Command("bash", "-c", "ioreg -rd1 -c IOPlatformExpertDevice | awk '/IOPlatformUUID/ { split($0, line, \"\\\"\"); printf(\"%s\\n\", line[4]); }'").Output()
	}
	if goos == "windows" {
		println("System is Windows")
		// fmt.Println("Hello World")
		// getName, _ = exec.Command("cmd", "/C", "hostname").Output()
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
		// getName, _ = exec.Command("bash", "-c", "cat /proc/sys/kernel/hostname").Output()
		getUserName, _ = exec.Command("bash", "-c", "whoami").Output()
		getPublicIp, _ = exec.Command("bash", "-c", "hostname -I").Output()
		getPrivateIp, _ = exec.Command("bash", "-c", "hostname -i").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname -A").Output()
		getMachineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()

		// getTimeZone, _ = exec.Command("bash", "-c", "date +'%Z %z'").Output()
		// getDisk, _ = exec.Command("bash", "-c", "pwd").Output()

		getImageName, _ = exec.Command("bash", "-c", "cat /proc/cmdline").Output()
		getPublicIp = Getdata(strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)), "INET")
	}

	// SerialID := uuid.New().String()
	serverInfo := model.Servers{
		// SerialID,
		// strings.TrimSpace(fmt.Sprintf("%s", getName)),
		strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
		// strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
		// goos,
		netIP,
		// strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)),
		strings.TrimSpace(fmt.Sprintf("%s", getPrivateIp)),
		// strings.TrimSpace(fmt.Sprintf("%s", getPrivateIp)),
		// strings.TrimSpace(fmt.Sprintf("%s", getTimeZone)),
		// strings.TrimSpace(fmt.Sprintf("%s", getDisk)),
		strings.TrimSpace(fmt.Sprintf("%s", getImageName)),
		strings.Title(goos),
		strings.TrimSpace(fmt.Sprintf("%s", getUserName)),
		// 0,
		// 0,
		// 0,
		// "", //MissingPatches
		// "", //InstalledPatches
		// "", //PatchDependenciesList
		// 0,
		// "",         //AmiID
		// "",         //AmiCreationDetail
		// "",         //PatchCommandID
		// "",         //InstallingPatches
		// 0,          //PatchInitiatedBy
		// time.Now(), //PatchInstalledDate
		// "",         //IntervalsEmailDateTime
		// time.Now(), //PatchScannedDate
		strings.TrimSpace(fmt.Sprintf("%s", getHostName)),
		// "", //ResourceGroup
		// 0,  //ResourceGroupID
		// "", //SupportedAppsData
		// time.Now(),
		// activationNumber, //AgentActivationID
		model.Activation_Id,
		// "a49eaf2a-04a2-454e-9e11-6d1c7a37fcf6",
		// "53616d706c652041637469766174696f6e",
		model.Activation_Code,
		// time.Now(),
	}

	jsonReq, _ := json.Marshal(serverInfo)
	out := model.Response{}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(configHelper.GetString("ManagerURL")+"registration/serverinfo", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Error("Error in resistering agent", err)
		return err
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//Get output data from the response
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		logger.Error("Error Unmarshaling Data", err)
		return err
	}
	logger.Info(out)
	str := "Agent Already Resistered"
	if out.Status && out.Data == str {
		err := UpdateAgentInfo(serverInfo)
		if err != nil {
			logger.Error("Error Updating Server Data", err)
			return err
		}
	}
	return nil
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
	case "INET":
		res := strings.Split(str, " ")
		result = res[0]
	default:
		fmt.Println("default")
	}
	return result
}

//Update server info  on restart or network change
func UpdateAgentInfo(server model.Servers) error {
	logger.Info("IN:UpdateAgentInfo")
	logger.Info(server)
	client := http.Client{}
	out := model.Response{}
	jsonReq, _ := json.Marshal(server)
	resp, err := client.Post(configHelper.GetString("ManagerURL")+"update/serverinfo", "application/json; charset=utf-8", bytes.NewBuffer([]byte(jsonReq)))
	if err != nil {
		logger.Error("Error updating agent public IP", err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//Get output data from the response
	err = json.Unmarshal(bodyBytes, &out)
	if err != nil {
		logger.Error("Error Unmarshaling Data", err)
		return err
	}
	if !out.Status {
		logger.Error("Error Updating Server Data")
		return err
	}
	return nil
}
