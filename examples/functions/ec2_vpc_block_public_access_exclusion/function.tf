# arn:aws:ec2:ap-northeast-1:111111111111:vpc-block-public-access-exclusion/vpc-block-public-access-exclusion-id
output "ec2_vpc_block_public_access_exclusion" {
  value = provider::arn::ec2_vpc_block_public_access_exclusion("vpc-block-public-access-exclusion-id")
}
