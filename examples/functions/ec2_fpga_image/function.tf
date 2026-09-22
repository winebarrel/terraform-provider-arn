# arn:aws:ec2:ap-northeast-1:111111111111:fpga-image/fpga-image-id
output "ec2_fpga_image" {
  value = provider::arn::ec2_fpga_image("fpga-image-id")
}
