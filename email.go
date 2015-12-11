package tempmail

// Email represents received email.
type Email struct {
	UniqueID  string `json:"mail_unique_id"`
	ID        string `json:"mail_id"`
	AddressID string `json:"mail_address_id"`
	From      string `json:"mail_from"`
	Subject   string `json:"mail_subject"`
	Preview   string `json:"mail_preview"`
	TextOnly  string `json:"mail_text_only"`
	Text      string `json:"mail_text"`
	HTML      string `json:"mail_html"`
}

// Delete deletes current email.
func (e Email) Delete() error {
	return nil
}
