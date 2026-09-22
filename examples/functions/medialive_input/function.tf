# arn:aws:medialive:ap-northeast-1:111111111111:input:input-id
output "medialive_input" {
  value = provider::arn::medialive_input("input-id")
}
