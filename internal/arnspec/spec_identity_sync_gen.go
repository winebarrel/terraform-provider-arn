// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: identity-sync
// Source: https://servicereference.us-east-1.amazonaws.com/v1/identity-sync/identity-sync.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "identity_sync_sync_profile_resource", Service: "identity-sync", Resource: "SyncProfileResource", Template: "arn:${Partition}:identity-sync:${Region}:${Account}:profile/${SyncProfileName}"},
		{Name: "identity_sync_sync_target_resource", Service: "identity-sync", Resource: "SyncTargetResource", Template: "arn:${Partition}:identity-sync:${Region}:${Account}:target/${SyncProfileName}/${SyncTargetName}"},
	})
}
