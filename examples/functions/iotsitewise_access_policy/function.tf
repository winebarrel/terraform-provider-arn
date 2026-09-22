# arn:aws:iotsitewise:ap-northeast-1:111111111111:access-policy/access-policy-id
output "iotsitewise_access_policy" {
  value = provider::arn::iotsitewise_access_policy("access-policy-id")
}
