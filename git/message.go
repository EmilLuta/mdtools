package git

// Message is a structure for keeping git messages
type Message struct {
	commit      string
	authorName  string
	authorEmail string
	date        string
	body        string
}
