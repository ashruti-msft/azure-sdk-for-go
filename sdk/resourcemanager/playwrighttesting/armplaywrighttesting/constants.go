//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armplaywrighttesting

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/playwrighttesting/armplaywrighttesting"
	moduleVersion = "v1.0.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// EnablementStatus - This property sets the connection region for Playwright client workers to cloud-hosted browsers. If
// enabled, workers connect to browsers in the closest Azure region, ensuring lower latency. If
// disabled, workers connect to browsers in the Azure region in which the workspace was initially created.
type EnablementStatus string

const (
	// EnablementStatusDisabled - The feature is Disabled.
	EnablementStatusDisabled EnablementStatus = "Disabled"
	// EnablementStatusEnabled - The feature is Enabled.
	EnablementStatusEnabled EnablementStatus = "Enabled"
)

// PossibleEnablementStatusValues returns the possible values for the EnablementStatus const type.
func PossibleEnablementStatusValues() []EnablementStatus {
	return []EnablementStatus{
		EnablementStatusDisabled,
		EnablementStatusEnabled,
	}
}

// FreeTrialState - The free-trial state.
type FreeTrialState string

const (
	// FreeTrialStateActive - The free-trial is Active.
	FreeTrialStateActive FreeTrialState = "Active"
	// FreeTrialStateExpired - The free-trial is Expired.
	FreeTrialStateExpired FreeTrialState = "Expired"
	// FreeTrialStateNotEligible - The free-trial is Not Eligible.
	FreeTrialStateNotEligible FreeTrialState = "NotEligible"
	// FreeTrialStateNotRegistered - The free-trial is Not Registered.
	FreeTrialStateNotRegistered FreeTrialState = "NotRegistered"
)

// PossibleFreeTrialStateValues returns the possible values for the FreeTrialState const type.
func PossibleFreeTrialStateValues() []FreeTrialState {
	return []FreeTrialState{
		FreeTrialStateActive,
		FreeTrialStateExpired,
		FreeTrialStateNotEligible,
		FreeTrialStateNotRegistered,
	}
}

// OfferingType - Offering type state.
type OfferingType string

const (
	// OfferingTypeGeneralAvailability - The offeringType is GeneralAvailability.
	OfferingTypeGeneralAvailability OfferingType = "GeneralAvailability"
	// OfferingTypeNotApplicable - The offeringType is NotApplicable.
	OfferingTypeNotApplicable OfferingType = "NotApplicable"
	// OfferingTypePrivatePreview - The offeringType is PrivatePreview.
	OfferingTypePrivatePreview OfferingType = "PrivatePreview"
	// OfferingTypePublicPreview - The offeringType is PublicPreview.
	OfferingTypePublicPreview OfferingType = "PublicPreview"
)

// PossibleOfferingTypeValues returns the possible values for the OfferingType const type.
func PossibleOfferingTypeValues() []OfferingType {
	return []OfferingType{
		OfferingTypeGeneralAvailability,
		OfferingTypeNotApplicable,
		OfferingTypePrivatePreview,
		OfferingTypePublicPreview,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Change accepted for processing..
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateCreating - Creation in progress..
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - Deletion in progress..
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

type QuotaNames string

const (
	// QuotaNamesReporting - The quota details for reporting feature. When enabled, Playwright client will be able to upload and
	// display test results, including artifacts like traces and screenshots, in the Playwright portal. This enables faster and
	// more efficient troubleshooting.
	QuotaNamesReporting QuotaNames = "Reporting"
	// QuotaNamesScalableExecution - The quota details for scalable execution feature. When enabled, Playwright client workers
	// can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing
	// test completion durations.
	QuotaNamesScalableExecution QuotaNames = "ScalableExecution"
)

// PossibleQuotaNamesValues returns the possible values for the QuotaNames const type.
func PossibleQuotaNamesValues() []QuotaNames {
	return []QuotaNames{
		QuotaNamesReporting,
		QuotaNamesScalableExecution,
	}
}
