# arn:aws:iotwireless:ap-northeast-1:111111111111:Destination/destination-name
output "iotwireless_destination" {
  value = provider::arn::iotwireless_destination("destination-name")
}
