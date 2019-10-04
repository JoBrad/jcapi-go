/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsSystemControls struct {

	CollectionTime string `json:"collection_time,omitempty"`

	ConfigValue string `json:"config_value,omitempty"`

	CurrentValue string `json:"current_value,omitempty"`

	FieldName string `json:"field_name,omitempty"`

	Name string `json:"name,omitempty"`

	Oid string `json:"oid,omitempty"`

	Subsystem string `json:"subsystem,omitempty"`

	SystemId string `json:"system_id,omitempty"`

	Type_ string `json:"type,omitempty"`
}