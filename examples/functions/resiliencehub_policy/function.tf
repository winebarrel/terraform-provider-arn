# arn:aws:resiliencehub:ap-northeast-1:111111111111:policy/policy-id
output "resiliencehub_policy" {
  value = provider::arn::resiliencehub_policy("policy-id")
}
