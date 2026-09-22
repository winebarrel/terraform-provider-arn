# arn:aws:iam::111111111111:server-certificate/certificate-name-with-path
output "iam_server_certificate" {
  value = provider::arn::iam_server_certificate("certificate-name-with-path")
}
