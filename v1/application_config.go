/*
 * JumpCloud APIs
 *
 *  JumpCloud's V1 API. This set of endpoints allows JumpCloud customers to manage commands, systems, & system users.
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1

type ApplicationConfig struct {

	IdpEntityId *ApplicationConfigIdpEntityId `json:"idpEntityId,omitempty"`

	IdpCertificate *ApplicationConfigIdpEntityId `json:"idpCertificate,omitempty"`

	SpEntityId *ApplicationConfigIdpEntityId `json:"spEntityId,omitempty"`

	AcsUrl *ApplicationConfigIdpEntityId `json:"acsUrl,omitempty"`

	ConstantAttributes *ApplicationConfigConstantAttributes `json:"constantAttributes,omitempty"`

	DatabaseAttributes *ApplicationConfigDatabaseAttributes `json:"databaseAttributes,omitempty"`

	IdpPrivateKey *ApplicationConfigIdpEntityId `json:"idpPrivateKey,omitempty"`
}
