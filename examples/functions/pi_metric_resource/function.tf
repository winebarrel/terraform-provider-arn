# arn:aws:pi:ap-northeast-1:111111111111:metrics/service-type/identifier
output "pi_metric_resource" {
  value = provider::arn::pi_metric_resource("service-type", "identifier")
}
