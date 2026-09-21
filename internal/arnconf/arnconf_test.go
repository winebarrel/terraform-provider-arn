package arnconf_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/terraform-provider-arn/internal/arnconf"
)

func write(t *testing.T, body string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), ".arn.hcl")
	require.NoError(t, os.WriteFile(p, []byte(body), 0o600))
	return p
}

func TestLoadResolveDefaults(t *testing.T) {
	c, err := arnconf.Load(write(t, `
account_id = "111111111111"
region     = "ap-northeast-1"
`))
	require.NoError(t, err)

	v, err := c.Resolve(arnconf.Opts{})
	require.NoError(t, err)
	assert.Equal(t, "111111111111", v.AccountID)
	assert.Equal(t, "ap-northeast-1", v.Region)
	assert.Equal(t, "aws", v.Partition, "partition defaults to aws")
}

func TestNamedAccountInheritsRegion(t *testing.T) {
	c, err := arnconf.Load(write(t, `
account_id = "111111111111"
region     = "ap-northeast-1"

account "foo" {
  account_id = "222222222222"
}

account "bar" {
  account_id = "333333333333"
  region     = "us-east-1"
}
`))
	require.NoError(t, err)

	foo, err := c.Resolve(arnconf.Opts{Account: "foo"})
	require.NoError(t, err)
	assert.Equal(t, "222222222222", foo.AccountID)
	assert.Equal(t, "ap-northeast-1", foo.Region, "unset region falls back to the top level")

	bar, err := c.Resolve(arnconf.Opts{Account: "bar"})
	require.NoError(t, err)
	assert.Equal(t, "333333333333", bar.AccountID)
	assert.Equal(t, "us-east-1", bar.Region)
}

func TestOptsOverrideAccountBlock(t *testing.T) {
	c, err := arnconf.Load(write(t, `
account_id = "111111111111"
region     = "ap-northeast-1"

account "foo" {
  account_id = "222222222222"
  region     = "us-east-1"
}
`))
	require.NoError(t, err)

	v, err := c.Resolve(arnconf.Opts{Account: "foo", Region: "eu-west-1", Partition: "aws-cn"})
	require.NoError(t, err)
	assert.Equal(t, "222222222222", v.AccountID)
	assert.Equal(t, "eu-west-1", v.Region)
	assert.Equal(t, "aws-cn", v.Partition)
}

func TestUnknownAccountListsKnownNames(t *testing.T) {
	c, err := arnconf.Load(write(t, `
account "foo" {
  account_id = "222222222222"
}

account "bar" {
  account_id = "333333333333"
}
`))
	require.NoError(t, err)

	_, err = c.Resolve(arnconf.Opts{Account: "baz"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `unknown account "baz"`)
	assert.Contains(t, err.Error(), "bar, foo")
}

func TestMissingFileIsAnError(t *testing.T) {
	_, err := arnconf.Load(filepath.Join(t.TempDir(), "nope.hcl"))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nope.hcl not found")
	assert.Contains(t, err.Error(), arnconf.EnvConfig)
}

// An empty file is enough for ARNs that need nothing from it. The file is
// required, but its contents are not.
func TestEmptyFileIsUsable(t *testing.T) {
	c, err := arnconf.Load(write(t, ""))
	require.NoError(t, err)

	v, err := c.Resolve(arnconf.Opts{})
	require.NoError(t, err)
	assert.Empty(t, v.AccountID)
	assert.Equal(t, "aws", v.Partition)
}

func TestMalformedFileIsAnError(t *testing.T) {
	_, err := arnconf.Load(write(t, `account_id = `))
	require.Error(t, err)
}

func TestDuplicateAccountBlock(t *testing.T) {
	_, err := arnconf.Load(write(t, `
account "foo" {
  account_id = "222222222222"
}

account "foo" {
  account_id = "333333333333"
}
`))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate account block")
}

func TestDefaultPath(t *testing.T) {
	t.Setenv(arnconf.EnvConfig, "")
	assert.Equal(t, ".arn.hcl", arnconf.DefaultPath())
	t.Setenv(arnconf.EnvConfig, "/tmp/other.hcl")
	assert.Equal(t, "/tmp/other.hcl", arnconf.DefaultPath())
}

// The cache reads its path once, and keeps reporting the same result.
func TestCacheReadsOnce(t *testing.T) {
	path := write(t, `account_id = "111111111111"`)
	c := arnconf.NewCache(path)

	first, err := c.Get()
	require.NoError(t, err)

	require.NoError(t, os.WriteFile(path, []byte(`account_id = "222222222222"`), 0o600))

	second, err := c.Get()
	require.NoError(t, err)
	assert.Same(t, first, second, "the file is read once per process")

	v, err := second.Resolve(arnconf.Opts{})
	require.NoError(t, err)
	assert.Equal(t, "111111111111", v.AccountID)
}

// A load failure is cached too, so a malformed file reports the same
// diagnostic on every call instead of being retried per call site.
func TestCacheRemembersFailure(t *testing.T) {
	c := arnconf.NewCache(filepath.Join(t.TempDir(), "absent.hcl"))

	_, err1 := c.Get()
	require.Error(t, err1)
	_, err2 := c.Get()
	assert.Equal(t, err1, err2)
}

func TestResolveValidatesAccountID(t *testing.T) {
	for _, id := range []string{"abc", "12345", "111111111111111111", "11111111111a", "111111111111:evil", "AWS"} {
		c, err := arnconf.Load(write(t, `account_id = "`+id+`"`))
		require.NoError(t, err)
		_, err = c.Resolve(arnconf.Opts{})
		require.ErrorContains(t, err, "invalid account id", id)
	}

	// "aws" is what AWS-managed policies carry, and "*" is how an ARN
	// written for an IAM policy wildcards the field.
	for _, id := range []string{"111111111111", "aws", "*"} {
		c, err := arnconf.Load(write(t, `account_id = "`+id+`"`))
		require.NoError(t, err)
		_, err = c.Resolve(arnconf.Opts{})
		require.NoError(t, err, id)
	}
}

func TestResolveValidatesRegion(t *testing.T) {
	c, err := arnconf.Load(write(t, `account_id = "111111111111"`))
	require.NoError(t, err)

	// Every partition's naming scheme has to survive the check.
	for _, r := range []string{
		"us-east-1", "ap-northeast-1", "eu-central-1", "us-gov-west-1",
		"cn-north-1", "us-iso-east-1", "us-isob-east-1", "il-central-1",
		"ap-southeast-7", "*",
	} {
		_, err := c.Resolve(arnconf.Opts{Region: r})
		require.NoError(t, err, r)
	}

	for _, r := range []string{"not-a-region", "AP-NORTHEAST-1", "useast1", "us-east", "us-east-1:x", "aws-global"} {
		_, err := c.Resolve(arnconf.Opts{Region: r})
		require.ErrorContains(t, err, "invalid region", r)
	}
}

func TestResolveValidatesPartition(t *testing.T) {
	c, err := arnconf.Load(write(t, `account_id = "111111111111"`))
	require.NoError(t, err)

	// IAM Access Analyzer reports the supported values as "*, aws, aws-cn,
	// aws-us-gov", so the wildcard is one of them.
	for _, p := range []string{"aws", "aws-cn", "aws-us-gov", "aws-iso", "aws-iso-b", "*"} {
		_, err := c.Resolve(arnconf.Opts{Partition: p})
		require.NoError(t, err, p)
	}

	for _, p := range []string{"nonsense", "AWS", "aws:cn", "aws-", "azure"} {
		_, err := c.Resolve(arnconf.Opts{Partition: p})
		require.ErrorContains(t, err, "invalid partition", p)
	}
}

// The file is checked by the same pass as the options, so a typo there is
// reported even when the call site overrides nothing.
func TestResolveValidatesTheFileItself(t *testing.T) {
	c, err := arnconf.Load(write(t, `
account_id = "111111111111"
region     = "nihon"
`))
	require.NoError(t, err)
	_, err = c.Resolve(arnconf.Opts{})
	require.ErrorContains(t, err, `invalid region "nihon"`)
}
