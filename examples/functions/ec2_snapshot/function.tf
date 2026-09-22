# arn:aws:ec2:ap-northeast-1::snapshot/snapshot-id
output "ec2_snapshot" {
  value = provider::arn::ec2_snapshot("snapshot-id")
}
