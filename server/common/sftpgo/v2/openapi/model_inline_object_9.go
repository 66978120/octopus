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

// InlineObject9 struct for InlineObject9
type InlineObject9 struct {
	// File modification time as unix timestamp in milliseconds
	ModificationTime *int32 `json:"modification_time,omitempty"`
}

// NewInlineObject9 instantiates a new InlineObject9 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject9() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// NewInlineObject9WithDefaults instantiates a new InlineObject9 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject9WithDefaults() *InlineObject9 {
	this := InlineObject9{}
	return &this
}

// GetModificationTime returns the ModificationTime field value if set, zero value otherwise.
func (o *InlineObject9) GetModificationTime() int32 {
	if o == nil || o.ModificationTime == nil {
		var ret int32
		return ret
	}
	return *o.ModificationTime
}

// GetModificationTimeOk returns a tuple with the ModificationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject9) GetModificationTimeOk() (*int32, bool) {
	if o == nil || o.ModificationTime == nil {
		return nil, false
	}
	return o.ModificationTime, true
}

// HasModificationTime returns a boolean if a field has been set.
func (o *InlineObject9) HasModificationTime() bool {
	if o != nil && o.ModificationTime != nil {
		return true
	}

	return false
}

// SetModificationTime gets a reference to the given int32 and assigns it to the ModificationTime field.
func (o *InlineObject9) SetModificationTime(v int32) {
	o.ModificationTime = &v
}

func (o InlineObject9) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ModificationTime != nil {
		toSerialize["modification_time"] = o.ModificationTime
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject9 struct {
	value *InlineObject9
	isSet bool
}

func (v NullableInlineObject9) Get() *InlineObject9 {
	return v.value
}

func (v *NullableInlineObject9) Set(val *InlineObject9) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject9) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject9) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject9(val *InlineObject9) *NullableInlineObject9 {
	return &NullableInlineObject9{value: val, isSet: true}
}

func (v NullableInlineObject9) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject9) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

