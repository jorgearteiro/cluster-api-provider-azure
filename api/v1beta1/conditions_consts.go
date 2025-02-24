/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"

// AzureCluster Conditions and Reasons.
const (
	// NetworkInfrastructureReadyCondition reports of current status of cluster infrastructure.
	NetworkInfrastructureReadyCondition clusterv1.ConditionType = "NetworkInfrastructureReady"
	// NamespaceNotAllowedByIdentity used to indicate cluster in a namespace not allowed by identity.
	NamespaceNotAllowedByIdentity = "NamespaceNotAllowedByIdentity"
)

// AzureMachine Conditions and Reasons.
const (
	// VMRunningCondition reports on current status of the Azure VM.
	VMRunningCondition clusterv1.ConditionType = "VMRunning"
	// VMCreatingReason used when the vm creation is in progress.
	VMCreatingReason = "VMCreating"
	// VMUpdatingReason used when the vm updating is in progress.
	VMUpdatingReason = "VMUpdating"
	// VMDeletingReason used when the vm is in a deleting state.
	VMDeletingReason = "VMDeleting"
	// VMProvisionFailedReason used for failures during vm provisioning.
	VMProvisionFailedReason = "VMProvisionFailed"
	// WaitingForClusterInfrastructureReason used when machine is waiting for cluster infrastructure to be ready before proceeding.
	WaitingForClusterInfrastructureReason = "WaitingForClusterInfrastructure"
	// WaitingForBootstrapDataReason used when machine is waiting for bootstrap data to be ready before proceeding.
	WaitingForBootstrapDataReason = "WaitingForBootstrapData"
	// BootstrapSucceededCondition reports the result of the execution of the boostrap data on the machine.
	BootstrapSucceededCondition = "BoostrapSucceeded"
	// BootstrapInProgressReason is used to indicate the bootstrap data has not finished executing.
	BootstrapInProgressReason = "BootstrapInProgress"
	// BootstrapFailedReason is used to indicate the bootstrap process ran into an error.
	BootstrapFailedReason = "BootstrapFailed"
)

// AzureMachinePool Conditions and Reasons.
const (
	// ScaleSetRunningCondition reports on current status of the Azure Scale Set.
	ScaleSetRunningCondition clusterv1.ConditionType = "ScaleSetRunning"
	// ScaleSetCreatingReason used when the scale set creation is in progress.
	ScaleSetCreatingReason = "ScaleSetCreating"
	// ScaleSetUpdatingReason used when the scale set updating is in progress.
	ScaleSetUpdatingReason = "ScaleSetUpdating"
	// ScaleSetDeletingReason used when the scale set is in a deleting state.
	ScaleSetDeletingReason = "ScaleSetDeleting"
	// ScaleSetProvisionFailedReason used for failures during scale set provisioning.
	ScaleSetProvisionFailedReason = "ScaleSetProvisionFailed"

	// ScaleSetDesiredReplicasCondition reports on the scaling state of the machine pool.
	ScaleSetDesiredReplicasCondition clusterv1.ConditionType = "ScaleSetDesiredReplicas"
	// ScaleSetScaleUpReason describes the machine pool scaling up.
	ScaleSetScaleUpReason = "ScaleSetScalingUp"
	// ScaleSetScaleDownReason describes the machine pool scaling down.
	ScaleSetScaleDownReason = "ScaleSetScalingDown"

	// ScaleSetModelUpdatedCondition reports on the model state of the pool.
	ScaleSetModelUpdatedCondition clusterv1.ConditionType = "ScaleSetModelUpdated"
	// ScaleSetModelOutOfDateReason describes the machine pool model being out of date.
	ScaleSetModelOutOfDateReason = "ScaleSetModelOutOfDate"
)

// Azure Services Conditions and Reasons.
const (
	// ResourceGroupReadyCondition means the resource group exists and is ready to be used.
	ResourceGroupReadyCondition clusterv1.ConditionType = "ResourceGroupReady"
	// VNetReadyCondition means the virtual network exists and is ready to be used.
	VNetReadyCondition clusterv1.ConditionType = "VNetReady"
	// SecurityGroupsReadyCondition means the security groups exist and are ready to be used.
	SecurityGroupsReadyCondition clusterv1.ConditionType = "SecurityGroupsReady"
	// RouteTablesReadyCondition means the route tables exist and are ready to be used.
	RouteTablesReadyCondition clusterv1.ConditionType = "RouteTablesReady"
	// PublicIPsReadyCondition means the public IPs exist and are ready to be used.
	PublicIPsReadyCondition clusterv1.ConditionType = "PublicIPsReady"
	// NATGatewaysReadyCondition means the NAT gateways exist and are ready to be used.
	NATGatewaysReadyCondition clusterv1.ConditionType = "NATGatewaysReady"
	// SubnetsReadyCondition means the subnets exist and are ready to be used.
	SubnetsReadyCondition clusterv1.ConditionType = "SubnetsReady"
	// LoadBalancersReadyCondition means the load balancers exist and are ready to be used.
	LoadBalancersReadyCondition clusterv1.ConditionType = "LoadBalancersReady"
	// PrivateDNSReadyCondition means the private DNS exists and is ready to be used.
	PrivateDNSReadyCondition clusterv1.ConditionType = "PrivateDNSReady"
	// BastionHostReadyCondition means the bastion host exists and is ready to be used.
	BastionHostReadyCondition clusterv1.ConditionType = "BastionHostReady"
	// InboundNATRulesReadyCondition means the inbound NAT rules exist and are ready to be used.
	InboundNATRulesReadyCondition clusterv1.ConditionType = "InboundNATRulesReady"
	// AvailabilitySetReadyCondition means the availability set exists and is ready to be used.
	AvailabilitySetReadyCondition clusterv1.ConditionType = "AvailabilitySetReady"
	// RoleAssignmentReadyCondition means the role assignment exists and is ready to be used.
	RoleAssignmentReadyCondition clusterv1.ConditionType = "RoleAssignmentReady"

	// CreatingReason means the resource is being created.
	CreatingReason = "Creating"
	// FailedReason means the resource failed to be created.
	FailedReason = "Failed"
	// DeletingReason means the resource is being deleted.
	DeletingReason = "Deleting"
	// DeletedReason means the resource was deleted.
	DeletedReason = "Deleted"
	// DeletionFailedReason means the resource failed to be deleted.
	DeletionFailedReason = "DeletionFailed"
	// UpdatingReason means the resource is being updated.
	UpdatingReason = "Updating"
)
