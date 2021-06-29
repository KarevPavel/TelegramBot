package models

//This object represents an incoming update.At most one of the optional parameters can be present in any given update.
type Update struct {
	UpdateId           int                 `json:"update_id,omitempty"`  //The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially. This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order. If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	Message            *Message            `json:"message"`              //Optional. New incoming message of any kind — text, photo, sticker, etc.
	EditedMessage      *Message            `json:"edited_message"`       //Optional. New version of a message that is known to the bot and was edited
	ChannelPost        *Message            `json:"channel_post"`         //Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	EditedChannelPost  *Message            `json:"edited_channel_post"`  //Optional. New version of a channel post that is known to the bot and was edited
	InlineQuery        *InlineQuery        `json:"inline_query"`         //Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"` //Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	CallbackQuery      *CallbackQuery      `json:"callback_query"`       //Optional. New incoming callback query
	ShippingQuery      *ShippingQuery      `json:"shipping_query"`       //Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery   *PreCheckoutQuery   `json:"pre_checkout_query"`   //Optional. New incoming pre-checkout query. Contains full information about checkout
	Poll               *Poll               `json:"poll"`                 //Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	PollAnswer         *PollAnswer         `json:"poll_answer"`          //Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
}

//Use this method to receive incoming updates using long polling (wiki). An Array of Update objects is returned.
type GetUpdates struct {
	Offset         int      `json:"offset"`          //Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers of previously received updates. By default, updates starting with the earliest unconfirmed update are returned. An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id. The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue. All previous updates will forgotten.
	Limit          int      `json:"limit"`           //Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Timeout        int      `json:"timeout"`         //Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	AllowedUpdates []string `json:"allowed_updates"` //A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
}

//Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
type SetWebhook struct {
	Url            string   `json:"url,omitempty"`   //HTTPS url to send updates to. Use an empty string to remove webhook integration
	Certificate    string   `json:"certificate"`     //Upload your public key certificate so that the root certificate in use can be checked. See our self-signed guide for details.
	MaxConnections int      `json:"max_connections"` //Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100. Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values to increase your bot’s throughput.
	AllowedUpdates []string `json:"allowed_updates"` //A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list of available update types. Specify an empty list to receive all updates regardless of type (default). If not specified, the previous setting will be used.Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted updates may be received for a short period of time.
}

//Contains information about the current status of a webhook.
type WebhookInfo struct {
	Url                  string   `json:"url,omitempty"`                    //Webhook URL, may be empty if webhook is not set up
	HasCustomCertificate bool     `json:"has_custom_certificate,omitempty"` //True, if a custom certificate was provided for webhook certificate checks
	PendingUpdateCount   int      `json:"pending_update_count,omitempty"`   //Number of updates awaiting delivery
	LastErrorDate        int      `json:"last_error_date"`                  //Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorMessage     string   `json:"last_error_message"`               //Optional. Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook
	MaxConnections       int      `json:"max_connections"`                  //Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	AllowedUpdates       []string `json:"allowed_updates"`                  //Optional. A list of update types the bot is subscribed to. Defaults to all update types
}

//This object represents a Telegram user or bot.
type User struct {
	Id                      int    `json:"id,omitempty"`                //Unique identifier for this user or bot
	IsBot                   bool   `json:"is_bot,omitempty"`            //True, if this user is a bot
	FirstName               string `json:"first_name,omitempty"`        //User‘s or bot’s first name
	LastName                string `json:"last_name"`                   //Optional. User‘s or bot’s last name
	Username                string `json:"username"`                    //Optional. User‘s or bot’s username
	LanguageCode            string `json:"language_code"`               //Optional. IETF language tag of the user's language
	CanJoinGroups           bool   `json:"can_join_groups"`             //Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"` //Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`     //Optional. True, if the bot supports inline queries. Returned only in getMe.
}

//This object represents a chat.
type Chat struct {
	Id               uint64           `json:"id,omitempty"`        //Unique identifier for this chat. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	Type             string           `json:"type,omitempty"`      //Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title            string           `json:"title"`               //Optional. Title, for supergroups, channels and group chats
	Username         string           `json:"username"`            //Optional. Username, for private chats, supergroups and channels if available
	FirstName        string           `json:"first_name"`          //Optional. First name of the other party in a private chat
	LastName         string           `json:"last_name"`           //Optional. Last name of the other party in a private chat
	Photo            *ChatPhoto       `json:"photo"`               //Optional. Chat photo. Returned only in getChat.
	Description      string           `json:"description"`         //Optional. Description, for groups, supergroups and channel chats. Returned only in getChat.
	InviteLink       string           `json:"invite_link"`         //Optional. Chat invite link, for groups, supergroups and channel chats. Each administrator in a chat generates their own invite links, so the bot must first generate the link using exportChatInviteLink. Returned only in getChat.
	PinnedMessage    *Message         `json:"pinned_message"`      //Optional. Pinned message, for groups, supergroups and channels. Returned only in getChat.
	Permissions      *ChatPermissions `json:"permissions"`         //Optional. Default chat member permissions, for groups and supergroups. Returned only in getChat.
	SlowModeDelay    int              `json:"slow_mode_delay"`     //Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user. Returned only in getChat.
	StickerSetName   string           `json:"sticker_set_name"`    //Optional. For supergroups, name of group sticker set. Returned only in getChat.
	CanSetStickerSet bool             `json:"can_set_sticker_set"` //Optional. True, if the bot can change the group sticker set. Returned only in getChat.
}

//This object represents a message.
type Message struct {
	MessageId             int                  `json:"message_id,omitempty"`    //Unique message identifier inside this chat
	From                  *User                `json:"from"`                    //Optional. Sender, empty for messages sent to channels
	Date                  int                  `json:"date,omitempty"`          //Date the message was sent in Unix time
	Chat                  *Chat                `json:"chat,omitempty"`          //Conversation the message belongs to
	ForwardFrom           *User                `json:"forward_from"`            //Optional. For forwarded messages, sender of the original message
	ForwardFromChat       *Chat                `json:"forward_from_chat"`       //Optional. For messages forwarded from channels, information about the original channel
	ForwardFromMessageId  int                  `json:"forward_from_message_id"` //Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardSignature      string               `json:"forward_signature"`       //Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSenderName     string               `json:"forward_sender_name"`     //Optional. Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages
	ForwardDate           int                  `json:"forward_date"`            //Optional. For forwarded messages, date the original message was sent in Unix time
	ReplyToMessage        *Message             `json:"reply_to_message"`        //Optional. For replies, the original message. Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	EditDate              int                  `json:"edit_date"`               //Optional. Date the message was last edited in Unix time
	MediaGroupId          string               `json:"media_group_id"`          //Optional. The unique identifier of a media message group this message belongs to
	AuthorSignature       string               `json:"author_signature"`        //Optional. Signature of the post author for messages in channels
	Text                  string               `json:"text"`                    //Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Entities              []MessageEntity      `json:"entities"`                //Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	CaptionEntities       []MessageEntity      `json:"caption_entities"`        //Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	Audio                 *Audio               `json:"audio"`                   //Optional. Message is an audio file, information about the file
	Document              *Document            `json:"document"`                //Optional. Message is a general file, information about the file
	Animation             *Animation           `json:"animation"`               //Optional. Message is an animation, information about the animation. For backward compatibility, when this field is set, the document field will also be set
	Game                  *Game                `json:"game"`                    //Optional. Message is a game, information about the game. More about games »
	Photo                 []PhotoSize          `json:"photo"`                   //Optional. Message is a photo, available sizes of the photo
	Sticker               *Sticker             `json:"sticker"`                 //Optional. Message is a sticker, information about the sticker
	Video                 *Video               `json:"video"`                   //Optional. Message is a video, information about the video
	Voice                 *Voice               `json:"voice"`                   //Optional. Message is a voice message, information about the file
	VideoNote             *VideoNote           `json:"video_note"`              //Optional. Message is a video note, information about the video message
	Caption               string               `json:"caption"`                 //Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Contact               *Contact             `json:"contact"`                 //Optional. Message is a shared contact, information about the contact
	Location              *Location            `json:"location"`                //Optional. Message is a shared location, information about the location
	Venue                 *Venue               `json:"venue"`                   //Optional. Message is a venue, information about the venue
	Poll                  *Poll                `json:"poll"`                    //Optional. Message is a native poll, information about the poll
	NewChatMembers        []User               `json:"new_chat_members"`        //Optional. New members that were added to the group or supergroup and information about them (the bot itself may be one of these members)
	LeftChatMember        *User                `json:"left_chat_member"`        //Optional. A member was removed from the group, information about them (this member may be the bot itself)
	NewChatTitle          string               `json:"new_chat_title"`          //Optional. A chat title was changed to this value
	NewChatPhoto          []PhotoSize          `json:"new_chat_photo"`          //Optional. A chat photo was change to this value
	DeleteChatPhoto       bool                 `json:"delete_chat_photo"`       //Optional. Service message: the chat photo was deleted
	GroupChatCreated      bool                 `json:"group_chat_created"`      //Optional. Service message: the group has been created
	SupergroupChatCreated bool                 `json:"supergroup_chat_created"` //Optional. Service message: the supergroup has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	ChannelChatCreated    bool                 `json:"channel_chat_created"`    //Optional. Service message: the channel has been created. This field can‘t be received in a message coming through updates, because bot can’t be a member of a channel when it is created. It can only be found in reply_to_message if someone replies to a very first message in a channel.
	MigrateToChatId       int                  `json:"migrate_to_chat_id"`      //Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	MigrateFromChatId     int                  `json:"migrate_from_chat_id"`    //Optional. The supergroup has been migrated from a group with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	PinnedMessage         *Message             `json:"pinned_message"`          //Optional. Specified message was pinned. Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	Invoice               *Invoice             `json:"invoice"`                 //Optional. Message is an invoice for a payment, information about the invoice. More about payments »
	SuccessfulPayment     *SuccessfulPayment   `json:"successful_payment"`      //Optional. Message is a service message about a successful payment, information about the payment. More about payments »
	ConnectedWebsite      string               `json:"connected_website"`       //Optional. The domain name of the website on which the user has logged in. More about Telegram Login »
	PassportData          *PassportData        `json:"passport_data"`           //Optional. Telegram Passport data
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`            //Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
}

//This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Type     string `json:"type,omitempty"`   //Type of the entity. Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.org), “email” (do-not-reply@telegram.org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames)
	Offset   int    `json:"offset,omitempty"` //Offset in UTF-16 code units to the start of the entity
	Length   int    `json:"length,omitempty"` //Length of the entity in UTF-16 code units
	Url      string `json:"url"`              //Optional. For “text_link” only, url that will be opened after user taps on the text
	User     *User  `json:"user"`             //Optional. For “text_mention” only, the mentioned user
	Language string `json:"language"`         //Optional. For “pre” only, the programming language of the entity text
}

//This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	FileId       string `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int    `json:"width,omitempty"`          //Photo width
	Height       int    `json:"height,omitempty"`         //Photo height
	FileSize     int    `json:"file_size"`                //Optional. File size
}

//This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	FileId       string     `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int        `json:"duration,omitempty"`       //Duration of the audio in seconds as defined by sender
	Performer    string     `json:"performer"`                //Optional. Performer of the audio as defined by sender or by audio tags
	Title        string     `json:"title"`                    //Optional. Title of the audio as defined by sender or by audio tags
	MimeType     string     `json:"mime_type"`                //Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`                //Optional. File size
	Thumb        *PhotoSize `json:"thumb"`                    //Optional. Thumbnail of the album cover to which the music file belongs
}

//This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	FileId       string     `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Thumb        *PhotoSize `json:"thumb"`                    //Optional. Document thumbnail as defined by sender
	FileName     string     `json:"file_name"`                //Optional. Original filename as defined by sender
	MimeType     string     `json:"mime_type"`                //Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`                //Optional. File size
}

//This object represents a video file.
type Video struct {
	FileId       string     `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width,omitempty"`          //Video width as defined by sender
	Height       int        `json:"height,omitempty"`         //Video height as defined by sender
	Duration     int        `json:"duration,omitempty"`       //Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`                    //Optional. Video thumbnail
	MimeType     string     `json:"mime_type"`                //Optional. Mime type of a file as defined by sender
	FileSize     int        `json:"file_size"`                //Optional. File size
}

//This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	FileId       string     `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int        `json:"width,omitempty"`          //Video width as defined by sender
	Height       int        `json:"height,omitempty"`         //Video height as defined by sender
	Duration     int        `json:"duration,omitempty"`       //Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`                    //Optional. Animation thumbnail as defined by sender
	FileName     string     `json:"file_name"`                //Optional. Original animation filename as defined by sender
	MimeType     string     `json:"mime_type"`                //Optional. MIME type of the file as defined by sender
	FileSize     int        `json:"file_size"`                //Optional. File size
}

//This object represents a voice note.
type Voice struct {
	FileId       string `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Duration     int    `json:"duration,omitempty"`       //Duration of the audio in seconds as defined by sender
	MimeType     string `json:"mime_type"`                //Optional. MIME type of the file as defined by sender
	FileSize     int    `json:"file_size"`                //Optional. File size
}

//This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	FileId       string     `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string     `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Length       int        `json:"length,omitempty"`         //Video width and height (diameter of the video message) as defined by sender
	Duration     int        `json:"duration,omitempty"`       //Duration of the video in seconds as defined by sender
	Thumb        *PhotoSize `json:"thumb"`                    //Optional. Video thumbnail
	FileSize     int        `json:"file_size"`                //Optional. File size
}

//This object represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number,omitempty"` //Contact's phone number
	FirstName   string `json:"first_name,omitempty"`   //Contact's first name
	LastName    string `json:"last_name"`              //Optional. Contact's last name
	UserId      int    `json:"user_id"`                //Optional. Contact's user identifier in Telegram
	Vcard       string `json:"vcard"`                  //Optional. Additional data about the contact in the form of a vCard
}

//This object represents a point on the map.
type Location struct {
	Longitude float32 `json:"longitude,omitempty"` //Longitude as defined by sender
	Latitude  float32 `json:"latitude,omitempty"`  //Latitude as defined by sender
}

//This object represents a venue.
type Venue struct {
	Location       *Location `json:"location,omitempty"` //Venue location
	Title          string    `json:"title,omitempty"`    //Name of the venue
	Address        string    `json:"address,omitempty"`  //Address of the venue
	FoursquareId   string    `json:"foursquare_id"`      //Optional. Foursquare identifier of the venue
	FoursquareType string    `json:"foursquare_type"`    //Optional. Foursquare type of the venue. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
}

//This object contains information about one answer option in a poll.
type PollOption struct {
	Text       string `json:"text,omitempty"`        //Option text, 1-100 characters
	VoterCount int    `json:"voter_count,omitempty"` //Number of users that voted for this option
}

//This object represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollId    string `json:"poll_id,omitempty"`    //Unique poll identifier
	User      *User  `json:"user,omitempty"`       //The user, who changed the answer to the poll
	OptionIds []int  `json:"option_ids,omitempty"` //0-based identifiers of answer options, chosen by the user. May be empty if the user retracted their vote.
}

//This object contains information about a poll.
type Poll struct {
	Id                    string       `json:"id,omitempty"`                      //Unique poll identifier
	Question              string       `json:"question,omitempty"`                //Poll question, 1-255 characters
	Options               []PollOption `json:"options,omitempty"`                 //List of poll options
	TotalVoterCount       int          `json:"total_voter_count,omitempty"`       //Total number of users that voted in the poll
	IsClosed              bool         `json:"is_closed,omitempty"`               //True, if the poll is closed
	IsAnonymous           bool         `json:"is_anonymous,omitempty"`            //True, if the poll is anonymous
	Type                  string       `json:"type,omitempty"`                    //Poll type, currently can be “regular” or “quiz”
	AllowsMultipleAnswers bool         `json:"allows_multiple_answers,omitempty"` //True, if the poll allows multiple answers
	CorrectOptionId       int          `json:"correct_option_id"`                 //Optional. 0-based identifier of the correct answer option. Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
}

//This object represent a user's profile pictures.
type UserProfilePhotos struct {
	TotalCount int         `json:"total_count,omitempty"` //Total number of profile pictures the target user has
	Photos     []PhotoSize `json:"photos,omitempty"`      //Requested profile pictures (in up to 4 sizes each)
}

//This object represents a file ready to be downloaded. The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile.
type File struct {
	FileId       string `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size"`                //Optional. File size, if known
	FilePath     string `json:"file_path"`                //Optional. File path. Use https://api.telegram.org/file/bot<token>/<file_path> to get the file.
}

//This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	Keyboard        []KeyboardButton `json:"keyboard,omitempty"` //Array of button rows, each represented by an Array of KeyboardButton objects
	ResizeKeyboard  bool             `json:"resize_keyboard"`    //Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	OneTimeKeyboard bool             `json:"one_time_keyboard"`  //Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat – the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	Selective       bool             `json:"selective"`          //Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select the new language. Other users in the group don’t see the keyboard.
}

//This object represents one button of the reply keyboard. For simple text buttons String can be used instead of this object to specify text of the button. Optional fields request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	Text            string                  `json:"text,omitempty"`   //Text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed
	RequestContact  bool                    `json:"request_contact"`  //Optional. If True, the user's phone number will be sent as a contact when the button is pressed. Available in private chats only
	RequestLocation bool                    `json:"request_location"` //Optional. If True, the user's current location will be sent when the button is pressed. Available in private chats only
	RequestPoll     *KeyboardButtonPollType `json:"request_poll"`     //Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed. Available in private chats only
}

//This object represents type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	Type string `json:"type"` //Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed, only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
}

//Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot. An exception is made for one-time keyboards that are hidden immediately after the user presses a button (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"` //Requests clients to remove the custom keyboard (user will not be able to summon this keyboard; if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	Selective      bool `json:"selective"`                 //Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard for that user, while still showing the keyboard with poll options to users who haven't voted yet.
}

//This object represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	InlineKeyboard []InlineKeyboardButton `json:"inline_keyboard,omitempty"` //Array of button rows, each represented by an Array of InlineKeyboardButton objects
}

//This object represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	Text                         string    `json:"text,omitempty"`                   //Label text on the button
	Url                          string    `json:"url"`                              //Optional. HTTP or tg:// url to be opened when button is pressed
	LoginUrl                     *LoginUrl `json:"login_url"`                        //Optional. An HTTP URL used to automatically authorize the user. Can be used as a replacement for the Telegram Login Widget.
	CallbackData                 string    `json:"callback_data"`                    //Optional. Data to be sent in a callback query to the bot when button is pressed, 1-64 bytes
	SwitchInlineQuery            string    `json:"switch_inline_query"`              //Optional. If set, pressing the button will prompt the user to select one of their chats, open that chat and insert the bot‘s username and the specified inline query in the input field. Can be empty, in which case just the bot’s username will be inserted.Note: This offers an easy way for users to start using your bot in inline mode when they are currently in a private chat with it. Especially useful when combined with switch_pm… actions – in this case the user will be automatically returned to the chat they switched from, skipping the chat selection screen.
	SwitchInlineQueryCurrentChat string    `json:"switch_inline_query_current_chat"` //Optional. If set, pressing the button will insert the bot‘s username and the specified inline query in the current chat's input field. Can be empty, in which case only the bot’s username will be inserted.This offers a quick way for the user to open your bot in inline mode in the same chat – good for selecting something from multiple options.
	CallbackGame                 string    //dont know what is that `json:"callback_game"` //Optional. Description of the game that will be launched when the user presses the button.NOTE: This type of button must always be the first button in the first row.
	Pay                          bool      `json:"pay"` //Optional. Specify True, to send a Pay button.NOTE: This type of button must always be the first button in the first row.
}

//This object represents a parameter of the inline keyboard button used to automatically authorize a user. Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram. All the user needs to do is tap/click a button and confirm that they want to log in:
type LoginUrl struct {
	Url                string `json:"url,omitempty"`        //An HTTP URL to be opened with user authorization data added to the query string when the button is pressed. If the user refuses to provide authorization data, the original URL without information about the user will be opened. The data added is the same as described in Receiving authorization data.NOTE: You must always check the hash of the received data to verify the authentication and the integrity of the data as described in Checking authorization.
	ForwardText        string `json:"forward_text"`         //Optional. New text of the button in forwarded messages.
	BotUsername        string `json:"bot_username"`         //Optional. Username of a bot, which will be used for user authorization. See Setting up a bot for more details. If not specified, the current bot's username will be assumed. The url's domain must be the same as the domain linked with the bot. See Linking your domain to the bot for more details.
	RequestWriteAccess bool   `json:"request_write_access"` //Optional. Pass True to request the permission for your bot to send messages to the user.
}

//This object represents an incoming callback query from a callback button in an inline keyboard. If the button that originated the query was attached to a message sent by the bot, the field message will be present. If the button was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly one of the fields data or game_short_name will be present.
type CallbackQuery struct {
	Id              string   `json:"id,omitempty"`            //Unique identifier for this query
	From            *User    `json:"from,omitempty"`          //Sender
	Message         *Message `json:"message"`                 //Optional. Message with the callback button that originated the query. Note that message content and message date will not be available if the message is too old
	InlineMessageId string   `json:"inline_message_id"`       //Optional. Identifier of the message sent via the bot in inline mode, that originated the query.
	ChatInstance    string   `json:"chat_instance,omitempty"` //Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent. Useful for high scores in games.
	Data            string   `json:"data"`                    //Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary data in this field.
	GameShortName   string   `json:"game_short_name"`         //Optional. Short name of a Game to be returned, serves as the unique identifier for the game
}

//Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
type ForceReply struct {
	ForceReply bool `json:"force_reply,omitempty"` //Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	Selective  bool `json:"selective"`             //Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id), sender of the original message.
}

//This object represents a chat photo.
type ChatPhoto struct {
	SmallFileId       string `json:"small_file_id,omitempty"`        //File identifier of small (160x160) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileUniqueId string `json:"small_file_unique_id,omitempty"` //Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	BigFileId         string `json:"big_file_id,omitempty"`          //File identifier of big (640x640) chat photo. This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileUniqueId   string `json:"big_file_unique_id,omitempty"`   //Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
}



//Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`         //Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMediaMessages  bool `json:"can_send_media_messages"`   //Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages
	CanSendPolls          bool `json:"can_send_polls"`            //Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendOtherMessages  bool `json:"can_send_other_messages"`   //Optional. True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"` //Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanChangeInfo         bool `json:"can_change_info"`           //Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	CanInviteUsers        bool `json:"can_invite_users"`          //Optional. True, if the user is allowed to invite new users to the chat
	CanPinMessages        bool `json:"can_pin_messages"`          //Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
}

//Contains information about why a request was unsuccessful.
type ResponseParameters struct {
	MigrateToChatId int `json:"migrate_to_chat_id"` //Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	RetryAfter      int `json:"retry_after"`        //Optional. In case of exceeding flood control, the number of seconds left to wait before the request can be repeated
}

//Represents a photo to be sent.
type InputMediaPhoto struct {
	Type      string `json:"type,omitempty"`  //Type of the result, must be photo
	Media     string `json:"media,omitempty"` //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Caption   string `json:"caption"`         //Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode string `json:"parse_mode"`      //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
}

//Represents a video to be sent.
type InputMediaVideo struct {
	Type              string `json:"type,omitempty"`     //Type of the result, must be video
	Media             string `json:"media,omitempty"`    //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb             string `json:"thumb"`              //Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption           string `json:"caption"`            //Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode         string `json:"parse_mode"`         //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Width             int    `json:"width"`              //Optional. Video width
	Height            int    `json:"height"`             //Optional. Video height
	Duration          int    `json:"duration"`           //Optional. Video duration
	SupportsStreaming bool   `json:"supports_streaming"` //Optional. Pass True, if the uploaded video is suitable for streaming
}

//Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Type      string `json:"type,omitempty"`  //Type of the result, must be animation
	Media     string `json:"media,omitempty"` //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb     string `json:"thumb"`           //Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption   string `json:"caption"`         //Optional. Caption of the animation to be sent, 0-1024 characters after entities parsing
	ParseMode string `json:"parse_mode"`      //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Width     int    `json:"width"`           //Optional. Animation width
	Height    int    `json:"height"`          //Optional. Animation height
	Duration  int    `json:"duration"`        //Optional. Animation duration
}

//Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	Type      string `json:"type,omitempty"`  //Type of the result, must be audio
	Media     string `json:"media,omitempty"` //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb     string `json:"thumb"`           //Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption   string `json:"caption"`         //Optional. Caption of the audio to be sent, 0-1024 characters after entities parsing
	ParseMode string `json:"parse_mode"`      //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration  int    `json:"duration"`        //Optional. Duration of the audio in seconds
	Performer string `json:"performer"`       //Optional. Performer of the audio
	Title     string `json:"title"`           //Optional. Title of the audio
}

//Represents a general file to be sent.
type InputMediaDocument struct {
	Type      string `json:"type,omitempty"`  //Type of the result, must be document
	Media     string `json:"media,omitempty"` //File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name. More info on Sending Files »
	Thumb     string `json:"thumb"`           //Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption   string `json:"caption"`         //Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode string `json:"parse_mode"`      //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
}

//Use this method to send text messages. On success, the sent Message is returned.
type SendMessage struct {
	ChatId                uint64               `json:"chat_id,omitempty"`        //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Text                  string               `json:"text,omitempty"`           //Text of the message to be sent, 1-4096 characters after entities parsing
	ParseMode             string               `json:"parse_mode"`               //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"` //Disables link previews for links in this message
	DisableNotification   bool                 `json:"disable_notification"`     //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId      int                  `json:"reply_to_message_id"`      //If the message is a reply, ID of the original message
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`             //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to forward messages of any kind. On success, the sent Message is returned.
type ForwardMessage struct {
	ChatId              string `json:"chat_id,omitempty"`      //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	FromChatId          string `json:"from_chat_id,omitempty"` //Unique identifier for the chat where the original message was sent (or channel username in the format @channelusername)
	DisableNotification bool   `json:"disable_notification"`   //Sends the message silently. Users will receive a notification with no sound.
	MessageId           int    `json:"message_id,omitempty"`   //Message identifier in the chat specified in from_chat_id
}

//Use this method to send photos. On success, the sent Message is returned.
type SendPhoto struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo               string               `json:"photo,omitempty"`      //Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using multipart/form-data. More info on Sending Files »
	Caption             string               `json:"caption"`              //Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send audio files, if you want Telegram clients to display them in the music player. Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently send audio files of up to 50 MB in size, this limit may be changed in the future.
type SendAudio struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Audio               string               `json:"audio,omitempty"`      //Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption             string               `json:"caption"`              //Audio caption, 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            int                  `json:"duration"`             //Duration of the audio in seconds
	Performer           string               `json:"performer"`            //Performer
	Title               string               `json:"title"`                //Track name
	Thumb               string               `json:"thumb"`                //Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send general files. On success, the sent Message is returned. Bots can currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
type SendDocument struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Document            string               `json:"document,omitempty"`   //File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Thumb               string               `json:"thumb"`                //Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string               `json:"caption"`              //Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document). On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size, this limit may be changed in the future.
type SendVideo struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Video               string               `json:"video,omitempty"`      //Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using multipart/form-data. More info on Sending Files »
	Duration            int                  `json:"duration"`             //Duration of sent video in seconds
	Width               int                  `json:"width"`                //Video width
	Height              int                  `json:"height"`               //Video height
	Thumb               string               `json:"thumb"`                //Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string               `json:"caption"`              //Video caption (may also be used when resending videos by file_id), 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	SupportsStreaming   bool                 `json:"supports_streaming"`   //Pass True, if the uploaded video is suitable for streaming
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed in the future.
type SendAnimation struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Animation           string               `json:"animation,omitempty"`  //Animation to send. Pass a file_id as String to send an animation that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet, or upload a new animation using multipart/form-data. More info on Sending Files »
	Duration            int                  `json:"duration"`             //Duration of sent animation in seconds
	Width               int                  `json:"width"`                //Animation width
	Height              int                  `json:"height"`               //Animation height
	Thumb               string               `json:"thumb"`                //Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	Caption             string               `json:"caption"`              //Animation caption (may also be used when resending animation by file_id), 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB in size, this limit may be changed in the future.
type SendVoice struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Voice               string               `json:"voice,omitempty"`      //Audio file to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Caption             string               `json:"caption"`              //Voice message caption, 0-1024 characters after entities parsing
	ParseMode           string               `json:"parse_mode"`           //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Duration            int                  `json:"duration"`             //Duration of the voice message in seconds
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this method to send video messages. On success, the sent Message is returned.
type SendVideoNote struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	VideoNote           string               `json:"video_note,omitempty"` //Video note to send. Pass a file_id as String to send a video note that exists on the Telegram servers (recommended) or upload a new video using multipart/form-data. More info on Sending Files ». Sending video notes by a URL is currently unsupported
	Duration            int                  `json:"duration"`             //Duration of sent video in seconds
	Length              int                  `json:"length"`               //Video width and height, i.e. diameter of the video message
	Thumb               string               `json:"thumb"`                //Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>. More info on Sending Files »
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send a group of photos or videos as an album. On success, an array of the sent Messages is returned.
type SendMediaGroup struct {
	ChatId              string            `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Media               []InputMediaVideo `json:"media,omitempty"`      //A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	DisableNotification bool              `json:"disable_notification"` //Sends the messages silently. Users will receive a notification with no sound.
	ReplyToMessageId    int               `json:"reply_to_message_id"`  //If the messages are a reply, ID of the original message
}

//Use this method to send point on the map. On success, the sent Message is returned.
type SendLocation struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float32              `json:"latitude,omitempty"`   //Latitude of the location
	Longitude           float32              `json:"longitude,omitempty"`  //Longitude of the location
	LivePeriod          int                  `json:"live_period"`          //Period in seconds for which the location will be updated (see Live Locations, should be between 60 and 86400.
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to edit live location messages. A location can be edited until its live_period expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageLiveLocation struct {
	ChatId          string               `json:"chat_id"`             //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                  `json:"message_id"`          //Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string               `json:"inline_message_id"`   //Required if chat_id and message_id are not specified. Identifier of the inline message
	Latitude        float32              `json:"latitude,omitempty"`  //Latitude of new location
	Longitude       float32              `json:"longitude,omitempty"` //Longitude of new location
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`        //A JSON-serialized object for a new inline keyboard.
}

//Use this method to stop updating a live location message before live_period expires. On success, if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
type StopMessageLiveLocation struct {
	ChatId          string               `json:"chat_id"`           //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                  `json:"message_id"`        //Required if inline_message_id is not specified. Identifier of the message with live location to stop
	InlineMessageId string               `json:"inline_message_id"` //Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`      //A JSON-serialized object for a new inline keyboard.
}

//Use this method to send information about a venue. On success, the sent Message is returned.
type SendVenue struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Latitude            float32              `json:"latitude,omitempty"`   //Latitude of the venue
	Longitude           float32              `json:"longitude,omitempty"`  //Longitude of the venue
	Title               string               `json:"title,omitempty"`      //Name of the venue
	Address             string               `json:"address,omitempty"`    //Address of the venue
	FoursquareId        string               `json:"foursquare_id"`        //Foursquare identifier of the venue
	FoursquareType      string               `json:"foursquare_type"`      //Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to send phone contacts. On success, the sent Message is returned.
type SendContact struct {
	ChatId              string               `json:"chat_id,omitempty"`      //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	PhoneNumber         string               `json:"phone_number,omitempty"` //Contact's phone number
	FirstName           string               `json:"first_name,omitempty"`   //Contact's first name
	LastName            string               `json:"last_name"`              //Contact's last name
	Vcard               string               `json:"vcard"`                  //Additional data about the contact in the form of a vCard, 0-2048 bytes
	DisableNotification bool                 `json:"disable_notification"`   //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`    //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`           //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove keyboard or to force a reply from the user.
}

//Use this method to send a native poll. On success, the sent Message is returned.
type SendPoll struct {
	ChatId                uint64               `json:"chat_id,omitempty"`       //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Question              string               `json:"question,omitempty"`      //Poll question, 1-255 characters
	Options               []string             `json:"options,omitempty"`       //A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	IsAnonymous           bool                 `json:"is_anonymous"`            //True, if the poll needs to be anonymous, defaults to True
	Type                  string               `json:"type"`                    //Poll type, “quiz” or “regular”, defaults to “regular”
	AllowsMultipleAnswers bool                 `json:"allows_multiple_answers"` //True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	CorrectOptionId       int                  `json:"correct_option_id"`       //0-based identifier of the correct answer option, required for polls in quiz mode
	IsClosed              bool                 `json:"is_closed"`               //Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	DisableNotification   bool                 `json:"disable_notification"`    //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId      int                  `json:"reply_to_message_id"`     //If the message is a reply, ID of the original message
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`            //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method when you need to tell the user that something is happening on the bot's side. The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its typing status). Returns True on success.
type SendChatAction struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Action string `json:"action,omitempty"`  //Type of action to broadcast. Choose one, depending on what the user is about to receive: typing for text messages, upload_photo for photos, record_video or upload_video for videos, record_audio or upload_audio for audio files, upload_document for general files, find_location for location data, record_video_note or upload_video_note for video notes.
}

//Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
type GetUserProfilePhotos struct {
	UserId int `json:"user_id,omitempty"` //Unique identifier of the target user
	Offset int `json:"offset"`            //Sequential number of the first photo to be returned. By default, all photos are returned.
	Limit  int `json:"limit"`             //Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
}

//Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be requested by calling getFile again.
type GetFile struct {
	FileId string `json:"file_id,omitempty"` //File identifier to get info about
}

//Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels, the user will not be able to return to the group on their own using invite links, etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type KickChatMember struct {
	ChatId    string `json:"chat_id,omitempty"` //Unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername)
	UserId    int    `json:"user_id,omitempty"` //Unique identifier of the target user
	UntilDate int    `json:"until_date"`        //Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less than 30 seconds from the current time they are considered to be banned forever
}

//Use this method to unban a previously kicked user in a supergroup or channel. The user will not return to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator for this to work. Returns True on success.
type UnbanChatMember struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target group or username of the target supergroup or channel (in the format @username)
	UserId int    `json:"user_id,omitempty"` //Unique identifier of the target user
}

//Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user. Returns True on success.
type RestrictChatMember struct {
	ChatId      string           `json:"chat_id,omitempty"`     //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int              `json:"user_id,omitempty"`     //Unique identifier of the target user
	Permissions *ChatPermissions `json:"permissions,omitempty"` //New user permissions
	UntilDate   int              `json:"until_date"`            //Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever
}

//Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters to demote a user. Returns True on success.
type PromoteChatMember struct {
	ChatId             string `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	UserId             int    `json:"user_id,omitempty"`    //Unique identifier of the target user
	CanChangeInfo      bool   `json:"can_change_info"`      //Pass True, if the administrator can change chat title, photo and other settings
	CanPostMessages    bool   `json:"can_post_messages"`    //Pass True, if the administrator can create channel posts, channels only
	CanEditMessages    bool   `json:"can_edit_messages"`    //Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanDeleteMessages  bool   `json:"can_delete_messages"`  //Pass True, if the administrator can delete messages of other users
	CanInviteUsers     bool   `json:"can_invite_users"`     //Pass True, if the administrator can invite new users to the chat
	CanRestrictMembers bool   `json:"can_restrict_members"` //Pass True, if the administrator can restrict, ban or unban chat members
	CanPinMessages     bool   `json:"can_pin_messages"`     //Pass True, if the administrator can pin messages, supergroups only
	CanPromoteMembers  bool   `json:"can_promote_members"`  //Pass True, if the administrator can add new administrators with a subset of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by him)
}

//Use this method to set a custom title for an administrator in a supergroup promoted by the bot. Returns True on success.
type SetChatAdministratorCustomTitle struct {
	ChatId      string `json:"chat_id,omitempty"`      //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	UserId      int    `json:"user_id,omitempty"`      //Unique identifier of the target user
	CustomTitle string `json:"custom_title,omitempty"` //New custom title for the administrator; 0-16 characters, emoji are not allowed
}

//Use this method to set default chat permissions for all members. The bot must be an administrator in the group or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success.
type SetChatPermissions struct {
	ChatId      string           `json:"chat_id,omitempty"`     //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	Permissions *ChatPermissions `json:"permissions,omitempty"` //New default chat permissions
}

//Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns the new invite link as String on success.
type ExportChatInviteLink struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

//Use this method to set a new profile photo for the chat. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatPhoto struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Photo  string `json:"photo,omitempty"`   //New chat photo, uploaded using multipart/form-data
}

//Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type DeleteChatPhoto struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

//Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatTitle struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Title  string `json:"title,omitempty"`   //New chat title, 1-255 characters
}

//Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Returns True on success.
type SetChatDescription struct {
	ChatId      string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Description string `json:"description"`       //New chat description, 0-255 characters
}

//Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type PinChatMessage struct {
	ChatId              string `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId           int    `json:"message_id,omitempty"` //Identifier of a message to pin
	DisableNotification bool   `json:"disable_notification"` //Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
}

//Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
type UnpinChatMessage struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
}

//Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
type LeaveChat struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

//Use this method to get up to date information about the chat (current name of the user for one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on success.
type GetChat struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

//Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots. If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministrators struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

//Use this method to get the number of members in a chat. Returns Int on success.
type GetChatMembersCount struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
}

//Use this method to get information about a member of a chat. Returns a ChatMember object on success.
type GetChatMember struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup or channel (in the format @channelusername)
	UserId int    `json:"user_id,omitempty"` //Unique identifier of the target user
}

//Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type SetChatStickerSet struct {
	ChatId         string `json:"chat_id,omitempty"`          //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	StickerSetName string `json:"sticker_set_name,omitempty"` //Name of the sticker set to be set as the group sticker set
}

//Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True on success.
type DeleteChatStickerSet struct {
	ChatId string `json:"chat_id,omitempty"` //Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
}

//Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
type AnswerCallbackQuery struct {
	CallbackQueryId string `json:"callback_query_id,omitempty"` //Unique identifier for the query to be answered
	Text            string `json:"text"`                        //Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	ShowAlert       bool   `json:"show_alert"`                  //If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	Url             string `json:"url"`                         //URL that will be opened by the user's client. If you have created a Game and accepted the conditions via @Botfather, specify the URL that opens your game – note that this will only work if the query comes from a callback_game button.Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	CacheTime       int    `json:"cache_time"`                  //The maximum amount of time in seconds that the result of the callback query may be cached client-side. Telegram apps will support caching starting in version 3.14. Defaults to 0.
}

//Use this method to edit text and game messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageText struct {
	ChatId                string               `json:"chat_id"`                  //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId             int                  `json:"message_id"`               //Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId       string               `json:"inline_message_id"`        //Required if chat_id and message_id are not specified. Identifier of the inline message
	Text                  string               `json:"text,omitempty"`           //New text of the message, 1-4096 characters after entities parsing
	ParseMode             string               `json:"parse_mode"`               //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"` //Disables link previews for links in this message
	ReplyMarkup           InlineKeyboardMarkup `json:"reply_markup"`             //A JSON-serialized object for an inline keyboard.
}

//Use this method to edit captions of messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageCaption struct {
	ChatId          string               `json:"chat_id"`           //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                  `json:"message_id"`        //Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string               `json:"inline_message_id"` //Required if chat_id and message_id are not specified. Identifier of the inline message
	Caption         string               `json:"caption"`           //New caption of the message, 0-1024 characters after entities parsing
	ParseMode       string               `json:"parse_mode"`        //Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`      //A JSON-serialized object for an inline keyboard.
}

//Use this method to edit animation, audio, document, photo, or video messages. If a message is a part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageMedia struct {
	ChatId              string               `json:"chat_id"`           //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId           int                  `json:"message_id"`        //Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId     string               `json:"inline_message_id"` //Required if chat_id and message_id are not specified. Identifier of the inline message
	Media               InputMediaDocument   `json:"media,omitempty"`   //A JSON-serialized object for a new media content of the message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`      //A JSON-serialized object for a new inline keyboard.
	InputMediaAnimation *InputMediaAnimation `json:"media"`             //Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
	InputMediaAudio     *InputMediaAudio     `json:"media"`             //Represents an audio file to be treated as music to be sent.
	InputMediaPhoto     *InputMediaPhoto     `json:"media"`             //Represents a photo to be sent.
	InputMediaVideo     *InputMediaVideo     `json:"media"`             //Represents a video to be sent.
}

//Use this method to edit only the reply markup of messages. On success, if edited message is sent by the bot, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkup struct {
	ChatId          string               `json:"chat_id"`           //Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId       int                  `json:"message_id"`        //Required if inline_message_id is not specified. Identifier of the message to edit
	InlineMessageId string               `json:"inline_message_id"` //Required if chat_id and message_id are not specified. Identifier of the inline message
	ReplyMarkup     InlineKeyboardMarkup `json:"reply_markup"`      //A JSON-serialized object for an inline keyboard.
}

//Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the final results is returned.
type StopPoll struct {
	ChatId      string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId   int                  `json:"message_id,omitempty"` //Identifier of the original message with the poll
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`         //A JSON-serialized object for a new message inline keyboard.
}

//Use this method to delete a message, including service messages, with the following limitations:- A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.- Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is an administrator of a group, it can delete any message there.- If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.Returns True on success.
type DeleteMessage struct {
	ChatId    string `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	MessageId int    `json:"message_id,omitempty"` //Identifier of the message to delete
}

//This object represents a sticker.
type Sticker struct {
	FileId       string        `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string        `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	Width        int           `json:"width,omitempty"`          //Sticker width
	Height       int           `json:"height,omitempty"`         //Sticker height
	IsAnimated   bool          `json:"is_animated,omitempty"`    //True, if the sticker is animated
	Thumb        *PhotoSize    `json:"thumb"`                    //Optional. Sticker thumbnail in the .webp or .jpg format
	Emoji        string        `json:"emoji"`                    //Optional. Emoji associated with the sticker
	SetName      string        `json:"set_name"`                 //Optional. Name of the sticker set to which the sticker belongs
	MaskPosition *MaskPosition `json:"mask_position"`            //Optional. For mask stickers, the position where the mask should be placed
	FileSize     int           `json:"file_size"`                //Optional. File size
}

//This object represents a sticker set.
type StickerSet struct {
	Name          string    `json:"name,omitempty"`           //Sticker set name
	Title         string    `json:"title,omitempty"`          //Sticker set title
	IsAnimated    bool      `json:"is_animated,omitempty"`    //True, if the sticker set contains animated stickers
	ContainsMasks bool      `json:"contains_masks,omitempty"` //True, if the sticker set contains masks
	Stickers      []Sticker `json:"stickers,omitempty"`       //List of all set stickers
}

//This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	Point  string  `json:"point,omitempty"`   //The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	XShift float32 `json:"x_shift,omitempty"` //Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
	YShift float32 `json:"y_shift,omitempty"` //Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
	Scale  float32 `json:"scale,omitempty"`   //Mask scaling coefficient. For example, 2.0 means double size.
}

//Use this method to send static .WEBP or animated .TGS stickers. On success, the sent Message is returned.
type SendSticker struct {
	ChatId              string               `json:"chat_id,omitempty"`    //Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	Sticker             string               `json:"sticker,omitempty"`    //Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .webp file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	DisableNotification bool                 `json:"disable_notification"` //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`  //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`         //Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove reply keyboard or to force a reply from the user.
}

//Use this method to get a sticker set. On success, a StickerSet object is returned.
type GetStickerSet struct {
	Name string `json:"name,omitempty"` //Name of the sticker set
}

//Use this method to upload a .png file with a sticker for later use in createNewStickerSet and addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
type UploadStickerFile struct {
	UserId     int    `json:"user_id,omitempty"`     //User identifier of sticker file owner
	PngSticker string `json:"png_sticker,omitempty"` //Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. More info on Sending Files »
}

//Use this method to create new sticker set owned by a user. The bot will be able to edit the created sticker set. Returns True on success.
type CreateNewStickerSet struct {
	UserId        int           `json:"user_id,omitempty"`     //User identifier of created sticker set owner
	Name          string        `json:"name,omitempty"`        //Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Title         string        `json:"title,omitempty"`       //Sticker set title, 1-64 characters
	PngSticker    string        `json:"png_sticker,omitempty"` //Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Emojis        string        `json:"emojis,omitempty"`      //One or more emoji corresponding to the sticker
	ContainsMasks bool          `json:"contains_masks"`        //Pass True, if a set of mask stickers should be created
	MaskPosition  *MaskPosition `json:"mask_position"`         //A JSON-serialized object for position where the mask should be placed on faces
}

//Use this method to add a new sticker to a set created by the bot. Returns True on success.
type AddStickerToSet struct {
	UserId       int           `json:"user_id,omitempty"`     //User identifier of sticker set owner
	Name         string        `json:"name,omitempty"`        //Sticker set name
	PngSticker   string        `json:"png_sticker,omitempty"` //Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px, and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More info on Sending Files »
	Emojis       string        `json:"emojis,omitempty"`      //One or more emoji corresponding to the sticker
	MaskPosition *MaskPosition `json:"mask_position"`         //A JSON-serialized object for position where the mask should be placed on faces
}

//Use this method to move a sticker in a set created by the bot to a specific position . Returns True on success.
type SetStickerPositionInSet struct {
	Sticker  string `json:"sticker,omitempty"`  //File identifier of the sticker
	Position int    `json:"position,omitempty"` //New sticker position in the set, zero-based
}

//Use this method to delete a sticker from a set created by the bot. Returns True on success.
type DeleteStickerFromSet struct {
	Sticker string `json:"sticker,omitempty"` //File identifier of the sticker
}

//This object represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	Id       string    `json:"id,omitempty"`     //Unique identifier for this query
	From     *User     `json:"from,omitempty"`   //Sender
	Location *Location `json:"location"`         //Optional. Sender location, only for bots that request user location
	Query    string    `json:"query,omitempty"`  //Text of the query (up to 256 characters)
	Offset   string    `json:"offset,omitempty"` //Offset of the results to be returned, can be controlled by the bot
}

//Use this method to send answers to an inline query. On success, True is returned.No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
	InlineQueryId     string                         `json:"inline_query_id,omitempty"` //Unique identifier for the answered query
	Results           []InlineQueryResultCachedAudio `json:"results,omitempty"`         //A JSON-serialized array of results for the inline query
	CacheTime         int                            `json:"cache_time"`                //The maximum amount of time in seconds that the result of the inline query may be cached on the server. Defaults to 300.
	IsPersonal        bool                           `json:"is_personal"`               //Pass True, if results may be cached on the server side only for the user that sent the query. By default, results may be returned to any user who sends the same query
	NextOffset        string                         `json:"next_offset"`               //Pass the offset that a client should send in the next query with the same text to receive more results. Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	SwitchPmText      string                         `json:"switch_pm_text"`            //If passed, clients will display a button with specified text that switches the user to a private chat with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmParameter string                         `json:"switch_pm_parameter"`       //Deep-linking parameter for the /start message sent to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube account to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button above the results, or even before showing any. The user presses the button, switches to a private chat with the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done, the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted to use the bot's inline capabilities.
}

//Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	Type                        string                       `json:"type,omitempty"`                  //Type of the result, must be article
	Id                          string                       `json:"id,omitempty"`                    //Unique identifier for this result, 1-64 Bytes
	Title                       string                       `json:"title,omitempty"`                 //Title of the result
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content,omitempty"` //Content of the message to be sent
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`                    //Optional. Inline keyboard attached to the message
	Url                         string                       `json:"url"`                             //Optional. URL of the result
	HideUrl                     bool                         `json:"hide_url"`                        //Optional. Pass True, if you don't want the URL to be shown in the message
	Description                 string                       `json:"description"`                     //Optional. Short description of the result
	ThumbUrl                    string                       `json:"thumb_url"`                       //Optional. Url of the thumbnail for the result
	ThumbWidth                  int                          `json:"thumb_width"`                     //Optional. Thumbnail width
	ThumbHeight                 int                          `json:"thumb_height"`                    //Optional. Thumbnail height
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`           //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`           //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`           //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be photo
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	PhotoUrl                    string                       `json:"photo_url,omitempty"`   //A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	ThumbUrl                    string                       `json:"thumb_url,omitempty"`   //URL of the thumbnail for the photo
	PhotoWidth                  int                          `json:"photo_width"`           //Optional. Width of the photo
	PhotoHeight                 int                          `json:"photo_height"`          //Optional. Height of the photo
	Title                       string                       `json:"title"`                 //Optional. Title for the result
	Description                 string                       `json:"description"`           //Optional. Short description of the result
	Caption                     string                       `json:"caption"`               //Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the photo
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be gif
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	GifUrl                      string                       `json:"gif_url,omitempty"`     //A valid URL for the GIF file. File size must not exceed 1MB
	GifWidth                    int                          `json:"gif_width"`             //Optional. Width of the GIF
	GifHeight                   int                          `json:"gif_height"`            //Optional. Height of the GIF
	GifDuration                 int                          `json:"gif_duration"`          //Optional. Duration of the GIF
	ThumbUrl                    string                       `json:"thumb_url,omitempty"`   //URL of the static thumbnail for the result (jpeg or gif)
	Title                       string                       `json:"title"`                 //Optional. Title for the result
	Caption                     string                       `json:"caption"`               //Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the GIF animation
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be mpeg4_gif
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	Mpeg4Url                    string                       `json:"mpeg4_url,omitempty"`   //A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4Width                  int                          `json:"mpeg4_width"`           //Optional. Video width
	Mpeg4Height                 int                          `json:"mpeg4_height"`          //Optional. Video height
	Mpeg4Duration               int                          `json:"mpeg4_duration"`        //Optional. Video duration
	ThumbUrl                    string                       `json:"thumb_url,omitempty"`   //URL of the static thumbnail (jpeg or gif) for the result
	Title                       string                       `json:"title"`                 //Optional. Title for the result
	Caption                     string                       `json:"caption"`               //Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the video animation
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultVideo struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be video
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	VideoUrl                    string                       `json:"video_url,omitempty"`   //A valid URL for the embedded video player or video file
	MimeType                    string                       `json:"mime_type,omitempty"`   //Mime type of the content of video url, “text/html” or “video/mp4”
	ThumbUrl                    string                       `json:"thumb_url,omitempty"`   //URL of the thumbnail (jpeg only) for the video
	Title                       string                       `json:"title,omitempty"`       //Title for the result
	Caption                     string                       `json:"caption"`               //Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	VideoWidth                  int                          `json:"video_width"`           //Optional. Video width
	VideoHeight                 int                          `json:"video_height"`          //Optional. Video height
	VideoDuration               int                          `json:"video_duration"`        //Optional. Video duration in seconds
	Description                 string                       `json:"description"`           //Optional. Short description of the result
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the video. This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultAudio struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be audio
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	AudioUrl                    string                       `json:"audio_url,omitempty"`   //A valid URL for the audio file
	Title                       string                       `json:"title,omitempty"`       //Title
	Caption                     string                       `json:"caption"`               //Optional. Caption, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	Performer                   string                       `json:"performer"`             //Optional. Performer
	AudioDuration               int                          `json:"audio_duration"`        //Optional. Audio duration in seconds
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the audio
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a voice recording in an .ogg container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
type InlineQueryResultVoice struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be voice
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	VoiceUrl                    string                       `json:"voice_url,omitempty"`   //A valid URL for the voice recording
	Title                       string                       `json:"title,omitempty"`       //Recording title
	Caption                     string                       `json:"caption"`               //Optional. Caption, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	VoiceDuration               int                          `json:"voice_duration"`        //Optional. Recording duration in seconds
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the voice recording
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only .PDF and .ZIP files can be sent using this method.
type InlineQueryResultDocument struct {
	Type                        string                       `json:"type,omitempty"`         //Type of the result, must be document
	Id                          string                       `json:"id,omitempty"`           //Unique identifier for this result, 1-64 bytes
	Title                       string                       `json:"title,omitempty"`        //Title for the result
	Caption                     string                       `json:"caption"`                //Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`             //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	DocumentUrl                 string                       `json:"document_url,omitempty"` //A valid URL for the file
	MimeType                    string                       `json:"mime_type,omitempty"`    //Mime type of the content of the file, either “application/pdf” or “application/zip”
	Description                 string                       `json:"description"`            //Optional. Short description of the result
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`           //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`  //Optional. Content of the message to be sent instead of the file
	ThumbUrl                    string                       `json:"thumb_url"`              //Optional. URL of the thumbnail (jpeg only) for the file
	ThumbWidth                  int                          `json:"thumb_width"`            //Optional. Thumbnail width
	ThumbHeight                 int                          `json:"thumb_height"`           //Optional. Thumbnail height
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be location
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 Bytes
	Latitude                    float32                      `json:"latitude,omitempty"`    //Location latitude in degrees
	Longitude                   float32                      `json:"longitude,omitempty"`   //Location longitude in degrees
	Title                       string                       `json:"title,omitempty"`       //Location title
	LivePeriod                  int                          `json:"live_period"`           //Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the location
	ThumbUrl                    string                       `json:"thumb_url"`             //Optional. Url of the thumbnail for the result
	ThumbWidth                  int                          `json:"thumb_width"`           //Optional. Thumbnail width
	ThumbHeight                 int                          `json:"thumb_height"`          //Optional. Thumbnail height
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
type InlineQueryResultVenue struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be venue
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 Bytes
	Latitude                    float32                      `json:"latitude,omitempty"`    //Latitude of the venue location in degrees
	Longitude                   float32                      `json:"longitude,omitempty"`   //Longitude of the venue location in degrees
	Title                       string                       `json:"title,omitempty"`       //Title of the venue
	Address                     string                       `json:"address,omitempty"`     //Address of the venue
	FoursquareId                string                       `json:"foursquare_id"`         //Optional. Foursquare identifier of the venue if known
	FoursquareType              string                       `json:"foursquare_type"`       //Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the venue
	ThumbUrl                    string                       `json:"thumb_url"`             //Optional. Url of the thumbnail for the result
	ThumbWidth                  int                          `json:"thumb_width"`           //Optional. Thumbnail width
	ThumbHeight                 int                          `json:"thumb_height"`          //Optional. Thumbnail height
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
type InlineQueryResultContact struct {
	Type                        string                       `json:"type,omitempty"`         //Type of the result, must be contact
	Id                          string                       `json:"id,omitempty"`           //Unique identifier for this result, 1-64 Bytes
	PhoneNumber                 string                       `json:"phone_number,omitempty"` //Contact's phone number
	FirstName                   string                       `json:"first_name,omitempty"`   //Contact's first name
	LastName                    string                       `json:"last_name"`              //Optional. Contact's last name
	Vcard                       string                       `json:"vcard"`                  //Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`           //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`  //Optional. Content of the message to be sent instead of the contact
	ThumbUrl                    string                       `json:"thumb_url"`              //Optional. Url of the thumbnail for the result
	ThumbWidth                  int                          `json:"thumb_width"`            //Optional. Thumbnail width
	ThumbHeight                 int                          `json:"thumb_height"`           //Optional. Thumbnail height
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`  //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a Game.
type InlineQueryResultGame struct {
	Type          string               `json:"type,omitempty"`            //Type of the result, must be game
	Id            string               `json:"id,omitempty"`              //Unique identifier for this result, 1-64 bytes
	GameShortName string               `json:"game_short_name,omitempty"` //Short name of the game
	ReplyMarkup   InlineKeyboardMarkup `json:"reply_markup"`              //Optional. Inline keyboard attached to the message
}

//Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	Type                        string                       `json:"type,omitempty"`          //Type of the result, must be photo
	Id                          string                       `json:"id,omitempty"`            //Unique identifier for this result, 1-64 bytes
	PhotoFileId                 string                       `json:"photo_file_id,omitempty"` //A valid file identifier of the photo
	Title                       string                       `json:"title"`                   //Optional. Title for the result
	Description                 string                       `json:"description"`             //Optional. Short description of the result
	Caption                     string                       `json:"caption"`                 //Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`              //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`            //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`   //Optional. Content of the message to be sent instead of the photo
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	Type                        string                       `json:"type,omitempty"`        //Type of the result, must be gif
	Id                          string                       `json:"id,omitempty"`          //Unique identifier for this result, 1-64 bytes
	GifFileId                   string                       `json:"gif_file_id,omitempty"` //A valid file identifier for the GIF file
	Title                       string                       `json:"title"`                 //Optional. Title for the result
	Caption                     string                       `json:"caption"`               //Optional. Caption of the GIF file to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`            //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`          //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"` //Optional. Content of the message to be sent instead of the GIF animation
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"` //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	Type                        string                       `json:"type,omitempty"`          //Type of the result, must be mpeg4_gif
	Id                          string                       `json:"id,omitempty"`            //Unique identifier for this result, 1-64 bytes
	Mpeg4FileId                 string                       `json:"mpeg4_file_id,omitempty"` //A valid file identifier for the MP4 file
	Title                       string                       `json:"title"`                   //Optional. Title for the result
	Caption                     string                       `json:"caption"`                 //Optional. Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`              //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`            //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`   //Optional. Content of the message to be sent instead of the video animation
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultCachedSticker struct {
	Type                        string                       `json:"type,omitempty"`            //Type of the result, must be sticker
	Id                          string                       `json:"id,omitempty"`              //Unique identifier for this result, 1-64 bytes
	StickerFileId               string                       `json:"sticker_file_id,omitempty"` //A valid file identifier of the sticker
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`              //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`     //Optional. Content of the message to be sent instead of the sticker
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`     //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`     //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`     //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
type InlineQueryResultCachedDocument struct {
	Type                        string                       `json:"type,omitempty"`             //Type of the result, must be document
	Id                          string                       `json:"id,omitempty"`               //Unique identifier for this result, 1-64 bytes
	Title                       string                       `json:"title,omitempty"`            //Title for the result
	DocumentFileId              string                       `json:"document_file_id,omitempty"` //A valid file identifier for the file
	Description                 string                       `json:"description"`                //Optional. Short description of the result
	Caption                     string                       `json:"caption"`                    //Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`                 //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`               //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`      //Optional. Content of the message to be sent instead of the file
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`      //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`      //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`      //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	Type                        string                       `json:"type,omitempty"`          //Type of the result, must be video
	Id                          string                       `json:"id,omitempty"`            //Unique identifier for this result, 1-64 bytes
	VideoFileId                 string                       `json:"video_file_id,omitempty"` //A valid file identifier for the video file
	Title                       string                       `json:"title,omitempty"`         //Title for the result
	Description                 string                       `json:"description"`             //Optional. Short description of the result
	Caption                     string                       `json:"caption"`                 //Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`              //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`            //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`   //Optional. Content of the message to be sent instead of the video
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
type InlineQueryResultCachedVoice struct {
	Type                        string                       `json:"type,omitempty"`          //Type of the result, must be voice
	Id                          string                       `json:"id,omitempty"`            //Unique identifier for this result, 1-64 bytes
	VoiceFileId                 string                       `json:"voice_file_id,omitempty"` //A valid file identifier for the voice message
	Title                       string                       `json:"title,omitempty"`         //Voice message title
	Caption                     string                       `json:"caption"`                 //Optional. Caption, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`              //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`            //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`   //Optional. Content of the message to be sent instead of the voice message
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
}

//Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
type InlineQueryResultCachedAudio struct {
	Type                        string                       `json:"type,omitempty"`          //Type of the result, must be audio
	Id                          string                       `json:"id,omitempty"`            //Unique identifier for this result, 1-64 bytes
	AudioFileId                 string                       `json:"audio_file_id,omitempty"` //A valid file identifier for the audio file
	Caption                     string                       `json:"caption"`                 //Optional. Caption, 0-1024 characters after entities parsing
	ParseMode                   string                       `json:"parse_mode"`              //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in the media caption.
	ReplyMarkup                 InlineKeyboardMarkup         `json:"reply_markup"`            //Optional. Inline keyboard attached to the message
	InputMessageContent         *InputContactMessageContent  `json:"input_message_content"`   //Optional. Content of the message to be sent instead of the audio
	InputTextMessageContent     *InputTextMessageContent     `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputLocationMessageContent *InputLocationMessageContent `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
	InputVenueMessageContent    *InputVenueMessageContent    `json:"input_message_content"`   //Represents the content of a text message to be sent as the result of an inline query
}

//Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	MessageText           string `json:"message_text,omitempty"`   //Text of the message to be sent, 1-4096 characters
	ParseMode             string `json:"parse_mode"`               //Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs in your bot's message.
	DisableWebPagePreview bool   `json:"disable_web_page_preview"` //Optional. Disables link previews for links in the sent message
}

//Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Latitude   float32 `json:"latitude,omitempty"`  //Latitude of the location in degrees
	Longitude  float32 `json:"longitude,omitempty"` //Longitude of the location in degrees
	LivePeriod int     `json:"live_period"`         //Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
}

//Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	Latitude       float32 `json:"latitude,omitempty"`  //Latitude of the venue in degrees
	Longitude      float32 `json:"longitude,omitempty"` //Longitude of the venue in degrees
	Title          string  `json:"title,omitempty"`     //Name of the venue
	Address        string  `json:"address,omitempty"`   //Address of the venue
	FoursquareId   string  `json:"foursquare_id"`       //Optional. Foursquare identifier of the venue, if known
	FoursquareType string  `json:"foursquare_type"`     //Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”.)
}

//Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number,omitempty"` //Contact's phone number
	FirstName   string `json:"first_name,omitempty"`   //Contact's first name
	LastName    string `json:"last_name"`              //Optional. Contact's last name
	Vcard       string `json:"vcard"`                  //Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
}

//Represents a result of an inline query that was chosen by the user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultId        string    `json:"result_id,omitempty"` //The unique identifier for the result that was chosen
	From            *User     `json:"from,omitempty"`      //The user that chose the result
	Location        *Location `json:"location"`            //Optional. Sender location, only for bots that require user location
	InlineMessageId string    `json:"inline_message_id"`   //Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached to the message. Will be also received in callback queries and can be used to edit the message.
	Query           string    `json:"query,omitempty"`     //The query that was used to obtain the result
}

//Use this method to send invoices. On success, the sent Message is returned.
type SendInvoice struct {
	ChatId                    int                  `json:"chat_id,omitempty"`             //Unique identifier for the target private chat
	Title                     string               `json:"title,omitempty"`               //Product name, 1-32 characters
	Description               string               `json:"description,omitempty"`         //Product description, 1-255 characters
	Payload                   string               `json:"payload,omitempty"`             //Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use for your internal processes.
	ProviderToken             string               `json:"provider_token,omitempty"`      //Payments provider token, obtained via Botfather
	StartParameter            string               `json:"start_parameter,omitempty"`     //Unique deep-linking parameter that can be used to generate this invoice when used as a start parameter
	Currency                  string               `json:"currency,omitempty"`            //Three-letter ISO 4217 currency code, see more on currencies
	Prices                    []LabeledPrice       `json:"prices,omitempty"`              //Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	ProviderData              string               `json:"provider_data"`                 //JSON-encoded data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
	PhotoUrl                  string               `json:"photo_url"`                     //URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
	PhotoSize                 int                  `json:"photo_size"`                    //Photo size
	PhotoWidth                int                  `json:"photo_width"`                   //Photo width
	PhotoHeight               int                  `json:"photo_height"`                  //Photo height
	NeedName                  bool                 `json:"need_name"`                     //Pass True, if you require the user's full name to complete the order
	NeedPhoneNumber           bool                 `json:"need_phone_number"`             //Pass True, if you require the user's phone number to complete the order
	NeedEmail                 bool                 `json:"need_email"`                    //Pass True, if you require the user's email address to complete the order
	NeedShippingAddress       bool                 `json:"need_shipping_address"`         //Pass True, if you require the user's shipping address to complete the order
	SendPhoneNumberToProvider bool                 `json:"send_phone_number_to_provider"` //Pass True, if user's phone number should be sent to provider
	SendEmailToProvider       bool                 `json:"send_email_to_provider"`        //Pass True, if user's email address should be sent to provider
	IsFlexible                bool                 `json:"is_flexible"`                   //Pass True, if the final price depends on the shipping method
	DisableNotification       bool                 `json:"disable_notification"`          //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId          int                  `json:"reply_to_message_id"`           //If the message is a reply, ID of the original message
	ReplyMarkup               InlineKeyboardMarkup `json:"reply_markup"`                  //A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
}

//If you sent an invoice requesting a shipping address and the parameter is_flexible was specified, the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to shipping queries. On success, True is returned.
type AnswerShippingQuery struct {
	ShippingQueryId string           `json:"shipping_query_id,omitempty"` //Unique identifier for the query to be answered
	Ok              bool             `json:"ok,omitempty"`                //Specify True if delivery to the specified address is possible and False if there are any problems (for example, if delivery to the specified address is not possible)
	ShippingOptions []ShippingOption `json:"shipping_options"`            //Required if ok is True. A JSON-serialized array of available shipping options.
	ErrorMessage    string           `json:"error_message"`               //Required if ok is False. Error message in human readable form that explains why it is impossible to complete the order (e.g. "Sorry, delivery to your desired address is unavailable'). Telegram will display this message to the user.
}

//Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	PreCheckoutQueryId string `json:"pre_checkout_query_id,omitempty"` //Unique identifier for the query to be answered
	Ok                 bool   `json:"ok,omitempty"`                    //Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	ErrorMessage       string `json:"error_message"`                   //Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout (e.g. "Sorry, somebody just bought the last of our amazing black T-shirts while you were busy filling out your payment details. Please choose a different color or garment!"). Telegram will display this message to the user.
}

//This object represents a portion of the price for goods or services.
type LabeledPrice struct {
	Label  string `json:"label,omitempty"`  //Portion label
	Amount int    `json:"amount,omitempty"` //Price of the product in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

//This object contains basic information about an invoice.
type Invoice struct {
	Title          string `json:"title,omitempty"`           //Product name
	Description    string `json:"description,omitempty"`     //Product description
	StartParameter string `json:"start_parameter,omitempty"` //Unique bot deep-linking parameter that can be used to generate this invoice
	Currency       string `json:"currency,omitempty"`        //Three-letter ISO 4217 currency code
	TotalAmount    int    `json:"total_amount,omitempty"`    //Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
}

//This object represents a shipping address.
type ShippingAddress struct {
	CountryCode string `json:"country_code,omitempty"` //ISO 3166-1 alpha-2 country code
	State       string `json:"state,omitempty"`        //State, if applicable
	City        string `json:"city,omitempty"`         //City
	StreetLine1 string `json:"street_line1,omitempty"` //First line for the address
	StreetLine2 string `json:"street_line2,omitempty"` //Second line for the address
	PostCode    string `json:"post_code,omitempty"`    //Address post code
}

//This object represents information about an order.
type OrderInfo struct {
	Name            string           `json:"name"`             //Optional. User name
	PhoneNumber     string           `json:"phone_number"`     //Optional. User's phone number
	Email           string           `json:"email"`            //Optional. User email
	ShippingAddress *ShippingAddress `json:"shipping_address"` //Optional. User shipping address
}

//This object represents one shipping option.
type ShippingOption struct {
	Id     string         `json:"id,omitempty"`     //Shipping option identifier
	Title  string         `json:"title,omitempty"`  //Option title
	Prices []LabeledPrice `json:"prices,omitempty"` //List of price portions
}

//This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	Currency                string     `json:"currency,omitempty"`                   //Three-letter ISO 4217 currency code
	TotalAmount             int        `json:"total_amount,omitempty"`               //Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload          string     `json:"invoice_payload,omitempty"`            //Bot specified invoice payload
	ShippingOptionId        string     `json:"shipping_option_id"`                   //Optional. Identifier of the shipping option chosen by the user
	OrderInfo               *OrderInfo `json:"order_info"`                           //Optional. Order info provided by the user
	TelegramPaymentChargeId string     `json:"telegram_payment_charge_id,omitempty"` //Telegram payment identifier
	ProviderPaymentChargeId string     `json:"provider_payment_charge_id,omitempty"` //Provider payment identifier
}

//This object contains information about an incoming shipping query.
type ShippingQuery struct {
	Id              string           `json:"id,omitempty"`               //Unique query identifier
	From            *User            `json:"from,omitempty"`             //User who sent the query
	InvoicePayload  string           `json:"invoice_payload,omitempty"`  //Bot specified invoice payload
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"` //User specified shipping address
}

//This object contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	Id               string     `json:"id,omitempty"`              //Unique query identifier
	From             *User      `json:"from,omitempty"`            //User who sent the query
	Currency         string     `json:"currency,omitempty"`        //Three-letter ISO 4217 currency code
	TotalAmount      int        `json:"total_amount,omitempty"`    //Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	InvoicePayload   string     `json:"invoice_payload,omitempty"` //Bot specified invoice payload
	ShippingOptionId string     `json:"shipping_option_id"`        //Optional. Identifier of the shipping option chosen by the user
	OrderInfo        *OrderInfo `json:"order_info"`                //Optional. Order info provided by the user
}

//Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data,omitempty"`        //Array with information about documents and other Telegram Passport elements that was shared with the bot
	Credentials *EncryptedCredentials      `json:"credentials,omitempty"` //Encrypted credentials required to decrypt the data
}

//This object represents a file uploaded to Telegram Passport. Currently all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	FileId       string `json:"file_id,omitempty"`        //Identifier for this file, which can be used to download or reuse the file
	FileUniqueId string `json:"file_unique_id,omitempty"` //Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
	FileSize     int    `json:"file_size,omitempty"`      //File size
	FileDate     int    `json:"file_date,omitempty"`      //Unix time when the file was uploaded
}

//Contains information about documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	Type        string         `json:"type,omitempty"` //Element type. One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Data        string         `json:"data"`           //Optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types. Can be decrypted and verified using the accompanying EncryptedCredentials.
	PhoneNumber string         `json:"phone_number"`   //Optional. User's verified phone number, available only for “phone_number” type
	Email       string         `json:"email"`          //Optional. User's verified email address, available only for “email” type
	Files       []PassportFile `json:"files"`          //Optional. Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide   *PassportFile  `json:"front_side"`     //Optional. Encrypted file with the front side of the document, provided by the user. Available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile  `json:"reverse_side"`   //Optional. Encrypted file with the reverse side of the document, provided by the user. Available for “driver_license” and “identity_card”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie      *PassportFile  `json:"selfie"`         //Optional. Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation"`    //Optional. Array of encrypted files with translated versions of documents provided by the user. Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Hash        string         `json:"hash,omitempty"` //Base64-encoded element hash for using in PassportElementErrorUnspecified
}

//Contains data required for decrypting and authenticating EncryptedPassportElement. See the Telegram Passport Documentation for a complete description of the data decryption and authentication processes.
type EncryptedCredentials struct {
	Data   string `json:"data,omitempty"`   //Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes and secrets required for EncryptedPassportElement decryption and authentication
	Hash   string `json:"hash,omitempty"`   //Base64-encoded data hash for data authentication
	Secret string `json:"secret,omitempty"` //Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption
}

//Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.
type SetPassportDataErrors struct {
	UserId int                             `json:"user_id,omitempty"` //User identifier
	Errors []PassportElementErrorDataField `json:"errors,omitempty"`  //A JSON-serialized array describing the errors
}

//Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
type PassportElementErrorDataField struct {
	Source    string `json:"source,omitempty"`     //Error source, must be data
	Type      string `json:"type,omitempty"`       //The section of the user's Telegram Passport which has the error, one of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”
	FieldName string `json:"field_name,omitempty"` //Name of the data field which has the error
	DataHash  string `json:"data_hash,omitempty"`  //Base64-encoded data hash
	Message   string `json:"message,omitempty"`    //Error message
}

//Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	Source   string `json:"source,omitempty"`    //Error source, must be front_side
	Type     string `json:"type,omitempty"`      //The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	FileHash string `json:"file_hash,omitempty"` //Base64-encoded hash of the file with the front side of the document
	Message  string `json:"message,omitempty"`   //Error message
}

//Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	Source   string `json:"source,omitempty"`    //Error source, must be reverse_side
	Type     string `json:"type,omitempty"`      //The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”
	FileHash string `json:"file_hash,omitempty"` //Base64-encoded hash of the file with the reverse side of the document
	Message  string `json:"message,omitempty"`   //Error message
}

//Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	Source   string `json:"source,omitempty"`    //Error source, must be selfie
	Type     string `json:"type,omitempty"`      //The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”
	FileHash string `json:"file_hash,omitempty"` //Base64-encoded hash of the file with the selfie
	Message  string `json:"message,omitempty"`   //Error message
}

//Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	Source   string `json:"source,omitempty"`    //Error source, must be file
	Type     string `json:"type,omitempty"`      //The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash,omitempty"` //Base64-encoded file hash
	Message  string `json:"message,omitempty"`   //Error message
}

//Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	Source     string   `json:"source,omitempty"`      //Error source, must be files
	Type       string   `json:"type,omitempty"`        //The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHashes []string `json:"file_hashes,omitempty"` //List of base64-encoded file hashes
	Message    string   `json:"message,omitempty"`     //Error message
}

//Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	Source   string `json:"source,omitempty"`    //Error source, must be translation_file
	Type     string `json:"type,omitempty"`      //Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHash string `json:"file_hash,omitempty"` //Base64-encoded file hash
	Message  string `json:"message,omitempty"`   //Error message
}

//Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	Source     string   `json:"source,omitempty"`      //Error source, must be translation_files
	Type       string   `json:"type,omitempty"`        //Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”
	FileHashes []string `json:"file_hashes,omitempty"` //List of base64-encoded file hashes
	Message    string   `json:"message,omitempty"`     //Error message
}

//Represents an issue in an unspecified place. The error is considered resolved when new data is added.
type PassportElementErrorUnspecified struct {
	Source      string `json:"source,omitempty"`       //Error source, must be unspecified
	Type        string `json:"type,omitempty"`         //Type of element of the user's Telegram Passport which has the issue
	ElementHash string `json:"element_hash,omitempty"` //Base64-encoded element hash
	Message     string `json:"message,omitempty"`      //Error message
}

//Use this method to send a game. On success, the sent Message is returned.
type SendGame struct {
	ChatId              int                  `json:"chat_id,omitempty"`         //Unique identifier for the target chat
	GameShortName       string               `json:"game_short_name,omitempty"` //Short name of the game, serves as the unique identifier for the game. Set up your games via Botfather.
	DisableNotification bool                 `json:"disable_notification"`      //Sends the message silently. Users will receive a notification with no sound.
	ReplyToMessageId    int                  `json:"reply_to_message_id"`       //If the message is a reply, ID of the original message
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`              //A JSON-serialized object for an inline keyboard. If empty, one ‘Play game_title’ button will be shown. If not empty, the first button must launch the game.
}

//This object represents a game. Use BotFather to create and edit games, their short names will act as unique identifiers.
type Game struct {
	Title        string          `json:"title,omitempty"`       //Title of the game
	Description  string          `json:"description,omitempty"` //Description of the game
	Photo        []PhotoSize     `json:"photo,omitempty"`       //Photo that will be displayed in the game message in chats.
	Text         string          `json:"text"`                  //Optional. Brief description of the game or high scores included in the game message. Can be automatically edited to include current high scores for the game when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	TextEntities []MessageEntity `json:"text_entities"`         //Optional. Special entities that appear in text, such as usernames, URLs, bot commands, etc.
	Animation    *Animation      `json:"animation"`             //Optional. Animation that will be displayed in the game message in chats. Upload via BotFather
}

//Use this method to set the score of the specified user in a game. On success, if the message was sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new score is not greater than the user's current score in the chat and force is False.
type SetGameScore struct {
	UserId             int    `json:"user_id,omitempty"`    //User identifier
	Score              int    `json:"score,omitempty"`      //New score, must be non-negative
	Force              bool   `json:"force"`                //Pass True, if the high score is allowed to decrease. This can be useful when fixing mistakes or banning cheaters
	DisableEditMessage bool   `json:"disable_edit_message"` //Pass True, if the game message should not be automatically edited to include the current scoreboard
	ChatId             int    `json:"chat_id"`              //Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId          int    `json:"message_id"`           //Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId    string `json:"inline_message_id"`    //Required if chat_id and message_id are not specified. Identifier of the inline message
}

//Use this method to get data for high score tables. Will return the score of the specified user and several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
type GetGameHighScores struct {
	UserId          int    `json:"user_id,omitempty"` //Target user id
	ChatId          int    `json:"chat_id"`           //Required if inline_message_id is not specified. Unique identifier for the target chat
	MessageId       int    `json:"message_id"`        //Required if inline_message_id is not specified. Identifier of the sent message
	InlineMessageId string `json:"inline_message_id"` //Required if chat_id and message_id are not specified. Identifier of the inline message
}

//This object represents one row of the high scores table for a game.
type GameHighScore struct {
	Position int   `json:"position,omitempty"` //Position in high score table for the game
	User     *User `json:"user,omitempty"`     //User
	Score    int   `json:"score,omitempty"`    //Score
}
