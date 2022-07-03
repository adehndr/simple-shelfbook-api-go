package web

type WebResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type WebResponseId struct {
	BookId string `json:"bookId"`
}

type WebResponseGetAll struct {
	Books []WebResponseGet `json:"books"`
}

type WebResponseGetById struct {
	Book BookResponse `json:"book"`
}

type WebResponseGet struct {
	BookId        string `json:"id"`
	BookName      string `json:"name"`
	BookPublisher string `json:"publisher"`
}
