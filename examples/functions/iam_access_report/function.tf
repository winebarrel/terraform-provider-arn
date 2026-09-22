# arn:aws:iam::111111111111:access-report/entity-path
output "iam_access_report" {
  value = provider::arn::iam_access_report("entity-path")
}
