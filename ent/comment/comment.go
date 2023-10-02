// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the comment type in the database.
	Label = "comment"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAuthor holds the string denoting the author field in the database.
	FieldAuthor = "author"
	// FieldPostid holds the string denoting the postid field in the database.
	FieldPostid = "postid"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldLikes holds the string denoting the likes field in the database.
	FieldLikes = "likes"
	// Table holds the table name of the comment in the database.
	Table = "comments"
)

// Columns holds all SQL columns for comment fields.
var Columns = []string{
	FieldID,
	FieldAuthor,
	FieldPostid,
	FieldContent,
	FieldLikes,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Comment queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByAuthor orders the results by the author field.
func ByAuthor(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAuthor, opts...).ToFunc()
}

// ByPostid orders the results by the postid field.
func ByPostid(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPostid, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}

// ByLikes orders the results by the likes field.
func ByLikes(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLikes, opts...).ToFunc()
}
