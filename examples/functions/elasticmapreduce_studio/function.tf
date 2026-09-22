# arn:aws:elasticmapreduce:ap-northeast-1:111111111111:studio/studio-id
output "elasticmapreduce_studio" {
  value = provider::arn::elasticmapreduce_studio("studio-id")
}
