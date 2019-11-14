/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsWindowsCrashes struct {

	Datetime string `json:"datetime,omitempty"`

	Module string `json:"module,omitempty"`

	Path string `json:"path,omitempty"`

	Pid string `json:"pid,omitempty"`

	Tid string `json:"tid,omitempty"`

	Version string `json:"version,omitempty"`

	ProcessUptime string `json:"process_uptime,omitempty"`

	StackTrace string `json:"stack_trace,omitempty"`

	ExceptionCode string `json:"exception_code,omitempty"`

	ExceptionMessage string `json:"exception_message,omitempty"`

	ExceptionAddress string `json:"exception_address,omitempty"`

	Registers string `json:"registers,omitempty"`

	CommandLine string `json:"command_line,omitempty"`

	CurrentDirectory string `json:"current_directory,omitempty"`

	Username string `json:"username,omitempty"`

	MachineName string `json:"machine_name,omitempty"`

	MajorVersion int32 `json:"major_version,omitempty"`

	MinorVersion int32 `json:"minor_version,omitempty"`

	BuildNumber int32 `json:"build_number,omitempty"`

	Type_ string `json:"type,omitempty"`

	CrashPath string `json:"crash_path,omitempty"`
}
