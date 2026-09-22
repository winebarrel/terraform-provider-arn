# arn:aws:lightsail:ap-northeast-1:111111111111:CloudFormationStackRecord/id
output "lightsail_cloud_formation_stack_record" {
  value = provider::arn::lightsail_cloud_formation_stack_record("id")
}
