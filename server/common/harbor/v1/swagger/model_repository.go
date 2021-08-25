/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Repository struct {
	// The ID of repository.
	Id int32 `json:"id,omitempty"`
	// The name of repository.
	Name string `json:"name,omitempty"`
	// The project ID of repository.
	ProjectId int32 `json:"project_id,omitempty"`
	// The description of repository.
	Description string `json:"description,omitempty"`
	// The pull count of repository.
	PullCount int32 `json:"pull_count,omitempty"`
	// The star count of repository.
	StarCount int32 `json:"star_count,omitempty"`
	// The tags count of repository.
	TagsCount int32 `json:"tags_count,omitempty"`
	// The creation time of repository.
	CreationTime string `json:"creation_time,omitempty"`
	// The update time of repository.
	UpdateTime string `json:"update_time,omitempty"`
}
