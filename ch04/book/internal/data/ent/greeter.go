// Code generated by entc, DO NOT EDIT.

package ent

import (
	"book/internal/data/ent/greeter"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Greeter is the model entity for the Greeter schema.
type Greeter struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Greeter) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case greeter.FieldID:
			values[i] = new(sql.NullInt64)
		case greeter.FieldName, greeter.FieldDescription:
			values[i] = new(sql.NullString)
		case greeter.FieldCreatedAt, greeter.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Greeter", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Greeter fields.
func (gr *Greeter) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case greeter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gr.ID = int64(value.Int64)
		case greeter.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				gr.Name = value.String
			}
		case greeter.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				gr.Description = value.String
			}
		case greeter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gr.CreatedAt = value.Time
			}
		case greeter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gr.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Greeter.
// Note that you need to call Greeter.Unwrap() before calling this method if this Greeter
// was returned from a transaction, and the transaction was committed or rolled back.
func (gr *Greeter) Update() *GreeterUpdateOne {
	return (&GreeterClient{config: gr.config}).UpdateOne(gr)
}

// Unwrap unwraps the Greeter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gr *Greeter) Unwrap() *Greeter {
	tx, ok := gr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Greeter is not a transactional entity")
	}
	gr.config.driver = tx.drv
	return gr
}

// String implements the fmt.Stringer.
func (gr *Greeter) String() string {
	var builder strings.Builder
	builder.WriteString("Greeter(")
	builder.WriteString(fmt.Sprintf("id=%v", gr.ID))
	builder.WriteString(", name=")
	builder.WriteString(gr.Name)
	builder.WriteString(", description=")
	builder.WriteString(gr.Description)
	builder.WriteString(", created_at=")
	builder.WriteString(gr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(gr.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Greeters is a parsable slice of Greeter.
type Greeters []*Greeter

func (gr Greeters) config(cfg config) {
	for _i := range gr {
		gr[_i].config = cfg
	}
}
