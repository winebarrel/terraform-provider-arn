# arn:aws:neptune-db:ap-northeast-1:111111111111:cluster-resource-id/*
output "neptune_db_database" {
  value = provider::arn::neptune_db_database("cluster-resource-id")
}
