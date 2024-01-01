package stores

type Merchant struct {
	Id          uint   `json:"id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	LogoUrl     string `json:"logoUrl"`
	CallbackUrl string `json:"callbackUrl"`
}
