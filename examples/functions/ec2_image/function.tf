# arn:aws:ec2:ap-northeast-1::image/image-id
output "ec2_image" {
  value = provider::arn::ec2_image("image-id")
}
