// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticDataQuality Data-source quality for one diagnostic check.
type DiagnosticDataQuality struct {
	// Data quality status, such as complete, partial, failed, or not_applicable.
	Status *string `json:"status,omitempty"`
	// Human-readable reason when data is partial, failed, or not applicable.
	Reason *string `json:"reason,omitempty"`
	// Missing or unavailable data sources.
	MissingSources []string `json:"missingSources,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticDataQuality instantiates a new DiagnosticDataQuality object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticDataQuality() *DiagnosticDataQuality {
	this := DiagnosticDataQuality{}
	return &this
}

// NewDiagnosticDataQualityWithDefaults instantiates a new DiagnosticDataQuality object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticDataQualityWithDefaults() *DiagnosticDataQuality {
	this := DiagnosticDataQuality{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DiagnosticDataQuality) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDataQuality) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DiagnosticDataQuality) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DiagnosticDataQuality) SetStatus(v string) {
	o.Status = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *DiagnosticDataQuality) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDataQuality) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *DiagnosticDataQuality) HasReason() bool {
	return o != nil && o.Reason != nil
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *DiagnosticDataQuality) SetReason(v string) {
	o.Reason = &v
}

// GetMissingSources returns the MissingSources field value if set, zero value otherwise.
func (o *DiagnosticDataQuality) GetMissingSources() []string {
	if o == nil || o.MissingSources == nil {
		var ret []string
		return ret
	}
	return o.MissingSources
}

// GetMissingSourcesOk returns a tuple with the MissingSources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticDataQuality) GetMissingSourcesOk() (*[]string, bool) {
	if o == nil || o.MissingSources == nil {
		return nil, false
	}
	return &o.MissingSources, true
}

// HasMissingSources returns a boolean if a field has been set.
func (o *DiagnosticDataQuality) HasMissingSources() bool {
	return o != nil && o.MissingSources != nil
}

// SetMissingSources gets a reference to the given []string and assigns it to the MissingSources field.
func (o *DiagnosticDataQuality) SetMissingSources(v []string) {
	o.MissingSources = v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticDataQuality) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	if o.MissingSources != nil {
		toSerialize["missingSources"] = o.MissingSources
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticDataQuality) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Status         *string  `json:"status,omitempty"`
		Reason         *string  `json:"reason,omitempty"`
		MissingSources []string `json:"missingSources,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"status", "reason", "missingSources"})
	} else {
		return err
	}
	o.Status = all.Status
	o.Reason = all.Reason
	o.MissingSources = all.MissingSources

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
