# arn:aws:repostspace:ap-northeast-1:111111111111:space/resource-id
output "repostspace_space" {
  value = provider::arn::repostspace_space("resource-id")
}
