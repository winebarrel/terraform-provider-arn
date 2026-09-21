// Package arnspec holds the ARN templates every generated function builds
// from, and the logic that turns a template plus a set of values into an ARN.
//
// The templates come from the AWS service reference feed
// (https://servicereference.us-east-1.amazonaws.com), which publishes an
// ARNFormats list per resource type, for example
//
//	arn:${Partition}:iam::${Account}:role/${RoleNameWithPath}
//
// cmd/gen turns that feed into spec_gen.go. Nothing here reaches the network:
// a provider must be able to declare its functions at startup, before any
// network call would be allowed to fail.
package arnspec

import (
	"fmt"
	"strings"
)

// Spec is one generated function: its name, where it came from, and the
// template it fills in.
type Spec struct {
	// Name is the Terraform function name, e.g. "iam_role".
	Name string
	// Service and Resource are the feed's own names, kept for documentation
	// and for error messages that need to point back at the source data.
	Service  string
	Resource string
	// Template is the raw ARNFormats entry.
	Template string

	// Args are the placeholders the caller supplies, in template order, as
	// snake_case parameter names. Derived from Template by Parse.
	Args []string
	// argsRaw are the same placeholders under their original feed names,
	// used when reporting which part of the template was left empty.
	argsRaw []string

	// argField is the ARN field each argument lands in, counted in colons:
	// 1 partition, 2 service, 3 region, 4 account, 5 and beyond the resource
	// part. Which checks apply to an argument follows from this, not from
	// what the feed happens to have named the placeholder. See Build.
	argField []int

	NeedsRegion  bool
	NeedsAccount bool

	// parts and holes describe Template in substitution order: the ARN is
	// parts[0] + holes[0] + parts[1] + ... with one more part than hole.
	parts []string
	holes []string
}

// Parse fills in the derived fields of a spec from its Template. It is called
// once per spec at init.
func Parse(s *Spec) error {
	parts, holes, err := split(s.Template)
	if err != nil {
		return fmt.Errorf("%s: %w", s.Name, err)
	}
	s.parts, s.holes = parts, holes

	// A placeholder is classified by the ARN field it sits in, not by what
	// the feed calls it. The feed is not consistent about the names: the
	// account field is ${Account} in most templates but ${AccountId} in
	// chime, datasync and sso, ${ManagementAccountId} in account and
	// ${VpcOwnerAccount} in kafka. Going by position covers all of them, and
	// it does not mistake the ${AccountId} that organizations and sso put in
	// the resource part for the account field.
	//
	// colons counts the field separators seen so far. An ARN is
	// arn:partition:service:region:account:resource, so everything from the
	// fifth colon on is the resource part.
	colons := 0
	for i, h := range holes {
		colons += strings.Count(parts[i], ":")
		switch colons {
		case partitionFieldIndex:
			// Always available: it defaults to "aws".
		case regionFieldIndex:
			s.NeedsRegion = true
		case accountFieldIndex:
			s.NeedsAccount = true
		default:
			s.argsRaw = append(s.argsRaw, h)
			s.Args = append(s.Args, SnakeCase(h))
			s.argField = append(s.argField, colons)
		}
	}
	return nil
}

// Colon counts that identify an ARN's fields:
// arn:partition:service:region:account:resource. The service field has no
// entry because nothing is substituted into it from the configuration.
const (
	partitionFieldIndex = 1
	regionFieldIndex    = 3
	accountFieldIndex   = 4
	resourceFieldIndex  = 5
)

// split breaks a template into literal parts and placeholder names.
func split(tmpl string) (parts, holes []string, err error) {
	rest := tmpl
	for {
		i := strings.Index(rest, "${")
		if i < 0 {
			parts = append(parts, rest)
			return parts, holes, nil
		}
		j := strings.Index(rest[i:], "}")
		if j < 0 {
			return nil, nil, fmt.Errorf("unterminated placeholder in %q", tmpl)
		}
		parts = append(parts, rest[:i])
		holes = append(holes, rest[i+2:i+j])
		rest = rest[i+j+1:]
	}
}

// Values are the configuration-supplied substitutions.
type Values struct {
	Partition string
	Region    string
	AccountID string
}

// Build substitutes args and values into the template. args must line up with
// s.Args; the caller is expected to have checked the count already, but Build
// re-checks so a programming error surfaces as a diagnostic rather than a
// panic.
//
// An empty value is rejected rather than interpolated. A silently truncated
// ARN like "arn:aws:iam:::role/foo" would be accepted by Terraform and only
// fail much later, at apply time, with an error that points nowhere near the
// missing configuration.
func (s *Spec) Build(v Values, args []string) (string, error) {
	if len(args) != len(s.Args) {
		return "", fmt.Errorf("%s takes %d argument(s), got %d", s.Name, len(s.Args), len(args))
	}
	if v.Partition == "" {
		return "", fmt.Errorf("%s needs a partition: set partition in the configuration file, or pass { partition = ... }", s.Name)
	}
	if s.NeedsRegion && v.Region == "" {
		return "", fmt.Errorf("%s needs a region: set region in the configuration file, or pass { region = ... }", s.Name)
	}
	if s.NeedsAccount && v.AccountID == "" {
		return "", fmt.Errorf("%s needs an account id: set account_id in the configuration file, or pass { account = ... }", s.Name)
	}

	var b strings.Builder
	b.Grow(len(s.Template) + 32)
	arg, colons := 0, 0
	for i, p := range s.parts {
		b.WriteString(p)
		if i >= len(s.holes) {
			break
		}
		colons += strings.Count(p, ":")
		switch h := s.holes[i]; colons {
		case partitionFieldIndex:
			b.WriteString(v.Partition)
		case regionFieldIndex:
			b.WriteString(v.Region)
		case accountFieldIndex:
			b.WriteString(v.AccountID)
		default:
			if args[arg] == "" {
				return "", fmt.Errorf("%s: argument %s (%s) is empty", s.Name, s.Args[arg], h)
			}
			// A colon in one of the five structural fields would shift every
			// field after it, so the ARN would name something else entirely.
			// One template puts an argument there: backup_recovery_point
			// parameterises the service field.
			//
			// Inside the resource part a colon is just a character, and what
			// it separates is up to the service. S3 object keys may contain
			// one, so rejecting it there would block an ARN that is perfectly
			// valid.
			if s.argField[arg] < resourceFieldIndex && strings.Contains(args[arg], ":") {
				return "", fmt.Errorf("%s: argument %s (%s) is an ARN field and cannot contain a colon: %q",
					s.Name, s.Args[arg], h, args[arg])
			}
			b.WriteString(args[arg])
			arg++
		}
	}
	return b.String(), nil
}

// SnakeCase normalizes a feed name into a Terraform identifier.
//
// The feed is not internally consistent: resource names appear as "role",
// "certificate-authority", "IntegrationResponse", "gatewayRoute" and
// "CIS Scan Configuration", sometimes several styles within one service. The
// rules below collapse all of them to snake_case, splitting runs of capitals
// so that "TLSInspectionConfiguration" becomes
// "tls_inspection_configuration" rather than one unreadable word.
//
// Where AWS itself ran words together ("Suitedefinition") there is nothing to
// split on, and the output keeps the run.
func SnakeCase(s string) string {
	var b strings.Builder
	b.Grow(len(s) + 8)

	r := []rune(s)
	for i, c := range r {
		switch {
		case c >= 'A' && c <= 'Z':
			// Insert a break before this capital when it starts a new word:
			// after a lowercase or digit ("gatewayRoute"), or at the end of a
			// run of capitals followed by a lowercase ("TLSInspection").
			prevLower := i > 0 && (isLowerOrDigit(r[i-1]))
			runEnd := i > 0 && isUpper(r[i-1]) && i+1 < len(r) && isLower(r[i+1])
			if prevLower || runEnd {
				b.WriteByte('_')
			}
			b.WriteRune(c - 'A' + 'a')
		case isLower(c) || (c >= '0' && c <= '9'):
			b.WriteRune(c)
		default:
			// Every other rune, including '-', '/', '.', ':' and spaces, is a
			// separator. Runs collapse below.
			b.WriteByte('_')
		}
	}

	out := b.String()
	for strings.Contains(out, "__") {
		out = strings.ReplaceAll(out, "__", "_")
	}
	return strings.Trim(out, "_")
}

func isUpper(c rune) bool        { return c >= 'A' && c <= 'Z' }
func isLower(c rune) bool        { return c >= 'a' && c <= 'z' }
func isLowerOrDigit(c rune) bool { return isLower(c) || (c >= '0' && c <= '9') }

// ArgRaw returns the feed's own placeholder name for argument i, e.g.
// "RoleNameWithPath" where Args[i] is "role_name_with_path". Documentation
// and diagnostics use it to point back at the template.
func (s *Spec) ArgRaw(i int) string { return s.argsRaw[i] }

// ArgIsStructural reports whether argument i lands in one of the ARN's five
// structural fields rather than in the resource part.
func (s *Spec) ArgIsStructural(i int) bool { return s.argField[i] < resourceFieldIndex }
