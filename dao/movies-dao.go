package dao

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/ujjkumsi/docker-go/models"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var Session *gocql.Session

const (
	COLLECTION = "movies"
)

// Establish a connection to database
func (m *MoviesDAO) Connect() {
	log.Println("connecting to cassandra")

	cluster := gocql.NewCluster(m.Server)
	cluster.ProtoVersion = 4
	cluster.Keyspace = m.Database

	session, err := cluster.CreateSession()

	Session = session

	if err != nil {
		log.Println("connection to cassndra failed")
		log.Fatal(err)
	}

	log.Println("connected to cassandra")
}

// Find list of movies
func (m *MoviesDAO) FindAll() ([]models.Movie, error) {
	var movies []models.Movie
	return movies, nil
}

// Find a movie by its id
func (m *MoviesDAO) FindById(id string) (models.Movie, error) {
	var movie models.Movie
	return movie, nil
}

// Insert a movie into database
func (m *MoviesDAO) Insert(movie models.Movie) error {
	if err := Session.Query(
		`INSERT INTO moviedb.movie (id, name, cover_image, description) VALUES (?, ?, ?, ?)`,
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
