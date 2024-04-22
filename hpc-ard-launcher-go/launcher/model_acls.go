/*
Launcher API

The Launcher API is the execution layer for the Capsules framework.  It handles all the details of launching and monitoring runtime environments.

API version: 3.3.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package launcher

import (
	"encoding/json"
)

// ACLS struct for ACLS
type ACLS struct {
	Read *ACLInfo `json:"read,omitempty"`
	Write *ACLInfo `json:"write,omitempty"`
	Admin *ACLInfo `json:"admin,omitempty"`
	AdditionalPropertiesField *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// NewACLS instantiates a new ACLS object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLS() *ACLS {
	this := ACLS{}
	return &this
}

// NewACLSWithDefaults instantiates a new ACLS object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLSWithDefaults() *ACLS {
	this := ACLS{}
	return &this
}

// GetRead returns the Read field value if set, zero value otherwise.
func (o *ACLS) GetRead() ACLInfo {
	if o == nil || o.Read == nil {
		var ret ACLInfo
		return ret
	}
	return *o.Read
}

// GetReadOk returns a tuple with the Read field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLS) GetReadOk() (*ACLInfo, bool) {
	if o == nil || o.Read == nil {
		return nil, false
	}
	return o.Read, true
}

// HasRead returns a boolean if a field has been set.
func (o *ACLS) HasRead() bool {
	if o != nil && o.Read != nil {
		return true
	}

	return false
}

// SetRead gets a reference to the given ACLInfo and assigns it to the Read field.
func (o *ACLS) SetRead(v ACLInfo) {
	o.Read = &v
}

// GetWrite returns the Write field value if set, zero value otherwise.
func (o *ACLS) GetWrite() ACLInfo {
	if o == nil || o.Write == nil {
		var ret ACLInfo
		return ret
	}
	return *o.Write
}

// GetWriteOk returns a tuple with the Write field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLS) GetWriteOk() (*ACLInfo, bool) {
	if o == nil || o.Write == nil {
		return nil, false
	}
	return o.Write, true
}

// HasWrite returns a boolean if a field has been set.
func (o *ACLS) HasWrite() bool {
	if o != nil && o.Write != nil {
		return true
	}

	return false
}

// SetWrite gets a reference to the given ACLInfo and assigns it to the Write field.
func (o *ACLS) SetWrite(v ACLInfo) {
	o.Write = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *ACLS) GetAdmin() ACLInfo {
	if o == nil || o.Admin == nil {
		var ret ACLInfo
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLS) GetAdminOk() (*ACLInfo, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *ACLS) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given ACLInfo and assigns it to the Admin field.
func (o *ACLS) SetAdmin(v ACLInfo) {
	o.Admin = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *ACLS) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || o.AdditionalPropertiesField == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLS) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalPropertiesField == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *ACLS) HasAdditionalPropertiesField() bool {
	if o != nil && o.AdditionalPropertiesField != nil {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *ACLS) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = &v
}

func (o ACLS) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Read != nil {
		toSerialize["read"] = o.Read
	}
	if o.Write != nil {
		toSerialize["write"] = o.Write
	}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.AdditionalPropertiesField != nil {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return json.Marshal(toSerialize)
}

type NullableACLS struct {
	value *ACLS
	isSet bool
}

func (v NullableACLS) Get() *ACLS {
	return v.value
}

func (v *NullableACLS) Set(val *ACLS) {
	v.value = val
	v.isSet = true
}

func (v NullableACLS) IsSet() bool {
	return v.isSet
}

func (v *NullableACLS) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLS(val *ACLS) *NullableACLS {
	return &NullableACLS{value: val, isSet: true}
}

func (v NullableACLS) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLS) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


