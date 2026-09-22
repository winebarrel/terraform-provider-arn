# arn:aws:rekognition:ap-northeast-1:111111111111:streamprocessor/streamprocessor-id
output "rekognition_streamprocessor" {
  value = provider::arn::rekognition_streamprocessor("streamprocessor-id")
}
