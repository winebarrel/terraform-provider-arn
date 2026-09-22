# arn:aws:aws-marketplace:ap-northeast-1:111111111111:tax-compliance-profile-change-task/resource-id
output "aws_marketplace_tax_compliance_profile_change_task" {
  value = provider::arn::aws_marketplace_tax_compliance_profile_change_task("resource-id")
}
