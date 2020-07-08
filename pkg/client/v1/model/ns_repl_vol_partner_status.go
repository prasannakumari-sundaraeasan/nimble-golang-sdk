// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsReplVolPartnerStatus Enum.

type NsReplVolPartnerStatus string

const (
	cNsReplVolPartnerStatusDataTransferTerminated             NsReplVolPartnerStatus = "Data transfer terminated"
	cNsReplVolPartnerStatusInitiatingContentBasedReplication  NsReplVolPartnerStatus = "Initiating content based replication"
	cNsReplVolPartnerStatusContentBasedReplicationComplete    NsReplVolPartnerStatus = "Content based replication complete"
	cNsReplVolPartnerStatusDataTransferComplete               NsReplVolPartnerStatus = "Data transfer complete"
	cNsReplVolPartnerStatusCreatingSnapshotOnPartner          NsReplVolPartnerStatus = "Creating snapshot on partner"
	cNsReplVolPartnerStatusSynchronizingPartnerConfiguration  NsReplVolPartnerStatus = "Synchronizing partner configuration"
	cNsReplVolPartnerStatusDataTransferInProgress             NsReplVolPartnerStatus = "Data transfer in progress"
	cNsReplVolPartnerStatusIdle                               NsReplVolPartnerStatus = "Idle"
	cNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner NsReplVolPartnerStatus = "Identifying common ancestry on partner"
	cNsReplVolPartnerStatusContentBasedReplicationInProgress  NsReplVolPartnerStatus = "Content based replication in progress"
	cNsReplVolPartnerStatusTerminatingDataTransfer            NsReplVolPartnerStatus = "Terminating data transfer"
	cNsReplVolPartnerStatusPaused                             NsReplVolPartnerStatus = "Paused"
	cNsReplVolPartnerStatusInitiatingDataTransfer             NsReplVolPartnerStatus = "Initiating data transfer"
)

var pNsReplVolPartnerStatusDataTransferTerminated NsReplVolPartnerStatus
var pNsReplVolPartnerStatusInitiatingContentBasedReplication NsReplVolPartnerStatus
var pNsReplVolPartnerStatusContentBasedReplicationComplete NsReplVolPartnerStatus
var pNsReplVolPartnerStatusDataTransferComplete NsReplVolPartnerStatus
var pNsReplVolPartnerStatusCreatingSnapshotOnPartner NsReplVolPartnerStatus
var pNsReplVolPartnerStatusSynchronizingPartnerConfiguration NsReplVolPartnerStatus
var pNsReplVolPartnerStatusDataTransferInProgress NsReplVolPartnerStatus
var pNsReplVolPartnerStatusIdle NsReplVolPartnerStatus
var pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner NsReplVolPartnerStatus
var pNsReplVolPartnerStatusContentBasedReplicationInProgress NsReplVolPartnerStatus
var pNsReplVolPartnerStatusTerminatingDataTransfer NsReplVolPartnerStatus
var pNsReplVolPartnerStatusPaused NsReplVolPartnerStatus
var pNsReplVolPartnerStatusInitiatingDataTransfer NsReplVolPartnerStatus

// Export
var NsReplVolPartnerStatusDataTransferTerminated *NsReplVolPartnerStatus
var NsReplVolPartnerStatusInitiatingContentBasedReplication *NsReplVolPartnerStatus
var NsReplVolPartnerStatusContentBasedReplicationComplete *NsReplVolPartnerStatus
var NsReplVolPartnerStatusDataTransferComplete *NsReplVolPartnerStatus
var NsReplVolPartnerStatusCreatingSnapshotOnPartner *NsReplVolPartnerStatus
var NsReplVolPartnerStatusSynchronizingPartnerConfiguration *NsReplVolPartnerStatus
var NsReplVolPartnerStatusDataTransferInProgress *NsReplVolPartnerStatus
var NsReplVolPartnerStatusIdle *NsReplVolPartnerStatus
var NsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner *NsReplVolPartnerStatus
var NsReplVolPartnerStatusContentBasedReplicationInProgress *NsReplVolPartnerStatus
var NsReplVolPartnerStatusTerminatingDataTransfer *NsReplVolPartnerStatus
var NsReplVolPartnerStatusPaused *NsReplVolPartnerStatus
var NsReplVolPartnerStatusInitiatingDataTransfer *NsReplVolPartnerStatus

func init() {
	pNsReplVolPartnerStatusDataTransferTerminated = cNsReplVolPartnerStatusDataTransferTerminated
	NsReplVolPartnerStatusDataTransferTerminated = &pNsReplVolPartnerStatusDataTransferTerminated

	pNsReplVolPartnerStatusInitiatingContentBasedReplication = cNsReplVolPartnerStatusInitiatingContentBasedReplication
	NsReplVolPartnerStatusInitiatingContentBasedReplication = &pNsReplVolPartnerStatusInitiatingContentBasedReplication

	pNsReplVolPartnerStatusContentBasedReplicationComplete = cNsReplVolPartnerStatusContentBasedReplicationComplete
	NsReplVolPartnerStatusContentBasedReplicationComplete = &pNsReplVolPartnerStatusContentBasedReplicationComplete

	pNsReplVolPartnerStatusDataTransferComplete = cNsReplVolPartnerStatusDataTransferComplete
	NsReplVolPartnerStatusDataTransferComplete = &pNsReplVolPartnerStatusDataTransferComplete

	pNsReplVolPartnerStatusCreatingSnapshotOnPartner = cNsReplVolPartnerStatusCreatingSnapshotOnPartner
	NsReplVolPartnerStatusCreatingSnapshotOnPartner = &pNsReplVolPartnerStatusCreatingSnapshotOnPartner

	pNsReplVolPartnerStatusSynchronizingPartnerConfiguration = cNsReplVolPartnerStatusSynchronizingPartnerConfiguration
	NsReplVolPartnerStatusSynchronizingPartnerConfiguration = &pNsReplVolPartnerStatusSynchronizingPartnerConfiguration

	pNsReplVolPartnerStatusDataTransferInProgress = cNsReplVolPartnerStatusDataTransferInProgress
	NsReplVolPartnerStatusDataTransferInProgress = &pNsReplVolPartnerStatusDataTransferInProgress

	pNsReplVolPartnerStatusIdle = cNsReplVolPartnerStatusIdle
	NsReplVolPartnerStatusIdle = &pNsReplVolPartnerStatusIdle

	pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner = cNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner
	NsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner = &pNsReplVolPartnerStatusIdentifyingCommonAncestryOnPartner

	pNsReplVolPartnerStatusContentBasedReplicationInProgress = cNsReplVolPartnerStatusContentBasedReplicationInProgress
	NsReplVolPartnerStatusContentBasedReplicationInProgress = &pNsReplVolPartnerStatusContentBasedReplicationInProgress

	pNsReplVolPartnerStatusTerminatingDataTransfer = cNsReplVolPartnerStatusTerminatingDataTransfer
	NsReplVolPartnerStatusTerminatingDataTransfer = &pNsReplVolPartnerStatusTerminatingDataTransfer

	pNsReplVolPartnerStatusPaused = cNsReplVolPartnerStatusPaused
	NsReplVolPartnerStatusPaused = &pNsReplVolPartnerStatusPaused

	pNsReplVolPartnerStatusInitiatingDataTransfer = cNsReplVolPartnerStatusInitiatingDataTransfer
	NsReplVolPartnerStatusInitiatingDataTransfer = &pNsReplVolPartnerStatusInitiatingDataTransfer

}
