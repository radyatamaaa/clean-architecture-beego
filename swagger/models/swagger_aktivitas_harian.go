package models

type SwaggerProsesKunjunganSummarySuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       struct {
		Akuisisi struct{
			Procentage float64 `json:"procentage" example:"100"`
			CountKunjungan int `json:"count_kunjungan" example:"1"`
			CountMaxKunjungan int `json:"count_max_kunjungan" example:"1"`
			IsFinished bool `json:"is_finished" example:"true"`
		} `json:"akuisisi"`
		Upgrade struct{
			Procentage float64 `json:"procentage" example:"100"`
			CountKunjungan int `json:"count_kunjungan" example:"1"`
			CountMaxKunjungan int `json:"count_max_kunjungan" example:"1"`
			IsFinished bool `json:"is_finished" example:"true"`
		} `json:"upgrade"`
		Aktivasi struct{
			Procentage float64 `json:"procentage" example:"100"`
			CountKunjungan int `json:"count_kunjungan" example:"1"`
			CountMaxKunjungan int `json:"count_max_kunjungan" example:"1"`
			IsFinished bool `json:"is_finished" example:"true"`
		} `json:"aktivasi"`
		ModalKunjunganMingguLalu struct{
			CountKunjungan int `json:"count_kunjungan" example:"1"`
			CountRealisasi int `json:"count_realisasi" example:"1"`
			CountMaxKunjungan int `json:"count_max_kunjungan" example:"1"`
		} `json:"modal_kunjungan_minggu_lalu"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerPresentaseMingguanSummarySuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       *struct {
		Rekomendasi string `json:"permission" example:"Akusisi"`
		TotalKunjungan int `json:"total_kunjungan" example:"1"`
		TotalMaxKunjungan int `json:"total_max_kunjungan" example:"1"`
		TotalRealisasi int `json:"total_realisasi" example:"1"`
		JumlahKunjunganHarian []struct{
			Hari int `json:"hari" example:"1"`
			JumlahKunjungan int `json:"jumlah_kunjungan" example:"1"`
		} `json:"jumlah_kunjungan_harian"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerPresentaseBulananSummarySuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       *struct {
		Rekomendasi string `json:"permission" example:"Akusisi"`
		TotalKunjungan int `json:"total_kunjungan" example:"1"`
		TotalMaxKunjungan int `json:"total_max_kunjungan" example:"1"`
		TotalRealisasi int `json:"total_realisasi" example:"1"`
		JumlahKunjunganBulanan []struct{
			Bulan string `json:"bulan" example:"Jan"`
			IsSelected bool `json:"is_selected" example:"false"`
			JumlahKunjungan int `json:"jumlah_kunjungan" example:"1"`
		} `json:"jumlah_kunjungan_bulanan"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}


//request
type SwaggerRequestPresentaseMingguan struct {
	Rekomendasi string `json:"permission" example:"Akuisisi"`
	StartDate string `json:"start_date" example:"2021-03-09"`
	EndDate string `json:"end_date" example:"2021-03-13"`
}
type SwaggerRequestPresentaseBulanan struct {
	Rekomendasi string `json:"permission" example:"Akuisisi"`
	Month string `json:"month" example:"2021-03"`
}

