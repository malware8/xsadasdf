package token

type TokenRequest struct {
	Settings Settings `json:"settings,omitempty"`
	User     User     `json:"user,omitempty"`
}

type Settings struct {
	ProjectID      int64          `json:"project_id,omitempty"`
	Currency       string         `json:"currency,omitempty"`
	ExternalID     string         `json:"external_id,omitempty"`
	Language       Language       `json:"language,omitempty"`
	Mode           string         `json:"mode,omitempty"` // default - "sandbox"
	PaymentMethod  int64          `json:"payment_method,omitempty"`
	PaymentWidget  PaymentWidget  `json:"payment_widget,omitempty"`
	RedirectPolicy RedirectPolicy `json:"redirect_policy,omitempty"`
	ReturnURL      string         `json:"return_url,omitempty"`
	UI             UI             `json:"ui,omitempty"`
}

type Language string

const (
	EN Language = "en" // Английский (США)
	AR Language = "ar" // Арабский
	BG Language = "bg" // Болгарский
	CN Language = "cn" // Китайский упрощенный
	TW Language = "tw" // Китайский традиционный
	CS Language = "cs" // Чешский
	FR Language = "fr" // Французский
	DE Language = "de" // Немецкий
	HE Language = "he" // Иврит
	IT Language = "it" // Итальянский
	JA Language = "ja" // Японский
	KO Language = "ko" // Корейский
	PL Language = "pl" // Польский
	PT Language = "pt" // Португальский (Бразилия)
	RO Language = "ro" // Румынский
	RU Language = "ru" // Русский (Россия)
	ES Language = "es" // Испанский (Испания)
	TH Language = "th" // Тайский
	TR Language = "tr" // Турецкий
	VI Language = "vi" // Вьетнамский
)

type PaymentWidget string

const (
	Paybycash PaymentWidget = "paybycash" // PayByCash
	Giftcard  PaymentWidget = "giftcard"  // GiftCard
)

type RedirectPolicy struct {
	AutoredirectFromStatusPage bool                    `json:"autoredirect_from_status_page,omitempty"`
	Delay                      int                     `json:"delay,omitempty"`
	ManualRedirectionAction    ManualRedirectionAction `json:"manual_redirection_action,omitempty"`
	RedirectButtonCaption      string                  `json:"redirect_button_caption,omitempty"`
	RedirectConditions         RedirectConditions      `json:"redirect_conditions,omitempty"`
	StatusForManualRedirection RedirectConditions      `json:"status_for_manual_redirection,omitempty"`
}

type ManualRedirectionAction string

const (
	Redirect    ManualRedirectionAction = "redirect"
	Postmessage ManualRedirectionAction = "postmessage"
)

type RedirectConditions string

const (
	None                 RedirectConditions = "none"
	Any                  RedirectConditions = "any"
	Successful           RedirectConditions = "successful"
	SuccessfulOrCanceled RedirectConditions = "successful_or_canceled"
)

type UI struct {
	Components                Components  `json:"components,omitempty"`
	Desktop                   Desktop     `json:"desktop,omitempty"`
	IsPaymentMethodsListMode  bool        `json:"is_payment_methods_list_mode,omitempty"`
	IsPreventExternalLinkOpen bool        `json:"is_prevent_external_link_open,omitempty"`
	Mode                      string      `json:"mode,omitempty"`
	Theme                     string      `json:"theme,omitempty"`
	UserAccount               UserAccount `json:"user_account,omitempty"`
}

type Components struct {
	Subscriptions   Subscriptions   `json:"subscriptions,omitempty"`
	VirtualCurrency VirtualCurrency `json:"virtual_currency,omitempty"`
	VirtualItems    VirtualItems    `json:"virtual_items,omitempty"`
}

type Subscriptions struct {
	Hidden bool  `json:"hidden,omitempty"`
	Order  int64 `json:"order,omitempty"`
}

type VirtualCurrency struct {
	CustomAmount bool  `json:"custom_amount,omitempty"`
	Hidden       bool  `json:"hidden,omitempty"`
	Order        int64 `json:"order,omitempty"`
}

type VirtualItems struct {
	Hidden        bool   `json:"hidden,omitempty"`
	Order         int64  `json:"order,omitempty"`
	SelectedGroup string `json:"selected_group,omitempty"`
	SelectedItem  string `json:"selected_item,omitempty"`
}

type Desktop struct {
	Header           DesktopHeader    `json:"header,omitempty"`
	SubscriptionList SubscriptionList `json:"subscription_list,omitempty"`
}

type DesktopHeader struct {
	CloseButton     bool              `json:"close_button,omitempty"`
	IsVisible       bool              `json:"is_visible,omitempty"`
	Type            DesktopHeaderType `json:"type,omitempty"`
	VisibleLogo     bool              `json:"visible_logo,omitempty"`
	VisibleName     bool              `json:"visible_name,omitempty"`
	VisiblePurchase bool              `json:"visible_purchase,omitempty"`
}

type DesktopHeaderType string

const (
	Compact DesktopHeaderType = "compact"
	Normanl DesktopHeaderType = "normal"
)

type SubscriptionList struct {
	Description       string `json:"description,omitempty"`
	DisplayLocalPrice bool   `json:"display_local_price,omitempty"`
}

type Header struct {
	VisibleVirtualCurrencyBalance bool `json:"visible_virtual_currency_balance,omitempty"`
}

type UserAccount struct {
	Info            Info                     `json:"info,omitempty"`
	PaymentAccounts PaymentAccounts          `json:"payment_accounts,omitempty"`
	Subscriptions   UserAccountSubscriptions `json:"subscriptions,omitempty"`
}

type Info struct {
	Enable bool  `json:"enable,omitempty"`
	Order  int64 `json:"order,omitempty"`
}

type PaymentAccounts struct {
	Enable bool  `json:"enable,omitempty"`
	Order  int64 `json:"order,omitempty"`
}

type UserAccountSubscriptions struct {
	Enable bool  `json:"enable,omitempty"`
	Order  int64 `json:"order,omitempty"`
}

type User struct {
	ID         ID                `json:"id,omitempty"`
	Age        int64             `json:"age,omitempty"`
	Attributes map[string]string `json:"attributes,omitempty"`
	Country    Country           `json:"country,omitempty"`
	Email      Email             `json:"email,omitempty"`
	IsLegal    bool              `json:"is_legal,omitempty"`
	Legal      Legal             `json:"legal,omitempty"`
	Name       Name              `json:"name,omitempty"`
	Phone      Phone             `json:"phone,omitempty"`
	PublicID   PublicID          `json:"public_id,omitempty"`
	SteamID    SteamID           `json:"steam_id,omitempty"`
	TrackingID TrackingID        `json:"tracking_id,omitempty"`
	UTM        UTM               `json:"utm,omitempty"`
}

type ID struct {
	Value string `json:"value,omitempty"`
}

type Country struct {
	AllowModify bool   `json:"allow_modify,omitempty"`
	Value       string `json:"value,omitempty"`
}

type Email struct {
	Value       string `json:"value,omitempty"`
	AllowModify bool   `json:"allow_modify,omitempty"`
}

type Legal struct {
	Address string `json:"address,omitempty"`
	Country string `json:"country,omitempty"`
	Name    string `json:"name,omitempty"`
	VatID   string `json:"vat_id,omitempty"`
}

type Name struct {
	AllowModify bool   `json:"allow_modify,omitempty"`
	Value       string `json:"value,omitempty"`
}

type Phone struct {
	Value string `json:"value,omitempty"`
}

type PublicID struct {
	Value string `json:"value,omitempty"`
}

type SteamID struct {
	Value string `json:"value,omitempty"`
}

type TrackingID struct {
	Value string `json:"value,omitempty"` // 32 chars
}

type UTM struct {
	Source   string `json:"utm_source,omitempty"`
	Medium   string `json:"utm_medium,omitempty"`
	Campaign string `json:"utm_campaign,omitempty"`
	Content  string `json:"utm_content,omitempty"`
	Term     string `json:"utm_term,omitempty"`
}

type TokenResponse struct {
	Token           string          `json:"token,omitempty"`
	ExtendedMessage ExtendedMessage `json:"extended_message,omitempty"`
	HttpStatusCode  int64           `json:"http_status_code,omitempty"`
	Message         string          `json:"message,omitempty"`
	RequestID       string          `json:"request_id,omitempty"`
}

type ExtendedMessage struct {
	GlobalErrors   []string            `json:"global_errors,omitempty"`
	PropertyErrors map[string][]string `json:"property_errors,omitempty"`
	Token          string              `json:"token,omitempty"`
}
