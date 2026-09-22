# arn:aws:a4b:ap-northeast-1:111111111111:conference-provider/resource-id
output "a4b_conferenceprovider" {
  value = provider::arn::a4b_conferenceprovider("resource-id")
}
