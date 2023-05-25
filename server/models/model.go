package model

import "time"

type InstanceInfo struct {
	Name      any       `json:"name"`
	UserName  string    `json:"user_name"`
	MachineID string    `json:"machine_id"`
	PublicIP  string    `json:"public_ip"`
	HostName  string    `json:"host_name"`
	OS        string    `json:"os_name"`
	CreatedAt time.Time `json:"createdAt"`
	Status    string    `json:"status"`
}

type RunCommand struct {
	MachineID string `json:"machine_id" db:"machine_id" validate:"required"`
	Command   string `json:"command" validate:"required"`
}

type Executable struct {
	Script     []byte
	Permission string
}
