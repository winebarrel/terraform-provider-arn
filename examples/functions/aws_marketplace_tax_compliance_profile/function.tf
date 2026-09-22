# arn:aws:aws-marketplace:ap-northeast-1:111111111111:tax-compliance-profile/resource-id
output "aws_marketplace_tax_compliance_profile" {
  value = provider::arn::aws_marketplace_tax_compliance_profile("resource-id")
}
