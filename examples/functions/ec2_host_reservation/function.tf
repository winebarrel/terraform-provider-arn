# arn:aws:ec2:ap-northeast-1:111111111111:host-reservation/host-reservation-id
output "ec2_host_reservation" {
  value = provider::arn::ec2_host_reservation("host-reservation-id")
}
