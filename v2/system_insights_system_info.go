/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsSystemInfo struct {

	Hostname string `json:"hostname,omitempty"`

	Uuid string `json:"uuid,omitempty"`

	CpuType string `json:"cpu_type,omitempty"`

	CpuSubtype string `json:"cpu_subtype,omitempty"`

	CpuBrand string `json:"cpu_brand,omitempty"`

	CpuPhysicalCores int32 `json:"cpu_physical_cores,omitempty"`

	CpuLogicalCores int32 `json:"cpu_logical_cores,omitempty"`

	CpuMicrocode string `json:"cpu_microcode,omitempty"`

	PhysicalMemory string `json:"physical_memory,omitempty"`

	HardwareVendor string `json:"hardware_vendor,omitempty"`

	HardwareModel string `json:"hardware_model,omitempty"`

	HardwareVersion string `json:"hardware_version,omitempty"`

	HardwareSerial string `json:"hardware_serial,omitempty"`

	ComputerName string `json:"computer_name,omitempty"`

	LocalHostname string `json:"local_hostname,omitempty"`

	CollectionTime string `json:"collection_time,omitempty"`

	SystemId string `json:"system_id,omitempty"`
}
