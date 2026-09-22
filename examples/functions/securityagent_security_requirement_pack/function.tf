# arn:aws:securityagent:ap-northeast-1:111111111111:security-requirement-pack/security-requirement-pack-id
output "securityagent_security_requirement_pack" {
  value = provider::arn::securityagent_security_requirement_pack("security-requirement-pack-id")
}
