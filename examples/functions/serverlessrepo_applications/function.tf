# arn:aws:serverlessrepo:ap-northeast-1:111111111111:applications/resource-id
output "serverlessrepo_applications" {
  value = provider::arn::serverlessrepo_applications("resource-id")
}
