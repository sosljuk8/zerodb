// Code generated by ent, DO NOT EDIT.

package ent

import (
	"composeapp/ent/settings"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Settings is the model entity for the Settings schema.
type Settings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Value holds the value of the "value" field.
	Value string `json:"value,omitempty"`
	// UID holds the value of the "uid" field.
	UID int `json:"uid,omitempty"`
	// Created holds the value of the "created" field.
	Created int `json:"created,omitempty"`
	// Actuality holds the value of the "actuality" field.
	Actuality    int `json:"actuality,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldID, settings.FieldCategoryID, settings.FieldUID, settings.FieldCreated, settings.FieldActuality:
			values[i] = new(sql.NullInt64)
		case settings.FieldName, settings.FieldValue:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settings fields.
func (s *Settings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settings.FieldCategoryID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field category_id", values[i])
			} else if value.Valid {
				s.CategoryID = int(value.Int64)
			}
		case settings.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				s.Name = value.String
			}
		case settings.FieldValue:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				s.Value = value.String
			}
		case settings.FieldUID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uid", values[i])
			} else if value.Valid {
				s.UID = int(value.Int64)
			}
		case settings.FieldCreated:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created", values[i])
			} else if value.Valid {
				s.Created = int(value.Int64)
			}
		case settings.FieldActuality:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field actuality", values[i])
			} else if value.Valid {
				s.Actuality = int(value.Int64)
			}
		default:
			s.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// GetValue returns the ent.Value that was dynamically selected and assigned to the Settings.
// This includes values selected through modifiers, order, etc.
func (s *Settings) GetValue(name string) (ent.Value, error) {
	return s.selectValues.Get(name)
}

// Update returns a builder for updating this Settings.
// Note that you need to call Settings.Unwrap() before calling this method if this Settings
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settings) Update() *SettingsUpdateOne {
	return NewSettingsClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Settings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settings) Unwrap() *Settings {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settings is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settings) String() string {
	var builder strings.Builder
	builder.WriteString("Settings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("category_id=")
	builder.WriteString(fmt.Sprintf("%v", s.CategoryID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(s.Name)
	builder.WriteString(", ")
	builder.WriteString("value=")
	builder.WriteString(s.Value)
	builder.WriteString(", ")
	builder.WriteString("uid=")
	builder.WriteString(fmt.Sprintf("%v", s.UID))
	builder.WriteString(", ")
	builder.WriteString("created=")
	builder.WriteString(fmt.Sprintf("%v", s.Created))
	builder.WriteString(", ")
	builder.WriteString("actuality=")
	builder.WriteString(fmt.Sprintf("%v", s.Actuality))
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings
