# arn:aws:cloudformation:ap-northeast-1:111111111111:changeSet/change-set-name/id
output "cloudformation_changeset" {
  value = provider::arn::cloudformation_changeset("change-set-name", "id")
}
