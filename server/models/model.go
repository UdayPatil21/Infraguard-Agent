package model

import (
	"time"

	"github.com/google/uuid"
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
	MachineID string `json:"machine_id" db:"machine_id" validate:"required"`
	Command   string `json:"command" validate:"required"`
}

type Executable struct {
	Script     []byte
	Permission string
}

type Activation struct {
	Id               int       `json:"id" db:"id"`
	Serial_Id        uuid.UUID `json:"serial_id" db:"serial_id"`
	Activation_Id    uuid.UUID `json:"activation_id" db:"activation_id"`
	Actiovation_Code string    `json:"activation_code" db:"activation_code"`
	Activation_Name  string    `json:"activation_name" db:"activation_name"`
	Server_Count     int       `json:"number_of_servers" db:"number_of_servers"`
	Created_By       string    `json:"created_by" db:"created_by"`
	Created_Date     time.Time `json:"created_date" db:"created_date"`
}
