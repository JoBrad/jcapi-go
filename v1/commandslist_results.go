/*
 * JumpCloud APIs
 *
 * V1 & V2 versions of JumpCloud's API. The previous version of JumpCloud's API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type CommandslistResults struct {

	// The name of the Command.
	Name string `json:"name,omitempty"`

	// The Command to execute.
	Command string `json:"command,omitempty"`

	// The Command OS.
	CommandType string `json:"commandType,omitempty"`

	// How the Command is excecuted.
	LaunchType string `json:"launchType,omitempty"`

	// 
	ListensTo string `json:"listensTo,omitempty"`

	// A crontab that consists of: [ (seconds) (minutes) (hours) (days of month) (months) (weekdays) ] or [ immediate ]. If you send this as an empty string, it will run immediately. 
	Schedule string `json:"schedule,omitempty"`

	// trigger to execute command.
	Trigger string `json:"trigger,omitempty"`

	// When the command will repeat.
	ScheduleRepeatType string `json:"scheduleRepeatType,omitempty"`

	// The ID of the Organization.
	Organization string `json:"organization,omitempty"`

	// The ID of the command.
	Id string `json:"_id,omitempty"`
}
