# arn:aws:ssm:ap-northeast-1:111111111111:association/association-id
output "ssm_association" {
  value = provider::arn::ssm_association("association-id")
}
