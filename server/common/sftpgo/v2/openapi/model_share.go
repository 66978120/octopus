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

// Share struct for Share
type Share struct {
	// auto-generated unique share identifier
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	// optional description
	Description *string `json:"description,omitempty"`
	Scope *ShareScope `json:"scope,omitempty"`
	// paths to files or directories, for share scope write this array must contain exactly one directory. Paths will not be validated on save so you can also create them after creating the share
	Paths []string `json:"paths,omitempty"`
	Username *string `json:"username,omitempty"`
	// creation time as unix timestamp in milliseconds
	CreatedAt *int64 `json:"created_at,omitempty"`
	// last update time as unix timestamp in milliseconds
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// last use time as unix timestamp in milliseconds
	LastUseAt *int64 `json:"last_use_at,omitempty"`
	// optional share expiration, as unix timestamp in milliseconds. 0 means no expiration
	ExpiresAt *int64 `json:"expires_at,omitempty"`
	// optional password to protect the share. The special value \"[**redacted**]\" means that a password has been set, you can use this value if you want to preserve the current password when you update a share
	Password *string `json:"password,omitempty"`
	// maximum allowed access tokens. 0 means no limit
	MaxTokens *int32 `json:"max_tokens,omitempty"`
	UsedTokens *int32 `json:"used_tokens,omitempty"`
	// Limit the share availability to these IP/Mask. IP/Mask must be in CIDR notation as defined in RFC 4632 and RFC 4291, for example \"192.0.2.0/24\" or \"2001:db8::/32\". An empty list means no restrictions
	AllowFrom []string `json:"allow_from,omitempty"`
}

// NewShare instantiates a new Share object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShare() *Share {
	this := Share{}
	return &this
}

// NewShareWithDefaults instantiates a new Share object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShareWithDefaults() *Share {
	this := Share{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Share) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Share) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Share) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Share) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Share) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Share) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Share) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Share) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Share) SetDescription(v string) {
	o.Description = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *Share) GetScope() ShareScope {
	if o == nil || o.Scope == nil {
		var ret ShareScope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetScopeOk() (*ShareScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *Share) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given ShareScope and assigns it to the Scope field.
func (o *Share) SetScope(v ShareScope) {
	o.Scope = &v
}

// GetPaths returns the Paths field value if set, zero value otherwise.
func (o *Share) GetPaths() []string {
	if o == nil || o.Paths == nil {
		var ret []string
		return ret
	}
	return o.Paths
}

// GetPathsOk returns a tuple with the Paths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetPathsOk() ([]string, bool) {
	if o == nil || o.Paths == nil {
		return nil, false
	}
	return o.Paths, true
}

// HasPaths returns a boolean if a field has been set.
func (o *Share) HasPaths() bool {
	if o != nil && o.Paths != nil {
		return true
	}

	return false
}

// SetPaths gets a reference to the given []string and assigns it to the Paths field.
func (o *Share) SetPaths(v []string) {
	o.Paths = v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Share) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Share) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Share) SetUsername(v string) {
	o.Username = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Share) GetCreatedAt() int64 {
	if o == nil || o.CreatedAt == nil {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetCreatedAtOk() (*int64, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Share) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *Share) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Share) GetUpdatedAt() int64 {
	if o == nil || o.UpdatedAt == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetUpdatedAtOk() (*int64, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Share) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given int64 and assigns it to the UpdatedAt field.
func (o *Share) SetUpdatedAt(v int64) {
	o.UpdatedAt = &v
}

// GetLastUseAt returns the LastUseAt field value if set, zero value otherwise.
func (o *Share) GetLastUseAt() int64 {
	if o == nil || o.LastUseAt == nil {
		var ret int64
		return ret
	}
	return *o.LastUseAt
}

// GetLastUseAtOk returns a tuple with the LastUseAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetLastUseAtOk() (*int64, bool) {
	if o == nil || o.LastUseAt == nil {
		return nil, false
	}
	return o.LastUseAt, true
}

// HasLastUseAt returns a boolean if a field has been set.
func (o *Share) HasLastUseAt() bool {
	if o != nil && o.LastUseAt != nil {
		return true
	}

	return false
}

// SetLastUseAt gets a reference to the given int64 and assigns it to the LastUseAt field.
func (o *Share) SetLastUseAt(v int64) {
	o.LastUseAt = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *Share) GetExpiresAt() int64 {
	if o == nil || o.ExpiresAt == nil {
		var ret int64
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetExpiresAtOk() (*int64, bool) {
	if o == nil || o.ExpiresAt == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *Share) HasExpiresAt() bool {
	if o != nil && o.ExpiresAt != nil {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int64 and assigns it to the ExpiresAt field.
func (o *Share) SetExpiresAt(v int64) {
	o.ExpiresAt = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *Share) GetPassword() string {
	if o == nil || o.Password == nil {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetPasswordOk() (*string, bool) {
	if o == nil || o.Password == nil {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *Share) HasPassword() bool {
	if o != nil && o.Password != nil {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *Share) SetPassword(v string) {
	o.Password = &v
}

// GetMaxTokens returns the MaxTokens field value if set, zero value otherwise.
func (o *Share) GetMaxTokens() int32 {
	if o == nil || o.MaxTokens == nil {
		var ret int32
		return ret
	}
	return *o.MaxTokens
}

// GetMaxTokensOk returns a tuple with the MaxTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetMaxTokensOk() (*int32, bool) {
	if o == nil || o.MaxTokens == nil {
		return nil, false
	}
	return o.MaxTokens, true
}

// HasMaxTokens returns a boolean if a field has been set.
func (o *Share) HasMaxTokens() bool {
	if o != nil && o.MaxTokens != nil {
		return true
	}

	return false
}

// SetMaxTokens gets a reference to the given int32 and assigns it to the MaxTokens field.
func (o *Share) SetMaxTokens(v int32) {
	o.MaxTokens = &v
}

// GetUsedTokens returns the UsedTokens field value if set, zero value otherwise.
func (o *Share) GetUsedTokens() int32 {
	if o == nil || o.UsedTokens == nil {
		var ret int32
		return ret
	}
	return *o.UsedTokens
}

// GetUsedTokensOk returns a tuple with the UsedTokens field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetUsedTokensOk() (*int32, bool) {
	if o == nil || o.UsedTokens == nil {
		return nil, false
	}
	return o.UsedTokens, true
}

// HasUsedTokens returns a boolean if a field has been set.
func (o *Share) HasUsedTokens() bool {
	if o != nil && o.UsedTokens != nil {
		return true
	}

	return false
}

// SetUsedTokens gets a reference to the given int32 and assigns it to the UsedTokens field.
func (o *Share) SetUsedTokens(v int32) {
	o.UsedTokens = &v
}

// GetAllowFrom returns the AllowFrom field value if set, zero value otherwise.
func (o *Share) GetAllowFrom() []string {
	if o == nil || o.AllowFrom == nil {
		var ret []string
		return ret
	}
	return o.AllowFrom
}

// GetAllowFromOk returns a tuple with the AllowFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Share) GetAllowFromOk() ([]string, bool) {
	if o == nil || o.AllowFrom == nil {
		return nil, false
	}
	return o.AllowFrom, true
}

// HasAllowFrom returns a boolean if a field has been set.
func (o *Share) HasAllowFrom() bool {
	if o != nil && o.AllowFrom != nil {
		return true
	}

	return false
}

// SetAllowFrom gets a reference to the given []string and assigns it to the AllowFrom field.
func (o *Share) SetAllowFrom(v []string) {
	o.AllowFrom = v
}

func (o Share) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.Paths != nil {
		toSerialize["paths"] = o.Paths
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if o.LastUseAt != nil {
		toSerialize["last_use_at"] = o.LastUseAt
	}
	if o.ExpiresAt != nil {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if o.Password != nil {
		toSerialize["password"] = o.Password
	}
	if o.MaxTokens != nil {
		toSerialize["max_tokens"] = o.MaxTokens
	}
	if o.UsedTokens != nil {
		toSerialize["used_tokens"] = o.UsedTokens
	}
	if o.AllowFrom != nil {
		toSerialize["allow_from"] = o.AllowFrom
	}
	return json.Marshal(toSerialize)
}

type NullableShare struct {
	value *Share
	isSet bool
}

func (v NullableShare) Get() *Share {
	return v.value
}

func (v *NullableShare) Set(val *Share) {
	v.value = val
	v.isSet = true
}

func (v NullableShare) IsSet() bool {
	return v.isSet
}

func (v *NullableShare) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShare(val *Share) *NullableShare {
	return &NullableShare{value: val, isSet: true}
}

func (v NullableShare) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShare) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

