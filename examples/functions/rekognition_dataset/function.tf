# arn:aws:rekognition:ap-northeast-1:111111111111:project/project-name/dataset/dataset-type/creation-timestamp
output "rekognition_dataset" {
  value = provider::arn::rekognition_dataset("project-name", "dataset-type", "creation-timestamp")
}
