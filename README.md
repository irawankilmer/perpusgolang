# Perpustakaan Online Sederhana
## Desain Database
```bash
Users

ID: UUID (Primary Key)
Name: VARCHAR(255)
Email: VARCHAR(255) (Unique)
PasswordHash: VARCHAR(255)
Role: ENUM('admin', 'librarian', 'member')
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Profiles

ID: UUID (Primary Key)
UserID: UUID (Foreign Key)
Address: VARCHAR(255)
PhoneNumber: VARCHAR(20)
DateOfBirth: DATE
Verified: BOOLEAN (default FALSE)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Foreign Key (UserID) References Users(ID)
Authors

ID: UUID (Primary Key)
Name: VARCHAR(255)
Biography: TEXT
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Publishers

ID: UUID (Primary Key)
Name: VARCHAR(255)
Address: VARCHAR(255)
ContactNumber: VARCHAR(20)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Categories

ID: UUID (Primary Key)
Name: VARCHAR(255)
Description: TEXT
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Books

ID: UUID (Primary Key)
Title: VARCHAR(255)
AuthorID: UUID (Foreign Key)
PublisherID: UUID (Foreign Key)
PublishedDate: DATE
ISBN: VARCHAR(13) (Unique)
CategoryID: UUID (Foreign Key)
Description: TEXT
CoverImageURL: VARCHAR(255)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Inventory

ID: UUID (Primary Key)
BookID: UUID (Foreign Key)
LibraryID: UUID (Foreign Key)
TotalCopies: INT
AvailableCopies: INT
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Borrows

ID: UUID (Primary Key)
UserID: UUID (Foreign Key)
BookID: UUID (Foreign Key)
BorrowDate: DATE
ReturnDate: DATE
Status: ENUM('borrowed', 'returned', 'overdue')
Fine: DECIMAL(10, 2)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Reservations

ID: UUID (Primary Key)
UserID: UUID (Foreign Key)
BookID: UUID (Foreign Key)
ReservationDate: DATE
Status: ENUM('reserved', 'cancelled')
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Libraries

ID: UUID (Primary Key)
Name: VARCHAR(255)
Address: VARCHAR(255)
ContactNumber: VARCHAR(20)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Notifications

ID: UUID (Primary Key)
UserID: UUID (Foreign Key)
Message: TEXT
Status: ENUM('unread', 'read')
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Transactions

ID: UUID (Primary Key)
BorrowID: UUID (Foreign Key)
Type: ENUM('borrow', 'return')
Date: TIMESTAMP
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Reviews

ID: UUID (Primary Key)
BookID: UUID (Foreign Key)
UserID: UUID (Foreign Key)
Rating: INTEGER (1 to 5)
Comment: TEXT
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
Fines

ID: UUID (Primary Key)
DaysLate: INT
FineAmount: DECIMAL(10, 2)
CreatedAt: TIMESTAMP
UpdatedAt: TIMESTAMP
```

## Endpoint API
```bash
Autentikasi dan Profil Pengguna
Register User

Method: POST
URL: /api/register
Description: Mendaftarkan pengguna baru.
Roles: Public
Login User

Method: POST
URL: /api/login
Description: Login untuk mendapatkan token JWT.
Roles: Public
Get User Profile

Method: GET
URL: /api/profile
Description: Mendapatkan profil pengguna yang sedang login.
Roles: Admin, Librarian, Member
Update User Profile

Method: PUT
URL: /api/profile
Description: Memperbarui profil pengguna yang sedang login.
Roles: Admin, Librarian, Member
Manajemen Buku
Create Book

Method: POST
URL: /api/books
Description: Menambahkan buku baru ke dalam sistem.
Roles: Admin, Librarian
Get All Books

Method: GET
URL: /api/books
Description: Mendapatkan daftar semua buku.
Roles: Admin, Librarian, Member
Get Book by ID

Method: GET
URL: /api/books/{id}
Description: Mendapatkan detail buku berdasarkan ID.
Roles: Admin, Librarian, Member
Update Book

Method: PUT
URL: /api/books/{id}
Description: Memperbarui detail buku.
Roles: Admin, Librarian
Delete Book

Method: DELETE
URL: /api/books/{id}
Description: Menghapus buku dari sistem.
Roles: Admin, Librarian
Manajemen Peminjaman dan Pengembalian
Borrow Book

Method: POST
URL: /api/borrows
Description: Meminjam buku.
Roles: Admin, Librarian, Member (hanya jika profil sudah terverifikasi)
Return Book

Method: POST
URL: /api/returns
Description: Mengembalikan buku.
Roles: Admin, Librarian, Member
Get Borrowed Books

Method: GET
URL: /api/borrows
Description: Mendapatkan daftar buku yang sedang dipinjam oleh pengguna.
Roles: Admin, Librarian, Member
Manajemen Denda
Get Fines

Method: GET
URL: /api/fines
Description: Mendapatkan daftar denda yang belum dibayar.
Roles: Admin, Librarian, Member
Pay Fine

Method: POST
URL: /api/fines/pay
Description: Membayar denda.
Roles: Admin, Librarian, Member
Manajemen Reservasi
Reserve Book

Method: POST
URL: /api/reservations
Description: Reservasi buku.
Roles: Admin, Librarian, Member
Get Reservations

Method: GET
URL: /api/reservations
Description: Mendapatkan daftar reservasi buku.
Roles: Admin, Librarian, Member
Cancel Reservation

Method: DELETE
URL: /api/reservations/{id}
Description: Membatalkan reservasi buku.
Roles: Admin, Librarian, Member
Manajemen Kategori dan Penulis
Create Category

Method: POST
URL: /api/categories
Description: Menambahkan kategori buku baru.
Roles: Admin, Librarian
Get All Categories

Method: GET
URL: /api/categories
Description: Mendapatkan daftar semua kategori buku.
Roles: Admin, Librarian, Member
Create Author

Method: POST
URL: /api/authors
Description: Menambahkan penulis baru.
Roles: Admin, Librarian
Get All Authors

Method: GET
URL: /api/authors
Description: Mendapatkan daftar semua penulis.
Roles: Admin, Librarian, Member

Get Author by ID

Method: GET
URL: /api/authors/{id}
Description: Mendapatkan detail penulis berdasarkan ID.
Roles: Admin, Librarian, Member
Update Author

Method: PUT
URL: /api/authors/{id}
Description: Memperbarui detail penulis.
Roles: Admin, Librarian
Delete Author

Method: DELETE
URL: /api/authors/{id}
Description: Menghapus penulis dari sistem.
Roles: Admin, Librarian
Manajemen Penerbit
Create Publisher

Method: POST
URL: /api/publishers
Description: Menambahkan penerbit baru.
Roles: Admin, Librarian
Get All Publishers

Method: GET
URL: /api/publishers
Description: Mendapatkan daftar semua penerbit.
Roles: Admin, Librarian, Member
Get Publisher by ID

Method: GET
URL: /api/publishers/{id}
Description: Mendapatkan detail penerbit berdasarkan ID.
Roles: Admin, Librarian, Member
Update Publisher

Method: PUT
URL: /api/publishers/{id}
Description: Memperbarui detail penerbit.
Roles: Admin, Librarian
Delete Publisher

Method: DELETE
URL: /api/publishers/{id}
Description: Menghapus penerbit dari sistem.
Roles: Admin, Librarian
Manajemen Profil Pengguna
Get All Users

Method: GET
URL: /api/users
Description: Mendapatkan daftar semua pengguna.
Roles: Admin, Librarian
Get User by ID

Method: GET
URL: /api/users/{id}
Description: Mendapatkan detail pengguna berdasarkan ID.
Roles: Admin, Librarian
Update User

Method: PUT
URL: /api/users/{id}
Description: Memperbarui detail pengguna.
Roles: Admin, Librarian
Delete User

Method: DELETE
URL: /api/users/{id}
Description: Menghapus pengguna dari sistem.
Roles: Admin
Manajemen Pustakawan
Create Librarian

Method: POST
URL: /api/librarians
Description: Menambahkan pustakawan baru.
Roles: Admin
Get All Librarians

Method: GET
URL: /api/librarians
Description: Mendapatkan daftar semua pustakawan.
Roles: Admin
Get Librarian by ID

Method: GET
URL: /api/librarians/{id}
Description: Mendapatkan detail pustakawan berdasarkan ID.
Roles: Admin
Update Librarian

Method: PUT
URL: /api/librarians/{id}
Description: Memperbarui detail pustakawan.
Roles: Admin
Delete Librarian

Method: DELETE
URL: /api/librarians/{id}
Description: Menghapus pustakawan dari sistem.
Roles: Admin
Manajemen Kategori
Create Category

Method: POST
URL: /api/categories
Description: Menambahkan kategori buku baru.
Roles: Admin, Librarian
Get All Categories

Method: GET
URL: /api/categories
Description: Mendapatkan daftar semua kategori buku.
Roles: Admin, Librarian, Member
Manajemen Denda
Get Fines

Method: GET
URL: /api/fines
Description: Mendapatkan daftar denda yang belum dibayar.
Roles: Admin, Librarian
Pay Fine

Method: POST
URL: /api/fines/pay
Description: Membayar denda.
Roles: Admin, Librarian, Member
Manajemen Reservasi
Reserve Book

Method: POST
URL: /api/reservations
Description: Reservasi buku.
Roles: Admin, Librarian, Member
Get Reservations

Method: GET
URL: /api/reservations
Description: Mendapatkan daftar reservasi buku.
Roles: Admin, Librarian, Member
Cancel Reservation

Method: DELETE
URL: /api/reservations/{id}
Description: Membatalkan reservasi buku.
Roles: Admin, Librarian, Member
```