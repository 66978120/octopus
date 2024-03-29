/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type GeneralInfo struct {
	// If the Harbor instance is deployed with nested notary.
	WithNotary bool `json:"with_notary,omitempty"`
	// If the Harbor instance is deployed with nested clair.
	WithClair bool `json:"with_clair,omitempty"`
	// If the Harbor instance is deployed with Admiral.
	WithAdmiral bool `json:"with_admiral,omitempty"`
	// The url of the endpoint of admiral instance.
	AdmiralEndpoint string `json:"admiral_endpoint,omitempty"`
	// The auth mode of current Harbor instance.
	AuthMode string `json:"auth_mode,omitempty"`
	// Indicate who can create projects, it could be 'adminonly' or 'everyone'.
	ProjectCreationRestriction string `json:"project_creation_restriction,omitempty"`
	// Indicate whether the Harbor instance enable user to register himself.
	SelfRegistration bool `json:"self_registration,omitempty"`
	// Indicate whether there is a ca root cert file ready for download in the file system.
	HasCaRoot bool `json:"has_ca_root,omitempty"`
	// The build version of Harbor.
	HarborVersion string `json:"harbor_version,omitempty"`
	// The UTC time in milliseconds, after which user can call scanAll API to scan all images.
	NextScanAll int32 `json:"next_scan_all,omitempty"`
	// The status of vulnerability data of Clair.
	ClairVulnerabilityStatus interface{} `json:"clair_vulnerability_status,omitempty"`
}
