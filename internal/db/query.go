package db

import (
	"fmt"
	"meli_challenge_back/internal/entity"
	"strconv"
	"time"
)

const selectEventsQuery = `SELECT tp.description as type, ev.description, TO_CHAR(ev.created_at, 'HH24-MI-SS-DD-MM-YYYY') as created_at, co.description as country
	FROM challenge.tbl_event ev
	INNER JOIN challenge.tbl_country co ON ev.country_id = co.id
	INNER JOIN challenge.tbl_type tp ON ev.type_id = tp.id`

// Función para crear un nuevo evento
func CreateEvent(event entity.EventRequest, dbEntity entity.DataBase) error {

	db, err := GetConnection(dbEntity)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO challenge.tbl_event(description, created_at, type_id, country_id) VALUES ($1, $2, $3, $4)",
		event.Description, time.Now(), event.TypeId, event.CountryId)
	if err != nil {
		return err
	}

	return nil
}

// Función genérica para listar eventos de manera paginada según una consulta
func ListEventsWithQuery(query string, cfgDb entity.DataBase, args ...interface{}) ([]entity.EventResponse, error) {

	db, err := GetConnection(cfgDb)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []entity.EventResponse
	for rows.Next() {
		event := entity.EventResponse{}
		err := rows.Scan(&event.Type, &event.Description, &event.CreatedAt, &event.Country)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}

// Función para listar todos los eventos de manera paginada
func ListEvents(page, pageSize string, cfgDb entity.DataBase) ([]entity.EventResponse, error) {
	query := fmt.Sprintf("%s ORDER BY created_at LIMIT %s OFFSET %s", selectEventsQuery,
		pageSize, calculateOffset(page, pageSize))

	return ListEventsWithQuery(query, cfgDb)
}

// Función para listar eventos por país de manera paginada
func ListEventsByCountry(country, page, pageSize string, cfgDb entity.DataBase) ([]entity.EventResponse, error) {
	query := fmt.Sprintf("%s WHERE co.description LIKE $1 ORDER BY created_at LIMIT %s OFFSET %s", selectEventsQuery,
		pageSize, calculateOffset(page, pageSize))

	return ListEventsWithQuery(query, cfgDb, "%"+country+"%")
}

// Función para listar eventos por typo de manera paginada
func ListEventsByType(typeEvent, page, pageSize string, cfgDb entity.DataBase) ([]entity.EventResponse, error) {
	query := fmt.Sprintf("%s WHERE tp.description LIKE $1 ORDER BY created_at LIMIT %s OFFSET %s", selectEventsQuery,
		pageSize, calculateOffset(page, pageSize))

	return ListEventsWithQuery(query, cfgDb, "%"+typeEvent+"%")
}

// Función para listar  eventos por tipo de evento de manera paginada
func ListEventsByDescription(description, page, pageSize string, cfgDb entity.DataBase) ([]entity.EventResponse, error) {
	query := fmt.Sprintf("%s WHERE ev.description LIKE $1 ORDER BY created_at LIMIT %s OFFSET %s", selectEventsQuery,
		pageSize, calculateOffset(page, pageSize))

	return ListEventsWithQuery(query, cfgDb, "%"+description+"%")
}

// Obtener la lista de países desde la base de datos
func GetCountries(cfgDb entity.DataBase) ([]entity.Data, error) {
	query := "SELECT id, description FROM challenge.tbl_country"
	return ListDatasWithQuery(query, cfgDb)
}

// Obtener la lista de tipos de evento desde la base de datos
func GetEventTypes(cfgDb entity.DataBase) ([]entity.Data, error) {
	query := "SELECT id, description FROM challenge.tbl_type"
	return ListDatasWithQuery(query, cfgDb)
}

// Función genérica para listar
func ListDatasWithQuery(query string, cfgDb entity.DataBase, args ...interface{}) ([]entity.Data, error) {

	db, err := GetConnection(cfgDb)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var datas []entity.Data
	for rows.Next() {
		data := entity.Data{}
		err := rows.Scan(&data.ID, &data.Description)
		if err != nil {
			return nil, err
		}
		datas = append(datas, data)
	}

	return datas, nil
}

// Obtener las métricas de los países con mayor cantidad de eventos en los ultimos 3 meses
func GetTopCountriesMetrics(dbEntity entity.DataBase) ([]entity.CountryMetric, error) {
	db, err := GetConnection(dbEntity)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Calcula la fecha límite para los últimos 3 mes
	lastMonth := time.Now().AddDate(0, -3, 0)

	query := `SELECT co.description as country, COUNT(*) as event_count
	FROM challenge.tbl_event ev
	INNER JOIN challenge.tbl_country co ON ev.country_id = co.id
	WHERE ev.created_at >= $1
	GROUP BY co.description
	ORDER BY event_count DESC
	LIMIT 3`
	rows, err := db.Query(query, lastMonth)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var metrics []entity.CountryMetric
	for rows.Next() {
		metric := entity.CountryMetric{}
		err := rows.Scan(&metric.Country, &metric.EventCount)
		if err != nil {
			return nil, err
		}
		metrics = append(metrics, metric)
	}

	return metrics, nil
}

// Función auxiliar para calcular el desplazamiento (offset) de la paginación
func calculateOffset(page, pageSize string) string {
	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)
	offset := (pageInt - 1) * pageSizeInt
	return strconv.Itoa(offset)
}
