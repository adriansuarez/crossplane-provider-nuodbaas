// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DatabaseInitParameters struct {

	// (String, Sensitive) The password for the DBA user
	// The password for the DBA user
	DbaPasswordSecretRef *v1.SecretKeySelector `json:"dbaPasswordSecretRef,omitempty" tf:"-"`

	// defined labels attached to the resource that can be used for filtering
	// User-defined labels attached to the resource that can be used for filtering
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The service tier for the database. If omitted, the project service tier is inherited.
	// The service tier for the database. If omitted, the project service tier is inherited.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type DatabaseObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// defined labels attached to the resource that can be used for filtering
	// User-defined labels attached to the resource that can be used for filtering
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the database
	// The name of the database
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) The organization that the database belongs to
	// The organization that the database belongs to
	Organization *string `json:"organization,omitempty" tf:"organization,omitempty"`

	// (String) The project that the database belongs to
	// The project that the database belongs to
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// (String) The service tier for the database. If omitted, the project service tier is inherited.
	// The service tier for the database. If omitted, the project service tier is inherited.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type DatabaseParameters struct {

	// (String, Sensitive) The password for the DBA user
	// The password for the DBA user
	// +kubebuilder:validation:Optional
	DbaPasswordSecretRef *v1.SecretKeySelector `json:"dbaPasswordSecretRef,omitempty" tf:"-"`

	// defined labels attached to the resource that can be used for filtering
	// User-defined labels attached to the resource that can be used for filtering
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// (String) The name of the database
	// The name of the database
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// (String) The organization that the database belongs to
	// The organization that the database belongs to
	// +kubebuilder:validation:Required
	Organization *string `json:"organization" tf:"organization,omitempty"`

	// (String) The project that the database belongs to
	// The project that the database belongs to
	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// (String) The service tier for the database. If omitted, the project service tier is inherited.
	// The service tier for the database. If omitted, the project service tier is inherited.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

// DatabaseSpec defines the desired state of Database
type DatabaseSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DatabaseParameters `json:"forProvider"`
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
	InitProvider DatabaseInitParameters `json:"initProvider,omitempty"`
}

// DatabaseStatus defines the observed state of Database.
type DatabaseStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DatabaseObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Database is the Schema for the Databases API. Resource for managing NuoDB databases created using the DBaaS Control Plane
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,nuodbaas}
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseSpec   `json:"spec"`
	Status            DatabaseStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DatabaseList contains a list of Databases
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}

// Repository type metadata.
var (
	Database_Kind             = "Database"
	Database_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Database_Kind}.String()
	Database_KindAPIVersion   = Database_Kind + "." + CRDGroupVersion.String()
	Database_GroupVersionKind = CRDGroupVersion.WithKind(Database_Kind)
)

func init() {
	SchemeBuilder.Register(&Database{}, &DatabaseList{})
}
