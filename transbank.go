package transbank

type Transaction interface {
	InitTransaction(params InitTransaction) (*InitTransactionResponse, error)
	GetTransactionResult(token string) (*TransactionResultResponse, error)
}

// OneCLick WIP: define all OneCLick service methods
type OneCLick interface {
	InitInscription(params interface{})
	FinishInscription(token string)
	Authorize(params interface{})
	ReverseTransaction(buyorder string)
	RemoveUser(tbkUser string, username string)
}

// OneCLickMall WIP: represents the "initTransactionRequest" to SOAP server webpay plus normal
type OneCLickMall interface{}

// InitTransaction represents the "initTransactionRequest" to SOAP server webpay plus normal,
// its the base params to use on plus services
type InitTransaction struct {
	Amount    float64 `json:"amount"`
	SessionID string  `json:"session_id"`
	ReturnURL string  `json:"return_url"`
	FinalURL  string  `json:"final_url,omitempty"`
	BuyOrder  string  `json:"buy_order"`

	Stores *[]Store `json:"stores"`

	WPMDetail *WPMDetail `json:"wpm_detail"`
}

// Store WIP: ...
type Store struct {
	CommerceCode string  `json:"commerce_code"`
	Amount       float64 `json:"amount"`
	BuyOrder     string  `json:"buy_order"`
}

// WPMDetail respresent the "WPMDetail" with user's inscription data
type WPMDetail struct {
	ServiceID           string `json:"service_id"`
	CardHolderID        string `json:"card_holder_id"`
	CardHolderName      string `json:"card_holder_name"`
	CardHolderLastName1 string `json:"card_holder_last_name_1"`
	CardHolderLastName2 string `json:"card_holder_last_name_2"`
	CardHolderMail      string `json:"card_holder_mail"`
	CellPhoneNumber     string `json:"cell_phone_number"`
	ExpirationDate      string `json:"expiration_date"`
	CommerceMail        string `json:"commerc_mail"`
	UfFlag              bool   `json:"uf_flag"`
}

// InitTransactionResponse represents the "initTransactionResponse" from SOAP server
type InitTransactionResponse struct {
	URL   string `json:"url"`
	Token string `json:"token"`
}

// TransactionResultResponse represents the "getTransactionResultResponse" SOAP server response
type TransactionResultResponse struct {
	AccountingDate  string       `json:"accounting_date"`
	BuyOrder        string       `json:"buy_order"`
	CardDetail      CardDetail   `json:"card_detail"`
	DetailOutput    DetailOutput `json:"detail_output"`
	SessionID       string       `json:"session_id"`
	TransactionDate string       `json:"transaction_date"`
	URLRedirection  string       `json:"url_redirection"`
	VCI             string       `json:"vci"`
}

// CardDetail represent the values of customer card
type CardDetail struct {
	CardNumber         string `json:"card_number"`
	CardExpirationDate string `json:"card_expiration_date"`
}

// DetailOutput represent transaction details values
type DetailOutput struct {
	SharesNumber      int     `json:"shares_number"`
	Amount            float64 `json:"amount"`
	CommerceCode      string  `json:"commerce_code"`
	BuyOrder          string  `json:"buy_order"`
	AuthorizationCode string  `json:"authorization_code"`
	PaymentTypeCode   string  `json:"payment_type_code"`
	ResponseCode      string  `json:"response_code"`
}
