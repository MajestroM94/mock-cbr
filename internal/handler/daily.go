package handler

import (
	"mock-cbr/internal/service"
	"mock-cbr/internal/storage"
	"net/http"
	"time"
)

// GetDailyHandler godoc
// @Summary Получение XML курса валют
// @Description Генерирует XML по дате. Если date=error, возвращает 500
// @Tags currency
// @Produce xml
// @Param date query string true "Дата в формате DD-MM-YYYY"
// @Success 200 {string} string "Пример ответа: XML Response"
// @Failure 500 {string} string "Пример ответа: Internal Server Error"
// @Router /daily [get]
func GetDailyHandler(w http.ResponseWriter, r *http.Request) {
	date := r.URL.Query().Get("date")

	if date == "error" {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	xmlResponse, err := service.GenerateRandomXML(date)
	if err != nil {
		http.Error(w, "Failed to generate XML", http.StatusInternalServerError)
		return
	}

	storage.DB.Create(&storage.Request{
		DateParam:  date,
		XMLContent: xmlResponse,
		CreatedAt:  time.Now(),
	})

	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(xmlResponse))
}
