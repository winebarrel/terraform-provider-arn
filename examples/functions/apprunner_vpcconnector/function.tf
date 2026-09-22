# arn:aws:apprunner:ap-northeast-1:111111111111:vpcconnector/vpc-connector-name/vpc-connector-version/vpc-connector-id
output "apprunner_vpcconnector" {
  value = provider::arn::apprunner_vpcconnector("vpc-connector-name", "vpc-connector-version", "vpc-connector-id")
}
