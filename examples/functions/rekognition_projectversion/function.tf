# arn:aws:rekognition:ap-northeast-1:111111111111:project/project-name/version/version-name/creation-timestamp
output "rekognition_projectversion" {
  value = provider::arn::rekognition_projectversion("project-name", "version-name", "creation-timestamp")
}
