// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package kbcloud

import "github.com/apecloud/kb-cloud-client-go/api/common"

type RbmqExchange struct {
	// Exchange name
	Name *string `json:"name,omitempty"`
	// VHost name
	Vhost *string `json:"vhost,omitempty"`
	// Exchange type
	Type *string `json:"type,omitempty"`
	// Durable
	Durable *bool `json:"durable,omitempty"`
	// Auto delete
	AutoDelete *bool `json:"auto_delete,omitempty"`
	// Internal
	Internal *bool `json:"internal,omitempty"`
	// Arguments
	Arguments    map[string]interface{} `json:"arguments,omitempty"`
	MessageStats *RbmqMessageStats      `json:"message_stats,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewRbmqExchange instantiates a new RbmqExchange object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewRbmqExchange() *RbmqExchange {
	this := RbmqExchange{}
	return &this
}

// NewRbmqExchangeWithDefaults instantiates a new RbmqExchange object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewRbmqExchangeWithDefaults() *RbmqExchange {
	this := RbmqExchange{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RbmqExchange) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RbmqExchange) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RbmqExchange) SetName(v string) {
	o.Name = &v
}

// GetVhost returns the Vhost field value if set, zero value otherwise.
func (o *RbmqExchange) GetVhost() string {
	if o == nil || o.Vhost == nil {
		var ret string
		return ret
	}
	return *o.Vhost
}

// GetVhostOk returns a tuple with the Vhost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetVhostOk() (*string, bool) {
	if o == nil || o.Vhost == nil {
		return nil, false
	}
	return o.Vhost, true
}

// HasVhost returns a boolean if a field has been set.
func (o *RbmqExchange) HasVhost() bool {
	return o != nil && o.Vhost != nil
}

// SetVhost gets a reference to the given string and assigns it to the Vhost field.
func (o *RbmqExchange) SetVhost(v string) {
	o.Vhost = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RbmqExchange) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RbmqExchange) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RbmqExchange) SetType(v string) {
	o.Type = &v
}

// GetDurable returns the Durable field value if set, zero value otherwise.
func (o *RbmqExchange) GetDurable() bool {
	if o == nil || o.Durable == nil {
		var ret bool
		return ret
	}
	return *o.Durable
}

// GetDurableOk returns a tuple with the Durable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetDurableOk() (*bool, bool) {
	if o == nil || o.Durable == nil {
		return nil, false
	}
	return o.Durable, true
}

// HasDurable returns a boolean if a field has been set.
func (o *RbmqExchange) HasDurable() bool {
	return o != nil && o.Durable != nil
}

// SetDurable gets a reference to the given bool and assigns it to the Durable field.
func (o *RbmqExchange) SetDurable(v bool) {
	o.Durable = &v
}

// GetAutoDelete returns the AutoDelete field value if set, zero value otherwise.
func (o *RbmqExchange) GetAutoDelete() bool {
	if o == nil || o.AutoDelete == nil {
		var ret bool
		return ret
	}
	return *o.AutoDelete
}

// GetAutoDeleteOk returns a tuple with the AutoDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetAutoDeleteOk() (*bool, bool) {
	if o == nil || o.AutoDelete == nil {
		return nil, false
	}
	return o.AutoDelete, true
}

// HasAutoDelete returns a boolean if a field has been set.
func (o *RbmqExchange) HasAutoDelete() bool {
	return o != nil && o.AutoDelete != nil
}

// SetAutoDelete gets a reference to the given bool and assigns it to the AutoDelete field.
func (o *RbmqExchange) SetAutoDelete(v bool) {
	o.AutoDelete = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *RbmqExchange) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *RbmqExchange) HasInternal() bool {
	return o != nil && o.Internal != nil
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *RbmqExchange) SetInternal(v bool) {
	o.Internal = &v
}

// GetArguments returns the Arguments field value if set, zero value otherwise.
func (o *RbmqExchange) GetArguments() map[string]interface{} {
	if o == nil || o.Arguments == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Arguments
}

// GetArgumentsOk returns a tuple with the Arguments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetArgumentsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Arguments == nil {
		return nil, false
	}
	return &o.Arguments, true
}

// HasArguments returns a boolean if a field has been set.
func (o *RbmqExchange) HasArguments() bool {
	return o != nil && o.Arguments != nil
}

// SetArguments gets a reference to the given map[string]interface{} and assigns it to the Arguments field.
func (o *RbmqExchange) SetArguments(v map[string]interface{}) {
	o.Arguments = v
}

// GetMessageStats returns the MessageStats field value if set, zero value otherwise.
func (o *RbmqExchange) GetMessageStats() RbmqMessageStats {
	if o == nil || o.MessageStats == nil {
		var ret RbmqMessageStats
		return ret
	}
	return *o.MessageStats
}

// GetMessageStatsOk returns a tuple with the MessageStats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbmqExchange) GetMessageStatsOk() (*RbmqMessageStats, bool) {
	if o == nil || o.MessageStats == nil {
		return nil, false
	}
	return o.MessageStats, true
}

// HasMessageStats returns a boolean if a field has been set.
func (o *RbmqExchange) HasMessageStats() bool {
	return o != nil && o.MessageStats != nil
}

// SetMessageStats gets a reference to the given RbmqMessageStats and assigns it to the MessageStats field.
func (o *RbmqExchange) SetMessageStats(v RbmqMessageStats) {
	o.MessageStats = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o RbmqExchange) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Vhost != nil {
		toSerialize["vhost"] = o.Vhost
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Durable != nil {
		toSerialize["durable"] = o.Durable
	}
	if o.AutoDelete != nil {
		toSerialize["auto_delete"] = o.AutoDelete
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.Arguments != nil {
		toSerialize["arguments"] = o.Arguments
	}
	if o.MessageStats != nil {
		toSerialize["message_stats"] = o.MessageStats
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *RbmqExchange) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Name         *string                `json:"name,omitempty"`
		Vhost        *string                `json:"vhost,omitempty"`
		Type         *string                `json:"type,omitempty"`
		Durable      *bool                  `json:"durable,omitempty"`
		AutoDelete   *bool                  `json:"auto_delete,omitempty"`
		Internal     *bool                  `json:"internal,omitempty"`
		Arguments    map[string]interface{} `json:"arguments,omitempty"`
		MessageStats *RbmqMessageStats      `json:"message_stats,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"name", "vhost", "type", "durable", "auto_delete", "internal", "arguments", "message_stats"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Name = all.Name
	o.Vhost = all.Vhost
	o.Type = all.Type
	o.Durable = all.Durable
	o.AutoDelete = all.AutoDelete
	o.Internal = all.Internal
	o.Arguments = all.Arguments
	if all.MessageStats != nil && all.MessageStats.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.MessageStats = all.MessageStats

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
