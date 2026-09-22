# arn:aws:aoss:ap-northeast-1:111111111111:collection-group/collection-group-id
output "aoss_collection_group" {
  value = provider::arn::aoss_collection_group("collection-group-id")
}
