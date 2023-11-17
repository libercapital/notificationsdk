package notificationsdk

type WebhookNotifyRequest struct {
	Events         []string    `json:"events"`
	VendorUUID     string      `json:"vendor_uuid"`
	IntegratorUUID *string     `json:"integrator_uuid"`
	Content        interface{} `json:"content"`
	SNSParams      SNSParams   `json:"-"`
	AccessToken    string      `json:"-"`
}

type WhatsappNotifyRequest struct {
	To          string    `json:"to"`
	Metadata    []string  `json:"metadata"`
	Template    string    `json:"template"`
	SNSParams   SNSParams `json:"-"`
	AccessToken string    `json:"-"`
}

type EmailRequest struct {
	To          Address                `json:"to"`
	From        *Address               `json:"from,omitempty"`
	Template    string                 `json:"template"`
	MetaData    map[string]interface{} `json:"meta_data"`
	Attachments []Attachment           `json:"attachments"`
	SNSParams   SNSParams              `json:"-"`
	AccessToken string                 `json:"-"`
}

type Address struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Attachment struct {
	FileName    string `json:"file_name"`
	FileType    string `json:"file_type"`
	FileContent string `json:"file_content"`
}

type SNSParams struct {
	DeduplicationID string
	GroupID         string
	TopicArn        string
}
