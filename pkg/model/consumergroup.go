package model

type ConsumerGroup struct {
	ConsumerGroupID    string
	PartitionsAssignor string
	State              string
	Consumers          []Consumer
	LagSummary         LagSummary
	Lags               []Lag
}

type Consumer struct {
	ConsumerId string
	InstanceId string
	ClientId   string
}

type LagSummary struct {
	Max_Lag              interface{}
	Total_Lag            interface{}
	Max_Lag_Topic_Name   string
	Max_Lag_Partition_Id interface{}
	Max_Lag_Consumer_Id  string
}

type Lag struct {
	TopicName     string
	Partition     interface{}
	ConsumerId    string
	InstanceId    string
	ClientId      string
	CurrentOffset interface{}
	LogEndOffset  interface{}
	Lag           interface{}
}
