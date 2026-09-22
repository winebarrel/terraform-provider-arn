# arn:aws:elasticfilesystem:ap-northeast-1:111111111111:access-point/access-point-id
output "elasticfilesystem_access_point" {
  value = provider::arn::elasticfilesystem_access_point("access-point-id")
}
