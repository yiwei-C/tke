/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_APIKey = map[string]string{
	"":     "APIKey contains expiration time used to apply the api key.",
	"spec": "Spec defines the desired identities of APIkey in this set.",
}

func (APIKey) SwaggerDoc() map[string]string {
	return map_APIKey
}

var map_APIKeyList = map[string]string{
	"":      "APIKeyList is the whole list of all identities.",
	"items": "List of api keys.",
}

func (APIKeyList) SwaggerDoc() map[string]string {
	return map_APIKeyList
}

var map_APIKeyReq = map[string]string{
	"":            "APIKeyReq contains expiration time used to apply the api key.",
	"expire":      "Expire is required, holds the duration of the api key become invalid. By default, 168h(= seven days)",
	"description": "Description describes api keys usage.",
}

func (APIKeyReq) SwaggerDoc() map[string]string {
	return map_APIKeyReq
}

var map_APIKeyReqPassword = map[string]string{
	"":            "APIKeyReqPassword contains userinfo and expiration time used to apply the api key.",
	"tenantID":    "TenantID for user",
	"username":    "Username",
	"password":    "Password (encoded by base64)",
	"description": "Description describes api keys usage.",
	"expire":      "Expire holds the duration of the api key become invalid. By default, 168h(= seven days)",
}

func (APIKeyReqPassword) SwaggerDoc() map[string]string {
	return map_APIKeyReqPassword
}

var map_APIKeySpec = map[string]string{
	"":            "APIKeySpec is a description of an apiKey.",
	"apiKey":      "APIkey is the jwt token used to authenticate user, and contains user info and sign.",
	"username":    "Username is creator",
	"description": "Description describes api keys usage.",
	"issue_at":    "IssueAt is the created time for api key",
	"expire_at":   "ExpireAt is the expire time for api key",
}

func (APIKeySpec) SwaggerDoc() map[string]string {
	return map_APIKeySpec
}

var map_APIKeyStatus = map[string]string{
	"":         "APIKeyStatus is a description of an api key status.",
	"disabled": "Disabled represents whether the apikey has been disabled.",
	"expired":  "Expired represents whether the apikey has been expired.",
}

func (APIKeyStatus) SwaggerDoc() map[string]string {
	return map_APIKeyStatus
}

var map_APISigningKey = map[string]string{
	"": "APISigningKey hold encryption and signing key.",
}

func (APISigningKey) SwaggerDoc() map[string]string {
	return map_APISigningKey
}

var map_APISigningKeyList = map[string]string{
	"":      "APISigningKeyList is the whole list of all signing key.",
	"items": "List of keys.",
}

func (APISigningKeyList) SwaggerDoc() map[string]string {
	return map_APISigningKeyList
}

var map_Action = map[string]string{
	"":            "Action defines a action verb for authorization.",
	"name":        "Name represents user access review request verb.",
	"description": "Description describes the action.",
}

func (Action) SwaggerDoc() map[string]string {
	return map_Action
}

var map_AllowedStatus = map[string]string{
	"":                "AllowedStatus includes the resource access request and response.",
	"resource":        "Resource is the resource of request",
	"web":             "Verb is the verb of request",
	"allowed":         "Allowed is required. True if the action would be allowed, false otherwise.",
	"denied":          "Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.",
	"reason":          "Reason is optional.  It indicates why a request was allowed or denied.",
	"evaluationError": "EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.",
}

func (AllowedStatus) SwaggerDoc() map[string]string {
	return map_AllowedStatus
}

var map_Binding = map[string]string{
	"":       "Binding is used to bind or unbind the subjects to or from the policy,role or group.",
	"users":  "Users holds references to the objects the policy applies to.",
	"groups": "Groups holds references to the groups the policy applies to.",
}

func (Binding) SwaggerDoc() map[string]string {
	return map_Binding
}

var map_Category = map[string]string{
	"": "Category defines a category of actions for policy.",
}

func (Category) SwaggerDoc() map[string]string {
	return map_Category
}

var map_CategoryList = map[string]string{
	"":      "CategoryList is the whole list of policy Category.",
	"items": "List of category.",
}

func (CategoryList) SwaggerDoc() map[string]string {
	return map_CategoryList
}

var map_CategorySpec = map[string]string{
	"":            "CategorySpec is a description of category.",
	"displayName": "DisplayName used to display category name",
	"actions":     "Actions represents a series of actions work on the policy category",
}

func (CategorySpec) SwaggerDoc() map[string]string {
	return map_CategorySpec
}

var map_Client = map[string]string{
	"":     "Client represents an OAuth2 client.",
	"spec": "Spec defines the desired identities of identity provider in this set.",
}

func (Client) SwaggerDoc() map[string]string {
	return map_Client
}

var map_ClientList = map[string]string{
	"":      "ClientList is the whole list of OAuth2 client.",
	"items": "List of identity providers.",
}

func (ClientList) SwaggerDoc() map[string]string {
	return map_ClientList
}

var map_ClientSpec = map[string]string{
	"":              "ClientSpec is a description of an client.",
	"trusted_peers": "TrustedPeers are a list of peers which can issue tokens on this client's behalf using the dynamic \"oauth2:server:client_id:(client_id)\" scope.",
	"public":        "Public clients must use either use a redirectURL 127.0.0.1:X or \"urn:ietf:wg:oauth:2.0:oob\".",
}

func (ClientSpec) SwaggerDoc() map[string]string {
	return map_ClientSpec
}

var map_ConfigMap = map[string]string{
	"":           "ConfigMap holds configuration data for tke to consume.",
	"data":       "Data contains the configuration data. Each key must consist of alphanumeric characters, '-', '_' or '.'. Values with non-UTF-8 byte sequences must use the BinaryData field. The keys stored in Data must not overlap with the keys in the BinaryData field, this is enforced during validation process.",
	"binaryData": "BinaryData contains the binary data. Each key must consist of alphanumeric characters, '-', '_' or '.'. BinaryData can contain byte sequences that are not in the UTF-8 range. The keys stored in BinaryData must not overlap with the ones in the Data field, this is enforced during validation process.",
}

func (ConfigMap) SwaggerDoc() map[string]string {
	return map_ConfigMap
}

var map_ConfigMapList = map[string]string{
	"":      "ConfigMapList is a resource containing a list of ConfigMap objects.",
	"items": "Items is the list of ConfigMaps.",
}

func (ConfigMapList) SwaggerDoc() map[string]string {
	return map_ConfigMapList
}

var map_Group = map[string]string{
	"":     "Group is an object that contains the metadata about identify about tke local idp or third-party idp.",
	"spec": "Spec defines the desired identities of group in this set.",
}

func (Group) SwaggerDoc() map[string]string {
	return map_Group
}

var map_GroupList = map[string]string{
	"":      "GroupList is the whole list of all groups.",
	"items": "List of Group.",
}

func (GroupList) SwaggerDoc() map[string]string {
	return map_GroupList
}

var map_GroupSpec = map[string]string{
	"": "GroupSpec is a description of an Group.",
}

func (GroupSpec) SwaggerDoc() map[string]string {
	return map_GroupSpec
}

var map_IdentityProvider = map[string]string{
	"":     "IdentityProvider is an object that contains the metadata about identify provider used to login to TKE.",
	"spec": "Spec defines the desired identities of identity provider in this set.",
}

func (IdentityProvider) SwaggerDoc() map[string]string {
	return map_IdentityProvider
}

var map_IdentityProviderList = map[string]string{
	"":      "IdentityProviderList is the whole list of all identity providers.",
	"items": "List of identity providers.",
}

func (IdentityProviderList) SwaggerDoc() map[string]string {
	return map_IdentityProviderList
}

var map_IdentityProviderSpec = map[string]string{
	"":       "IdentityProviderSpec is a description of an identity provider.",
	"name":   "The Name of the connector that is used when displaying it to the end user.",
	"type":   "The type of the connector. E.g. 'oidc' or 'ldap'",
	"config": "Config holds all the configuration information specific to the connector type. Since there no generic struct we can use for this purpose, it is stored as a json string.",
}

func (IdentityProviderSpec) SwaggerDoc() map[string]string {
	return map_IdentityProviderSpec
}

var map_LocalGroup = map[string]string{
	"":     "LocalGroup represents a group of users.",
	"spec": "Spec defines the desired identities of group document in this set.",
}

func (LocalGroup) SwaggerDoc() map[string]string {
	return map_LocalGroup
}

var map_LocalGroupList = map[string]string{
	"":      "LocalGroupList is the whole list of all groups.",
	"items": "List of LocalGroup.",
}

func (LocalGroupList) SwaggerDoc() map[string]string {
	return map_LocalGroupList
}

var map_LocalGroupSpec = map[string]string{
	"": "LocalGroupSpec is a description of group.",
}

func (LocalGroupSpec) SwaggerDoc() map[string]string {
	return map_LocalGroupSpec
}

var map_LocalGroupStatus = map[string]string{
	"":      "LocalGroupStatus represents information about the status of a group.",
	"users": "Users represents the members of the group.",
}

func (LocalGroupStatus) SwaggerDoc() map[string]string {
	return map_LocalGroupStatus
}

var map_LocalIdentity = map[string]string{
	"":     "LocalIdentity is an object that contains the metadata about identify used to login to TKE.",
	"spec": "Spec defines the desired identities of identity in this set.",
}

func (LocalIdentity) SwaggerDoc() map[string]string {
	return map_LocalIdentity
}

var map_LocalIdentityList = map[string]string{
	"":      "LocalIdentityList is the whole list of all identities.",
	"items": "List of identities.",
}

func (LocalIdentityList) SwaggerDoc() map[string]string {
	return map_LocalIdentityList
}

var map_LocalIdentitySpec = map[string]string{
	"": "LocalIdentitySpec is a description of an identity.",
}

func (LocalIdentitySpec) SwaggerDoc() map[string]string {
	return map_LocalIdentitySpec
}

var map_LocalIdentityStatus = map[string]string{
	"":               "LocalIdentityStatus is a description of an identity status.",
	"LastUpdateTime": "The last time the local identity was updated.",
}

func (LocalIdentityStatus) SwaggerDoc() map[string]string {
	return map_LocalIdentityStatus
}

var map_NonResourceAttributes = map[string]string{
	"":     "NonResourceAttributes includes the authorization attributes available for non-resource requests to the Authorizer interface",
	"path": "Path is the URL path of the request",
	"verb": "Verb is the standard HTTP verb",
}

func (NonResourceAttributes) SwaggerDoc() map[string]string {
	return map_NonResourceAttributes
}

var map_PasswordReq = map[string]string{
	"": "PasswordReq contains info to update password for a localIdentity",
}

func (PasswordReq) SwaggerDoc() map[string]string {
	return map_PasswordReq
}

var map_Policy = map[string]string{
	"":     "Policy represents a policy document for access control.",
	"spec": "Spec defines the desired identities of policy document in this set.",
}

func (Policy) SwaggerDoc() map[string]string {
	return map_Policy
}

var map_PolicyBinding = map[string]string{
	"":         "PolicyBinding references the request to bind or unbind policies to the role.",
	"policies": "Policies holds the policies will bind or unbind to the role.",
}

func (PolicyBinding) SwaggerDoc() map[string]string {
	return map_PolicyBinding
}

var map_PolicyList = map[string]string{
	"":      "PolicyList is the whole list of all policies.",
	"items": "List of policies.",
}

func (PolicyList) SwaggerDoc() map[string]string {
	return map_PolicyList
}

var map_PolicySpec = map[string]string{
	"": "PolicySpec is a description of a policy.",
}

func (PolicySpec) SwaggerDoc() map[string]string {
	return map_PolicySpec
}

var map_PolicyStatus = map[string]string{
	"":       "PolicyStatus represents information about the status of a policy.",
	"users":  "Users represents the users the policy applies to.",
	"groups": "Groups represents the groups the policy applies to.",
}

func (PolicyStatus) SwaggerDoc() map[string]string {
	return map_PolicyStatus
}

var map_ResourceAttributes = map[string]string{
	"":            "ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface",
	"namespace":   "Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces \"\" (empty) is defaulted for LocalSubjectAccessReviews \"\" (empty) is empty for cluster-scoped resources \"\" (empty) means \"all\" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview",
	"verb":        "Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  \"*\" means all.",
	"group":       "Group is the API Group of the Resource.  \"*\" means all.",
	"version":     "Version is the API Version of the Resource.  \"*\" means all.",
	"resource":    "Resource is one of the existing resource types.  \"*\" means all.",
	"subresource": "Subresource is one of the existing resource types.  \"\" means none.",
	"name":        "Name is the name of the resource being requested for a \"get\" or deleted for a \"delete\". \"\" (empty) means all.",
}

func (ResourceAttributes) SwaggerDoc() map[string]string {
	return map_ResourceAttributes
}

var map_Role = map[string]string{
	"":     "Role is a collection with multiple policies.",
	"spec": "Spec defines the desired identities of role document in this set.",
}

func (Role) SwaggerDoc() map[string]string {
	return map_Role
}

var map_RoleList = map[string]string{
	"":      "RoleList is the whole list of policy.",
	"items": "List of rules.",
}

func (RoleList) SwaggerDoc() map[string]string {
	return map_RoleList
}

var map_RoleSpec = map[string]string{
	"":         "RoleSpec is a description of role.",
	"username": "Username is Creator",
}

func (RoleSpec) SwaggerDoc() map[string]string {
	return map_RoleSpec
}

var map_RoleStatus = map[string]string{
	"":       "RoleStatus represents information about the status of a role.",
	"users":  "Users represents the users the role applies to.",
	"groups": "Groups represents the groups the role applies to.",
}

func (RoleStatus) SwaggerDoc() map[string]string {
	return map_RoleStatus
}

var map_Rule = map[string]string{
	"":     "Rule represents a rule document for access control.",
	"spec": "Spec defines the desired identities of policy document in this set.",
}

func (Rule) SwaggerDoc() map[string]string {
	return map_Rule
}

var map_RuleList = map[string]string{
	"":      "RuleList is the whole list of all rules.",
	"items": "List of rules.",
}

func (RuleList) SwaggerDoc() map[string]string {
	return map_RuleList
}

var map_RuleSpec = map[string]string{
	"": "RuleSpec is a description of a rule.",
}

func (RuleSpec) SwaggerDoc() map[string]string {
	return map_RuleSpec
}

var map_Statement = map[string]string{
	"":       "Statement defines a series of action on resource can be done or not.",
	"effect": "Effect indicates action on the resource is allowed or not, can be \"allow\" or \"deny\"",
}

func (Statement) SwaggerDoc() map[string]string {
	return map_Statement
}

var map_Subject = map[string]string{
	"": "Subject references a user can specify by id or name.",
}

func (Subject) SwaggerDoc() map[string]string {
	return map_Subject
}

var map_SubjectAccessReview = map[string]string{
	"":       "SubjectAccessReview checks whether or not a user or group can perform an action.  Not filling in a spec.namespace means \"in all namespaces\".",
	"spec":   "Spec holds information about the request being evaluated",
	"status": "Status is filled in by the server and indicates whether the request is allowed or not",
}

func (SubjectAccessReview) SwaggerDoc() map[string]string {
	return map_SubjectAccessReview
}

var map_SubjectAccessReviewSpec = map[string]string{
	"":                       "SubjectAccessReviewSpec is a description of the access request.  Exactly one of ResourceAttributes and NonResourceAttributes must be set",
	"resourceAttributes":     "ResourceAttributes describes information for a resource access request",
	"resourceAttributesList": "ResourceAttributesList describes information for multi resource access request.",
	"nonResourceAttributes":  "NonResourceAttributes describes information for a non-resource access request",
	"user":                   "User is the user you're testing for. If you specify \"User\" but not \"Groups\", then is it interpreted as \"What if User were not a member of any groups",
	"groups":                 "Groups is the groups you're testing for.",
	"extra":                  "Extra corresponds to the user.Info.GetExtra() method from the authenticator.  Since that is input to the authorizer it needs a reflection here.",
	"uid":                    "UID information about the requesting user.",
}

func (SubjectAccessReviewSpec) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewSpec
}

var map_SubjectAccessReviewStatus = map[string]string{
	"":                "SubjectAccessReviewStatus represents the current state of a SubjectAccessReview.",
	"allowed":         "Allowed is required. True if the action would be allowed, false otherwise.",
	"denied":          "Denied is optional. True if the action would be denied, otherwise false. If both allowed is false and denied is false, then the authorizer has no opinion on whether to authorize the action. Denied may not be true if Allowed is true.",
	"reason":          "Reason is optional.  It indicates why a request was allowed or denied.",
	"evaluationError": "EvaluationError is an indication that some error occurred during the authorization check. It is entirely possible to get an error and be able to continue determine authorization status in spite of it. For instance, RBAC can be missing a role, but enough roles are still present and bound to reason about the request.",
	"allowedList":     "AllowedList is the allowed response for batch authorization request.",
}

func (SubjectAccessReviewStatus) SwaggerDoc() map[string]string {
	return map_SubjectAccessReviewStatus
}

var map_User = map[string]string{
	"":     "User is an object that contains the metadata about identify about tke local idp or third-party idp.",
	"spec": "Spec defines the desired identities of identity in this set.",
}

func (User) SwaggerDoc() map[string]string {
	return map_User
}

var map_UserList = map[string]string{
	"":      "UserList is the whole list of all users.",
	"items": "List of User.",
}

func (UserList) SwaggerDoc() map[string]string {
	return map_UserList
}

var map_UserSpec = map[string]string{
	"":     "UserSpec is a description of an user.",
	"name": "Name must be unique in the same tenant.",
}

func (UserSpec) SwaggerDoc() map[string]string {
	return map_UserSpec
}

// AUTO-GENERATED FUNCTIONS END HERE
