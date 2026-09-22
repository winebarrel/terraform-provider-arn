# arn:aws:ssm:ap-northeast-1:111111111111:parameter/parameter-name-without-leading-slash
output "ssm_parameter" {
  value = provider::arn::ssm_parameter("parameter-name-without-leading-slash")
}
