package ai

type AI interface {
	Chat(text string) (content string, err error)
	StreamChat(text string) (ch <-chan string, err error)
}
