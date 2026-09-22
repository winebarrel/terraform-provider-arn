# arn:aws:iotevents:ap-northeast-1:111111111111:alarmModel/alarm-model-name
output "iotevents_alarm_model" {
  value = provider::arn::iotevents_alarm_model("alarm-model-name")
}
