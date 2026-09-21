// Package arnconf loads the .arn.hcl file that supplies the account, region
// and partition an ARN is built from.
//
// Provider-defined functions cannot read provider configuration, so the
// values have to come from somewhere the function itself can reach. A file
// in the Terraform project root keeps them visible and diffable, unlike an
// environment variable or an STS call made behind the practitioner's back.
package arnconf

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/winebarrel/terraform-provider-arn/internal/arnvalue"
)

const (
	// DefaultFilename is looked up relative to the process working
	// directory, which is where terraform was invoked, i.e. the project root.
	DefaultFilename = ".arn.hcl"

	// EnvConfig overrides the path to the configuration file.
	EnvConfig = "ARN_CONFIG"

	// DefaultPartition is used when neither the file nor the call site says
	// otherwise. Every commercial region lives in "aws"; "aws-cn" and
	// "aws-us-gov" have to be asked for.
	DefaultPartition = "aws"
)

// File is the decoded shape of .arn.hcl.
type File struct {
	AccountID *string   `hcl:"account_id"`
	Region    *string   `hcl:"region"`
	Partition *string   `hcl:"partition"`
	Accounts  []Account `hcl:"account,block"`
}

// Account is a named account block. Fields left unset fall back to the
// top-level defaults rather than to nothing, so a block that only overrides
// account_id keeps the default region.
type Account struct {
	Name      string  `hcl:"name,label"`
	AccountID *string `hcl:"account_id"`
	Region    *string `hcl:"region"`
	Partition *string `hcl:"partition"`
}

// Values is a fully resolved set of substitutions. An empty string means the
// value was never supplied; it is only an error if the ARN template needs it.
type Values struct {
	AccountID string
	Region    string
	Partition string
}

// Config is the loaded file plus the resolution rules. The zero value is not
// usable; call Load.
type Config struct {
	path     string
	file     File
	accounts map[string]Account
}

// ErrUnknownAccount is returned when an account name is not declared in the
// configuration file. It carries the known names so the diagnostic can list
// them.
type ErrUnknownAccount struct {
	Name  string
	Known []string
}

func (e *ErrUnknownAccount) Error() string {
	if len(e.Known) == 0 {
		return fmt.Sprintf("unknown account %q: no account blocks are declared", e.Name)
	}
	return fmt.Sprintf("unknown account %q: declared accounts are %s", e.Name, strings.Join(e.Known, ", "))
}

// Path returns the file path this configuration was loaded from, whether or
// not the file existed.
func (c *Config) Path() string { return c.path }

// Load reads the configuration file. The file is required: a few ARN shapes
// need nothing from it (an S3 bucket ARN carries neither account nor region),
// but letting those work without it would mean the provider behaves
// differently depending on which function a configuration happens to call
// first. Reporting the absence once, at load time, is easier to act on.
func Load(path string) (*Config, error) {
	c := &Config{path: path, accounts: map[string]Account{}}

	src, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s not found: create it, or point %s at another path", path, EnvConfig)
		}
		return nil, err
	}

	parser := hclparse.NewParser()
	f, diags := parser.ParseHCL(src, path)
	if diags.HasErrors() {
		return nil, diags
	}
	if diags := gohcl.DecodeBody(f.Body, nil, &c.file); diags.HasErrors() {
		return nil, diags
	}

	for _, a := range c.file.Accounts {
		if _, dup := c.accounts[a.Name]; dup {
			return nil, fmt.Errorf("%s: duplicate account block %q", path, a.Name)
		}
		c.accounts[a.Name] = a
	}
	return c, nil
}

// DefaultPath returns the configuration path: ARN_CONFIG when set, otherwise
// .arn.hcl in the working directory.
func DefaultPath() string {
	if p := os.Getenv(EnvConfig); p != "" {
		return p
	}
	return DefaultFilename
}

// Opts are the per-call overrides taken from a function's trailing options
// argument.
type Opts struct {
	Account   string // name of an account block
	AccountID string // literal account id, bypassing the account blocks
	Region    string
	Partition string
}

// Resolve layers the call-site options over the named account block over the
// top-level defaults. Precedence is narrowest-first: an explicit region in
// the options wins over the account block's region, which wins over the
// top-level region.
func (c *Config) Resolve(o Opts) (Values, error) {
	v := Values{
		AccountID: deref(c.file.AccountID),
		Region:    deref(c.file.Region),
		Partition: deref(c.file.Partition),
	}

	if o.Account != "" {
		a, ok := c.accounts[o.Account]
		if !ok {
			return Values{}, &ErrUnknownAccount{Name: o.Account, Known: c.accountNames()}
		}
		// An account block inherits every field it does not set, so a block
		// that only names an id still resolves the default region.
		override(&v.AccountID, a.AccountID)
		override(&v.Region, a.Region)
		override(&v.Partition, a.Partition)
	}

	if o.AccountID != "" {
		v.AccountID = o.AccountID
	}
	if o.Region != "" {
		v.Region = o.Region
	}
	if o.Partition != "" {
		v.Partition = o.Partition
	}
	if v.Partition == "" {
		v.Partition = DefaultPartition
	}
	if err := v.validate(); err != nil {
		return Values{}, err
	}
	return v, nil
}

// validate rejects values that cannot be part of a well-formed ARN. It runs
// after resolution, so it covers both the configuration file and the
// call-site options without having to check each separately.
func (v Values) validate() error {
	if v.AccountID != "" {
		if err := arnvalue.AccountID(v.AccountID); err != nil {
			return err
		}
	}
	if v.Region != "" {
		if err := arnvalue.Region(v.Region); err != nil {
			return err
		}
	}
	return arnvalue.Partition(v.Partition)
}

func (c *Config) accountNames() []string {
	names := make([]string, 0, len(c.accounts))
	for n := range c.accounts {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}

func deref(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}

func override(dst *string, src *string) {
	if src != nil {
		*dst = *src
	}
}

// Cache memoizes a single Config for the lifetime of the provider process.
// Terraform keeps the provider alive for the whole plan or apply, so the
// file is read once no matter how many ARNs a configuration builds.
//
// The path is fixed at construction rather than passed to Get, so that the
// memoized value cannot disagree with the path a caller asked for.
type Cache struct {
	path string
	once sync.Once
	cfg  *Config
	err  error
}

// NewCache returns a cache that will read path on first use.
func NewCache(path string) *Cache {
	return &Cache{path: path}
}

// Get loads the configuration on first use. The error, if any, is cached
// too: a malformed file should report the same diagnostic on every function
// call rather than being retried per call.
func (c *Cache) Get() (*Config, error) {
	c.once.Do(func() {
		c.cfg, c.err = Load(c.path)
	})
	return c.cfg, c.err
}
