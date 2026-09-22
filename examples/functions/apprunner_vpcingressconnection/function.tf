# arn:aws:apprunner:ap-northeast-1:111111111111:vpcingressconnection/vpc-ingress-connection-name/vpc-ingress-connection-id
output "apprunner_vpcingressconnection" {
  value = provider::arn::apprunner_vpcingressconnection("vpc-ingress-connection-name", "vpc-ingress-connection-id")
}
