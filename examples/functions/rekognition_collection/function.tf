# arn:aws:rekognition:ap-northeast-1:111111111111:collection/collection-id
output "rekognition_collection" {
  value = provider::arn::rekognition_collection("collection-id")
}
