# arn:aws:medialive:ap-northeast-1:111111111111:inputSecurityGroup:input-security-group-id
output "medialive_input_security_group" {
  value = provider::arn::medialive_input_security_group("input-security-group-id")
}
