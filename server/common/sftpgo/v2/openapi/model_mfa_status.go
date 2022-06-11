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

// MFAStatus struct for MFAStatus
type MFAStatus struct {
	IsActive *bool `json:"is_active,omitempty"`
	TotpConfigs []TOTPConfig `json:"totp_configs,omitempty"`
}

// NewMFAStatus instantiates a new MFAStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFAStatus() *MFAStatus {
	this := MFAStatus{}
	return &this
}

// NewMFAStatusWithDefaults instantiates a new MFAStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFAStatusWithDefaults() *MFAStatus {
	this := MFAStatus{}
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *MFAStatus) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAStatus) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *MFAStatus) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *MFAStatus) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetTotpConfigs returns the TotpConfigs field value if set, zero value otherwise.
func (o *MFAStatus) GetTotpConfigs() []TOTPConfig {
	if o == nil || o.TotpConfigs == nil {
		var ret []TOTPConfig
		return ret
	}
	return o.TotpConfigs
}

// GetTotpConfigsOk returns a tuple with the TotpConfigs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFAStatus) GetTotpConfigsOk() ([]TOTPConfig, bool) {
	if o == nil || o.TotpConfigs == nil {
		return nil, false
	}
	return o.TotpConfigs, true
}

// HasTotpConfigs returns a boolean if a field has been set.
func (o *MFAStatus) HasTotpConfigs() bool {
	if o != nil && o.TotpConfigs != nil {
		return true
	}

	return false
}

// SetTotpConfigs gets a reference to the given []TOTPConfig and assigns it to the TotpConfigs field.
func (o *MFAStatus) SetTotpConfigs(v []TOTPConfig) {
	o.TotpConfigs = v
}

func (o MFAStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.TotpConfigs != nil {
		toSerialize["totp_configs"] = o.TotpConfigs
	}
	return json.Marshal(toSerialize)
}

type NullableMFAStatus struct {
	value *MFAStatus
	isSet bool
}

func (v NullableMFAStatus) Get() *MFAStatus {
	return v.value
}

func (v *NullableMFAStatus) Set(val *MFAStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAStatus(val *MFAStatus) *NullableMFAStatus {
	return &NullableMFAStatus{value: val, isSet: true}
}

func (v NullableMFAStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

