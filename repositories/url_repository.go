package repositories

import (
	"github.com/gocql/gocql"
	"time"
)

type UrlRepository interface {
	Save(UrlData) error
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
