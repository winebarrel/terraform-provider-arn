// Code generated from the AWS service reference using cmd/gen; DO NOT EDIT.
//
// Service: apptest
// Source: https://servicereference.us-east-1.amazonaws.com/v1/apptest/apptest.json
// Functions: 4
//
// Regenerate with: make gen

package arnspec

func init() {
	register([]Spec{
		{Name: "apptest_test_case", Service: "apptest", Resource: "TestCase", Template: "arn:${Partition}:apptest:${Region}:${Account}:testcase/${TestCaseId}"},
		{Name: "apptest_test_configuration", Service: "apptest", Resource: "TestConfiguration", Template: "arn:${Partition}:apptest:${Region}:${Account}:testconfiguration/${TestConfigurationId}"},
		{Name: "apptest_test_run", Service: "apptest", Resource: "TestRun", Template: "arn:${Partition}:apptest:${Region}:${Account}:testrun/${TestRunId}"},
		{Name: "apptest_test_suite", Service: "apptest", Resource: "TestSuite", Template: "arn:${Partition}:apptest:${Region}:${Account}:testsuite/${TestSuiteId}"},
	})
}
