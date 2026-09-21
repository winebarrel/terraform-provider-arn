// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: route53profiles
// Source: https://servicereference.us-east-1.amazonaws.com/v1/route53profiles/route53profiles.json
// Functions: 2
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "route53profiles_profile", Service: "route53profiles", Resource: "profile", Template: "arn:${Partition}:route53profiles:${Region}:${Account}:profile/${ResourceId}"},
		{Name: "route53profiles_profile_association", Service: "route53profiles", Resource: "profile-association", Template: "arn:${Partition}:route53profiles:${Region}:${Account}:profile-association/${ResourceId}"},
	})
}
