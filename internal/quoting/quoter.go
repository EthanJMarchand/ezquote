package quoting

type QuotingWorker interface {
	GetAll() ([]Coin, error)
	GetQuote(string) (*Quote, error)
}

type QuotingService struct {
	Worker QuotingWorker
}

// TODO: Validate the worker
func NewAPIService(w QuotingWorker) *QuotingService {
	return &QuotingService{
		Worker: w,
	}
}
