# arn:aws:cleanrooms:ap-northeast-1:111111111111:collaboration/collaboration-id
output "cleanrooms_collaboration" {
  value = provider::arn::cleanrooms_collaboration("collaboration-id")
}
