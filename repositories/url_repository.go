package repositories

type UrlRepository interface {
	Save(UrlData) error
}

type UrlData struct {
	ShortUrl string
	LongUrl  string
}

type CassandraUrlRepository struct {
}

func (r CassandraUrlRepository) Save(data UrlData) error {
	return nil
}
