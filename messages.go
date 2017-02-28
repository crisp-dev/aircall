package aircall

type Request struct{}

type TransferCallRequest struct {
	UserID int `json:"user_id"`
}

type LinkCallRequest struct {
	Link string `json:"link"`
}

type Paginate struct {
	Page    int
	PerPage int
	Order   string
	From    int
	To      int
}

type Search struct {
	PhoneNumber string
	Email       string
}

type ContactRequest struct {
	FirstName    string        `json:"first_name"`
	LastName     string        `json:"last_name"`
	CompanyName  string        `json:"company_name"`
	Information  string        `json:"information"`
	PhoneNumbers []ContactInfo `json:"phone_numbers"`
	Emails       []ContactInfo `json:"emails"`
}

type Response struct {
}

type ResponseMeta struct {
	Count            int    `json:"count"`
	Total            int    `json:"total"`
	CurrentPage      int    `json:"current_page"`
	PerPage          int    `json:"per_page"`
	NextPageLink     string `json:"next_page_link"`
	PreviousPageLink string `json:"previous_page_link"`
}

type PingResponse struct {
	Ping string `json:"ping"`
}

type CompanyResponse struct {
	Company Company `json:"company"`
}

type UsersResponse struct {
	Meta  ResponseMeta `json:"meta"`
	Users []User       `json:"users"`
}

type UserResponse struct {
	User User `json:"user"`
}

type NumbersResponse struct {
	Meta    ResponseMeta `json:"meta"`
	Numbers []Number     `json:"numbers"`
}

type NumberResponse struct {
	Number Number `json:"number"`
}

type CallsResponse struct {
	Meta  ResponseMeta `json:"meta"`
	Calls []Call       `json:"calls"`
}

type CallResponse struct {
	Call Call `json:"call"`
}

type ContactsResponse struct {
	Meta     ResponseMeta `json:"meta"`
	Contacts []Contact    `json:"contacts"`
}

type ContactResponse struct {
	Contact Contact `json:"contact"`
}
