package models

import (
	"autobros/connection"
	"autobros/entities"
	"database/sql"
	"fmt"
)

type ModelKendaraan struct {
	c *sql.DB
}

func NewModelKendaraan() *ModelKendaraan {
	c, err := connection.CreateConnection()
	if err != nil {
		panic(err)
	}

	return &ModelKendaraan{
		c: c,
	}
}

func (p *ModelKendaraan) LoadData() ([]entities.Kendaraan, error) {

	rows, err := p.c.Query("select * from data_servis_kendaraan")
	if err != nil {
		return []entities.Kendaraan{}, err
	}
	defer rows.Close()

	var dataKendaraan []entities.Kendaraan
	for rows.Next() {
		var kendaraan entities.Kendaraan
		rows.Scan(&kendaraan.ID, &kendaraan.NamaPemilik, &kendaraan.NamaKendaraan, &kendaraan.JenisID, &kendaraan.NomorKendaraan, &kendaraan.DetailServis, &kendaraan.TanggalServis, &kendaraan.StatusServis, &kendaraan.EstimasiWaktu)
		if kendaraan.JenisID == "1" {
			kendaraan.JenisID = "Mobil"
		} else {
			kendaraan.JenisID = "Motor"
		}
		dataKendaraan = append(dataKendaraan, kendaraan)
	}

	return dataKendaraan, nil

}

func (p *ModelKendaraan) Create(kendaraan entities.Kendaraan) bool {

	result, err := p.c.Exec("INSERT INTO data_servis_kendaraan (nama_pemilik,nama_kendaraan,id_jenis,nomor_kendaraan,detail_servis,tanggal_servis,status_servis,estimasi_waktu) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)",
		kendaraan.NamaPemilik, kendaraan.NamaKendaraan, kendaraan.JenisID, kendaraan.NomorKendaraan, kendaraan.DetailServis, kendaraan.TanggalServis, kendaraan.StatusServis,kendaraan.EstimasiWaktu)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()
	return lastInsertId > 0
}

func (p *ModelKendaraan) GetData(id int64, kendaraan *entities.Kendaraan) error {

	return p.c.QueryRow("select id,nama_pemilik,nama_kendaraan, id_jenis, nomor_kendaraan,detail_servis,status_servis from data_servis_kendaraan where id = $1", id).Scan(
		&kendaraan.ID, &kendaraan.NamaPemilik, &kendaraan.NamaKendaraan, &kendaraan.JenisID, &kendaraan.NomorKendaraan, &kendaraan.DetailServis, &kendaraan.StatusServis)
}

func (p *ModelKendaraan) Update(kendaraan entities.Kendaraan) error {

	_, err := p.c.Exec(
		"update data_servis_kendaraan set nama_pemilik = $2, nama_kendaraan = $3, id_jenis = $4, nomor_kendaraan = $5, detail_servis = $6, status_servis = $7 where id = $1",
		kendaraan.NamaPemilik, kendaraan.NamaKendaraan, kendaraan.JenisID, kendaraan.NomorKendaraan, kendaraan.DetailServis, kendaraan.StatusServis)

	if err != nil {
		return err
	}

	return nil
}

func (p *ModelKendaraan) Delete(id int64) {

	p.c.Exec("delete from data_servis_kendaraan where id = $1", id)

}
