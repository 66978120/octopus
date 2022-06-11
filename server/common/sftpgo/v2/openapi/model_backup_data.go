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

// BackupData struct for BackupData
type BackupData struct {
	Users []User `json:"users,omitempty"`
	Folders []BaseVirtualFolder `json:"folders,omitempty"`
	Admins []Admin `json:"admins,omitempty"`
	ApiKeys []AuthAPIKey `json:"api_keys,omitempty"`
	Shares []Share `json:"shares,omitempty"`
	Version *int32 `json:"version,omitempty"`
}

// NewBackupData instantiates a new BackupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupData() *BackupData {
	this := BackupData{}
	return &this
}

// NewBackupDataWithDefaults instantiates a new BackupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupDataWithDefaults() *BackupData {
	this := BackupData{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *BackupData) GetUsers() []User {
	if o == nil || o.Users == nil {
		var ret []User
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetUsersOk() ([]User, bool) {
	if o == nil || o.Users == nil {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *BackupData) HasUsers() bool {
	if o != nil && o.Users != nil {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []User and assigns it to the Users field.
func (o *BackupData) SetUsers(v []User) {
	o.Users = v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *BackupData) GetFolders() []BaseVirtualFolder {
	if o == nil || o.Folders == nil {
		var ret []BaseVirtualFolder
		return ret
	}
	return o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetFoldersOk() ([]BaseVirtualFolder, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *BackupData) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []BaseVirtualFolder and assigns it to the Folders field.
func (o *BackupData) SetFolders(v []BaseVirtualFolder) {
	o.Folders = v
}

// GetAdmins returns the Admins field value if set, zero value otherwise.
func (o *BackupData) GetAdmins() []Admin {
	if o == nil || o.Admins == nil {
		var ret []Admin
		return ret
	}
	return o.Admins
}

// GetAdminsOk returns a tuple with the Admins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetAdminsOk() ([]Admin, bool) {
	if o == nil || o.Admins == nil {
		return nil, false
	}
	return o.Admins, true
}

// HasAdmins returns a boolean if a field has been set.
func (o *BackupData) HasAdmins() bool {
	if o != nil && o.Admins != nil {
		return true
	}

	return false
}

// SetAdmins gets a reference to the given []Admin and assigns it to the Admins field.
func (o *BackupData) SetAdmins(v []Admin) {
	o.Admins = v
}

// GetApiKeys returns the ApiKeys field value if set, zero value otherwise.
func (o *BackupData) GetApiKeys() []AuthAPIKey {
	if o == nil || o.ApiKeys == nil {
		var ret []AuthAPIKey
		return ret
	}
	return o.ApiKeys
}

// GetApiKeysOk returns a tuple with the ApiKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetApiKeysOk() ([]AuthAPIKey, bool) {
	if o == nil || o.ApiKeys == nil {
		return nil, false
	}
	return o.ApiKeys, true
}

// HasApiKeys returns a boolean if a field has been set.
func (o *BackupData) HasApiKeys() bool {
	if o != nil && o.ApiKeys != nil {
		return true
	}

	return false
}

// SetApiKeys gets a reference to the given []AuthAPIKey and assigns it to the ApiKeys field.
func (o *BackupData) SetApiKeys(v []AuthAPIKey) {
	o.ApiKeys = v
}

// GetShares returns the Shares field value if set, zero value otherwise.
func (o *BackupData) GetShares() []Share {
	if o == nil || o.Shares == nil {
		var ret []Share
		return ret
	}
	return o.Shares
}

// GetSharesOk returns a tuple with the Shares field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetSharesOk() ([]Share, bool) {
	if o == nil || o.Shares == nil {
		return nil, false
	}
	return o.Shares, true
}

// HasShares returns a boolean if a field has been set.
func (o *BackupData) HasShares() bool {
	if o != nil && o.Shares != nil {
		return true
	}

	return false
}

// SetShares gets a reference to the given []Share and assigns it to the Shares field.
func (o *BackupData) SetShares(v []Share) {
	o.Shares = v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *BackupData) GetVersion() int32 {
	if o == nil || o.Version == nil {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackupData) GetVersionOk() (*int32, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *BackupData) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *BackupData) SetVersion(v int32) {
	o.Version = &v
}

func (o BackupData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Users != nil {
		toSerialize["users"] = o.Users
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	if o.Admins != nil {
		toSerialize["admins"] = o.Admins
	}
	if o.ApiKeys != nil {
		toSerialize["api_keys"] = o.ApiKeys
	}
	if o.Shares != nil {
		toSerialize["shares"] = o.Shares
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableBackupData struct {
	value *BackupData
	isSet bool
}

func (v NullableBackupData) Get() *BackupData {
	return v.value
}

func (v *NullableBackupData) Set(val *BackupData) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupData) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupData(val *BackupData) *NullableBackupData {
	return &NullableBackupData{value: val, isSet: true}
}

func (v NullableBackupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

