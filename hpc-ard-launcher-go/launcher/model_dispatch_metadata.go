/*
Launcher API

The Launcher API is the execution layer for the Capsules framework.  It handles all the details of launching and monitoring runtime environments.

API version: 3.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package launcher

import (
	"encoding/json"
)

// DispatchMetadata struct for DispatchMetadata
type DispatchMetadata struct {
	Owner *string `json:"owner,omitempty"`
	Dispatcher *string `json:"dispatcher,omitempty"`
	Carriers *map[string]string `json:"carriers,omitempty"`
	Launched *string `json:"launched,omitempty"`
	Terminated *string `json:"terminated,omitempty"`
	UserInterfaces *[]UserInterface `json:"userInterfaces,omitempty"`
	AdditionalPropertiesField *map[string]interface{} `json:"additionalProperties,omitempty"`
}

// NewDispatchMetadata instantiates a new DispatchMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDispatchMetadata() *DispatchMetadata {
	this := DispatchMetadata{}
	return &this
}

// NewDispatchMetadataWithDefaults instantiates a new DispatchMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDispatchMetadataWithDefaults() *DispatchMetadata {
	this := DispatchMetadata{}
	return &this
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *DispatchMetadata) GetOwner() string {
	if o == nil || o.Owner == nil {
		var ret string
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetOwnerOk() (*string, bool) {
	if o == nil || o.Owner == nil {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *DispatchMetadata) HasOwner() bool {
	if o != nil && o.Owner != nil {
		return true
	}

	return false
}

// SetOwner gets a reference to the given string and assigns it to the Owner field.
func (o *DispatchMetadata) SetOwner(v string) {
	o.Owner = &v
}

// GetDispatcher returns the Dispatcher field value if set, zero value otherwise.
func (o *DispatchMetadata) GetDispatcher() string {
	if o == nil || o.Dispatcher == nil {
		var ret string
		return ret
	}
	return *o.Dispatcher
}

// GetDispatcherOk returns a tuple with the Dispatcher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetDispatcherOk() (*string, bool) {
	if o == nil || o.Dispatcher == nil {
		return nil, false
	}
	return o.Dispatcher, true
}

// HasDispatcher returns a boolean if a field has been set.
func (o *DispatchMetadata) HasDispatcher() bool {
	if o != nil && o.Dispatcher != nil {
		return true
	}

	return false
}

// SetDispatcher gets a reference to the given string and assigns it to the Dispatcher field.
func (o *DispatchMetadata) SetDispatcher(v string) {
	o.Dispatcher = &v
}

// GetCarriers returns the Carriers field value if set, zero value otherwise.
func (o *DispatchMetadata) GetCarriers() map[string]string {
	if o == nil || o.Carriers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Carriers
}

// GetCarriersOk returns a tuple with the Carriers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetCarriersOk() (*map[string]string, bool) {
	if o == nil || o.Carriers == nil {
		return nil, false
	}
	return o.Carriers, true
}

// HasCarriers returns a boolean if a field has been set.
func (o *DispatchMetadata) HasCarriers() bool {
	if o != nil && o.Carriers != nil {
		return true
	}

	return false
}

// SetCarriers gets a reference to the given map[string]string and assigns it to the Carriers field.
func (o *DispatchMetadata) SetCarriers(v map[string]string) {
	o.Carriers = &v
}

// GetLaunched returns the Launched field value if set, zero value otherwise.
func (o *DispatchMetadata) GetLaunched() string {
	if o == nil || o.Launched == nil {
		var ret string
		return ret
	}
	return *o.Launched
}

// GetLaunchedOk returns a tuple with the Launched field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetLaunchedOk() (*string, bool) {
	if o == nil || o.Launched == nil {
		return nil, false
	}
	return o.Launched, true
}

// HasLaunched returns a boolean if a field has been set.
func (o *DispatchMetadata) HasLaunched() bool {
	if o != nil && o.Launched != nil {
		return true
	}

	return false
}

// SetLaunched gets a reference to the given string and assigns it to the Launched field.
func (o *DispatchMetadata) SetLaunched(v string) {
	o.Launched = &v
}

// GetTerminated returns the Terminated field value if set, zero value otherwise.
func (o *DispatchMetadata) GetTerminated() string {
	if o == nil || o.Terminated == nil {
		var ret string
		return ret
	}
	return *o.Terminated
}

// GetTerminatedOk returns a tuple with the Terminated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetTerminatedOk() (*string, bool) {
	if o == nil || o.Terminated == nil {
		return nil, false
	}
	return o.Terminated, true
}

// HasTerminated returns a boolean if a field has been set.
func (o *DispatchMetadata) HasTerminated() bool {
	if o != nil && o.Terminated != nil {
		return true
	}

	return false
}

// SetTerminated gets a reference to the given string and assigns it to the Terminated field.
func (o *DispatchMetadata) SetTerminated(v string) {
	o.Terminated = &v
}

// GetUserInterfaces returns the UserInterfaces field value if set, zero value otherwise.
func (o *DispatchMetadata) GetUserInterfaces() []UserInterface {
	if o == nil || o.UserInterfaces == nil {
		var ret []UserInterface
		return ret
	}
	return *o.UserInterfaces
}

// GetUserInterfacesOk returns a tuple with the UserInterfaces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetUserInterfacesOk() (*[]UserInterface, bool) {
	if o == nil || o.UserInterfaces == nil {
		return nil, false
	}
	return o.UserInterfaces, true
}

// HasUserInterfaces returns a boolean if a field has been set.
func (o *DispatchMetadata) HasUserInterfaces() bool {
	if o != nil && o.UserInterfaces != nil {
		return true
	}

	return false
}

// SetUserInterfaces gets a reference to the given []UserInterface and assigns it to the UserInterfaces field.
func (o *DispatchMetadata) SetUserInterfaces(v []UserInterface) {
	o.UserInterfaces = &v
}

// GetAdditionalPropertiesField returns the AdditionalPropertiesField field value if set, zero value otherwise.
func (o *DispatchMetadata) GetAdditionalPropertiesField() map[string]interface{} {
	if o == nil || o.AdditionalPropertiesField == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.AdditionalPropertiesField
}

// GetAdditionalPropertiesFieldOk returns a tuple with the AdditionalPropertiesField field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DispatchMetadata) GetAdditionalPropertiesFieldOk() (*map[string]interface{}, bool) {
	if o == nil || o.AdditionalPropertiesField == nil {
		return nil, false
	}
	return o.AdditionalPropertiesField, true
}

// HasAdditionalPropertiesField returns a boolean if a field has been set.
func (o *DispatchMetadata) HasAdditionalPropertiesField() bool {
	if o != nil && o.AdditionalPropertiesField != nil {
		return true
	}

	return false
}

// SetAdditionalPropertiesField gets a reference to the given map[string]interface{} and assigns it to the AdditionalPropertiesField field.
func (o *DispatchMetadata) SetAdditionalPropertiesField(v map[string]interface{}) {
	o.AdditionalPropertiesField = &v
}

func (o DispatchMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Owner != nil {
		toSerialize["owner"] = o.Owner
	}
	if o.Dispatcher != nil {
		toSerialize["dispatcher"] = o.Dispatcher
	}
	if o.Carriers != nil {
		toSerialize["carriers"] = o.Carriers
	}
	if o.Launched != nil {
		toSerialize["launched"] = o.Launched
	}
	if o.Terminated != nil {
		toSerialize["terminated"] = o.Terminated
	}
	if o.UserInterfaces != nil {
		toSerialize["userInterfaces"] = o.UserInterfaces
	}
	if o.AdditionalPropertiesField != nil {
		toSerialize["additionalProperties"] = o.AdditionalPropertiesField
	}
	return json.Marshal(toSerialize)
}

type NullableDispatchMetadata struct {
	value *DispatchMetadata
	isSet bool
}

func (v NullableDispatchMetadata) Get() *DispatchMetadata {
	return v.value
}

func (v *NullableDispatchMetadata) Set(val *DispatchMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDispatchMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDispatchMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispatchMetadata(val *DispatchMetadata) *NullableDispatchMetadata {
	return &NullableDispatchMetadata{value: val, isSet: true}
}

func (v NullableDispatchMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispatchMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


