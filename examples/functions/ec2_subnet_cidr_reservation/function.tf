# arn:aws:ec2:ap-northeast-1:111111111111:subnet-cidr-reservation/subnet-cidr-reservation-id
output "ec2_subnet_cidr_reservation" {
  value = provider::arn::ec2_subnet_cidr_reservation("subnet-cidr-reservation-id")
}
