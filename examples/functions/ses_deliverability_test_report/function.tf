# arn:aws:ses:ap-northeast-1:111111111111:deliverability-test-report/report-id
output "ses_deliverability_test_report" {
  value = provider::arn::ses_deliverability_test_report("report-id")
}
