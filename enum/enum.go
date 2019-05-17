package enum

const (
	ServerStatusUnknown int32 = 0
	ServerStatusOK      int32 = 1
	ServerStatusDown    int32 = -1
	ServerStatusGFWed   int32 = -2

	PostKey = "Post-Key"
)
