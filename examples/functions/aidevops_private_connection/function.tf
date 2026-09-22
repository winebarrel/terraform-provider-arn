# arn:aws:aidevops:ap-northeast-1:111111111111:private-connection/name
output "aidevops_private_connection" {
  value = provider::arn::aidevops_private_connection("name")
}
