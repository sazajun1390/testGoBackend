// Code generated by ent, DO NOT EDIT.

package chatroom

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the chatroom type in the database.
	Label = "chat_room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCreator holds the string denoting the creator edge name in mutations.
	EdgeCreator = "creator"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// EdgeParticipants holds the string denoting the participants edge name in mutations.
	EdgeParticipants = "participants"
	// EdgeMemberships holds the string denoting the memberships edge name in mutations.
	EdgeMemberships = "memberships"
	// Table holds the table name of the chatroom in the database.
	Table = "chat_rooms"
	// CreatorTable is the table that holds the creator relation/edge.
	CreatorTable = "chat_rooms"
	// CreatorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatorInverseTable = "users"
	// CreatorColumn is the table column denoting the creator relation/edge.
	CreatorColumn = "user_chat_rooms"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "chat_room_messages"
	// ParticipantsTable is the table that holds the participants relation/edge. The primary key declared below.
	ParticipantsTable = "chat_room_members"
	// ParticipantsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ParticipantsInverseTable = "users"
	// MembershipsTable is the table that holds the memberships relation/edge.
	MembershipsTable = "chat_room_members"
	// MembershipsInverseTable is the table name for the ChatRoomMember entity.
	// It exists in this package in order to avoid circular dependency with the "chatroommember" package.
	MembershipsInverseTable = "chat_room_members"
	// MembershipsColumn is the table column denoting the memberships relation/edge.
	MembershipsColumn = "room_id"
)

// Columns holds all SQL columns for chatroom fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "chat_rooms"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_chat_rooms",
}

var (
	// ParticipantsPrimaryKey and ParticipantsColumn2 are the table columns denoting the
	// primary key for the participants relation (M2M).
	ParticipantsPrimaryKey = []string{"user_id", "room_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the ChatRoom queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatorField orders the results by creator field.
func ByCreatorField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCreatorStep(), sql.OrderByField(field, opts...))
	}
}

// ByMessagesCount orders the results by messages count.
func ByMessagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMessagesStep(), opts...)
	}
}

// ByMessages orders the results by messages terms.
func ByMessages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMessagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByParticipantsCount orders the results by participants count.
func ByParticipantsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newParticipantsStep(), opts...)
	}
}

// ByParticipants orders the results by participants terms.
func ByParticipants(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newParticipantsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByMembershipsCount orders the results by memberships count.
func ByMembershipsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newMembershipsStep(), opts...)
	}
}

// ByMemberships orders the results by memberships terms.
func ByMemberships(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newMembershipsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCreatorStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CreatorInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, CreatorTable, CreatorColumn),
	)
}
func newMessagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MessagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, MessagesTable, MessagesColumn),
	)
}
func newParticipantsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ParticipantsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ParticipantsTable, ParticipantsPrimaryKey...),
	)
}
func newMembershipsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(MembershipsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, MembershipsTable, MembershipsColumn),
	)
}