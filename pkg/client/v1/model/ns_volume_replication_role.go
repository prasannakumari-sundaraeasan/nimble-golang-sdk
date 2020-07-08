// Copyright 2020 Hewlett Packard Enterprise Development LP
package model

// Golang package for NsVolumeReplicationRole Enum.

type NsVolumeReplicationRole string

const (
	cNsVolumeReplicationRolePeriodicSnapshotDownstream NsVolumeReplicationRole = "periodic_snapshot_downstream"
	cNsVolumeReplicationRoleSynchronousUpstream        NsVolumeReplicationRole = "synchronous_upstream"
	cNsVolumeReplicationRoleSynchronousDownstream      NsVolumeReplicationRole = "synchronous_downstream"
	cNsVolumeReplicationRoleNoReplication              NsVolumeReplicationRole = "no_replication"
	cNsVolumeReplicationRolePeriodicSnapshotUpstream   NsVolumeReplicationRole = "periodic_snapshot_upstream"
)

var pNsVolumeReplicationRolePeriodicSnapshotDownstream NsVolumeReplicationRole
var pNsVolumeReplicationRoleSynchronousUpstream NsVolumeReplicationRole
var pNsVolumeReplicationRoleSynchronousDownstream NsVolumeReplicationRole
var pNsVolumeReplicationRoleNoReplication NsVolumeReplicationRole
var pNsVolumeReplicationRolePeriodicSnapshotUpstream NsVolumeReplicationRole

// Export
var NsVolumeReplicationRolePeriodicSnapshotDownstream *NsVolumeReplicationRole
var NsVolumeReplicationRoleSynchronousUpstream *NsVolumeReplicationRole
var NsVolumeReplicationRoleSynchronousDownstream *NsVolumeReplicationRole
var NsVolumeReplicationRoleNoReplication *NsVolumeReplicationRole
var NsVolumeReplicationRolePeriodicSnapshotUpstream *NsVolumeReplicationRole

func init() {
	pNsVolumeReplicationRolePeriodicSnapshotDownstream = cNsVolumeReplicationRolePeriodicSnapshotDownstream
	NsVolumeReplicationRolePeriodicSnapshotDownstream = &pNsVolumeReplicationRolePeriodicSnapshotDownstream

	pNsVolumeReplicationRoleSynchronousUpstream = cNsVolumeReplicationRoleSynchronousUpstream
	NsVolumeReplicationRoleSynchronousUpstream = &pNsVolumeReplicationRoleSynchronousUpstream

	pNsVolumeReplicationRoleSynchronousDownstream = cNsVolumeReplicationRoleSynchronousDownstream
	NsVolumeReplicationRoleSynchronousDownstream = &pNsVolumeReplicationRoleSynchronousDownstream

	pNsVolumeReplicationRoleNoReplication = cNsVolumeReplicationRoleNoReplication
	NsVolumeReplicationRoleNoReplication = &pNsVolumeReplicationRoleNoReplication

	pNsVolumeReplicationRolePeriodicSnapshotUpstream = cNsVolumeReplicationRolePeriodicSnapshotUpstream
	NsVolumeReplicationRolePeriodicSnapshotUpstream = &pNsVolumeReplicationRolePeriodicSnapshotUpstream

}
