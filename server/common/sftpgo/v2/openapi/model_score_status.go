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

// ScoreStatus struct for ScoreStatus
type ScoreStatus struct {
	// if 0 the host is not listed
	Score *int32 `json:"score,omitempty"`
}

// NewScoreStatus instantiates a new ScoreStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScoreStatus() *ScoreStatus {
	this := ScoreStatus{}
	return &this
}

// NewScoreStatusWithDefaults instantiates a new ScoreStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScoreStatusWithDefaults() *ScoreStatus {
	this := ScoreStatus{}
	return &this
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *ScoreStatus) GetScore() int32 {
	if o == nil || o.Score == nil {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScoreStatus) GetScoreOk() (*int32, bool) {
	if o == nil || o.Score == nil {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *ScoreStatus) HasScore() bool {
	if o != nil && o.Score != nil {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *ScoreStatus) SetScore(v int32) {
	o.Score = &v
}

func (o ScoreStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Score != nil {
		toSerialize["score"] = o.Score
	}
	return json.Marshal(toSerialize)
}

type NullableScoreStatus struct {
	value *ScoreStatus
	isSet bool
}

func (v NullableScoreStatus) Get() *ScoreStatus {
	return v.value
}

func (v *NullableScoreStatus) Set(val *ScoreStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableScoreStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableScoreStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScoreStatus(val *ScoreStatus) *NullableScoreStatus {
	return &NullableScoreStatus{value: val, isSet: true}
}

func (v NullableScoreStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScoreStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

