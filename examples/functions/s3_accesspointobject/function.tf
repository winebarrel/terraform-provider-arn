# arn:aws:s3:ap-northeast-1:111111111111:accesspoint/access-point-name/object/object-name
output "s3_accesspointobject" {
  value = provider::arn::s3_accesspointobject("access-point-name", "object-name")
}
