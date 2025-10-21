package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/francotraversa/siriusbackend/internal/auth"
	"github.com/francotraversa/siriusbackend/internal/types"
	"github.com/francotraversa/siriusbackend/internal/utils"
	"github.com/labstack/echo/v4"
)

func GetMessageByDateUseCase(c echo.Context) (*[]types.MessageFilter, error) {
	db := utils.DatabaseInstance{}.Instance()
	uid, err := auth.IdFromContext(c)
	if err != nil {
		return nil, err
	}
	loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	var from, to time.Time

	if b := c.QueryParam("between"); b != "" {
		ps := strings.SplitN(b, ",", 2)
		if len(ps) == 2 {
			if f, _ := parseDate(ps[0], loc); !f.IsZero() {
				from = f
			}
			if t, dateOnly := parseDate(ps[1], loc); !t.IsZero() {
				if dateOnly {
					t = t.AddDate(0, 0, 1)
				}
				to = t
			}
		}
	} else {
		if f, _ := parseDate(c.QueryParam("from"), loc); !f.IsZero() {
			from = f
		}
		if t, dateOnly := parseDate(c.QueryParam("to"), loc); !t.IsZero() {
			if dateOnly {
				t = t.AddDate(0, 0, 1)
			} // [from, to)
			to = t
		}
	}
	statuses := normalizarConsulta(c.QueryParam("status"))
	services := normalizarConsulta(c.QueryParam("service"))

	q := db.Table("message_destinations").
		Select(`
		message_destinations.id      AS dest_id,
		message_destinations.service AS service,
		message_destinations.status  AS status,
		messages.id                  AS message_id,
		messages.content             AS content,
		messages.created_at          AS created_at`).
		Joins("JOIN messages ON messages.id = message_destinations.message_id").
		Where("messages.user_id = ?", uid)

	if !from.IsZero() {
		q = q.Where("messages.created_at >= ?", from) // ← calificado
	}
	if !to.IsZero() {
		q = q.Where("messages.created_at < ?", to) // ← calificado
	}
	if len(statuses) > 0 {
		q = q.Where("message_destinations.status IN ?", statuses)
	}
	if len(services) > 0 {
		q = q.Where("message_destinations.service IN ?", services)
	}

	var mf []types.MessageFilter
	if err := q.
		Order("messages.created_at DESC, message_destinations.id ASC").
		Limit(50).
		Scan(&mf).Error; err != nil { // Scan a tu DTO
		return nil, fmt.Errorf("Error data base")
	}
	return &mf, nil

}

func parseDate(s string, loc *time.Location) (time.Time, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return time.Time{}, false
	}
	if t, err := time.Parse(time.RFC3339, s); err == nil {
		return t, false
	}
	if d, err := time.ParseInLocation("2006-01-02", s, loc); err == nil {
		return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, loc), true
	}
	return time.Time{}, false
}

func normalizarConsulta(s string) []string {
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		v := strings.ToLower(strings.TrimSpace(p))
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
