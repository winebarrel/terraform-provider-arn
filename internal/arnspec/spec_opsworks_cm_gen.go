// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: opsworks-cm
// Source: https://servicereference.us-east-1.amazonaws.com/v1/opsworks-cm/opsworks-cm.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "opsworks_cm_backup", Service: "opsworks-cm", Resource: "backup", Template: "arn:${Partition}:opsworks-cm::${Account}:backup/${ServerName}-{Date-and-Time-Stamp-of-Backup}"},
		{Name: "opsworks_cm_server", Service: "opsworks-cm", Resource: "server", Template: "arn:${Partition}:opsworks-cm::${Account}:server/${ServerName}/${UniqueId}"},
	})
}
