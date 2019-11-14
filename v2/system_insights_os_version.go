/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type SystemInsightsOsVersion struct {

	Name string `json:"name,omitempty"`

	Version string `json:"version,omitempty"`

	Major int32 `json:"major,omitempty"`

	Minor int32 `json:"minor,omitempty"`

	Patch int32 `json:"patch,omitempty"`

	Build string `json:"build,omitempty"`

	Platform string `json:"platform,omitempty"`

	PlatformLike string `json:"platform_like,omitempty"`

	Codename string `json:"codename,omitempty"`

	InstallDate string `json:"install_date,omitempty"`

	CollectionTime string `json:"collection_time,omitempty"`

	SystemId string `json:"system_id,omitempty"`
}
