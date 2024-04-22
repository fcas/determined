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

// Payload struct for Payload
type Payload struct {
	Name *string `json:"name,omitempty"`
	Id *string `json:"id,omitempty"`
	Version *string `json:"version,omitempty"`
	Carriers *[]string `json:"carriers,omitempty"`
	LaunchParameters *LaunchParameters `json:"launchParameters,omitempty"`
	ResourceRequirements *ResourceRequirements `json:"resourceRequirements,omitempty"`
	AdditionalPropertiesField *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// NewPayload instantiates a new Payload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayload() *Payload {
	this := Payload{}
	return &this
}

// NewPayloadWithDefaults instantiates a new Payload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPayloadWithDefaults() *Payload {
	this := Payload{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Payload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Payload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Payload) SetName(v string) {
	o.Name = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Payload) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Payload) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Payload) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Payload) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Payload) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Payload) SetVersion(v string) {
	o.Version = &v
}

// GetCarriers returns the Carriers field value if set, zero value otherwise.
func (o *Payload) GetCarriers() []string {
	if o == nil || o.Carriers == nil {
		var ret []string
		return ret
	}
	return *o.Carriers
}

// GetCarriersOk returns a tuple with the Carriers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetCarriersOk() (*[]string, bool) {
	if o == nil || o.Carriers == nil {
		return nil, false
	}
	return o.Carriers, true
}

// HasCarriers returns a boolean if a field has been set.
func (o *Payload) HasCarriers() bool {
	if o != nil && o.Carriers != nil {
		return true
	}

	return false
}

// SetCarriers gets a reference to the given []string and assigns it to the Carriers field.
func (o *Payload) SetCarriers(v []string) {
	o.Carriers = &v
}

// GetLaunchParameters returns the LaunchParameters field value if set, zero value otherwise.
func (o *Payload) GetLaunchParameters() LaunchParameters {
	if o == nil || o.LaunchParameters == nil {
		var ret LaunchParameters
		return ret
	}
	return *o.LaunchParameters
}

// GetLaunchParametersOk returns a tuple with the LaunchParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetLaunchParametersOk() (*LaunchParameters, bool) {
	if o == nil || o.LaunchParameters == nil {
		return nil, false
	}
	return o.LaunchParameters, true
}

// HasLaunchParameters returns a boolean if a field has been set.
func (o *Payload) HasLaunchParameters() bool {
	if o != nil && o.LaunchParameters != nil {
		return true
	}

	return false
}

// SetLaunchParameters gets a reference to the given LaunchParameters and assigns it to the LaunchParameters field.
func (o *Payload) SetLaunchParameters(v LaunchParameters) {
	o.LaunchParameters = &v
}

// GetResourceRequirements returns the ResourceRequirements field value if set, zero value otherwise.
func (o *Payload) GetResourceRequirements() ResourceRequirements {
	if o == nil || o.ResourceRequirements == nil {
		var ret ResourceRequirements
		return ret
	}
	return *o.ResourceRequirements
}

// GetResourceRequirementsOk returns a tuple with the ResourceRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetResourceRequirementsOk() (*ResourceRequirements, bool) {
	if o == nil || o.ResourceRequirements == nil {
		return nil, false
	}
	return o.ResourceRequirements, true
}

// HasResourceRequirements returns a boolean if a field has been set.
func (o *Payload) HasResourceRequirements() bool {
	if o != nil && o.ResourceRequirements != nil {
		return true
	}

	return false
}

// SetResourceRequirements gets a reference to the given ResourceRequirements and assigns it to the ResourceRequirements field.
func (o *Payload) SetResourceRequirements(v ResourceRequirements) {
	o.ResourceRequirements = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *Payload) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || o.AdditionalPropertiesField == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payload) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalPropertiesField == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *Payload) HasAdditionalPropertiesField() bool {
	if o != nil && o.AdditionalPropertiesField != nil {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *Payload) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = &v
}

func (o Payload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	if o.Carriers != nil {
		toSerialize["carriers"] = o.Carriers
	}
	if o.LaunchParameters != nil {
		toSerialize["launchParameters"] = o.LaunchParameters
	}
	if o.ResourceRequirements != nil {
		toSerialize["resourceRequirements"] = o.ResourceRequirements
	}
	if o.AdditionalPropertiesField != nil {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return json.Marshal(toSerialize)
}

type NullablePayload struct {
	value *Payload
	isSet bool
}

func (v NullablePayload) Get() *Payload {
	return v.value
}

func (v *NullablePayload) Set(val *Payload) {
	v.value = val
	v.isSet = true
}

func (v NullablePayload) IsSet() bool {
	return v.isSet
}

func (v *NullablePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayload(val *Payload) *NullablePayload {
	return &NullablePayload{value: val, isSet: true}
}

func (v NullablePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


