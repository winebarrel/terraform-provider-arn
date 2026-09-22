# arn:aws:s3-object-lambda:ap-northeast-1:111111111111:accesspoint/access-point-name
output "s3_object_lambda_objectlambdaaccesspoint" {
  value = provider::arn::s3_object_lambda_objectlambdaaccesspoint("access-point-name")
}
