package entities

type Kendaraan struct {
	ID             int32
	NamaPemilik    string `validate:"required" label:"nama_pemilik"`
	NamaKendaraan  string `validate:"required" label:"nama_kendaraan"`
	JenisID        string `label:"id_jenis"`
	NomorKendaraan string `label:"nomor_kendaraan"`
	DetailServis   string `validate:"required" label:"detail_servis"`
	TanggalServis  string `label:"tanggal_servis"`
	StatusServis   string `label:"status_servis"`
	EstimasiWaktu  string `label:"estimasi_waktu"`
}
