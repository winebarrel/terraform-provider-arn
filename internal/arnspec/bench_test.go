package arnspec_test

import (
	"testing"

	"github.com/winebarrel/terraform-provider-arn/internal/arnspec"
)

// BenchmarkAll measures the one-time cost paid behind the sync.Once: parsing
// every generated template into its parts. It is amortised over the life of
// the provider process, so it matters only at startup.
func BenchmarkAll(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		_ = arnspec.All()
	}
}

func BenchmarkBuild(b *testing.B) {
	s, ok := arnspec.Lookup("iam_role")
	if !ok {
		b.Fatal("iam_role not found")
	}
	v := arnspec.Values{Partition: "aws", AccountID: "111111111111"}
	args := []string{"my-role"}

	b.ReportAllocs()
	for b.Loop() {
		if _, err := s.Build(v, args); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkBuildWithRegionAndAccount(b *testing.B) {
	s, ok := arnspec.Lookup("sqs_queue")
	if !ok {
		b.Fatal("sqs_queue not found")
	}
	v := arnspec.Values{Partition: "aws", Region: "ap-northeast-1", AccountID: "111111111111"}
	args := []string{"my-queue"}

	b.ReportAllocs()
	for b.Loop() {
		if _, err := s.Build(v, args); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkLookup(b *testing.B) {
	b.ReportAllocs()
	for b.Loop() {
		if _, ok := arnspec.Lookup("iam_role"); !ok {
			b.Fatal("not found")
		}
	}
}

// BenchmarkParseAll measures the work All() does on its first call, without
// the sync.Once in the way. Each iteration parses every generated template
// from scratch, which is what a fresh provider process pays once.
func BenchmarkParseAll(b *testing.B) {
	specs := arnspec.All()
	fresh := make([]arnspec.Spec, len(specs))

	b.ReportAllocs()
	for b.Loop() {
		for i, s := range specs {
			fresh[i] = arnspec.Spec{Name: s.Name, Service: s.Service, Resource: s.Resource, Template: s.Template}
			if err := arnspec.Parse(&fresh[i]); err != nil {
				b.Fatal(err)
			}
		}
	}
}
