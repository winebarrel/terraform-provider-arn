# arn:aws:aiops:ap-northeast-1:111111111111:investigation-group/investigation-group-id
output "aiops_investigation_group" {
  value = provider::arn::aiops_investigation_group("investigation-group-id")
}
