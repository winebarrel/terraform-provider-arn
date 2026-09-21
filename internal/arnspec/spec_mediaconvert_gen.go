// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: mediaconvert
// Source: https://servicereference.us-east-1.amazonaws.com/v1/mediaconvert/mediaconvert.json
// Functions: 5
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "mediaconvert_certificate_association", Service: "mediaconvert", Resource: "CertificateAssociation", Template: "arn:${Partition}:mediaconvert:${Region}:${Account}:certificates/${CertificateArn}"},
		{Name: "mediaconvert_job", Service: "mediaconvert", Resource: "Job", Template: "arn:${Partition}:mediaconvert:${Region}:${Account}:jobs/${JobId}"},
		{Name: "mediaconvert_job_template", Service: "mediaconvert", Resource: "JobTemplate", Template: "arn:${Partition}:mediaconvert:${Region}:${Account}:jobTemplates/${JobTemplateName}"},
		{Name: "mediaconvert_preset", Service: "mediaconvert", Resource: "Preset", Template: "arn:${Partition}:mediaconvert:${Region}:${Account}:presets/${PresetName}"},
		{Name: "mediaconvert_queue", Service: "mediaconvert", Resource: "Queue", Template: "arn:${Partition}:mediaconvert:${Region}:${Account}:queues/${QueueName}"},
	})
}
