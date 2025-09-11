Di Go, huruf awal (capital atau kecil) pada nama struct, field, fungsi, variabel, atau method menentukan visibilitas (aksesibilitas):

📌 Aturan Penulisan:

1. Huruf besar (Capitalized / PascalCase)
Artinya Exported → bisa diakses dari package lain.
Contoh:
type Student struct {   // bisa dipakai di package lain
    Name  string        // bisa diakses di luar package
    Grade int           // bisa diakses di luar package
}

2. Huruf kecil (camelCase / lowercase)
Artinya Unexported → hanya bisa diakses di dalam package yang sama.
Contoh:
type student struct {   // hanya bisa dipakai di package ini
    name  string        // tidak bisa diakses langsung dari luar
    grade int
}
