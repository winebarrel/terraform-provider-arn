// Package arnvalue checks the values that go into an ARN's structural fields.
//
// The checks are on shape, not existence. Whether an account exists, or
// whether AWS has built a region yet, is not something a string can be asked
// about. What these catch is the value that could never be right, which would
// otherwise be interpolated into a syntactically valid ARN and fail much
// later, at apply time, somewhere that does not mention the configuration.
//
// The same field can be filled from the configuration file, from a call-site
// option, or from a positional argument where AWS spells it as a template
// placeholder, so the rules live here rather than in any one of those.
package arnvalue

import (
	"fmt"
	"regexp"
)

var (
	// An AWS account id is exactly twelve digits, and has been for the life
	// of the service. Two other values belong in the field:
	//
	//   - "aws", which is what AWS-managed policies carry:
	//     arn:aws:iam::aws:policy/AdministratorAccess.
	//   - "*", since an ARN written for an IAM policy may wildcard the field.
	reAccountID = regexp.MustCompile(`^([0-9]{12}|aws|\*)$`)

	// Region names run "xx-word-N", with extra words for the isolated
	// partitions: us-east-1, ap-northeast-1, us-gov-west-1, us-iso-east-1.
	// Global resources carry an empty region field rather than a name like
	// "aws-global", so there is no such case to allow through here.
	reRegion = regexp.MustCompile(`^([a-z]{2}(-[a-z]+)+-[0-9]+|\*)$`)

	// Every partition to date is "aws" or "aws-" plus one or more words:
	// aws-cn, aws-us-gov, aws-iso-b. Matching the shape rather than a fixed
	// list means a new partition works without a release here. IAM Access
	// Analyzer reports the supported values as "*, aws, aws-cn, aws-us-gov",
	// so the wildcard belongs here too.
	rePartition = regexp.MustCompile(`^(aws(-[a-z0-9]+)*|\*)$`)
)

// AccountID checks the ARN's account field.
func AccountID(s string) error {
	if !reAccountID.MatchString(s) {
		return fmt.Errorf("invalid account id %q: expected twelve digits, \"aws\" for an AWS-managed resource, or \"*\"", s)
	}
	return nil
}

// Region checks the ARN's region field.
func Region(s string) error {
	if !reRegion.MatchString(s) {
		return fmt.Errorf("invalid region %q: expected a name like ap-northeast-1, or \"*\"", s)
	}
	return nil
}

// Partition checks the ARN's partition field.
func Partition(s string) error {
	if !rePartition.MatchString(s) {
		return fmt.Errorf("invalid partition %q: expected aws, aws-cn, aws-us-gov, another aws- partition, or \"*\"", s)
	}
	return nil
}
