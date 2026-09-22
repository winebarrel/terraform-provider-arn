# arn:aws:elasticloadbalancing:ap-northeast-1:111111111111:truststore/trust-store-name/trust-store-id
output "elasticloadbalancing_truststore" {
  value = provider::arn::elasticloadbalancing_truststore("trust-store-name", "trust-store-id")
}
