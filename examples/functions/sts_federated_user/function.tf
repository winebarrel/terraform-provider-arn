# arn:aws:sts::111111111111:federated-user/federated-user-name
output "sts_federated_user" {
  value = provider::arn::sts_federated_user("federated-user-name")
}
