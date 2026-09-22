# arn:aws:dms:ap-northeast-1:111111111111:instance-profile:*
output "dms_instance_profile" {
  value = provider::arn::dms_instance_profile()
}
