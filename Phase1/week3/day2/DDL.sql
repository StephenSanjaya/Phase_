-- CREATE TABLE table_name (
--     column_name data_type constraint,
--     column_name data_type constraint
-- )

/*
Data Type :
- Varchar -> text, mirip seperti String. Max daya tampung cuman 255 karakter
- Text    -> daya tampungnya lebih besar
- Integer
- Double
- DATE      -> ini cuman menyimpan YYYY-MM-DD
- DATETIME  -> ini cuman menyimpan YYYY-MM-DD HH:MM:SS


Constraint -> RULES
- PRIMARY KEY : nilai unik yang tidak boleh ada samanya dalam 1 table DAN WAJIB TERISI
- UNIQUE      : nilai unik yang tidak boleh ada samanya dalam 1 table TETAPI TIDAK WAJIB TERISI
- NOT NULL    : memberitahu bahwa column tersebut WAJIB DIISI (TIDAK BOLEH KOSONG)
- AUTO INCREMENT (SERIAL) : meningkatkan nilai (default nya +1)
- FOREIGN KEY : column yang menghubungkan 2 buah table
*/

| Judul Buku     | Penulis       | Peminjam     | Tanggal Peminjaman | ID |
|----------------|---------------|--------------|--------------------| -- |
| Database 101   | John Doe      | Alice        | 2021-01-01         | 1  |
| Database 101   | John Doe      | Bob          | 2021-01-02         | 2  |
| Cooking for All| Jane Smith    | Alice        | 2021-01-03         | 3  |


CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(100)
)

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username varchar(100) NOT NULL,
    password varchar(100)
)

insert into users (username, "password") values 
('panji', 'testong'),
('budi', 'testong')

select * from users u 

/*
Indexing : digunakan untuk meningkatkan kecepatan dalam menampilkan data

CREATE INDEX index_name ON table_name(column_name)
*/

CREATE INDEX idx_username ON users(username)

2 Teknologi :
1. Elastic Search
2. Redis

punya 1.000.000 data
- tanpa index -> 1 detik
- index -> 0.5 detik
- redis -> 0.1 detik
---------------------------------

Untuk ngambil data dari beberapa table :
1. JOIN -> efeknya berat dan bisa jadi lebih lama prosesnya
2. VIEWS

---------------------------------

table mahasiswa
id | nama mhs | semester | id mata kuliah |
1  | panji    | 2        | 1              |
2  | budi     | 4        | 2              |




table mata kuliah
id | mata kuliah | sks | dosen pengajar
1  | database    | 3   | anto
2  | agama       | 2   | tini

CREATE TABLE mata_kuliah (
    id serial primary key,
    mata_kuliah varchar(100),
    sks int(2),
    dosen_pengajar varchar(100)
);

CREATE TABLE mahasiswa (
    id serial primary key,
    nama_mhs varchar(100),
    semester int(2),
    id_mata_kuliah int(2),
    -- FOREIGN KEY (column_name) REFERENCES table_name(column_name)
    FOREIGN KEY (id_mata_kuliah) REFERENCES mata_kuliah(id)
);

-----------------------------

| Nama Peserta | Email            | Sesi Diikuti       | Lokasi Sesi |
|--------------|------------------|--------------------|-------------|
| Alice        | alice@email.com  | Sesi 1, Sesi 2     | Ruang A, Ruang B |
| Bob          | bob@email.com    | Sesi 3             | Ruang C     |
| Charlie      | charlie@email.com| Sesi 2, Sesi 3     | Ruang B, Ruang C |



| Nama Peserta | Email            | Sesi Diikuti | Lokasi Sesi | ID Peserta |
|--------------|------------------|--------------|-------------|------------|
| Alice        | alice@email.com  | Sesi 1       | Ruang A     | 1
| Alice        | alice@email.com  | Sesi 2       | Ruang B     | 1
| Bob          | bob@email.com    | Sesi 3       | Ruang C     | 2
| Charlie      | charlie@email.com| Sesi 2       | Ruang B     | 3
| Charlie      | charlie@email.com| Sesi 3       | Ruang C     | 3

--------------------
Advance Query :

1. Concat -> Menggabungkan 2 buah / lebih string