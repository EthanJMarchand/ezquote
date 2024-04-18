package api

type APIWorker interface {
	GetAll() ([]byte, error)
	GetQuote(string) ([]byte, error)
}

type APIService struct {
	Worker APIWorker
}

func NewAPIService(w APIWorker) *APIService {
	return &APIService{
		Worker: w,
	}
}
