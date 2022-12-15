package model

import (
	"database/sql"
	"event-management/db"
	"fmt"
	"time"
)

type Events []Event

func (events *Events) FindAll() (err error) {
	users := UsersMap{}
	err = users.FindAll()
	if err != nil {
		return
	}

	tags := TagsMap{}
	err = tags.FindAll()
	if err != nil {
		return
	}

	var rows *sql.Rows

	sqlStatement := `SELECT * FROM events WHERE deleted_at IS NULL ORDER BY id`

	rows, err = db.Connection.Query(sqlStatement)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var event = Event{}
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.StartDate,
			&event.EndDate,
			&event.Description,
			&event.Location,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)

		if err != nil {
			return
		}

		event.ContactPerson = users[event.CreatedBy]

		eventTags := EventTags{}
		err = eventTags.FindByEventID(event.ID)
		if err != nil {
			return
		}

		tagsStr := []string{}
		for _, eventTag := range eventTags {
			tagsStr = append(tagsStr, tags[eventTag.TagID].Name)
		}

		event.Tags = tagsStr

		*events = append(*events, event)
	}

	return
}

func (events *Events) FindAllWithTag(tag string) (err error) {
	users := UsersMap{}
	err = users.FindAll()
	if err != nil {
		return
	}

	tags := TagsMap{}
	err = tags.FindAll()
	if err != nil {
		return
	}

	var rows *sql.Rows

	sqlStatement := `
		SELECT e.* FROM events e
		JOIN event_tags et ON et.event_id = e.id
		WHERE (SELECT COUNT(*) FROM tags t WHERE t.id = et.tag_id AND t."name" LIKE $1) > 0
		GROUP BY e.id
		ORDER BY id
	`

	rows, err = db.Connection.Query(sqlStatement, fmt.Sprintf(`%%%s%%`, tag))
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var event = Event{}
		err = rows.Scan(
			&event.ID,
			&event.Name,
			&event.StartDate,
			&event.EndDate,
			&event.Description,
			&event.Location,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)

		if err != nil {
			return
		}

		event.ContactPerson = users[event.CreatedBy]

		eventTags := EventTags{}
		err = eventTags.FindByEventID(event.ID)
		if err != nil {
			return
		}

		tagsStr := []string{}
		for _, eventTag := range eventTags {
			tagsStr = append(tagsStr, tags[eventTag.TagID].Name)
		}

		event.Tags = tagsStr

		*events = append(*events, event)
	}

	return
}

type Event struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	StartDate     time.Time `json:"start_date"`
	EndDate       time.Time `json:"end_date"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	CreatedBy     int       `json:"-"`
	ContactPerson User      `json:"contact_person"`
	Tags          []string  `json:"tags"`
	Timestamps
}

func (event *Event) Create() (err error) {
	sqlStatement := `
		INSERT INTO events (name, start_date, end_date, description, location, created_by) 
		VALUES ($1, $2, $3, $4, $5, $6)
		Returning *
	`

	err = db.Connection.
		QueryRow(
			sqlStatement,
			event.Name,
			event.StartDate,
			event.EndDate,
			event.Description,
			event.Location,
			event.CreatedBy,
		).
		Scan(
			&event.ID,
			&event.Name,
			&event.StartDate,
			&event.EndDate,
			&event.Description,
			&event.Location,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)

	if err != nil {
		return
	}

	return
}

func (event *Event) Find() (err error) {
	sqlStatement := `SELECT * FROM events WHERE deleted_at IS NULL AND id = $1`

	err = db.Connection.QueryRow(sqlStatement, event.ID).
		Scan(
			&event.ID,
			&event.Name,
			&event.StartDate,
			&event.EndDate,
			&event.Description,
			&event.Location,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)

	if err != nil {
		return
	}

	return
}

func (event *Event) Update() (err error) {
	sqlStatement := `
		UPDATE events
		SET name = $2, start_date = $3, end_date = $4, description = $5, location = $6, updated_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
		Returning *
	`

	err = db.Connection.
		QueryRow(
			sqlStatement,
			event.ID,
			event.Name,
			event.StartDate,
			event.EndDate,
			event.Description,
			event.Location,
		).
		Scan(
			&event.ID,
			&event.Name,
			&event.StartDate,
			&event.EndDate,
			&event.Description,
			&event.Location,
			&event.CreatedBy,
			&event.CreatedAt,
			&event.UpdatedAt,
			&event.DeletedAt,
		)

	if err != nil {
		return
	}

	return
}

func (event *Event) Delete() error {
	sqlStatement := `
		UPDATE events
		SET updated_at = NOW(), deleted_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
	`

	_, err := db.Connection.Exec(sqlStatement, event.ID)
	if err != nil {
		return err
	}

	return nil
}

type EventTags []EventTag

func (eventTags *EventTags) FindByEventID(eventID int) (err error) {
	var rows *sql.Rows

	sqlStatement := `SELECT * FROM event_tags WHERE event_id = $1`

	rows, err = db.Connection.Query(sqlStatement, eventID)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var eventTag = EventTag{}
		err = rows.Scan(
			&eventTag.EventID,
			&eventTag.TagID,
			&eventTag.CreatedAt,
		)

		if err != nil {
			return
		}

		*eventTags = append(*eventTags, eventTag)
	}

	return
}

type EventTag struct {
	EventID   int
	TagID     int
	CreatedAt time.Time
}

func (eventTag *EventTag) Create() (err error) {
	sqlStatement := `
		INSERT INTO event_tags (event_id, tag_id) 
		VALUES ($1, $2)
		Returning *
	`

	err = db.Connection.QueryRow(sqlStatement, eventTag.EventID, eventTag.TagID).
		Scan(&eventTag.EventID, &eventTag.TagID, &eventTag.CreatedAt)

	if err != nil {
		return
	}

	return
}

func (eventTag *EventTag) DeleteByEventID() error {
	sqlStatement := `DELETE FROM event_tags WHERE event_id = $1`

	_, err := db.Connection.Exec(sqlStatement, eventTag.EventID)
	if err != nil {
		return err
	}

	return nil
}

type EventParticipant struct {
	ID      int    `json:"-"`
	EventID int    `json:"-"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Status  string `json:"status"`
	Timestamps
}

func (ep *EventParticipant) Create() (err error) {
	sqlStatement := `
		INSERT INTO event_participants (event_id, name, email) 
		VALUES ($1, $2, $3)
		Returning *
	`

	err = db.Connection.QueryRow(sqlStatement, ep.EventID, ep.Name, ep.Email).
		Scan(&ep.ID, &ep.EventID, &ep.Name, &ep.Email, &ep.Status, &ep.CreatedAt, &ep.UpdatedAt, &ep.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (ep *EventParticipant) Find() (err error) {
	sqlStatement := `SELECT * FROM event_participants WHERE deleted_at IS NULL AND id = $1`

	err = db.Connection.QueryRow(sqlStatement, ep.ID).
		Scan(&ep.ID, &ep.EventID, &ep.Name, &ep.Email, &ep.Status, &ep.CreatedAt, &ep.UpdatedAt, &ep.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (ep *EventParticipant) FindByEmailAndEventID() (err error) {
	sqlStatement := `SELECT * FROM event_participants WHERE deleted_at IS NULL AND event_id = $1 AND email = $2`

	err = db.Connection.QueryRow(sqlStatement, ep.EventID, ep.Email).
		Scan(&ep.ID, &ep.EventID, &ep.Name, &ep.Email, &ep.Status, &ep.CreatedAt, &ep.UpdatedAt, &ep.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (ep *EventParticipant) Attend() (err error) {
	sqlStatement := `
		UPDATE event_participants
		SET status = 'ATTENDED', updated_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
	`

	err = db.Connection.QueryRow(sqlStatement, ep.ID).
		Scan(&ep.ID, &ep.EventID, &ep.Name, &ep.Email, &ep.Status, &ep.CreatedAt, &ep.UpdatedAt, &ep.DeletedAt)

	if err != nil {
		return
	}

	return
}

func (ep *EventParticipant) Cancel() (err error) {
	sqlStatement := `
		UPDATE event_participants
		SET status = 'CANCELED', updated_at = NOW(), deleted_at = NOW()
		WHERE deleted_at IS NULL AND id = $1
	`

	_, err = db.Connection.Exec(sqlStatement, ep.ID)

	if err != nil {
		return
	}

	return
}
