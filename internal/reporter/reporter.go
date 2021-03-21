package reporter

// Reporter is an intreface to collect data about different redis operations performance
// methods names correspond with redis operations this performance should be measured
type Reporter interface {
	SetGet() SetGetReporter
	Sdiff() SdiffReporter
	PubSub() PubSubReporter
	Smembers() SmembersReporter
}

// SetGetReporter is an interface to record the performance of Set and Get
type SetGetReporter interface {
	WillSet() error
	DidGet() error
}

// SdiffReporter is an interface to record the performance of Sdiff
type SdiffReporter interface {
	WillSdiff() error
	DidSdiff() error
}

// PubSubReporter is an interface to record the performance of Pub and Sub
type PubSubReporter interface {
	WillPub() error
	DidSub() error
}

// SmembersReporter is an interface to record the performance of Smembers
type SmembersReporter interface {
	WillSmembers(amount int) error
	DidSmembers() error
}

// ZremrangebyscoreReporter is an interface to record the performance of Zremrangebyscore
type ZremrangebyscoreReporter interface {
	WillZremrangeyscore(amount int) error
	DidZremrangeyscore() error
}
