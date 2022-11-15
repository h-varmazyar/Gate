package tgBotApi

type MessageType string

const (
	TextMessage      MessageType = "Text"
	VideoMessage     MessageType = "Video"
	AudioMessage     MessageType = "Audio"
	VoiceMessage     MessageType = "Voice"
	PhotoMessage     MessageType = "photo"
	ContactMessage   MessageType = "Contact"
	StickerMessage   MessageType = "Sticker"
	UnknownMessage   MessageType = "Unknown"
	LocationMessage  MessageType = "Location"
	DocumentMessage  MessageType = "Document"
	VideoNoteMessage MessageType = "VideoNote"
)

type Chattable struct {
	MetaData string
	FileId   string
	Text     string
	Type     MessageType
}
