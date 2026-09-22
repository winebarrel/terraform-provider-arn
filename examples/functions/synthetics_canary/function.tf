# arn:aws:synthetics:ap-northeast-1:111111111111:canary:canary-name
output "synthetics_canary" {
  value = provider::arn::synthetics_canary("canary-name")
}
