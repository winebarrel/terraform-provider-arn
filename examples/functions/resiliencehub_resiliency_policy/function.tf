# arn:aws:resiliencehub:ap-northeast-1:111111111111:resiliency-policy/resiliency-policy-id
output "resiliencehub_resiliency_policy" {
  value = provider::arn::resiliencehub_resiliency_policy("resiliency-policy-id")
}
