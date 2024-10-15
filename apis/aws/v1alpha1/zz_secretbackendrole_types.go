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

type SecretBackendRoleInitParameters struct {

	// The path the AWS secret backend is mounted at,
	// with no leading or trailing /s.
	// The path of the AWS Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of iam_user, assumed_role, or
	// federation_token.
	// Role credential type.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when credential_type is one of
	// assumed_role or federation_token.
	// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	DefaultStsTTL *float64 `json:"defaultStsTtl,omitempty" tf:"default_sts_ttl,omitempty"`

	// External ID to set for assume role creds.
	// Valid only when credential_type is set to assumed_role.
	// External ID to set for assume role creds.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of assumed_role or federation_token, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in iam_groups combined with the policy_document
	// and policy_arns parameters.
	// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	// +listType=set
	IAMGroups []*string `json:"iamGroups,omitempty" tf:"iam_groups,omitempty"`

	// A map of strings representing key/value pairs
	// to be used as tags for any IAM user that is created by this role.
	// A map of strings representing key/value pairs used as tags for any IAM user created by this role.
	// +mapType=granular
	IAMTags map[string]*string `json:"iamTags,omitempty" tf:"iam_tags,omitempty"`

	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is
	// one of assumed_role or federation_token.
	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	MaxStsTTL *float64 `json:"maxStsTtl,omitempty" tf:"max_sts_ttl,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// credential_type is iam_user. If not specified, then no permissions boundary
	// policy will be attached.
	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	PermissionsBoundaryArn *string `json:"permissionsBoundaryArn,omitempty" tf:"permissions_boundary_arn,omitempty"`

	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With iam_user, the policies will be
	// attached to IAM users when they are requested. With assumed_role and
	// federation_token, the policy ARNs will act as a filter on what the credentials
	// can do, similar to policy_document. When credential_type is iam_user or
	// federation_token, at least one of policy_document or policy_arns must
	// be specified.
	// ARN for an existing IAM policy the role should use.
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// The IAM policy document for the role. The
	// behavior depends on the credential type. With iam_user, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With assumed_role and federation_token, the policy document will
	// act as a filter on what the credentials can do, similar to policy_arns.
	// IAM policy the role should use in JSON format.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when credential_type is assumed_role and
	// prohibited otherwise.
	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'
	// +listType=set
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// A map of strings representing key/value pairs to be set
	// during assume role creds creation. Valid only when credential_type is set to
	// assumed_role.
	// Session tags to be set for assume role creds created.
	// +mapType=granular
	SessionTags map[string]*string `json:"sessionTags,omitempty" tf:"session_tags,omitempty"`

	// The path for the user name. Valid only when
	// credential_type is iam_user. Default is /.
	// The path for the user name. Valid only when credential_type is iam_user. Default is /
	UserPath *string `json:"userPath,omitempty" tf:"user_path,omitempty"`
}

type SecretBackendRoleObservation struct {

	// The path the AWS secret backend is mounted at,
	// with no leading or trailing /s.
	// The path of the AWS Secret Backend the role belongs to.
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of iam_user, assumed_role, or
	// federation_token.
	// Role credential type.
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when credential_type is one of
	// assumed_role or federation_token.
	// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	DefaultStsTTL *float64 `json:"defaultStsTtl,omitempty" tf:"default_sts_ttl,omitempty"`

	// External ID to set for assume role creds.
	// Valid only when credential_type is set to assumed_role.
	// External ID to set for assume role creds.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of assumed_role or federation_token, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in iam_groups combined with the policy_document
	// and policy_arns parameters.
	// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	// +listType=set
	IAMGroups []*string `json:"iamGroups,omitempty" tf:"iam_groups,omitempty"`

	// A map of strings representing key/value pairs
	// to be used as tags for any IAM user that is created by this role.
	// A map of strings representing key/value pairs used as tags for any IAM user created by this role.
	// +mapType=granular
	IAMTags map[string]*string `json:"iamTags,omitempty" tf:"iam_tags,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is
	// one of assumed_role or federation_token.
	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	MaxStsTTL *float64 `json:"maxStsTtl,omitempty" tf:"max_sts_ttl,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// credential_type is iam_user. If not specified, then no permissions boundary
	// policy will be attached.
	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	PermissionsBoundaryArn *string `json:"permissionsBoundaryArn,omitempty" tf:"permissions_boundary_arn,omitempty"`

	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With iam_user, the policies will be
	// attached to IAM users when they are requested. With assumed_role and
	// federation_token, the policy ARNs will act as a filter on what the credentials
	// can do, similar to policy_document. When credential_type is iam_user or
	// federation_token, at least one of policy_document or policy_arns must
	// be specified.
	// ARN for an existing IAM policy the role should use.
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// The IAM policy document for the role. The
	// behavior depends on the credential type. With iam_user, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With assumed_role and federation_token, the policy document will
	// act as a filter on what the credentials can do, similar to policy_arns.
	// IAM policy the role should use in JSON format.
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when credential_type is assumed_role and
	// prohibited otherwise.
	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'
	// +listType=set
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// A map of strings representing key/value pairs to be set
	// during assume role creds creation. Valid only when credential_type is set to
	// assumed_role.
	// Session tags to be set for assume role creds created.
	// +mapType=granular
	SessionTags map[string]*string `json:"sessionTags,omitempty" tf:"session_tags,omitempty"`

	// The path for the user name. Valid only when
	// credential_type is iam_user. Default is /.
	// The path for the user name. Valid only when credential_type is iam_user. Default is /
	UserPath *string `json:"userPath,omitempty" tf:"user_path,omitempty"`
}

type SecretBackendRoleParameters struct {

	// The path the AWS secret backend is mounted at,
	// with no leading or trailing /s.
	// The path of the AWS Secret Backend the role belongs to.
	// +kubebuilder:validation:Optional
	Backend *string `json:"backend,omitempty" tf:"backend,omitempty"`

	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of iam_user, assumed_role, or
	// federation_token.
	// Role credential type.
	// +kubebuilder:validation:Optional
	CredentialType *string `json:"credentialType,omitempty" tf:"credential_type,omitempty"`

	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when credential_type is one of
	// assumed_role or federation_token.
	// The default TTL in seconds for STS credentials. When a TTL is not specified when STS credentials are requested, and a default TTL is specified on the role, then this default TTL will be used. Valid only when credential_type is one of assumed_role or federation_token.
	// +kubebuilder:validation:Optional
	DefaultStsTTL *float64 `json:"defaultStsTtl,omitempty" tf:"default_sts_ttl,omitempty"`

	// External ID to set for assume role creds.
	// Valid only when credential_type is set to assumed_role.
	// External ID to set for assume role creds.
	// +kubebuilder:validation:Optional
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of assumed_role or federation_token, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in iam_groups combined with the policy_document
	// and policy_arns parameters.
	// A list of IAM group names. IAM users generated against this vault role will be added to these IAM Groups. For a credential type of assumed_role or federation_token, the policies sent to the corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the policies from each group in iam_groups combined with the policy_document and policy_arns parameters.
	// +kubebuilder:validation:Optional
	// +listType=set
	IAMGroups []*string `json:"iamGroups,omitempty" tf:"iam_groups,omitempty"`

	// A map of strings representing key/value pairs
	// to be used as tags for any IAM user that is created by this role.
	// A map of strings representing key/value pairs used as tags for any IAM user created by this role.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	IAMTags map[string]*string `json:"iamTags,omitempty" tf:"iam_tags,omitempty"`

	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is
	// one of assumed_role or federation_token.
	// The max allowed TTL in seconds for STS credentials (credentials TTL are capped to max_sts_ttl). Valid only when credential_type is one of assumed_role or federation_token.
	// +kubebuilder:validation:Optional
	MaxStsTTL *float64 `json:"maxStsTtl,omitempty" tf:"max_sts_ttl,omitempty"`

	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The namespace is always relative to the provider's configured namespace.
	// Available only for Vault Enterprise.
	// Target namespace. (requires Enterprise)
	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// credential_type is iam_user. If not specified, then no permissions boundary
	// policy will be attached.
	// The ARN of the AWS Permissions Boundary to attach to IAM users created in the role. Valid only when credential_type is iam_user. If not specified, then no permissions boundary policy will be attached.
	// +kubebuilder:validation:Optional
	PermissionsBoundaryArn *string `json:"permissionsBoundaryArn,omitempty" tf:"permissions_boundary_arn,omitempty"`

	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With iam_user, the policies will be
	// attached to IAM users when they are requested. With assumed_role and
	// federation_token, the policy ARNs will act as a filter on what the credentials
	// can do, similar to policy_document. When credential_type is iam_user or
	// federation_token, at least one of policy_document or policy_arns must
	// be specified.
	// ARN for an existing IAM policy the role should use.
	// +kubebuilder:validation:Optional
	// +listType=set
	PolicyArns []*string `json:"policyArns,omitempty" tf:"policy_arns,omitempty"`

	// The IAM policy document for the role. The
	// behavior depends on the credential type. With iam_user, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With assumed_role and federation_token, the policy document will
	// act as a filter on what the credentials can do, similar to policy_arns.
	// IAM policy the role should use in JSON format.
	// +kubebuilder:validation:Optional
	PolicyDocument *string `json:"policyDocument,omitempty" tf:"policy_document,omitempty"`

	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when credential_type is assumed_role and
	// prohibited otherwise.
	// ARNs of AWS roles allowed to be assumed. Only valid when credential_type is 'assumed_role'
	// +kubebuilder:validation:Optional
	// +listType=set
	RoleArns []*string `json:"roleArns,omitempty" tf:"role_arns,omitempty"`

	// A map of strings representing key/value pairs to be set
	// during assume role creds creation. Valid only when credential_type is set to
	// assumed_role.
	// Session tags to be set for assume role creds created.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	SessionTags map[string]*string `json:"sessionTags,omitempty" tf:"session_tags,omitempty"`

	// The path for the user name. Valid only when
	// credential_type is iam_user. Default is /.
	// The path for the user name. Valid only when credential_type is iam_user. Default is /
	// +kubebuilder:validation:Optional
	UserPath *string `json:"userPath,omitempty" tf:"user_path,omitempty"`
}

// SecretBackendRoleSpec defines the desired state of SecretBackendRole
type SecretBackendRoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecretBackendRoleParameters `json:"forProvider"`
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
	InitProvider SecretBackendRoleInitParameters `json:"initProvider,omitempty"`
}

// SecretBackendRoleStatus defines the observed state of SecretBackendRole.
type SecretBackendRoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecretBackendRoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// SecretBackendRole is the Schema for the SecretBackendRoles API. Creates a role on an AWS Secret Backend for Vault.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,vault}
type SecretBackendRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.backend) || (has(self.initProvider) && has(self.initProvider.backend))",message="spec.forProvider.backend is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.credentialType) || (has(self.initProvider) && has(self.initProvider.credentialType))",message="spec.forProvider.credentialType is a required parameter"
	Spec   SecretBackendRoleSpec   `json:"spec"`
	Status SecretBackendRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretBackendRoleList contains a list of SecretBackendRoles
type SecretBackendRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretBackendRole `json:"items"`
}

// Repository type metadata.
var (
	SecretBackendRole_Kind             = "SecretBackendRole"
	SecretBackendRole_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecretBackendRole_Kind}.String()
	SecretBackendRole_KindAPIVersion   = SecretBackendRole_Kind + "." + CRDGroupVersion.String()
	SecretBackendRole_GroupVersionKind = CRDGroupVersion.WithKind(SecretBackendRole_Kind)
)

func init() {
	SchemeBuilder.Register(&SecretBackendRole{}, &SecretBackendRoleList{})
}
