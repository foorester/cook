package transport

type (
	GetBooksRes struct {
		ServiceRes
		Books []GetBookReq
	}
)
