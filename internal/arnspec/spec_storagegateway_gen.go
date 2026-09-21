// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: storagegateway
// Source: https://servicereference.us-east-1.amazonaws.com/v1/storagegateway/storagegateway.json
// Functions: 9
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "storagegateway_cache_report", Service: "storagegateway", Resource: "cache-report", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:share/${ShareId}/cache-report/${CacheReportId}"},
		{Name: "storagegateway_device", Service: "storagegateway", Resource: "device", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:gateway/${GatewayId}/device/${Vtldevice}"},
		{Name: "storagegateway_fs_association", Service: "storagegateway", Resource: "fs-association", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:fs-association/${FsaId}"},
		{Name: "storagegateway_gateway", Service: "storagegateway", Resource: "gateway", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:gateway/${GatewayId}"},
		{Name: "storagegateway_share", Service: "storagegateway", Resource: "share", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:share/${ShareId}"},
		{Name: "storagegateway_tape", Service: "storagegateway", Resource: "tape", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:tape/${TapeBarcode}"},
		{Name: "storagegateway_tapepool", Service: "storagegateway", Resource: "tapepool", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:tapepool/${PoolId}"},
		{Name: "storagegateway_target", Service: "storagegateway", Resource: "target", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:gateway/${GatewayId}/target/${IscsiTarget}"},
		{Name: "storagegateway_volume", Service: "storagegateway", Resource: "volume", Template: "arn:${Partition}:storagegateway:${Region}:${Account}:gateway/${GatewayId}/volume/${VolumeId}"},
	})
}
