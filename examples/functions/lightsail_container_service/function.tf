# arn:aws:lightsail:ap-northeast-1:111111111111:ContainerService/id
output "lightsail_container_service" {
  value = provider::arn::lightsail_container_service("id")
}
