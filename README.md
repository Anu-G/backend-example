# livecode-gorm-wmb

## Requirement
Aplikasi dapat melakukan:

    Management data master menu (CRUD) 
    Management data master menu price (CRUD)
    Management data master table (CRUD)
    Management data master trans type (CRUD)
    Management data master discount (CRUD)
    Melakukan customer registration
    Melakukan aktivasi member customer yang sudah terdaftar sekaligus memberikan privilege discount
    Melakukan transaksi penjualan dengan validasi apabila meja sudah dipakai tidak bisa dibuat bill
    Mencetak bill berdasarkan bill_id, sekaligus melakukan update meja menjadi available
    Memberikan informasi total penjualan harian

Output
Dibuatkan urutan untuk melakukan run aplikasi boleh di app.go atau dibuatkan file tersendiri, dengan hasil sebagai berikut:

++++++ Soal No 1 ++++++

> Create Customer

> List / Get Customer

> dst…

## Run DB Migrate
``` go run . db:migrate ```

## Run DB Seeds
``` go run . db:seeds ```
