# arn:aws:iam::111111111111:federated-user/user-name
output "iam_federated_user" {
  value = provider::arn::iam_federated_user("user-name")
}
