# arn:aws:eks:ap-northeast-1:111111111111:access-entry/cluster-name/iam-identity-type/iam-identity-account-id/iam-identity-name/uuid
output "eks_access_entry" {
  value = provider::arn::eks_access_entry("cluster-name", "iam-identity-type", "iam-identity-account-id", "iam-identity-name", "uuid")
}
