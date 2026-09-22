# arn:aws:eks:ap-northeast-1:111111111111:identityproviderconfig/cluster-name/identity-provider-type/identity-provider-config-name/uuid
output "eks_identityproviderconfig" {
  value = provider::arn::eks_identityproviderconfig("cluster-name", "identity-provider-type", "identity-provider-config-name", "uuid")
}
