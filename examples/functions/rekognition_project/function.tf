# arn:aws:rekognition:ap-northeast-1:111111111111:project/project-name/creation-timestamp
output "rekognition_project" {
  value = provider::arn::rekognition_project("project-name", "creation-timestamp")
}
