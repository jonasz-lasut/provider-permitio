/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

)




type ConditionSetRuleInitParameters struct {


// The permission that will be granted to the userset on the resourceset. The permission can be either a resource action id, or {resource_key}:{action_key}, i.e: the "permission name".
Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

// The resourceset that represents the resources that are granted for access, i.e: all the resources matching this rule can be accessed by the userset to perform the granted permission
// +crossplane:generate:reference:type=github.com/jonasz-lasut/provider-permitio/apis/permitio/v1alpha1.ResourceSet
ResourceSet *string `json:"resourceSet,omitempty" tf:"resource_set,omitempty"`

// Reference to a ResourceSet in permitio to populate resourceSet.
// +kubebuilder:validation:Optional
ResourceSetRef *v1.Reference `json:"resourceSetRef,omitempty" tf:"-"`

// Selector for a ResourceSet in permitio to populate resourceSet.
// +kubebuilder:validation:Optional
ResourceSetSelector *v1.Selector `json:"resourceSetSelector,omitempty" tf:"-"`

// The userset that will be given permission, i.e: all the users matching this rule will be given the specified permission
// +crossplane:generate:reference:type=github.com/jonasz-lasut/provider-permitio/apis/permitio/v1alpha1.UserSet
UserSet *string `json:"userSet,omitempty" tf:"user_set,omitempty"`

// Reference to a UserSet in permitio to populate userSet.
// +kubebuilder:validation:Optional
UserSetRef *v1.Reference `json:"userSetRef,omitempty" tf:"-"`

// Selector for a UserSet in permitio to populate userSet.
// +kubebuilder:validation:Optional
UserSetSelector *v1.Selector `json:"userSetSelector,omitempty" tf:"-"`
}


type ConditionSetRuleObservation struct {


// Unique id of the environment that owns the condition set rule
EnvironmentID *string `json:"environmentId,omitempty" tf:"environment_id,omitempty"`

ID *string `json:"id,omitempty" tf:"id,omitempty"`

// Unique id of the organization that owns the condition set rule
OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

// The permission that will be granted to the userset on the resourceset. The permission can be either a resource action id, or {resource_key}:{action_key}, i.e: the "permission name".
Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

// Unique id of the project that owns the condition set rule
ProjectID *string `json:"projectId,omitempty" tf:"project_id,omitempty"`

// The resourceset that represents the resources that are granted for access, i.e: all the resources matching this rule can be accessed by the userset to perform the granted permission
ResourceSet *string `json:"resourceSet,omitempty" tf:"resource_set,omitempty"`

// The userset that will be given permission, i.e: all the users matching this rule will be given the specified permission
UserSet *string `json:"userSet,omitempty" tf:"user_set,omitempty"`
}


type ConditionSetRuleParameters struct {


// The permission that will be granted to the userset on the resourceset. The permission can be either a resource action id, or {resource_key}:{action_key}, i.e: the "permission name".
// +kubebuilder:validation:Optional
Permission *string `json:"permission,omitempty" tf:"permission,omitempty"`

// The resourceset that represents the resources that are granted for access, i.e: all the resources matching this rule can be accessed by the userset to perform the granted permission
// +crossplane:generate:reference:type=github.com/jonasz-lasut/provider-permitio/apis/permitio/v1alpha1.ResourceSet
// +kubebuilder:validation:Optional
ResourceSet *string `json:"resourceSet,omitempty" tf:"resource_set,omitempty"`

// Reference to a ResourceSet in permitio to populate resourceSet.
// +kubebuilder:validation:Optional
ResourceSetRef *v1.Reference `json:"resourceSetRef,omitempty" tf:"-"`

// Selector for a ResourceSet in permitio to populate resourceSet.
// +kubebuilder:validation:Optional
ResourceSetSelector *v1.Selector `json:"resourceSetSelector,omitempty" tf:"-"`

// The userset that will be given permission, i.e: all the users matching this rule will be given the specified permission
// +crossplane:generate:reference:type=github.com/jonasz-lasut/provider-permitio/apis/permitio/v1alpha1.UserSet
// +kubebuilder:validation:Optional
UserSet *string `json:"userSet,omitempty" tf:"user_set,omitempty"`

// Reference to a UserSet in permitio to populate userSet.
// +kubebuilder:validation:Optional
UserSetRef *v1.Reference `json:"userSetRef,omitempty" tf:"-"`

// Selector for a UserSet in permitio to populate userSet.
// +kubebuilder:validation:Optional
UserSetSelector *v1.Selector `json:"userSetSelector,omitempty" tf:"-"`
}

// ConditionSetRuleSpec defines the desired state of ConditionSetRule
type ConditionSetRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       ConditionSetRuleParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider       ConditionSetRuleInitParameters `json:"initProvider,omitempty"`
}

// ConditionSetRuleStatus defines the observed state of ConditionSetRule.
type ConditionSetRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          ConditionSetRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion


// ConditionSetRule is the Schema for the ConditionSetRules API. <no value>
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,permitio}
type ConditionSetRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.permission) || (has(self.initProvider) && has(self.initProvider.permission))",message="spec.forProvider.permission is a required parameter"
	Spec              ConditionSetRuleSpec   `json:"spec"`
	Status            ConditionSetRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConditionSetRuleList contains a list of ConditionSetRules
type ConditionSetRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConditionSetRule `json:"items"`
}

// Repository type metadata.
var (
	ConditionSetRule_Kind             = "ConditionSetRule"
	ConditionSetRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ConditionSetRule_Kind}.String()
	ConditionSetRule_KindAPIVersion   = ConditionSetRule_Kind + "." + CRDGroupVersion.String()
	ConditionSetRule_GroupVersionKind = CRDGroupVersion.WithKind(ConditionSetRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ConditionSetRule{}, &ConditionSetRuleList{})
}