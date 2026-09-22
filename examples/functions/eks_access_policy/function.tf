# arn:aws:eks::aws:cluster-access-policy/access-policy-name
output "eks_access_policy" {
  value = provider::arn::eks_access_policy("access-policy-name")
}
