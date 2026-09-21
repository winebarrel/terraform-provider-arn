// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: application-autoscaling
// Source: https://servicereference.us-east-1.amazonaws.com/v1/application-autoscaling/application-autoscaling.json
// Functions: 1
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "application_autoscaling_scalable_target", Service: "application-autoscaling", Resource: "ScalableTarget", Template: "arn:${Partition}:application-autoscaling:${Region}:${Account}:scalable-target/${ResourceId}"},
	})
}
