package old

import "encoding/json"

type GetMeResponse struct { //Deprecated
	Ok     bool   `json:"ok"`
	Result Result `json:"result"`
}

type Result struct { //Deprecated
	ID                      int    `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	Username                string `json:"username"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

type SendMessageRequest struct { //Deprecated
	ChatId                int64     `json:"chat_id,omitempty"`        //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string    `json:"text,omitempty"`           //Text of the message to be sent
	ParseMode             ParseMode `json:"parse_mode"`               //Optional 	Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool      `json:"disable_web_page_preview"` //Optional 	Disables link previews for links in this message
	DisableNotification   bool      `json:"disable_notification"`     //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId      int       `json:"reply_to_message_id"`      //If the message is a reply, ID of the original message
	//reply_markup
}

type PollRequest struct { //Deprecated
	Id                    int64        `json:"id"`
	ChatId                int64        `json:"chat_id,omitempty"`       //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question              string       `json:"question,omitempty"`      //PollRequest question, 1-255 characters
	Options               []string `json:"options,omitempty"`       //A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous           bool         `json:"is_anonymous"`            //True, if the poll needs to be anonymous, defaults to True
	Type                  string       `json:"type"`                    //PollRequest type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"` //True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId       int64        `json:"correct_option_id"`       //0-based identifier of the correct answer option, required for polls in quiz mode
	IsClosed              bool         `json:"is_closed"`               //Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool         `json:"disable_notification"`    //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId      int          `json:"reply_to_message_id"`     //If the message is a reply, ID of the original message
}


type PollResponse struct { //Deprecated
	Id                    string       `json:"id"`
	ChatId                int64        `json:"chat_id,omitempty"`       //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question              string       `json:"question,omitempty"`      //PollResponse question, 1-255 characters
	Options               []PollOption `json:"options,omitempty"`       //A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous           bool         `json:"is_anonymous"`            //True, if the poll needs to be anonymous, defaults to True
	Type                  string       `json:"type"`                    //PollResponse type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers"` //True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId       int64        `json:"correct_option_id"`       //0-based identifier of the correct answer option, required for polls in quiz mode
	IsClosed              bool         `json:"is_closed"`               //Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool         `json:"disable_notification"`    //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId      int          `json:"reply_to_message_id"`     //If the message is a reply, ID of the original message
}

type PollOption struct { //Deprecated
	Text       string `json:"text,omitempty"` //	Option text, 1-100 characters
	VoterCount string `json:"voter_count"`    //Number of users that voted for this option
}


type ParseMode string

const (
	MarkdownV2 ParseMode = "MarkdownV2"
	Markdown   ParseMode = "Markdown"
	HTML       ParseMode = "HTML"
)

// Message is returned by almost every request, and contains data about
// almost anything.
type Message struct { //Deprecated
	MessageID             int                `json:"message_id,omitempty"`
	Date                  int                `json:"date,omitempty"`
	Chat                  *Chat              `json:"chat,omitempty"`
	From                  *User              `json:"from"`                    // optional
	ForwardFrom           *User              `json:"forward_from"`            // optional
	ForwardFromChat       *Chat              `json:"forward_from_chat"`       // optional
	ForwardFromMessageID  int                `json:"forward_from_message_id"` // optional
	ForwardDate           int                `json:"forward_date"`            // optional
	Poll                  *PollResponse      `json:"poll"`                    // optionaloptional
	ReplyToMessage        *Message           `json:"reply_to_message"`        // optional
	EditDate              int                `json:"edit_date"`               // optional
	Text                  string             `json:"text"`                    // optional
	Entities              []MessageEntity    `json:"entities"`                // optional
	CaptionEntities       []MessageEntity    `json:"caption_entities"`        // optional
	Audio                 *Audio             `json:"audio"`                   // optional
	Document              *Document          `json:"document"`                // optional
	Animation             *ChatAnimation     `json:"animation"`               // optional
	Game                  *Game              `json:"game"`                    // optional
	Photo                 *[]PhotoSize       `json:"photo"`                   // optional
	Sticker               *Sticker           `json:"sticker"`                 // optional
	Video                 *Video             `json:"video"`                   // optional
	VideoNote             *VideoNote         `json:"video_note"`              // optional
	Voice                 *Voice             `json:"voice"`                   // optional
	Caption               string             `json:"caption"`                 // optional
	Contact               *Contact           `json:"contact"`                 // optional
	Location              *Location          `json:"location"`                // optional
	Venue                 *Venue             `json:"venue"`                   // optional
	NewChatMembers        *[]User            `json:"new_chat_members"`        // optional
	LeftChatMember        *User              `json:"left_chat_member"`        // optional
	NewChatTitle          string             `json:"new_chat_title"`          // optional
	NewChatPhoto          *[]PhotoSize       `json:"new_chat_photo"`          // optional
	DeleteChatPhoto       bool               `json:"delete_chat_photo"`       // optional
	GroupChatCreated      bool               `json:"group_chat_created"`      // optional
	SuperGroupChatCreated bool               `json:"supergroup_chat_created"` // optional
	ChannelChatCreated    bool               `json:"channel_chat_created"`    // optional
	MigrateToChatID       int64              `json:"migrate_to_chat_id"`      // optional
	MigrateFromChatID     int64              `json:"migrate_from_chat_id"`    // optional
	PinnedMessage         *Message           `json:"pinned_message"`          // optional
	Invoice               *Invoice           `json:"invoice"`                 // optional
	SuccessfulPayment     *SuccessfulPayment `json:"successful_payment"`      // optional
	PassportData          *PassportData      `json:"passport_data,omitempty"` // optional
}

// User is a user on Telegram.
type User struct { //Deprecated
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`     // optional
	UserName     string `json:"username"`      // optional
	LanguageCode string `json:"language_code"` // optional
	IsBot        bool   `json:"is_bot"`        // optional
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct { //Deprecated
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// Chat contains information about the place a message was sent.
type Chat struct { //Deprecated
	ID                  int64      `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`                          // optional
	UserName            string     `json:"username"`                       // optional
	FirstName           string     `json:"first_name"`                     // optional
	LastName            string     `json:"last_name"`                      // optional
	AllMembersAreAdmins bool       `json:"all_members_are_administrators"` // optional
	Photo               *ChatPhoto `json:"photo"`
	Description         string     `json:"description,omitempty"` // optional
	InviteLink          string     `json:"invite_link,omitempty"` // optional
	PinnedMessage       *Message   `json:"pinned_message"`        // optional
}

type MessageEntity struct { //Deprecated
	Type   string `json:"type"`
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`  // optional
	User   *User  `json:"user"` // optional
}

// Audio contains information about audio.
type Audio struct { //Deprecated
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"` // optional
	Title     string `json:"title"`     // optional
	MimeType  string `json:"mime_type"` // optional
	FileSize  int    `json:"file_size"` // optional
}

// PhotoSize contains information about photos.
type PhotoSize struct { //Deprecated
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"` // optional
}

// Document contains information about a document.
type Document struct { //Deprecated
	FileID    string     `json:"file_id"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileName  string     `json:"file_name"` // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// Sticker contains information about a sticker.
type Sticker struct { //Deprecated
	FileID     string     `json:"file_id"`
	Width      int        `json:"width"`
	Height     int        `json:"height"`
	Thumbnail  *PhotoSize `json:"thumb"`       // optional
	Emoji      string     `json:"emoji"`       // optional
	FileSize   int        `json:"file_size"`   // optional
	SetName    string     `json:"set_name"`    // optional
	IsAnimated bool       `json:"is_animated"` // optional
}

// ChatAnimation contains information about an animation.
type ChatAnimation struct { //Deprecated
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileName  string     `json:"file_name"` // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// Video contains information about a video.
type Video struct { //Deprecated
	FileID    string     `json:"file_id"`
	Width     int        `json:"width"`
	Height    int        `json:"height"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	MimeType  string     `json:"mime_type"` // optional
	FileSize  int        `json:"file_size"` // optional
}

// VideoNote contains information about a video.
type VideoNote struct { //Deprecated
	FileID    string     `json:"file_id"`
	Length    int        `json:"length"`
	Duration  int        `json:"duration"`
	Thumbnail *PhotoSize `json:"thumb"`     // optional
	FileSize  int        `json:"file_size"` // optional
}

// Voice contains information about a voice.
type Voice struct { //Deprecated
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"` // optional
	FileSize int    `json:"file_size"` // optional
}

// Contact contains information about a contact.
// Note that LastName and UserID may be empty.
type Contact struct { //Deprecated
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"` // optional
	UserID      int    `json:"user_id"`   // optional
}

// Location contains information about a place.
type Location struct { //Deprecated
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

// Venue contains information about a venue, including its Location.
type Venue struct { //Deprecated
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"` // optional
}

// Invoice contains basic information about an invoice.
type Invoice struct { //Deprecated
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

// Game is a game within Telegram.
type Game struct { //Deprecated
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct { //Deprecated
	Currency                string     `json:"currency"`
	TotalAmount             int        `json:"total_amount"`
	InvoicePayload          string     `json:"invoice_payload"`
	ShippingOptionID        string     `json:"shipping_option_id,omitempty"`
	OrderInfo               *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string     `json:"provider_payment_charge_id"`
}

// Animation is a GIF animation demonstrating the game.
type Animation struct { //Deprecated
	FileID   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

// OrderInfo represents information about an order.
type OrderInfo struct { //Deprecated
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct { //Deprecated
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

//APIResponse standard api response
type APIResponse struct { //Deprecated
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result"`
	ErrorCode   int                 `json:"error_code"`
	Description string              `json:"description"`
	Parameters  *ResponseParameters `json:"parameters"`
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct { //Deprecated
	MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
	RetryAfter      int   `json:"retry_after"`        // optional
}

type UpdateQueryParams struct { //Deprecated
	Offset         int
	Limit          int
	Timeout        int
	AllowedUpdates string
}

// PollAnswer pollAnswer
type PollAnswer struct { //Deprecated
	PollId    string `json:"poll_id"`
	User      *User  `json:"user"`
	OptionIds []int  `json:"option_ids"`
}

// UpdateResponse is an update response, from GetUpdates.
type UpdateResponse struct { //Deprecated
	UpdateID           int                 `json:"update_id"`
	Message            *Message            `json:"message"`
	PollAnswer         *PollAnswer         `json:"poll_answer"`
	EditedMessage      *Message            `json:"edited_message"`
	ChannelPost        *Message            `json:"channel_post"`
	EditedChannelPost  *Message            `json:"edited_channel_post"`
	InlineQuery        *InlineQuery        `json:"inline_query"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	CallbackQuery      *CallbackQuery      `json:"callback_query"`
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`
}

// InlineQuery is a Query from Telegram for an inline request.
type InlineQuery struct { //Deprecated
	ID       string    `json:"id"`
	From     *User     `json:"from"`
	Location *Location `json:"location"` // optional
	Query    string    `json:"query"`
	Offset   string    `json:"offset"`
}

// ChosenInlineResult is an inline query result chosen by a User
type ChosenInlineResult struct { //Deprecated
	ResultID        string    `json:"result_id"`
	From            *User     `json:"from"`
	Location        *Location `json:"location"`
	InlineMessageID string    `json:"inline_message_id"`
	Query           string    `json:"query"`
}

// CallbackQuery is data sent when a keyboard button with callback data
// is clicked.
type CallbackQuery struct { //Deprecated
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"message"`           // optional
	InlineMessageID string   `json:"inline_message_id"` // optional
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`            // optional
	GameShortName   string   `json:"game_short_name"` // optional
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct { //Deprecated
	ID              string           `json:"id"`
	From            *User            `json:"from"`
	InvoicePayload  string           `json:"invoice_payload"`
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct { //Deprecated
	ID               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int        `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionID string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info,omitempty"`
}

//InlineKeyboardMarkup This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct { //Deprecated
	InlineKeyboardButton InlineKeyboardButton `json:"inline_keyboard"` //Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

//InlineKeyboardButton This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct { //Deprecated
	Text              string `json:"text,omitempty"`      //Label text on the button
	Url               string `json:"text"`                //Optional. HTTP or tg:// url to be opened when button is pressed
	//LoginUrl          string `json:"login_url"`           //Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	CallbackData      string `json:"callback_data"`       //Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery string `json:"switch_inline_query"` //Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.
	//callback_game 	CallbackGame 	Optional. Description of the game that will be launched when the user presses the button.
	//pay 	Boolean 	Optional. Specify True, to send a Pay button.
}