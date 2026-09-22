# arn:aws:aws-marketplace:ap-northeast-1:111111111111:verification-type/verification-type/verification-evidence/resource-id
output "aws_marketplace_verification_evidence" {
  value = provider::arn::aws_marketplace_verification_evidence("verification-type", "resource-id")
}
