package dao

import (
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/ujjkumsi/docker-go/models"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var session *gocql.Session

const (
	COLLECTION = "movies"
)

type stop struct {
	error
}

func retry(attempts int, sleep time.Duration, server string, database string, fn func(server string, database string) error) error {
	if err := fn(server, database); err != nil {
		if s, ok := err.(stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return retry(attempts, 2*sleep, server, database, fn)
		}
		return err
	}
	return nil
}

/*
Connect function
This lets the application to connect with the database - cassandra
Its an extension function on MoviesDAO which expects server address and keyspace name
Here, we use default port : 9042 for cql queries
*/
func (m *MoviesDAO) Connect() {
	if err := retry(3, 60*1000*1000*1000, m.Server, m.Database, initSession); err != nil {
		log.Fatal("Unable to connect with database after several tries")
	}
	log.Println("Connected to database")
}

func initSession(server string, database string) error {
	var err error
	if session == nil || session.Closed() {
		if session, err = getCluster(server, database).CreateSession(); err != nil {
			return err
		}
	}
	return nil
}

func getCluster(server string, database string) *gocql.ClusterConfig {
	cluster := gocql.NewCluster(server)
	cluster.Port = 9042
	cluster.ProtoVersion = 4
	cluster.Keyspace = database
	return cluster
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	return movies, nil
}

// FindByID function lets you find movies by providing their ids
func (m *MoviesDAO) FindByID(id string) (models.Movie, error) {
	var movie models.Movie
	return movie, nil
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie models.Movie) error {
	if err := initSession(m.Server, m.Database); err != nil {
		return err
	}
	if err := session.Query(
		`INSERT INTO movieapi.movie (id, name, cover_image, description) VALUES (?, ?, ?, ?)`,
		movie.ID, movie.Name, movie.CoverImage, movie.Description).Exec(); err != nil {
		return err
	}

	return nil
}

// Delete an existing movie
func (m *MoviesDAO) Delete(movie models.Movie) error {
	return nil
}

// Update an existing movie
func (m *MoviesDAO) Update(movie models.Movie) error {
	return nil
}
