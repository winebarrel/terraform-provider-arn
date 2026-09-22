# arn:aws:interconnect:ap-northeast-1:111111111111:environment/id
output "interconnect_environment" {
  value = provider::arn::interconnect_environment("id")
}
