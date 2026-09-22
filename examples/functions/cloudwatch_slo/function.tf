# arn:aws:cloudwatch:ap-northeast-1:111111111111:slo/slo-name
output "cloudwatch_slo" {
  value = provider::arn::cloudwatch_slo("slo-name")
}
