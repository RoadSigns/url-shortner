package repositories

import (
	"github.com/gocql/gocql"
	"time"
)

type UrlRepository interface {
	Save(UrlData) error
	Get(shortUrl string) (urlData UrlData, err error)
}

type UrlData struct {
	ShortUrl string
	LongUrl  string
}

type CassandraUrlRepository struct {
	Cluster *gocql.ClusterConfig
}

func (r CassandraUrlRepository) Save(data UrlData) error {
	r.Cluster.Keyspace = "urls"
	r.Cluster.Consistency = gocql.Quorum
	r.Cluster.ProtoVersion = 4
	r.Cluster.ConnectTimeout = time.Second * 10
	r.Cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	session, sessionErr := r.Cluster.CreateSession()
	defer session.Close()

	if sessionErr != nil {
		return sessionErr
	}
	var uuid, _ = gocql.RandomUUID()
	if err := session.Query(`INSERT INTO urls (uuid, short_url, long_url) VALUES (?, ?, ?)`,
		uuid, data.ShortUrl, data.LongUrl).Exec(); err != nil {
		return err
	}

	return nil
}

func (r CassandraUrlRepository) Get(shortUrl string) (urlData UrlData, err error) {
	r.Cluster.Keyspace = "urls"
	r.Cluster.Consistency = gocql.Quorum
	r.Cluster.ProtoVersion = 4
	r.Cluster.ConnectTimeout = time.Second * 10
	r.Cluster.Authenticator = gocql.PasswordAuthenticator{Username: "cassandra", Password: "cassandra"}
	session, sessionErr := r.Cluster.CreateSession()
	defer session.Close()

	if sessionErr != nil {
		return UrlData{}, sessionErr
	}
	var id gocql.UUID
	var shorturl string
	var longurl string

	if err := session.Query(`SELECT uuid, short_url, long_url FROM urls WHERE short_url = ? LIMIT 1 ALLOW FILTERING`,
		shortUrl).Consistency(gocql.One).Scan(&id, &shorturl, &longurl); err != nil {
		return UrlData{}, err
	}

	return UrlData{
		ShortUrl: shorturl,
		LongUrl:  longurl,
	}, nil
}
