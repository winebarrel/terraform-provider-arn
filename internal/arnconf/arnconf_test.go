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

func TestMissingFileIsNotAnError(t *testing.T) {
	c, err := arnconf.Load(filepath.Join(t.TempDir(), "nope.hcl"))
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
