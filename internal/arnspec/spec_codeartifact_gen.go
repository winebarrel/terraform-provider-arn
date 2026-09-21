// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: codeartifact
// Source: https://servicereference.us-east-1.amazonaws.com/v1/codeartifact/codeartifact.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "codeartifact_domain", Service: "codeartifact", Resource: "domain", Template: "arn:${Partition}:codeartifact:${Region}:${Account}:domain/${DomainName}"},
		{Name: "codeartifact_package", Service: "codeartifact", Resource: "package", Template: "arn:${Partition}:codeartifact:${Region}:${Account}:package/${DomainName}/${RepositoryName}/${PackageFormat}/${PackageNamespace}/${PackageName}"},
		{Name: "codeartifact_package_group", Service: "codeartifact", Resource: "package-group", Template: "arn:${Partition}:codeartifact:${Region}:${Account}:package-group/${DomainName}${EncodedPackageGroupPattern}"},
		{Name: "codeartifact_repository", Service: "codeartifact", Resource: "repository", Template: "arn:${Partition}:codeartifact:${Region}:${Account}:repository/${DomainName}/${RepositoryName}"},
	})
}
