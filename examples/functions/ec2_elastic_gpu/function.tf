# arn:aws:ec2:ap-northeast-1:111111111111:elastic-gpu/elastic-gpu-id
output "ec2_elastic_gpu" {
  value = provider::arn::ec2_elastic_gpu("elastic-gpu-id")
}
