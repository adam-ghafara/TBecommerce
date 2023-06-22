package kendaraancontroller

import (
	"autobros/libraries"
	"html/template"
	"net/http"
	"strconv"
	"time"

	"autobros/models"

	"autobros/entities"
)

var validation = libraries.NewValidation()
var kendaraanModel = models.NewModelKendaraan()

func Dashboard(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("web/servis/index.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, r)
}

func Tables(w http.ResponseWriter, r *http.Request) {

	tables, _ := kendaraanModel.LoadData()

	data := map[string]interface{}{
		"tabel": tables,
	}

	temp, err := template.ParseFiles("web/servis/tables.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Forms(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("web/servis/forms.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {

		r.ParseForm()

		var kendaraan entities.Kendaraan
		kendaraan.NamaPemilik = r.Form.Get("nama_pemilik")
		kendaraan.NamaKendaraan = r.Form.Get("nama_kendaraan")
		kendaraan.JenisID = r.Form.Get("id_jenis")
		kendaraan.NomorKendaraan = r.Form.Get("nomor_kendaraan")
		kendaraan.DetailServis = r.Form.Get("detail_servis")
		Tanggal := time.Now()
		TgltoStr := Tanggal.String
		kendaraan.TanggalServis = TgltoStr()
		kendaraan.StatusServis = "Pending"
		kendaraan.EstimasiWaktu = r.Form.Get("estimasi_waktu")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(kendaraan)

		if vErrors != nil {
			data["tables"] = kendaraan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Kendaraan berhasil dibuat"
			kendaraanModel.Create(kendaraan)
		}

		temp, _ := template.ParseFiles("web/servis/forms.html")
		temp.Execute(w, data)
	}

}

func FormAdmin(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		queryString := r.URL.Query()
		id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

		var kendaraan entities.Kendaraan
		kendaraanModel.GetData(id, &kendaraan)

		data := map[string]interface{}{
			"admins": kendaraan,
		}

		temp, err := template.ParseFiles("web/servis/form_admin.html")
		if err != nil {
			panic(err)
		}
		temp.Execute(w, data)

	} else if r.Method == http.MethodPost {

		r.ParseForm()

		var kendaraan entities.Kendaraan
		kendaraan.NamaPemilik = r.Form.Get("nama_pemilik")
		kendaraan.NamaKendaraan = r.Form.Get("nama_kendaraan")
		kendaraan.JenisID = r.Form.Get("id_jenis")
		kendaraan.NomorKendaraan = r.Form.Get("nomor_kendaraan")
		kendaraan.DetailServis = r.Form.Get("detail_servis")
		kendaraan.StatusServis = r.Form.Get("status_servis")

		var data = make(map[string]interface{})

		vErrors := validation.Struct(kendaraan)

		if vErrors != nil {
			data["admins"] = kendaraan
			data["validation"] = vErrors
		} else {
			data["pesan"] = "Data Kendaraan berhasil di update."
			kendaraanModel.Update(kendaraan)
		}
		temp, _ := template.ParseFiles("web/servis/form_admin.html")
		temp.Execute(w, data)
	}

}

func Details(w http.ResponseWriter, r *http.Request) {

	kendaraan := &entities.Kendaraan{}
	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	err := kendaraanModel.GetData(id, kendaraan)
	if err != nil {
		// Handle the error
		panic(err)
	}

	data := map[string]interface{}{
		"detail": kendaraan,
	}

	temp, err := template.ParseFiles("web/servis/detail.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, data)
}

func Delete(w http.ResponseWriter, r *http.Request) {

	queryString := r.URL.Query()
	id, _ := strconv.ParseInt(queryString.Get("id"), 10, 64)

	kendaraanModel.Delete(id)

	http.Redirect(w, r, "/autobros/tables", http.StatusSeeOther)
}
