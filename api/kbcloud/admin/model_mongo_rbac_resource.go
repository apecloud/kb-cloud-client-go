// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at ApeCloud (https://www.apecloud.com/).
// Copyright 2022-Present ApeCloud Co., Ltd

package admin

import "github.com/apecloud/kb-cloud-client-go/api/common"

// MongoRBACResource Native MongoDB resource. cluster:true is system operations, not all database data. Empty db/collection are explicit wildcard scopes; system collections need explicit resources.
type MongoRBACResource struct {
	Cluster *bool `json:"cluster,omitempty"`
	// Preserve an explicit empty string for native wildcard scope; omit only for cluster resources.
	Db *string `json:"db,omitempty"`
	// Preserve an explicit empty string for native wildcard scope; omit only for cluster resources.
	Collection *string `json:"collection,omitempty"`
	// Read-only native evidence; write requests reject this internal resource.
	AnyResource *bool `json:"anyResource,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMongoRBACResource instantiates a new MongoRBACResource object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMongoRBACResource() *MongoRBACResource {
	this := MongoRBACResource{}
	return &this
}

// NewMongoRBACResourceWithDefaults instantiates a new MongoRBACResource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMongoRBACResourceWithDefaults() *MongoRBACResource {
	this := MongoRBACResource{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *MongoRBACResource) GetCluster() bool {
	if o == nil || o.Cluster == nil {
		var ret bool
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACResource) GetClusterOk() (*bool, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *MongoRBACResource) HasCluster() bool {
	return o != nil && o.Cluster != nil
}

// SetCluster gets a reference to the given bool and assigns it to the Cluster field.
func (o *MongoRBACResource) SetCluster(v bool) {
	o.Cluster = &v
}

// GetDb returns the Db field value if set, zero value otherwise.
func (o *MongoRBACResource) GetDb() string {
	if o == nil || o.Db == nil {
		var ret string
		return ret
	}
	return *o.Db
}

// GetDbOk returns a tuple with the Db field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACResource) GetDbOk() (*string, bool) {
	if o == nil || o.Db == nil {
		return nil, false
	}
	return o.Db, true
}

// HasDb returns a boolean if a field has been set.
func (o *MongoRBACResource) HasDb() bool {
	return o != nil && o.Db != nil
}

// SetDb gets a reference to the given string and assigns it to the Db field.
func (o *MongoRBACResource) SetDb(v string) {
	o.Db = &v
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *MongoRBACResource) GetCollection() string {
	if o == nil || o.Collection == nil {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACResource) GetCollectionOk() (*string, bool) {
	if o == nil || o.Collection == nil {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *MongoRBACResource) HasCollection() bool {
	return o != nil && o.Collection != nil
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *MongoRBACResource) SetCollection(v string) {
	o.Collection = &v
}

// GetAnyResource returns the AnyResource field value if set, zero value otherwise.
func (o *MongoRBACResource) GetAnyResource() bool {
	if o == nil || o.AnyResource == nil {
		var ret bool
		return ret
	}
	return *o.AnyResource
}

// GetAnyResourceOk returns a tuple with the AnyResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MongoRBACResource) GetAnyResourceOk() (*bool, bool) {
	if o == nil || o.AnyResource == nil {
		return nil, false
	}
	return o.AnyResource, true
}

// HasAnyResource returns a boolean if a field has been set.
func (o *MongoRBACResource) HasAnyResource() bool {
	return o != nil && o.AnyResource != nil
}

// SetAnyResource gets a reference to the given bool and assigns it to the AnyResource field.
func (o *MongoRBACResource) SetAnyResource(v bool) {
	o.AnyResource = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MongoRBACResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return common.Marshal(o.UnparsedObject)
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.Db != nil {
		toSerialize["db"] = o.Db
	}
	if o.Collection != nil {
		toSerialize["collection"] = o.Collection
	}
	if o.AnyResource != nil {
		toSerialize["anyResource"] = o.AnyResource
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return common.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MongoRBACResource) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Cluster     *bool   `json:"cluster,omitempty"`
		Db          *string `json:"db,omitempty"`
		Collection  *string `json:"collection,omitempty"`
		AnyResource *bool   `json:"anyResource,omitempty"`
	}{}
	if err = common.Unmarshal(bytes, &all); err != nil {
		return err
	}
	additionalProperties := make(map[string]interface{})
	if err = common.Unmarshal(bytes, &additionalProperties); err == nil {
		common.DeleteKeys(additionalProperties, &[]string{"cluster", "db", "collection", "anyResource"})
	} else {
		return err
	}
	o.Cluster = all.Cluster
	o.Db = all.Db
	o.Collection = all.Collection
	o.AnyResource = all.AnyResource

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
