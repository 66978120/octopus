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

// DataProviderStatus struct for DataProviderStatus
type DataProviderStatus struct {
	IsActive *bool `json:"is_active,omitempty"`
	Driver *string `json:"driver,omitempty"`
	Error *string `json:"error,omitempty"`
}

// NewDataProviderStatus instantiates a new DataProviderStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataProviderStatus() *DataProviderStatus {
	this := DataProviderStatus{}
	return &this
}

// NewDataProviderStatusWithDefaults instantiates a new DataProviderStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataProviderStatusWithDefaults() *DataProviderStatus {
	this := DataProviderStatus{}
	return &this
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *DataProviderStatus) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProviderStatus) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *DataProviderStatus) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *DataProviderStatus) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *DataProviderStatus) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProviderStatus) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *DataProviderStatus) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *DataProviderStatus) SetDriver(v string) {
	o.Driver = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *DataProviderStatus) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataProviderStatus) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *DataProviderStatus) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *DataProviderStatus) SetError(v string) {
	o.Error = &v
}

func (o DataProviderStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IsActive != nil {
		toSerialize["is_active"] = o.IsActive
	}
	if o.Driver != nil {
		toSerialize["driver"] = o.Driver
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableDataProviderStatus struct {
	value *DataProviderStatus
	isSet bool
}

func (v NullableDataProviderStatus) Get() *DataProviderStatus {
	return v.value
}

func (v *NullableDataProviderStatus) Set(val *DataProviderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDataProviderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDataProviderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataProviderStatus(val *DataProviderStatus) *NullableDataProviderStatus {
	return &NullableDataProviderStatus{value: val, isSet: true}
}

func (v NullableDataProviderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataProviderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

