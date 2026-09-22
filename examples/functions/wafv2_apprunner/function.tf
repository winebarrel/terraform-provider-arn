# arn:aws:apprunner:ap-northeast-1:111111111111:service/service-name/service-id
output "wafv2_apprunner" {
  value = provider::arn::wafv2_apprunner("service-name", "service-id")
}
