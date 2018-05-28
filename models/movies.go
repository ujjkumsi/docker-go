package models

import "github.com/gocql/gocql"

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Movie struct {
	ID          gocql.UUID `json:"id"`
	Name        string     `json:"name"`
	CoverImage  string     `json:"cover_image"`
	Description string     `json:"description"`
}
