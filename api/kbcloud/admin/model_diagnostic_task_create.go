// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// DiagnosticTaskCreate Request to start a diagnostic task for one cluster.
type DiagnosticTaskCreate struct {
	// Optional rule IDs to run. Empty means run the default P0 rule set.
	RuleIds []string `json:"ruleIds,omitempty"`
	// Whether to collect recent diagnostic log snippets.
	IncludeLogs *bool `json:"includeLogs,omitempty"`
	// Whether to collect Kubernetes events.
	IncludeEvents *bool `json:"includeEvents,omitempty"`
	// Optional human-readable reason for this diagnostic task.
	Description *string `json:"description,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewDiagnosticTaskCreate instantiates a new DiagnosticTaskCreate object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewDiagnosticTaskCreate() *DiagnosticTaskCreate {
	this := DiagnosticTaskCreate{}
	var includeLogs bool = true
	this.IncludeLogs = &includeLogs
	var includeEvents bool = true
	this.IncludeEvents = &includeEvents
	return &this
}

// NewDiagnosticTaskCreateWithDefaults instantiates a new DiagnosticTaskCreate object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewDiagnosticTaskCreateWithDefaults() *DiagnosticTaskCreate {
	this := DiagnosticTaskCreate{}
	var includeLogs bool = true
	this.IncludeLogs = &includeLogs
	var includeEvents bool = true
	this.IncludeEvents = &includeEvents
	return &this
}

// GetRuleIds returns the RuleIds field value if set, zero value otherwise.
func (o *DiagnosticTaskCreate) GetRuleIds() []string {
	if o == nil || o.RuleIds == nil {
		var ret []string
		return ret
	}
	return o.RuleIds
}

// GetRuleIdsOk returns a tuple with the RuleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTaskCreate) GetRuleIdsOk() (*[]string, bool) {
	if o == nil || o.RuleIds == nil {
		return nil, false
	}
	return &o.RuleIds, true
}

// HasRuleIds returns a boolean if a field has been set.
func (o *DiagnosticTaskCreate) HasRuleIds() bool {
	return o != nil && o.RuleIds != nil
}

// SetRuleIds gets a reference to the given []string and assigns it to the RuleIds field.
func (o *DiagnosticTaskCreate) SetRuleIds(v []string) {
	o.RuleIds = v
}

// GetIncludeLogs returns the IncludeLogs field value if set, zero value otherwise.
func (o *DiagnosticTaskCreate) GetIncludeLogs() bool {
	if o == nil || o.IncludeLogs == nil {
		var ret bool
		return ret
	}
	return *o.IncludeLogs
}

// GetIncludeLogsOk returns a tuple with the IncludeLogs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTaskCreate) GetIncludeLogsOk() (*bool, bool) {
	if o == nil || o.IncludeLogs == nil {
		return nil, false
	}
	return o.IncludeLogs, true
}

// HasIncludeLogs returns a boolean if a field has been set.
func (o *DiagnosticTaskCreate) HasIncludeLogs() bool {
	return o != nil && o.IncludeLogs != nil
}

// SetIncludeLogs gets a reference to the given bool and assigns it to the IncludeLogs field.
func (o *DiagnosticTaskCreate) SetIncludeLogs(v bool) {
	o.IncludeLogs = &v
}

// GetIncludeEvents returns the IncludeEvents field value if set, zero value otherwise.
func (o *DiagnosticTaskCreate) GetIncludeEvents() bool {
	if o == nil || o.IncludeEvents == nil {
		var ret bool
		return ret
	}
	return *o.IncludeEvents
}

// GetIncludeEventsOk returns a tuple with the IncludeEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTaskCreate) GetIncludeEventsOk() (*bool, bool) {
	if o == nil || o.IncludeEvents == nil {
		return nil, false
	}
	return o.IncludeEvents, true
}

// HasIncludeEvents returns a boolean if a field has been set.
func (o *DiagnosticTaskCreate) HasIncludeEvents() bool {
	return o != nil && o.IncludeEvents != nil
}

// SetIncludeEvents gets a reference to the given bool and assigns it to the IncludeEvents field.
func (o *DiagnosticTaskCreate) SetIncludeEvents(v bool) {
	o.IncludeEvents = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DiagnosticTaskCreate) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiagnosticTaskCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DiagnosticTaskCreate) HasDescription() bool {
	return o != nil && o.Description != nil
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DiagnosticTaskCreate) SetDescription(v string) {
	o.Description = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o DiagnosticTaskCreate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.RuleIds != nil {
		toSerialize["ruleIds"] = o.RuleIds
	}
	if o.IncludeLogs != nil {
		toSerialize["includeLogs"] = o.IncludeLogs
	}
	if o.IncludeEvents != nil {
		toSerialize["includeEvents"] = o.IncludeEvents
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *DiagnosticTaskCreate) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		RuleIds       []string `json:"ruleIds,omitempty"`
		IncludeLogs   *bool    `json:"includeLogs,omitempty"`
		IncludeEvents *bool    `json:"includeEvents,omitempty"`
		Description   *string  `json:"description,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"ruleIds", "includeLogs", "includeEvents", "description"})
	} else {
		return err
	}
	o.RuleIds = all.RuleIds
	o.IncludeLogs = all.IncludeLogs
	o.IncludeEvents = all.IncludeEvents
	o.Description = all.Description

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
