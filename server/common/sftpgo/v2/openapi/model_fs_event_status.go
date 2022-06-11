/*
SFTPGo

SFTPGo allows to securely share your files over SFTP, HTTP and optionally FTP/S and WebDAV as well. Several storage backends are supported and they are configurable per user, so you can serve a local directory for a user and an S3 bucket (or part of it) for another one. SFTPGo also supports virtual folders, a virtual folder can use any of the supported storage backends. So you can have, for example, an S3 user that exposes a GCS bucket (or part of it) on a specified path and an encrypted local filesystem on another one. Virtual folders can be private or shared among multiple users, for shared virtual folders you can define different quota limits for each user. SFTPGo allows to create HTTP/S links to externally share files and folders securely, by setting limits to the number of downloads/uploads, protecting the share with a password, limiting access by source IP address, setting an automatic expiration date. 

API version: 2.2.2-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// FsEventStatus Event status:   * `1` - no error   * `2` - generic error   * `3` - quota exceeded error 
type FsEventStatus int32

// List of FsEventStatus
const (
	FSEVENTSTATUS__1 FsEventStatus = 1
	FSEVENTSTATUS__2 FsEventStatus = 2
	FSEVENTSTATUS__3 FsEventStatus = 3
)

// All allowed values of FsEventStatus enum
var AllowedFsEventStatusEnumValues = []FsEventStatus{
	1,
	2,
	3,
}

func (v *FsEventStatus) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FsEventStatus(value)
	for _, existing := range AllowedFsEventStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FsEventStatus", value)
}

// NewFsEventStatusFromValue returns a pointer to a valid FsEventStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFsEventStatusFromValue(v int32) (*FsEventStatus, error) {
	ev := FsEventStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FsEventStatus: valid values are %v", v, AllowedFsEventStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FsEventStatus) IsValid() bool {
	for _, existing := range AllowedFsEventStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to FsEventStatus value
func (v FsEventStatus) Ptr() *FsEventStatus {
	return &v
}

type NullableFsEventStatus struct {
	value *FsEventStatus
	isSet bool
}

func (v NullableFsEventStatus) Get() *FsEventStatus {
	return v.value
}

func (v *NullableFsEventStatus) Set(val *FsEventStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableFsEventStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableFsEventStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFsEventStatus(val *FsEventStatus) *NullableFsEventStatus {
	return &NullableFsEventStatus{value: val, isSet: true}
}

func (v NullableFsEventStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFsEventStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
