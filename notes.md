Di Go, huruf awal (capital atau kecil) pada nama struct, field, fungsi, variabel, atau method menentukan visibilitas (aksesibilitas):

📌 Aturan Penulisan:
TODO: 1. Huruf besar (Capitalized / PascalCase)
Artinya Exported → bisa diakses dari package lain.
Contoh:
type Student struct {   // bisa dipakai di package lain
    Name  string        // bisa diakses di luar package
    Grade int           // bisa diakses di luar package
}

TODO: 2. Huruf kecil (camelCase / lowercase)
Artinya Unexported → hanya bisa diakses di dalam package yang sama.
Contoh:
type student struct {   // hanya bisa dipakai di package ini
    name  string        // tidak bisa diakses langsung dari luar
    grade int
}

TODO: Jadi, kapan pilih pointer vs value receiver?
TODO: Pointer receiver (*T):
Dipakai kalau method butuh mengubah data dalam struct.
Supaya tidak perlu copy struct besar (lebih efisien untuk struct besar).

TODO: Value receiver (T):
Cocok untuk struct kecil (misal cuma beberapa field).
Method tidak mengubah data struct.
