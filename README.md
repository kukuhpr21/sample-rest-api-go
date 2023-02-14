# SAMPLE REST API GO

## Penggunaan
1. Eksekusi file sql pada folder docs/db/[file sql] (ambil versi terakhir)
2. Copy dan rename file .env.example menjadi .env
3. Isi configurasi pada file .env sesuai kebutuhan
4. Lakukan perintah  **go mod tidy**
5. Running project anda **go run .**

## Endpoint
- Base URL **host:port/v1**
- Tambahkan [Authorization] [Bearer (access token)] pada header jika mengakses endpoint **products**
- Tambahkan [refresh-token] [(refresh_token)] pada header jika mengakses **auth/refresh-token**

| Method  | Path | Body | Response |
| ------- | ---- | ---- | -------- |
| POST    | auth/login  |{"email":"admin@gmail.com","password":"123456"} | {"code": 200,"status": "OK","data": {"access_token": "","refresh_token":""}} |
| GET     | auth/refresh-token  | |{"code": 200,"status": "OK","data": {"access_token": "","refresh_token":""}}|
| POST | products | {"name": ""} | {"code": 201,"status": "Created","data": {"id": "","name": ""}} |
| PUT | products/:id | {"name": ""} |{"code": 200,"status": "OK","data": {"id": "","name": ""}} |
| DELETE | products/:id | | {"code": 200,"status": "OK","data": {"id": ""}} |
| GET | products | | {"code": 200,"status": "OK","data": [{"id": "","name": ""}]} |
| GET | products/:id | | {"code": 200,"status": "OK","data": {"id": "","name": ""}} |