# arn:aws:ssm-contacts:ap-northeast-1:111111111111:rotation/rotation-id
output "ssm_contacts_rotation" {
  value = provider::arn::ssm_contacts_rotation("rotation-id")
}
