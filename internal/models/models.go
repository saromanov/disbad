package models

// Log defines inner representation of the log data
type Log struct {
	Operation string `json:"operation"`
	Key       []byte `json:"key"`
	Value     []byte `json:"value"`
}

type Node struct {
	ID          string
	GrpcAddress string
	RaftAddress string
	ClusterID   string
}
