# arn:aws:medialive:ap-northeast-1:111111111111:multiplex:multiplex-id
output "medialive_multiplex" {
  value = provider::arn::medialive_multiplex("multiplex-id")
}
