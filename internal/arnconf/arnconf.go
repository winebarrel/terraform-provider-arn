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
	"path/filepath"
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
	// Import names one other file to read before this one, as a path relative
	// to the directory of the file that names it, or an absolute path.
	Import *string `hcl:"import"`

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

// Config is the loaded configuration: the defaults every ARN starts from, and
// the named accounts. The zero value is not usable; call Load.
//
// The defaults are kept as their own fields rather than as a File, because a
// File is one file as written and these are the result of merging several.
// Reusing the type would leave Import and Accounts sitting here meaning
// nothing.
type Config struct {
	path string

	accountID *string
	region    *string
	partition *string

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

// Load reads the configuration file and the file it imports, if any. The file
// is required: a few ARN shapes need nothing from it (an S3 bucket ARN
// carries neither account nor region), but letting those work without it
// would mean the provider behaves differently depending on which function a
// configuration happens to call first. Reporting the absence once, at load
// time, is easier to act on.
//
// The imported file is applied first and the importing file second, so a
// value set closer to where you are reading wins. Only values that are
// actually set take part: a file that names an account id and nothing else
// leaves the region it inherited alone.
func Load(path string) (*Config, error) {
	root, err := parseFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s not found: create it, or point %s at another path", path, EnvConfig)
		}
		return nil, err
	}

	c := &Config{path: path, accounts: map[string]Account{}}

	if root.Import != nil {
		p := resolveImport(path, *root.Import)
		f, err := parseFile(p)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("%s: import %q not found", path, *root.Import)
			}
			return nil, err
		}
		// An import that imports is rejected rather than ignored. Ignoring it
		// would leave a file whose contents are quietly not being used, with
		// nothing to tell you why the account you declared is missing.
		if f.Import != nil {
			return nil, fmt.Errorf("%s: imported files cannot import, but %s does", path, p)
		}
		c.merge(f)
	}

	c.merge(root)
	return c, nil
}

// resolveImport turns an import path into a filesystem path. A relative path
// resolves against the directory of the file that named it, not the working
// directory, so a pair of files can be moved together.
func resolveImport(from, path string) string {
	if filepath.IsAbs(path) {
		return path
	}
	return filepath.Join(filepath.Dir(from), path)
}

// parseFile reads and decodes one file. Duplicate account blocks are an error
// here, unlike across files: within a single file there is nothing to
// override, so a repeated name is a typo.
func parseFile(path string) (File, error) {
	var f File

	src, err := os.ReadFile(path)
	if err != nil {
		// The caller says what a missing file means: the two cases want
		// different advice, and only one of them is about ARN_CONFIG.
		return f, err
	}

	parser := hclparse.NewParser()
	parsed, diags := parser.ParseHCL(src, path)
	if diags.HasErrors() {
		return f, diags
	}
	if diags := gohcl.DecodeBody(parsed.Body, nil, &f); diags.HasErrors() {
		return f, diags
	}

	seen := make(map[string]struct{}, len(f.Accounts))
	for _, a := range f.Accounts {
		if _, dup := seen[a.Name]; dup {
			return f, fmt.Errorf("%s: duplicate account block %q", path, a.Name)
		}
		seen[a.Name] = struct{}{}
	}
	return f, nil
}

// merge applies one file over what is already loaded. Set values win over
// unset ones, and an account block replaces one of the same name that the
// imported file declared, which is the point of importing a file and then
// adjusting it.
func (c *Config) merge(f File) {
	overridePtr(&c.accountID, f.AccountID)
	overridePtr(&c.region, f.Region)
	overridePtr(&c.partition, f.Partition)

	for _, a := range f.Accounts {
		c.accounts[a.Name] = a
	}
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
		AccountID: deref(c.accountID),
		Region:    deref(c.region),
		Partition: deref(c.partition),
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

// override sets dst from src when src says anything. The two shapes are
// separate because the two layers differ: merging files keeps "unset" as a
// nil pointer so a later file can still fill it in, while resolving has
// already collapsed unset to the empty string.
func override(dst *string, src *string) {
	if src != nil {
		*dst = *src
	}
}

func overridePtr(dst **string, src *string) {
	if src != nil {
		*dst = src
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
