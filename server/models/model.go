package model

import (
	"time"
)

var (
	Activation_Id   string
	Activation_Code string
)

type InstanceInfo struct {
	Name              string    `json:"name"`
	UserName          string    `json:"user_name"`
	MachineID         string    `json:"machine_id"`
	PublicIP          string    `json:"public_ip"`
	HostName          string    `json:"host_name"`
	OS                string    `json:"os_name"`
	CreatedAt         time.Time `json:"createdAt"`
	Status            string    `json:"status"`
	Activation_Number int       `json:"activation_number"`
}

type RunCommand struct {
	MachineID string `json:"machine_id" gorm:"machine_id" validate:"required"`
	Command   string `json:"command" validate:"required"`
}

type Executable struct {
	Script     []byte
	Permission string
}

// type Activation struct {
// 	Id               int       `json:"id" gorm:"id"`
// 	Serial_Id        uuid.UUID `json:"serial_id" gorm:"serial_id"`
// 	Activation_Id    uuid.UUID `json:"activation_id" gorm:"activation_id"`
// 	Actiovation_Code string    `json:"activation_code" gorm:"activation_code"`
// 	Activation_Name  string    `json:"activation_name" gorm:"activation_name"`
// 	Server_Count     int       `json:"number_of_servers" gorm:"number_of_servers"`
// 	Created_By       string    `json:"created_by" gorm:"created_by"`
// 	Created_Date     time.Time `json:"created_date" gorm:"created_date"`
// }

type AgentActivations struct {
	ID                   int       `gorm:"column:ID;AUTO_INCREMENT;NOT NULL"`
	SerialID             string    `gorm:"column:SerialID"`
	ActivationID         string    `gorm:"column:ActivationID"`
	ActivationCode       string    `gorm:"column:ActivationCode"`
	ActivationName       string    `gorm:"column:ActivationName"`
	TotalServers         int       `gorm:"column:TotalServers"`
	CreatedBy            int       `gorm:"column:CreatedBy"`
	CreatedDateTime      time.Time `gorm:"column:CreatedDateTime"`
	ModifiedBy           int       `gorm:"column:ModifiedBy;NOT NULL"`
	LastModifiedDateTime time.Time `gorm:"column:LastModifiedDateTime"`
	IsActive             string    `gorm:"column:IsActive;type:enum('Yes','No');default:Yes"`
}
type Servers struct {
	//ID                     int    `gorm:"column:ID;AUTO_INCREMENT;NOT NULL"`
	SerialID     string `gorm:"column:SerialID;NOT NULL"`
	Name         string `gorm:"column:Name;default:no-name-assigned"`
	InstanceID   string `gorm:"column:InstanceID;NOT NULL"`
	ServerID     string `gorm:"column:ServerID"`
	InstanceType string `gorm:"column:InstanceType"`
	//ServerType             string `gorm:"column:ServerType;type:enum('Hybrid','Non Hybrid');default:Non Hybrid"`
	//ServerTags             string `gorm:"column:ServerTags;type:MEDIUMTEXT;NOT NULL"`
	PublicIP   string `gorm:"column:PublicIP;NOT NULL"`
	PublicDNS  string `gorm:"column:PublicDNS;NOT NULL"`
	PrivateIP  string `gorm:"column:PrivateIP"`
	PrivateDNS string `gorm:"column:PrivateDNS"`
	// RegionID               int    `gorm:"column:RegionID;NOT NULL"`
	// OSID                   int    `gorm:"column:OSID;default:0;NOT NULL"`
	// OtherRegionName        string `gorm:"column:OtherRegionName;NOT NULL"`
	AvailabilityZone string `gorm:"column:AvailabilityZone"`
	//VPC                    string `gorm:"column:VPC"`
	//SubnetID               string `gorm:"column:SubnetID"`
	//SecurityGroups         string `gorm:"column:SecurityGroups"`
	MappedDisks string `gorm:"column:MappedDisks"`
	ImageName   string `gorm:"column:ImageName"`
	Platform    string `gorm:"column:Platform;NOT NULL"`
	//ADEnabled              string `gorm:"column:ADEnabled;type:enum('Yes','No');default:No;NOT NULL"`
	//OsVersion              string `gorm:"column:OsVersion;NOT NULL"`
	//ClusterID              int    `gorm:"column:ClusterID;NOT NULL"`
	//InstanceProfileARN     string `gorm:"column:InstanceProfileARN;NOT NULL"`
	// Tags                   string `gorm:"column:Tags;NOT NULL"`
	// AdditionalData         string `gorm:"column:AdditionalData;type:MEDIUMTEXT;NOT NULL"`
	ComputerName string `gorm:"column:ComputerName;NOT NULL"`
	ProviderID   int    `gorm:"column:ProviderID;NOT NULL"`
	ProjectID    int    `gorm:"column:ProjectID;NOT NULL"`
	CompanyID    int    `gorm:"column:CompanyID;NOT NULL"`
	// SSMStatus              string `gorm:"column:SSMStatus;type:enum('Yes','No');default:Yes"`
	// IsServerRunning        string `gorm:"column:IsServerRunning;type:enum('Yes','No');default:Yes"`
	// IsServerLocked         string `gorm:"column:IsServerLocked;type:enum('Yes','No');default:No"`
	MissingPatches        string `gorm:"column:MissingPatches;type:LONGTEXT"`
	InstalledPatches      string `gorm:"column:InstalledPatches;type:LONGTEXT;NOT NULL"`
	PatchDependenciesList string `gorm:"column:PatchDependenciesList"`
	// ComplianceStatus       string `gorm:"column:ComplianceStatus;type:enum('Yes','No');default:Yes"`
	TotalPatches      int    `gorm:"column:TotalPatches;NOT NULL"`
	AmiID             string `gorm:"column:AmiID;NOT NULL"`
	AmiCreationDetail string `gorm:"column:AmiCreationDetail;type:LONGTEXT;NOT NULL"`
	// IsPatchInstalled       string `gorm:"column:IsPatchInstalled;type:enum('Yes','No');default:Yes"`
	PatchCommandID         string `gorm:"column:PatchCommandID;NOT NULL"`
	InstallingPatches      string `gorm:"column:InstallingPatches;type:LONGTEXT;NOT NULL"`
	PatchInitiatedBy       int    `gorm:"column:PatchInitiatedBy;NOT NULL"`
	PatchInstalledDate     string `gorm:"column:PatchInstalledDate;NOT NULL"`
	IntervalsEmailDateTime string `gorm:"column:IntervalsEmailDateTime;NOT NULL"`
	PatchScannedDate       string `gorm:"column:PatchScannedDate;NOT NULL"`
	SiteHostName           string `gorm:"column:SiteHostName;NOT NULL"`
	ResourceGroup          string `gorm:"column:ResourceGroup;NOT NULL"`
	ResourceGroupID        int    `gorm:"column:ResourceGroupID;NOT NULL"`
	SupportedAppsData      string `gorm:"column:SupportedAppsData;type:LONGTEXT;NOT NULL"`
	// IsMasterKeyAssigned    string `gorm:"column:IsMasterKeyAssigned;type:enum('Yes','No');default:No"`
	// IsAbortedFromPolicy    string `gorm:"column:IsAbortedFromPolicy;type:enum('Yes','No');default:No"`
	LastHealthCheckAt string `gorm:"column:LastHealthCheckAt"`
	AgentActivationID int    `gorm:"column:AgentActivationID;default:0;NOT NULL"`
	CreatedDate       string `gorm:"column:CreatedDate;NOT NULL"`
	// IsActive          string `gorm:"column:IsActive;type:enum('Yes','No');default:No"`
	// IsTerminated      string `gorm:"column:IsTerminated;type:enum('Yes','No');default:No"`
	// IsDefault         string `gorm:"column:IsDefault;type:enum('Yes','No');default:No"`
}
