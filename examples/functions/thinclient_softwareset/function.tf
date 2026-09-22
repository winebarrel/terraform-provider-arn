# arn:aws:thinclient:ap-northeast-1:111111111111:softwareset/software-set-id
output "thinclient_softwareset" {
  value = provider::arn::thinclient_softwareset("software-set-id")
}
