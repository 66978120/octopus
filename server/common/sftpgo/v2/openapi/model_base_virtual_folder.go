/*
SFTPGo

SFTPGo allows to securely share your files over SFTP, HTTP and optionally FTP/S and WebDAV as well. Several storage backends are supported and they are configurable per user, so you can serve a local directory for a user and an S3 bucket (or part of it) for another one. SFTPGo also supports virtual folders, a virtual folder can use any of the supported storage backends. So you can have, for example, an S3 user that exposes a GCS bucket (or part of it) on a specified path and an encrypted local filesystem on another one. Virtual folders can be private or shared among multiple users, for shared virtual folders you can define different quota limits for each user. SFTPGo allows to create HTTP/S links to externally share files and folders securely, by setting limits to the number of downloads/uploads, protecting the share with a password, limiting access by source IP address, setting an automatic expiration date. 

API version: 2.2.2-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// BaseVirtualFolder Defines the filesystem for the virtual folder and the used quota limits. The same folder can be shared among multiple users and each user can have different quota limits or a different virtual path.
type BaseVirtualFolder struct {
	Id *int32 `json:"id,omitempty"`
	// unique name for this virtual folder
	Name *string `json:"name,omitempty"`
	// absolute filesystem path to use as virtual folder
	MappedPath *string `json:"mapped_path,omitempty"`
	// optional description
	Description *string `json:"description,omitempty"`
	UsedQuotaSize *int64 `json:"used_quota_size,omitempty"`
	UsedQuotaFiles *int32 `json:"used_quota_files,omitempty"`
	// Last quota update as unix timestamp in milliseconds
	LastQuotaUpdate *int64 `json:"last_quota_update,omitempty"`
	// list of usernames associated with this virtual folder
	Users []string `json:"users,omitempty"`
	Filesystem *FilesystemConfig `json:"filesystem,omitempty"`
}

// NewBaseVirtualFolder instantiates a new BaseVirtualFolder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseVirtualFolder() *BaseVirtualFolder {
	this := BaseVirtualFolder{}
	return &this
}

// NewBaseVirtualFolderWithDefaults instantiates a new BaseVirtualFolder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseVirtualFolderWithDefaults() *BaseVirtualFolder {
	this := BaseVirtualFolder{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetId() int32 {
	if o == nil || o.Id == nil {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetIdOk() (*int32, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *BaseVirtualFolder) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BaseVirtualFolder) SetName(v string) {
	o.Name = &v
}

// GetMappedPath returns the MappedPath field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetMappedPath() string {
	if o == nil || o.MappedPath == nil {
		var ret string
		return ret
	}
	return *o.MappedPath
}

// GetMappedPathOk returns a tuple with the MappedPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetMappedPathOk() (*string, bool) {
	if o == nil || o.MappedPath == nil {
		return nil, false
	}
	return o.MappedPath, true
}

// HasMappedPath returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasMappedPath() bool {
	if o != nil && o.MappedPath != nil {
		return true
	}

	return false
}

// SetMappedPath gets a reference to the given string and assigns it to the MappedPath field.
func (o *BaseVirtualFolder) SetMappedPath(v string) {
	o.MappedPath = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BaseVirtualFolder) SetDescription(v string) {
	o.Description = &v
}

// GetUsedQuotaSize returns the UsedQuotaSize field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetUsedQuotaSize() int64 {
	if o == nil || o.UsedQuotaSize == nil {
		var ret int64
		return ret
	}
	return *o.UsedQuotaSize
}

// GetUsedQuotaSizeOk returns a tuple with the UsedQuotaSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetUsedQuotaSizeOk() (*int64, bool) {
	if o == nil || o.UsedQuotaSize == nil {
		return nil, false
	}
	return o.UsedQuotaSize, true
}

// HasUsedQuotaSize returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasUsedQuotaSize() bool {
	if o != nil && o.UsedQuotaSize != nil {
		return true
	}

	return false
}

// SetUsedQuotaSize gets a reference to the given int64 and assigns it to the UsedQuotaSize field.
func (o *BaseVirtualFolder) SetUsedQuotaSize(v int64) {
	o.UsedQuotaSize = &v
}

// GetUsedQuotaFiles returns the UsedQuotaFiles field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetUsedQuotaFiles() int32 {
	if o == nil || o.UsedQuotaFiles == nil {
		var ret int32
		return ret
	}
	return *o.UsedQuotaFiles
}

// GetUsedQuotaFilesOk returns a tuple with the UsedQuotaFiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetUsedQuotaFilesOk() (*int32, bool) {
	if o == nil || o.UsedQuotaFiles == nil {
		return nil, false
	}
	return o.UsedQuotaFiles, true
}

// HasUsedQuotaFiles returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasUsedQuotaFiles() bool {
	if o != nil && o.UsedQuotaFiles != nil {
		return true
	}

	return false
}

// SetUsedQuotaFiles gets a reference to the given int32 and assigns it to the UsedQuotaFiles field.
func (o *BaseVirtualFolder) SetUsedQuotaFiles(v int32) {
	o.UsedQuotaFiles = &v
}

// GetLastQuotaUpdate returns the LastQuotaUpdate field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetLastQuotaUpdate() int64 {
	if o == nil || o.LastQuotaUpdate == nil {
		var ret int64
		return ret
	}
	return *o.LastQuotaUpdate
}

// GetLastQuotaUpdateOk returns a tuple with the LastQuotaUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetLastQuotaUpdateOk() (*int64, bool) {
	if o == nil || o.LastQuotaUpdate == nil {
		return nil, false
	}
	return o.LastQuotaUpdate, true
}

// HasLastQuotaUpdate returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasLastQuotaUpdate() bool {
	if o != nil && o.LastQuotaUpdate != nil {
		return true
	}

	return false
}

// SetLastQuotaUpdate gets a reference to the given int64 and assigns it to the LastQuotaUpdate field.
func (o *BaseVirtualFolder) SetLastQuotaUpdate(v int64) {
	o.LastQuotaUpdate = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetUsers() []string {
	if o == nil || o.Users == nil {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetUsersOk() ([]string, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *BaseVirtualFolder) SetUsers(v []string) {
	o.Users = v
}

// GetFilesystem returns the Filesystem field value if set, zero value otherwise.
func (o *BaseVirtualFolder) GetFilesystem() FilesystemConfig {
	if o == nil || o.Filesystem == nil {
		var ret FilesystemConfig
		return ret
	}
	return *o.Filesystem
}

// GetFilesystemOk returns a tuple with the Filesystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseVirtualFolder) GetFilesystemOk() (*FilesystemConfig, bool) {
	if o == nil || o.Filesystem == nil {
		return nil, false
	}
	return o.Filesystem, true
}

// HasFilesystem returns a boolean if a field has been set.
func (o *BaseVirtualFolder) HasFilesystem() bool {
	if o != nil && o.Filesystem != nil {
		return true
	}

	return false
}

// SetFilesystem gets a reference to the given FilesystemConfig and assigns it to the Filesystem field.
func (o *BaseVirtualFolder) SetFilesystem(v FilesystemConfig) {
	o.Filesystem = &v
}

func (o BaseVirtualFolder) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.MappedPath != nil {
		toSerialize["mapped_path"] = o.MappedPath
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.UsedQuotaSize != nil {
		toSerialize["used_quota_size"] = o.UsedQuotaSize
	}
	if o.UsedQuotaFiles != nil {
		toSerialize["used_quota_files"] = o.UsedQuotaFiles
	}
	if o.LastQuotaUpdate != nil {
		toSerialize["last_quota_update"] = o.LastQuotaUpdate
	}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Filesystem != nil {
		toSerialize["filesystem"] = o.Filesystem
	}
	return json.Marshal(toSerialize)
}

type NullableBaseVirtualFolder struct {
	value *BaseVirtualFolder
	isSet bool
}

func (v NullableBaseVirtualFolder) Get() *BaseVirtualFolder {
	return v.value
}

func (v *NullableBaseVirtualFolder) Set(val *BaseVirtualFolder) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseVirtualFolder) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseVirtualFolder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseVirtualFolder(val *BaseVirtualFolder) *NullableBaseVirtualFolder {
	return &NullableBaseVirtualFolder{value: val, isSet: true}
}

func (v NullableBaseVirtualFolder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseVirtualFolder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


