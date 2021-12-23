package notificationsdk

type WebhookNotifyRequest struct {
	Events     []string    `json:"events"`
	VendorUUID string      `json:"vendor_uuid"`
	Content    interface{} `json:"content"`
}

type WhatsappNotifyRequest struct {
	To       string   `json:"to"`
	Metadata []string `json:"metadata"`
	Template string   `json:"template"`
}
