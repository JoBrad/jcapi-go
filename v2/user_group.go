/*
 * JumpCloud APIs
 *
 *  JumpCloud's V2 API. This set of endpoints allows JumpCloud customers to manage objects, groupings and mappings and interact with the JumpCloud Graph.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v2

type UserGroup struct {

	Attributes *UserGroupAttributes `json:"attributes,omitempty"`

	// ObjectId uniquely identifying a User Group.
	Id string `json:"id,omitempty"`

	// The type of the group.
	Type_ string `json:"type,omitempty"`

	// Display name of a User Group.
	Name string `json:"name,omitempty"`
}
