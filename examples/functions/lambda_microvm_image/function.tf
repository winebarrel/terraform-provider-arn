# arn:aws:lambda:ap-northeast-1:111111111111:microvm-image:microvm-image-name
output "lambda_microvm_image" {
  value = provider::arn::lambda_microvm_image("microvm-image-name")
}
