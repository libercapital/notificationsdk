package notificationsdk

type WebhookNotifyRequest struct {
	Events         []string    `json:"events"`
	VendorUUID     string      `json:"vendor_uuid"`
	IntegratorUUID *string     `json:"integrator_uuid"`
	Content        interface{} `json:"content"`
}

type WhatsappNotifyRequest struct {
	To       string   `json:"to"`
	Metadata []string `json:"metadata"`
	Template string   `json:"template"`
}

type EmailRequest struct {
	To          Address                `json:"to"`
	From        *Address               `json:"from,omitempty"`
	Template    string                 `json:"template"`
	MetaData    map[string]interface{} `json:"meta_data"`
	Attachments []Attachment           `json:"attachments"`
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
