package models

type SwaggerGetRekomendasiSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       struct {
		Data []*struct {
			NamaLengkap       *string  `json:"nama_lengkap" example:"NURHAYATI"`
			IDNumberMasking   *string  `json:"id_number_masking" example:"32701234********"`
			TargetRekomendasi *string  `json:"target_rekomendasi" example:"Akuisisi"`
			KelasAgen         *string  `json:"kelas_agen" example:"Calon Agen"`
			LatitudeAgen      *float64 `json:"latitude_agen" example:"-6.30051"`
			LongitudeAgen     *float64 `json:"longitude_agen" example:"106.827"`
			StatusKunjungan   *string  `json:"status_kunjungan" example:"Belum"`
			Jarak             *string  `json:"jarak" example:"0m"`
			KelurahanDomisili             *string `json:"kelurahan_domisili" example:"Ancol"`
			NomorTelepon      *string  `json:"nomor_telepon" example:"088798798"`
			Alamat            *string  `json:"alamat" example:"Bogor"`
			PnPengelola       *string  `json:"pn_pengelola" example:"09789789"`
		} `json:"data"`
		DataTerakhirUpdated *string `json:"data_terakhir_updated" example:"17 August 2021"`
		Paginator           struct {
			CurrentPage  int `json:"current_page" example:"1"`
			PerPage      int `json:"limit_per_page" example:"1"`
			PreviousPage int `json:"back_page" example:"1"`
			NextPage     int `json:"next_page" example:"2"`
			TotalRecords int `json:"total_records" example:"1"`
			TotalPages   int `json:"total_pages" example:"1"`
		} `json:"paginator"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerGetKelurahanSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       struct {
		Data []*struct {
			KodeFilter int `json:"kode_filter" example:"1"`
			Filter string `json:"filter" example:"Ancol"`
		} `json:"data"`
		Paginator           struct {
			CurrentPage  int `json:"current_page" example:"1"`
			PerPage      int `json:"limit_per_page" example:"1"`
			PreviousPage int `json:"back_page" example:"1"`
			NextPage     int `json:"next_page" example:"2"`
			TotalRecords int `json:"total_records" example:"1"`
			TotalPages   int `json:"total_pages" example:"1"`
		} `json:"paginator"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerGetSuggestionLocationSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       *struct {
		Data []*struct {
			NamaLengkap       *string  `json:"nama_lengkap" example:"NURHAYATI"`
			IDNumberMasking   *string  `json:"id_number_masking" example:"32701234********"`
			TargetRekomendasi *string  `json:"target_rekomendasi" example:"Akuisisi"`
			KelasAgen         *string  `json:"kelas_agen" example:"Calon Agen"`
			LatitudeAgen      *float64 `json:"latitude_agen" example:"-6.30051"`
			LongitudeAgen     *float64 `json:"longitude_agen" example:"106.827"`
			StatusKunjungan   *string  `json:"status_kunjungan" example:"Belum"`
			Jarak             *string  `json:"jarak" example:"0m"`
			KelurahanDomisili             *string `json:"kelurahan_domisili" example:"Ancol"`
			NomorTelepon      *string  `json:"nomor_telepon" example:"088798798"`
			Alamat            *string  `json:"alamat" example:"Bogor"`
			PnPengelola       *string  `json:"pn_pengelola" example:"09789789"`
		} `json:"data"`
		TotalRecord          int `json:"total_record" example:"1"`
		TotalSudahDiKunjungi int `json:"total_sudah_di_kunjungi" example:"0"`
		TotalBelumDiKunjungi int `json:"total_belum_di_kunjungi" example:"1"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerGetRekomendasiByIdNumberSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       *struct {
		ID                    int      `json:"id" example:"1"`
		NamaLengkap           *string  `json:"nama_lengkap" example:"NURHAYATI"`
		NomorTelepon          *string  `json:"nomor_telepon" example:"081112131415"`
		Cifno                 *string  `json:"cifno" example:"EK774B43"`
		IDNumber              *string  `json:"id_number" example:"3270123456789010"`
		IDNumberMasking       *string  `json:"id_number_masking" example:"32701234********"`
		TargetRekomendasi     *string  `json:"target_rekomendasi" example:"Akuisisi"`
		KelasAgen             *string  `json:"kelas_agen" example:"Calon Agen"`
		PnPengelola           *string  `json:"pn_pengelola" example:"00999999"`
		Alamat                *string  `json:"alamat" example:"Jl. Bukit Kencana III, RT.006/RW.019, Jatirahayu, Kec. Pd. Melati, Kota Bks, Jawa Barat 17414"`
		LatitudeAgen          *float64 `json:"latitude_agen" example:"-6.30051"`
		LongitudeAgen         *float64 `json:"longitude_agen" example:"106.827"`
		StatusKunjungan       *string  `json:"status_kunjungan" example:"Belum"`
		KodeKunjungan         *int     `json:"kode_kunjungan" example:"0"`
		StatusKetertarikan    *string  `json:"status_ketertarikan" example:"Belum"`
		KodeKetertarikan      *int     `json:"kode_ketertarikan" example:"0"`
		DeskripsiKetertarikan *string  `json:"deskripsi_ketertarikan" example:"null"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerSubmitKunjunganSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       *struct {
		ID                    int      `json:"id" example:"1"`
		NamaLengkap           *string  `json:"nama_lengkap" example:"NURHAYATI"`
		NomorTelepon          *string  `json:"nomor_telepon" example:"081112131415"`
		Cifno                 *string  `json:"cifno" example:"EK774B43"`
		IDNumber              *string  `json:"id_number" example:"3270123456789010"`
		IDNumberMasking       *string  `json:"id_number_masking" example:"32701234********"`
		TargetRekomendasi     *string  `json:"target_rekomendasi" example:"Akuisisi"`
		KelasAgen             *string  `json:"kelas_agen" example:"Calon Agen"`
		PnPengelola           *string  `json:"pn_pengelola" example:"00999999"`
		Alamat                *string  `json:"alamat" example:"Jl. Bukit Kencana III, RT.006/RW.019, Jatirahayu, Kec. Pd. Melati, Kota Bks, Jawa Barat 17414"`
		LatitudeAgen          *float64 `json:"latitude_agen" example:"-6.30051"`
		LongitudeAgen         *float64 `json:"longitude_agen" example:"106.827"`
		StatusKunjungan       *string  `json:"status_kunjungan" example:"Sudah"`
		KodeKunjungan         *int     `json:"kode_kunjungan" example:"1"`
		StatusKetertarikan    *string  `json:"status_ketertarikan" example:"Tertarik"`
		KodeKetertarikan      *int     `json:"kode_ketertarikan" example:"1"`
		DeskripsiKetertarikan *string  `json:"deskripsi_ketertarikan" example:"sangat membantu"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerSubmitKunjunganBulkSuccess struct {
	StatusCode int    `json:"status_code" example:"200"`
	Status     string `json:"status_desc" example:"OK"`
	Msg        string `json:"message" example:"Success"`
	Data       []*struct {
		ID                    int      `json:"id" example:"1"`
		NamaLengkap           *string  `json:"nama_lengkap" example:"NURHAYATI"`
		NomorTelepon          *string  `json:"nomor_telepon" example:"081112131415"`
		Cifno                 *string  `json:"cifno" example:"EK774B43"`
		IDNumber              *string  `json:"id_number" example:"3270123456789010"`
		IDNumberMasking       *string  `json:"id_number_masking" example:"32701234********"`
		TargetRekomendasi     *string  `json:"target_rekomendasi" example:"Akuisisi"`
		KelasAgen             *string  `json:"kelas_agen" example:"Calon Agen"`
		PnPengelola           *string  `json:"pn_pengelola" example:"00999999"`
		Alamat                *string  `json:"alamat" example:"Jl. Bukit Kencana III, RT.006/RW.019, Jatirahayu, Kec. Pd. Melati, Kota Bks, Jawa Barat 17414"`
		LatitudeAgen          *float64 `json:"latitude_agen" example:"-6.30051"`
		LongitudeAgen         *float64 `json:"longitude_agen" example:"106.827"`
		StatusKunjungan       *string  `json:"status_kunjungan" example:"Sudah"`
		KodeKunjungan         *int     `json:"kode_kunjungan" example:"1"`
		StatusKetertarikan    *string  `json:"status_ketertarikan" example:"Tertarik"`
		KodeKetertarikan      *int     `json:"kode_ketertarikan" example:"1"`
		DeskripsiKetertarikan *string  `json:"deskripsi_ketertarikan" example:"sangat membantu"`
	} `json:"data"`
	Errors *string `json:"errors" example:"null"`
}

type SwaggerDetailInformationSuccess struct {
	Aktivasi struct {
		StatusCode int    `json:"status_code" example:"200"`
		Status     string `json:"status_desc" example:"OK"`
		Msg        string `json:"message" example:"Success"`
		Data       *struct {
			Card *struct {
				StatusKunjungan   *string `json:"status_kunjungan" example:"Sudah"`
				NamaLengkap       *string `json:"nama_lengkap" example:"AKBAR RIZKY"`
				IdNumberMasking   *string `json:"id_number_masking" example:"32701234********"`
				TargetRekomendasi *string `json:"target_rekomendasi" example:"Aktivasi"`
			} `json:"card"`
			Score *struct {
				Score                 *float64 `json:"score" example:"789"`
				KategoriRekomendasi   *string  `json:"kategori_rekomendasi" example:"Sangat Berpotensi"`
				KeteranganRekomendasi *string  `json:"keterangan_rekomendasi" example:"Agen ini masuk PRIORITAS 1 untuk dapat ditingkatkan transaksinya karena sales volume agen sudah cukup dan frekuensi transaksi kurang 624"`
				TargetAktivasi        *string  `json:"target_aktivasi" example:"Agen BEP"`
				TerakhirKunjungan     *string  `json:"terakhir_kunjungan" example:"6 September 2021"`
				StatusAktivasi        *string  `json:"status_aktivasi" example:"Belum Diaktivasi"`
			} `json:"score"`
			InformasiPersonal *struct {
				LamaMenjadiAgen    *string  `json:"lama_menjadi_agen" example:"1 Tahuni"`
				TanggalMenjadiAgen *string  `json:"tanggal_menjadi_agen" example:"1 March 2020"`
				IdAgen             *string  `json:"id_agen" example:"10064536"`
				MidCode            *string  `json:"mid_code" example:"1314960000"`
				TidCode            *string  `json:"tid_code" example:"21345708"`
				NomorTelepon       *string  `json:"nomor_telepon" example:"081112131416"`
				Alamat             *string  `json:"alamat" example:"Tlajung Udik, Gunung Putri, Bogor, West Java 16962"`
				LatitudeAgen       *float64 `json:"latitude_agen" example:"-6.447828249483351"`
				LongitudeAgen      *float64 `json:"longitude_agen" example:"106.90567042647541"`
			} `json:"informasi_personal"`
		} `json:"data"`
		Errors *string `json:"errors" example:"null"`
	} `json:"aktivasi"`
	Akuisisi struct {
		StatusCode int    `json:"status_code" example:"200"`
		Status     string `json:"status_desc" example:"OK"`
		Msg        string `json:"message" example:"Success"`
		Data       *struct {
			Card *struct {
				StatusKunjungan   *string `json:"status_kunjungan" example:"Belum"`
				NamaLengkap       *string `json:"nama_lengkap" example:"NURHAYATI"`
				IdNumberMasking   *string `json:"id_number_masking" example:"32701234********"`
				TargetRekomendasi *string `json:"target_rekomendasi" example:"Akuisisi"`
			} `json:"card"`
			Score *struct {
				Score                 *float64 `json:"score" example:"334"`
				KategoriRekomendasi   *string  `json:"kategori_rekomendasi" example:"Sangat Tidak Berpotensi"`
				KeteranganRekomendasi *string  `json:"keterangan_rekomendasi" example:"Agen ini masuk PRIORITAS 5 untuk dapat ditingkatkan transaksinya karena sales volume kurang 1009511499.00 dan frekuensi transaksi kurang 1887"`
				TerakhirKunjungan     *string  `json:"terakhir_kunjungan" example:"3 September 2021"`
			} `json:"score"`
			InformasiPersonal *struct {
				NomorTelepon  *string  `json:"nomor_telepon" example:"081112131415"`
				Alamat        *string  `json:"alamat" example:" Tlajung Udik, Gunung Putri, Bogor, West Java 16962"`
				LatitudeAgen  *float64 `json:"latitude_agen" example:"-6.4481454132681835"`
				LongitudeAgen *float64 `json:"longitude_agen" example:"106.90576028047671"`
			} `json:"informasi_personal"`
		} `json:"data"`
		Errors *string `json:"errors" example:"null"`
	} `json:"akuisisi"`
	Upgrade struct {
		StatusCode int    `json:"status_code" example:"200"`
		Status     string `json:"status_desc" example:"OK"`
		Msg        string `json:"message" example:"Success"`
		Data       *struct {
			KelasAgen *string `json:"kelas_agen" example:"Juragan"`
			Card      *struct {
				StatusKunjungan   *string `json:"status_kunjungan" example:"Sudah"`
				NamaLengkap       *string `json:"nama_lengkap" example:"RIZBAR ANANDA"`
				IdNumberMasking   *string `json:"id_number_masking" example:"32751112********"`
				TargetRekomendasi *string `json:"target_rekomendasi" example:"Upgrade"`
			} `json:"card"`
			Score *struct {
				Score                 *float64 `json:"score" example:"436"`
				KategoriRekomendasi   *string  `json:"kategori_rekomendasi" example:"Tidak Berpotensi"`
				KeteranganRekomendasi *string  `json:"keterangan_rekomendasi" example:"Agen ini masuk PRIORITAS 4 untuk dapat ditingkatkan transaksinya karena sales volume kurang 5145558526.00 dan frekuensi transaksi kurang 6333"`
				TargetUpgrade         *string  `json:"target_upgrade" example:"Juragan"`
				TerakhirKunjungan     *string  `json:"terakhir_kunjungan" example:"6 September 2021"`
			} `json:"score"`
			InformasiPersonal *struct {
				IdAgen        *string  `json:"id_agen" example:"10067301"`
				MidCode       *string  `json:"mid_code" example:"1326911000"`
				TidCode       *string  `json:"tid_code" example:"21151073"`
				NomorTelepon  *string  `json:"nomor_telepon" example:"081112131417"`
				Alamat        *string  `json:"alamat" example:"Jl. Tlajung Udik, RT.03/RW.07, Tlajung Udik, Kec. Gn. Putri, Bogor, Jawa Barat 16962"`
				LatitudeAgen  *float64 `json:"latitude_agen" example:"-6.447912204623168"`
				LongitudeAgen *float64 `json:"longitude_agen" example:"106.9065273922615"`
			} `json:"informasi_personal"`
		} `json:"data"`
		Errors *string `json:"errors" example:"null"`
	} `json:"upgrade"`
}

//error
type SwaggerSubmitKunjunganErrorConflict struct {
	StatusCode int       `json:"status_code" example:"409"`
	Status     string    `json:"status_desc" example:"Conflict"`
	Msg        string    `json:"message" example:"Rekomendasi Sudah Di Kunjungi"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"Your Item already exist or duplicate"`
}

type SwaggerGetRekomendasiByIdNumberErrorNotFound struct {
	StatusCode int       `json:"status_code" example:"404"`
	Status     string    `json:"status_desc" example:"Not Found"`
	Msg        string    `json:"message" example:"Rekomendasi Tidak Di temukan"`
	Data       *struct{} `json:"data" example:""`
	Errors     *string   `json:"errors" example:"Your requested Item is not found"`
}

type SwaggerGetRekomendasiByIdNumberErrorNotFoundAndNotFoundRoutes struct {
	ErrorNotFoundGetRekomendasiByIdNumber SwaggerGetRekomendasiByIdNumberErrorNotFound `json:"error_not_found_get_rekomendasi_by_id_number"`
	ErrorNotFoundRoutes                   SwaggerErrorNotFoundRoutes                   `json:"error_not_found_routes"`
}

//request
type SwaggerSubmitKunjunganRequest struct {
	IdNumber              *string `json:"id_number" example:"327511121314152000"`
	Ketertarikan          *int    `json:"ketertarikan" example:"1"`
	DeskripsiKetertarikan *string `json:"deskripsi_ketertarikan" example:"Sangat membantu"`
}
type SwaggerSubmitKunjunganBulkRequest struct {
	Submitted SwaggerSubmitKunjunganRequest `json:"submitted"`
}
type SwaggerRequestDetailInformation struct {
	IdRekomendasi     *int    `json:"id_rekomendasi" example:"30"`
	TargetRekomendasi *string `json:"target_rekomendasi" example:"Aktivasi"`
}

type SwaggerGetKelurahanRequest struct {
	Page              int      `json:"page" example:"1"`
	Limit             int      `json:"limit" example:"10"`
	Search            string   `json:"search" example:"Ancol"`
}