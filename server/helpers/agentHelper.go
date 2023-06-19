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
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
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

type Response struct {
	Data   model.Clusters
	Status string
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
	activationDetais, err := GetActivation()
	if err != nil {
		logger.Error("error getting activation info", err)
		return err
	}
	// activationId, _ := uuid.Parse(model.Activation_Id)
	//Validate activation details
	if activationDetais.ActivationID != model.Activation_Id && activationDetais.ActivationCode != model.Activation_Code {
		//panic server because no need of further execution
		panic("NO ACTIVATION AVAILABLE FOR PROVIDED DETAILS! PLEASE PROVIDE CORRECT ACTIVATION DETAILS")
	}
	//get activation ID from details and used as AgentActivationID
	activationNumber := activationDetais.ID
	goos := runtime.GOOS
	var getName any
	var getUserName any
	var getMachineId any
	var getPublicIp any
	var getHostName any
	var getPrivateIp any
	var getTimeZone any
	var getDisk any
	var getImageName any

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
		getPublicIp, _ = exec.Command("bash", "-c", "hostname -I").Output()
		getPrivateIp, _ = exec.Command("bash", "-c", "hostname -i").Output()
		getHostName, _ = exec.Command("bash", "-c", "hostname -A").Output()
		getMachineId, _ = exec.Command("bash", "-c", "cat /etc/machine-id").Output()

		getTimeZone, _ = exec.Command("bash", "-c", "date +'%Z %z'").Output()
		getDisk, _ = exec.Command("bash", "-c", "pwd").Output()

		getImageName, _ = exec.Command("bash", "-c", "cat /proc/cmdline").Output()
		getPublicIp = Getdata(strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)), "INET")
	}

	// Bind all data into the object
	// obj := model.InstanceInfo{strings.TrimSpace(fmt.Sprintf("%s", getName)), strings.TrimSpace(fmt.Sprintf("%s", getUserName)), strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
	// 	strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)), strings.TrimSpace(fmt.Sprintf("%s", getHostName)), goos, time.Now(), "Active", activationDetais.Id}

	//	1	36454d1c-aefc-4aa1-8411-eca0d221e1cb	i-00165452f0f332f5b	i-00165452f0f332f5b	i-00165452f0f332f5b	t2.micro	Non Hybrid						1	0		ap-south-1a	vpc-634bbc08	subnet-80e3ffe8	launch-wizard-1(sg-0cd76d8d5ceb54570)	/dev/xvda(vol-0f92a26a7be4f0165)	amzn2-ami-hvm-2.0.20210326.0-x86_64-gp2(ami-0bcf5425cdc1d8a85)	Linux	No		36					1	2	2	No	No	No	{"Categories":["General	Security Updates"],"Patches":[{"PatchName":"chrony	CurrentVersion":"3.5.1	Version":"4.0-3.amzn2.0.1	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"ec2-utils	CurrentVersion":"1.2	Version":"1.2-44.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"glibc	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"glibc-all-langpacks	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"glibc-common	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"glibc-locale-source	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"glibc-minimal-langpack	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"kernel	CurrentVersion":"4.14.231	Version":"4.14.232-176.381.amzn2	Category":"Security Updates	Severity":","DownloadSize":","PatchCurrentStatus":"Not Started	IsActive":true},{"PatchName":"kernel-tools	CurrentVersion":"4.14.231	Version":"4.14.232-176.381.amzn2	Category":"Security Updates	Severity":","DownloadSize":","PatchCurrentStatus":"Not Started	IsActive":true},{"PatchName":"libcrypt	CurrentVersion":"2.26	Version":"2.26-45.amzn2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"libjpeg-turbo	CurrentVersion":"1.2.90	Version":"2.0.90-2.amzn2.0.1	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true},{"PatchName":"python3	CurrentVersion":"3.7.9	Version":"3.7.9-1.amzn2.0.3	Category":"Security Updates	Severity":","DownloadSize":","PatchCurrentStatus":"Not Started	IsActive":true},{"PatchName":"python3-libs	CurrentVersion":"3.7.9	Version":"3.7.9-1.amzn2.0.3	Category":"Security Updates	Severity":","DownloadSize":","PatchCurrentStatus":"Not Started	IsActive":true},{"PatchName":"python3-pip	CurrentVersion":"9.0.3	Version":"20.2.2-1.amzn2.0.2	Category":"Security Updates	Severity":","DownloadSize":","PatchCurrentStatus":"Not Started	IsActive":true},{"PatchName":"python3-setuptools	CurrentVersion":"38.4.0	Version":"49.1.3-1.amzn2.0.2	Category":"General	Severity":","DownloadSize":","PatchCurrentStatus":","IsActive":true}]}		NULL	Yes	0			Yes			0	0000-00-00 00:00:00		0000-00-00 00:00:00			0		Yes	No	NULL	2021-05-07 16:53:56	No	Yes	No
	SerialID := uuid.New().String()
	serverInfo := model.Servers{
		SerialID,
		strings.TrimSpace(fmt.Sprintf("%s", getName)),
		strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
		strings.TrimSpace(fmt.Sprintf("%s", getMachineId)),
		goos,
		netIP,
		strings.TrimSpace(fmt.Sprintf("%s", getPublicIp)),
		strings.TrimSpace(fmt.Sprintf("%s", getPrivateIp)),
		strings.TrimSpace(fmt.Sprintf("%s", getPrivateIp)),
		strings.TrimSpace(fmt.Sprintf("%s", getTimeZone)),
		strings.TrimSpace(fmt.Sprintf("%s", getDisk)),
		strings.TrimSpace(fmt.Sprintf("%s", getImageName)),
		goos,
		strings.TrimSpace(fmt.Sprintf("%s", getUserName)),
		0,
		0,
		0,
		"", //MissingPatches
		"", //InstalledPatches
		"", //PatchDependenciesList
		0,
		"",         //AmiID
		"",         //AmiCreationDetail
		"",         //PatchCommandID
		"",         //InstallingPatches
		0,          //PatchInitiatedBy
		time.Now(), //PatchInstalledDate
		"",         //IntervalsEmailDateTime
		time.Now(), //PatchScannedDate
		strings.TrimSpace(fmt.Sprintf("%s", getHostName)),
		"", //ResourceGroup
		0,  //ResourceGroupID
		"", //SupportedAppsData
		time.Now(),
		activationNumber, //AgentActivationID
		time.Now(),
	}

	jsonReq, _ := json.Marshal(serverInfo)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Post(configHelper.GetString("ManagerURL")+"instance-info", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		logger.Error("Error in resistering agent", err)
		return err
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	server := model.UpdateServer{}

	str := "\"Agent Already Resistered\""
	if string(bodyBytes) == str {
		server.InstanceID = strings.TrimSpace(fmt.Sprintf("%s", getMachineId))
		server.NetIP = netIP
		UpdateNetworkIP(server)
	}
	log.Println(string(bodyBytes))
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

//Update server public ip on restart or network change
func UpdateNetworkIP(server model.UpdateServer) {
	logger.Info(server)
	client := http.Client{}
	jsonReq, _ := json.Marshal(server)
	_, err := client.Post(configHelper.GetString("ManagerURL")+"update-ip", "application/json; charset=utf-8", bytes.NewBuffer([]byte(jsonReq)))
	if err != nil {
		logger.Error("Error updating agent public IP", err)
	}
}
