package main

type DeviceInfo struct {
	Manufacturer string `json:"manufacturer"`
	Model        string `json:"model"`
	HWVersion    string `json:"hw_version"`
	SWVersion    string `json:"sw_version"`
}

type Device struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Description  string                 `json:"description"`
	Room         string                 `json:"room"`
	Type         string                 `json:"type"`
	CustomData   map[string]string      `json:"custom_data"`
	Capabilities map[string]interface{} `json:"capabilities"`
	Properties   map[string]interface{} `json:"properties"`
	DeviceInfo   DeviceInfo             `json:"device_info"`
}

type Payload struct {
	UserID  string   `json:"user_id"`
	Devices []Device `json:"devices"`
}

type Response struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}
