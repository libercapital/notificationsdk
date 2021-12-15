package notificationsdk

type WebhookNotifyRequest struct {
	Events     []string    `json:"events"`
	VendorUUID string      `json:"vendor_uuid"`
	Content    interface{} `json:"content"`
}
