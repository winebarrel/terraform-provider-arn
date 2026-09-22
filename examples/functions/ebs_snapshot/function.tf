# arn:aws:ec2:ap-northeast-1::snapshot/snapshot-id
output "ebs_snapshot" {
  value = provider::arn::ebs_snapshot("snapshot-id")
}
