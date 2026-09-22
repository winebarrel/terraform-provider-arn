# arn:aws:ec2:ap-northeast-1:111111111111:declarative-policies-report/declarative-policies-report-id
output "ec2_declarative_policies_report" {
  value = provider::arn::ec2_declarative_policies_report("declarative-policies-report-id")
}
