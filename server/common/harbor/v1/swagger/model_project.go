/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Project struct {
	// Project ID
	ProjectId int32 `json:"project_id,omitempty"`
	// The owner ID of the project always means the creator of the project.
	OwnerId int32 `json:"owner_id,omitempty"`
	// The name of the project.
	Name string `json:"name,omitempty"`
	// The creation time of the project.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of the project.
	UpdateTime string `json:"update_time,omitempty"`
	// A deletion mark of the project (1 means it's deleted, 0 is not)
	Deleted int32 `json:"deleted,omitempty"`
	// The owner name of the project.
	OwnerName string `json:"owner_name,omitempty"`
	// Correspond to the UI about whether the project's publicity is  updatable (for UI)
	Togglable bool `json:"Togglable,omitempty"`
	// The role ID of the current user who triggered the API (for UI)
	CurrentUserRoleId int32 `json:"current_user_role_id,omitempty"`
	// The number of the repositories under this project.
	RepoCount int32 `json:"repo_count,omitempty"`
	// The metadata of the project.
	Metadata interface{} `json:"metadata,omitempty"`
}
