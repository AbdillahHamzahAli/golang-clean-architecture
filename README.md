# Golang Clean Architecture REST API

Proyek ini adalah implementasi REST API menggunakan pola Clean Architecture di Golang. Stack utama:

- Gin sebagai HTTP framework (`github.com/gin-gonic/gin`)
- GORM sebagai ORM (`gorm.io/gorm`, `gorm.io/driver/postgres`)
- PostgreSQL sebagai database
- JWT untuk otentikasi

## Struktur Proyek

```text
c:/My App/golang-clean-architecture/
├─ cmd/
│  ├─ web/
│  │  └─ main.go                # Entry point aplikasi HTTP
│  └─ command.go                # Menjalankan perintah CLI (migrasi/seed, dll)
├─ internal/
│  ├─ config/
│  │  └─ gorm.go                # Koneksi database (env → DSN Postgres)
│  ├─ delivery/
│  │  └─ http/
│  │     ├─ user_controller.go  # Controller Auth/User (Register, Login)
│  │     ├─ post_controller.go  # Controller Post (CRUD by id)
│  │     ├─ route/
│  │     │  ├─ user_route.go    # Route auth
│  │     │  └─ post_route.go    # Route post + middleware
│  │     └─ middleware/
│  │        ├─ middleware.go    # Constructor Middleware(jwtUC, db)
│  │        └─ auth_middleware.go# Auth Bearer JWT
│  ├─ infrastructure/
│  │  └─ pgsql/
│  │     └─ post_pg.go          # Implementasi repository Post (GORM)
│  ├─ domain/
│  │  ├─ entity/                # Entitas domain (Post, User, ...)
│  │  └─ dto/                   # DTO request/response
│  ├─ usecase/                  # Bisnis use case (Auth, Post, Jwt)
│  └─ repository/               # Kontrak repository (interface)
├─ db/
│  └─ migrations/               # Migrasi/Seeder (opsional)
├─ .env                         # Variabel lingkungan (jangan commit)
├─ go.mod, go.sum
└─ README.md
```

## Clean Architecture (Ringkas)

- **Delivery (Interface Adapter)**: layer HTTP (`internal/delivery/http`) berisi controller, routing, dan middleware.
- **Usecase (Application Business Rules)**: `internal/usecase/` mengorkestrasi alur bisnis (auth, post).
- **Domain (Enterprise Business Rules)**: `internal/domain/` menyimpan `entity/` dan `dto/` yang agnostik terhadap framework.
- **Infrastructure**: `internal/infrastructure/pgsql/` sebagai implementasi repository menggunakan GORM & Postgres.
- **Repository (Abstraction)**: `internal/repository/` mendefinisikan kontrak yang diimplementasikan oleh infrastructure.

Ketergantungan mengalir dari luar ke dalam: Delivery → Usecase → Domain. Infrastructure mengimplementasi `repository` dan diinject ke usecase.

## Persiapan

1. Instal Go (versi sesuai `go.mod`).
2. Siapkan PostgreSQL.
3. Buat file `.env` di root dengan variabel berikut (contoh):

```env
ENV=development
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=my_db
```

`internal/config/gorm.go` akan memuat `.env` jika `ENV != "production"` dan membangun DSN Postgres.

## Menjalankan Aplikasi

```bash
# Jalankan HTTP server di :8080
go run cmd/web/main.go
```

Aplikasi menggunakan Gin default logger & recovery. Jika Anda menjalankan dengan argumen tambahan, `main.go` akan mengalihkan ke `cmd.Commands(db)` (mis. untuk migrasi/seed jika sudah diimplementasi).

## Otentikasi

- Menggunakan Bearer JWT.
- Middleware `Auth(m *Middleware)` di `internal/delivery/http/middleware/auth_middleware.go` memvalidasi header `Authorization: Bearer <token>` dan menyetel `user_id` ke context.
- Endpoint Post membutuhkan Bearer token yang valid.

## Endpoint

Base URL: `http://localhost:8080`

- **Auth** (`internal/delivery/http/user_controller.go` + `route/user_route.go`)
  - POST `/auth/register`
  - POST `/auth/login`

- **Post** (`internal/delivery/http/post_controller.go` + `route/post_route.go`)
  - POST `/post`
  - PUT `/post/:id`
  - DELETE `/post/:id`
  - GET `/post/:id`

Middleware JWT diterapkan pada semua route Post.

## Contoh Request

- Register

```bash
curl -X POST http://localhost:8080/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret"
  }'
```

- Login (ambil token)

```bash
curl -X POST http://localhost:8080/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "password": "secret"
  }'
```

Respons login diharapkan mengandung token JWT. Gunakan token ini untuk endpoint Post berikut.

- Create Post

```bash
curl -X POST http://localhost:8080/post \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Judul",
    "content": "Konten"
  }'
```

- Update Post

```bash
curl -X PUT http://localhost:8080/post/<POST_ID_UUID> \
  -H "Authorization: Bearer <JWT_TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Judul Baru",
    "content": "Konten Baru"
  }'
```

- Delete Post

```bash
curl -X DELETE http://localhost:8080/post/<POST_ID_UUID> \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

- Get Post By ID

```bash
curl -X GET http://localhost:8080/post/<POST_ID_UUID> \
  -H "Authorization: Bearer <JWT_TOKEN>"
```

## Catatan Implementasi Penting

- Pada repository Post (`internal/infrastructure/pgsql/post_pg.go`) pencarian dan penghapusan by ID menggunakan parameterized query:
  - `Where("id = ?", id).First(&post)`
  - `Delete(&entity.Post{}, "id = ?", id)`
  Hal ini menghindari error Postgres saat menggunakan UUID.

- `internal/config/gorm.go` memuat `.env` (non-production) dan membuka koneksi Postgres dengan TimeZone `Asia/Jakarta`.

## Pengembangan

- Tambahkan validasi DTO pada layer Delivery/Usecase sesuai kebutuhan.
- Lengkapi logging, error handling, dan pengujian unit pada Usecase.
- Tambahkan migrasi database di `db/migrations/` dan jalankan via perintah CLI di `cmd/` bila diperlukan.

## Lisensi

Bebas digunakan untuk pembelajaran dan pengembangan.
