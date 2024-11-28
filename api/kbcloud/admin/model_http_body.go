// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

// HttpBody Represents an HTTP request or response body.
type HttpBody struct {
	// The HTTP Content-Type header value specifying the content type of the body.
	ContentType *string `json:"contentType,omitempty"`
	// The HTTP request/response body as raw binary.
	Data *string `json:"data,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewHttpBody instantiates a new HttpBody object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewHttpBody() *HttpBody {
	this := HttpBody{}
	return &this
}

// NewHttpBodyWithDefaults instantiates a new HttpBody object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewHttpBodyWithDefaults() *HttpBody {
	this := HttpBody{}
	return &this
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *HttpBody) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpBody) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *HttpBody) HasContentType() bool {
	return o != nil && o.ContentType != nil
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *HttpBody) SetContentType(v string) {
	o.ContentType = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *HttpBody) GetData() string {
	if o == nil || o.Data == nil {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpBody) GetDataOk() (*string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *HttpBody) HasData() bool {
	return o != nil && o.Data != nil
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *HttpBody) SetData(v string) {
	o.Data = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o HttpBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.ContentType != nil {
		toSerialize["contentType"] = o.ContentType
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *HttpBody) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		ContentType *string `json:"contentType,omitempty"`
		Data        *string `json:"data,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return common.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"contentType", "data"})
	} else {
		return err
	}
	o.ContentType = all.ContentType
	o.Data = all.Data

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
