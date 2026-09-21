// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: nimble
// Source: https://servicereference.us-east-1.amazonaws.com/v1/nimble/nimble.json
// Functions: 8
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "nimble_eula", Service: "nimble", Resource: "eula", Template: "arn:${Partition}:nimble:${Region}:${Account}:eula/${EulaId}"},
		{Name: "nimble_eula_acceptance", Service: "nimble", Resource: "eula-acceptance", Template: "arn:${Partition}:nimble:${Region}:${Account}:eula-acceptance/${EulaAcceptanceId}"},
		{Name: "nimble_launch_profile", Service: "nimble", Resource: "launch-profile", Template: "arn:${Partition}:nimble:${Region}:${Account}:launch-profile/${LaunchProfileId}"},
		{Name: "nimble_streaming_image", Service: "nimble", Resource: "streaming-image", Template: "arn:${Partition}:nimble:${Region}:${Account}:streaming-image/${StreamingImageId}"},
		{Name: "nimble_streaming_session", Service: "nimble", Resource: "streaming-session", Template: "arn:${Partition}:nimble:${Region}:${Account}:streaming-session/${StreamingSessionId}"},
		{Name: "nimble_streaming_session_backup", Service: "nimble", Resource: "streaming-session-backup", Template: "arn:${Partition}:nimble:${Region}:${Account}:streaming-session-backup/${StreamingSessionBackupId}"},
		{Name: "nimble_studio", Service: "nimble", Resource: "studio", Template: "arn:${Partition}:nimble:${Region}:${Account}:studio/${StudioId}"},
		{Name: "nimble_studio_component", Service: "nimble", Resource: "studio-component", Template: "arn:${Partition}:nimble:${Region}:${Account}:studio-component/${StudioComponentId}"},
	})
}
