package digitaltwins

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// EndpointProvisioningState enumerates the values for endpoint provisioning state.
type EndpointProvisioningState string

const (
	// Canceled ...
	Canceled EndpointProvisioningState = "Canceled"
	// Deleted ...
	Deleted EndpointProvisioningState = "Deleted"
	// Deleting ...
	Deleting EndpointProvisioningState = "Deleting"
	// Disabled ...
	Disabled EndpointProvisioningState = "Disabled"
	// Failed ...
	Failed EndpointProvisioningState = "Failed"
	// Moving ...
	Moving EndpointProvisioningState = "Moving"
	// Provisioning ...
	Provisioning EndpointProvisioningState = "Provisioning"
	// Restoring ...
	Restoring EndpointProvisioningState = "Restoring"
	// Succeeded ...
	Succeeded EndpointProvisioningState = "Succeeded"
	// Suspending ...
	Suspending EndpointProvisioningState = "Suspending"
	// Warning ...
	Warning EndpointProvisioningState = "Warning"
)

// PossibleEndpointProvisioningStateValues returns an array of possible values for the EndpointProvisioningState const type.
func PossibleEndpointProvisioningStateValues() []EndpointProvisioningState {
	return []EndpointProvisioningState{Canceled, Deleted, Deleting, Disabled, Failed, Moving, Provisioning, Restoring, Succeeded, Suspending, Warning}
}

// EndpointType enumerates the values for endpoint type.
type EndpointType string

const (
	// EndpointTypeDigitalTwinsEndpointResourceProperties ...
	EndpointTypeDigitalTwinsEndpointResourceProperties EndpointType = "DigitalTwinsEndpointResourceProperties"
	// EndpointTypeEventGrid ...
	EndpointTypeEventGrid EndpointType = "EventGrid"
	// EndpointTypeEventHub ...
	EndpointTypeEventHub EndpointType = "EventHub"
	// EndpointTypeServiceBus ...
	EndpointTypeServiceBus EndpointType = "ServiceBus"
)

// PossibleEndpointTypeValues returns an array of possible values for the EndpointType const type.
func PossibleEndpointTypeValues() []EndpointType {
	return []EndpointType{EndpointTypeDigitalTwinsEndpointResourceProperties, EndpointTypeEventGrid, EndpointTypeEventHub, EndpointTypeServiceBus}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCanceled ...
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleted ...
	ProvisioningStateDeleted ProvisioningState = "Deleted"
	// ProvisioningStateDeleting ...
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed ...
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateMoving ...
	ProvisioningStateMoving ProvisioningState = "Moving"
	// ProvisioningStateProvisioning ...
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateRestoring ...
	ProvisioningStateRestoring ProvisioningState = "Restoring"
	// ProvisioningStateSucceeded ...
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateSuspending ...
	ProvisioningStateSuspending ProvisioningState = "Suspending"
	// ProvisioningStateWarning ...
	ProvisioningStateWarning ProvisioningState = "Warning"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{ProvisioningStateCanceled, ProvisioningStateDeleted, ProvisioningStateDeleting, ProvisioningStateFailed, ProvisioningStateMoving, ProvisioningStateProvisioning, ProvisioningStateRestoring, ProvisioningStateSucceeded, ProvisioningStateSuspending, ProvisioningStateWarning}
}

// Reason enumerates the values for reason.
type Reason string

const (
	// AlreadyExists ...
	AlreadyExists Reason = "AlreadyExists"
	// Invalid ...
	Invalid Reason = "Invalid"
)

// PossibleReasonValues returns an array of possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{AlreadyExists, Invalid}
}