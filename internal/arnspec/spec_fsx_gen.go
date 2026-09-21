// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: fsx
// Source: https://servicereference.us-east-1.amazonaws.com/v1/fsx/fsx.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "fsx_association", Service: "fsx", Resource: "association", Template: "arn:${Partition}:fsx:${Region}:${Account}:association/${FileSystemIdOrFileCacheId}/${DataRepositoryAssociationId}"},
		{Name: "fsx_backup", Service: "fsx", Resource: "backup", Template: "arn:${Partition}:fsx:${Region}:${Account}:backup/${BackupId}"},
		{Name: "fsx_file_cache", Service: "fsx", Resource: "file-cache", Template: "arn:${Partition}:fsx:${Region}:${Account}:file-cache/${FileCacheId}"},
		{Name: "fsx_file_system", Service: "fsx", Resource: "file-system", Template: "arn:${Partition}:fsx:${Region}:${Account}:file-system/${FileSystemId}"},
		{Name: "fsx_snapshot", Service: "fsx", Resource: "snapshot", Template: "arn:${Partition}:fsx:${Region}:${Account}:snapshot/${VolumeId}/${SnapshotId}"},
		{Name: "fsx_storage_virtual_machine", Service: "fsx", Resource: "storage-virtual-machine", Template: "arn:${Partition}:fsx:${Region}:${Account}:storage-virtual-machine/${FileSystemId}/${StorageVirtualMachineId}"},
		{Name: "fsx_task", Service: "fsx", Resource: "task", Template: "arn:${Partition}:fsx:${Region}:${Account}:task/${TaskId}"},
		{Name: "fsx_volume", Service: "fsx", Resource: "volume", Template: "arn:${Partition}:fsx:${Region}:${Account}:volume/${FileSystemId}/${VolumeId}"},
	})
}
