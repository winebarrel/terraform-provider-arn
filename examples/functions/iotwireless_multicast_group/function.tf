# arn:aws:iotwireless:ap-northeast-1:111111111111:MulticastGroup/multicast-group-id
output "iotwireless_multicast_group" {
  value = provider::arn::iotwireless_multicast_group("multicast-group-id")
}
