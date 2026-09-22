# arn:aws:aoss:ap-northeast-1:111111111111:collection/collection-id
output "aoss_collection" {
  value = provider::arn::aoss_collection("collection-id")
}
