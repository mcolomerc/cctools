package model

type ServiceResult struct {
}

type KafkaServiceResult struct {
	Topics         []Topic
	ConsumerGroups []ConsumerGroup
	ServiceResult
}