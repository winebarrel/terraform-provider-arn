# arn:aws:elastic-inference:ap-northeast-1:111111111111:elastic-inference-accelerator/accelerator-id
output "ec2_elastic_inference" {
  value = provider::arn::ec2_elastic_inference("accelerator-id")
}
