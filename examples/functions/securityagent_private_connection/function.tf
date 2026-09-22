# arn:aws:securityagent:ap-northeast-1:111111111111:private-connection/private-connection-name
output "securityagent_private_connection" {
  value = provider::arn::securityagent_private_connection("private-connection-name")
}
