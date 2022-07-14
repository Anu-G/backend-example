# livecode-gorm-wmb

## Requirement
[LIVECODE] WMB 2.0
1. Untuk memberikan opsi metoda bayar dan memudahkan pembayaran pelanggannya, Warung Makan Bahari akan bekerja sama dengan salah satu payment gateway Lopei. Tim Software Developer WMB diminta untuk melakukan pengembangan aplikasi yang sekarang untuk bisa terintegrasi dengan Lopei Payment Gateway. 
   Alur proses pembayaran yang baru adalah 
     - pelanggan ditawarkan opsi pembayaran, tunai atau Lopei (tidak bisa split payment, pilih salah satu), 
     - Untuk pembayaran Lopei, pelanggan diminta memasukkan ID yang terdaftar kemudian masukan jumlah tagihan makan.
     - Setelah sukses pembayaran, struk akan tercetak, tabel meja akan diubah status nya menjadi kosong
   Catatan : Gunakan gRPC Lopei Server sebagai server payment gateway, tambahkan table untuk kebutuhan informasi pembayaran di API WMB POS, lalu integrasikan gRPC client nya

2. Top Manajemen WMB juga sedang mempertimbangkan untuk membuat aplikasi mobile Self Service POS untuk pelanggannya, sehingga proses transaksi akan semakin mudah. Oleh karena itu untuk meningkatkan keamanan API WMB POS,  Implementasikan JWT untuk memproteksi API

3. Implementasikan Unit Testing untuk API WMB POS

repo: livecode-wmb-2
deadline: Kamis, 14 Juli 2022 | 21:00 WIB

## Key Notes
``` 
1. protobuf file adjustmen
int32 lopeiId -> string lopeiId
float amount -> double amount

2. grpc server data type adjustment
type Customer struct {
	LopeiId 	string
	Balance     float64
}

3. server repository addition
return error if id/phoneNumber not found
```

## ENV
``` change .env.example to .env ```

## Run DB Migrate
``` go run . db:migrate ```

## Run DB Seeds
``` go run . db:seeds ```

## API Spec

### Register User
- Request: POST
- Endpoint : `/auth/register`
- Body : 
```json
{
    "user_name":"angga21",
    "user_password":"passwordangga",
    "customer_name":"Angga",
    "mobile_phone_no":"081245340921",
    "email":"angga@mail.com"
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": {
        "ID": 6,
        "CreatedAt": "2022-07-14T23:14:07.852029088+07:00",
        "UpdatedAt": "2022-07-14T23:14:07.852029088+07:00",
        "DeletedAt": null,
        "CustomerName": "Angga",
        "MobilePhoneNo": "081245340921",
        "IsMember": false,
        "Discounts": null,
        "Bills": null,
        "UserCredential": {
            "ID": 0,
            "CreatedAt": "0001-01-01T00:00:00Z",
            "UpdatedAt": "0001-01-01T00:00:00Z",
            "DeletedAt": null,
            "UserName": "",
            "UserPassword": "",
            "Email": "",
            "CustomerID": 0
        }
    }
}
  ```

### Login
- Request: POST
- Endpoint : `/auth/login`
- Body : 
```json
{
    "UserName":"angga21",
    "UserPassword":"passwordangga"
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTc4MTU5MDYsImlhdCI6MTY1NzgxNTMwNiwiaXNzIjoiV01fQkFIQVJJIiwidXNlck5hbWUiOiIiLCJFbWFpbCI6IiIsImFjY2Vzc1VVSUQiOiJlZWZkYzE4MC0wMjgzLTQwNjYtYmQ5Zi0xZjNmZGM5MzI1Y2QifQ.po7yaLjPN8zvYkRlDUQ5WS4ootbaSIvF8G0kwh49c14"
}
  ```

### Get Menu
- Request: GET
- Endpoint : `/menu/:id`
- Variable : number
- Response:
```json
  "data": {
        "ID": 2,
        "CreatedAt": "2022-07-06T13:09:23.16975+07:00",
        "UpdatedAt": "2022-07-06T17:08:38.254381+07:00",
        "DeletedAt": null,
        "MenuName": "sayur sop spesial",
        "MenuPrices": [
            {
                "ID": 2,
                "CreatedAt": "2022-07-06T13:09:23.170598+07:00",
                "UpdatedAt": "2022-07-06T13:09:23.170598+07:00",
                "DeletedAt": null,
                "MenuID": 2,
                "Price": 2000,
                "Bills": null
            }
        ]
    }
  ```

### Update Menu
- Request: PUT
- Endpoint : `/menu/update`
- Scheme : update name, update price
- Body : 
```json
{
    "menu_id":2,
    "menu_name":"sayur sop bayam",
    "menu_price":1500
}
```
- Response:
```json
  "data": {
        "ID": 2,
        "CreatedAt": "2022-07-06T13:09:23.16975+07:00",
        "UpdatedAt": "2022-07-06T19:58:15.073783413+07:00",
        "DeletedAt": null,
        "MenuName": "sayur sop bayam",
        "MenuPrices": [
            {
                "ID": 6,
                "CreatedAt": "2022-07-06T19:58:15.076796127+07:00",
                "UpdatedAt": "2022-07-06T19:58:15.076796127+07:00",
                "DeletedAt": null,
                "MenuID": 2,
                "Price": 1500,
                "Bills": null
            }
        ]
    }
  ```

### Delete Menu
- Request: DELETE
- Endpoint : `/menu/:id`
- Variable : number

### Create Menu
- Request: POST
- Endpoint : `/menu/register`
- Body : 
```json
{
    "menu_name":"Nutrisari Jambu",
    "menu_price":3000
}
```
- Response:
```json
  "data": {
        "ID": 5,
        "CreatedAt": "2022-07-06T20:01:37.780429542+07:00",
        "UpdatedAt": "2022-07-06T20:01:37.780429542+07:00",
        "DeletedAt": null,
        "MenuName": "Nutrisari Jambu",
        "MenuPrices": [
            {
                "ID": 7,
                "CreatedAt": "2022-07-06T20:01:37.781663279+07:00",
                "UpdatedAt": "2022-07-06T20:01:37.781663279+07:00",
                "DeletedAt": null,
                "MenuID": 5,
                "Price": 3000,
                "Bills": null
            }
        ]
    }
  ```

### Update Customer
- Request: PUT
- Endpoint : `/customer/update`
- Scheme : update name, update phone number, activate member, add discount
- Body : 
```json
{
    "customer_id":5,
    "customer_name":"Abdul Kadir",
    "mobile_phone_no":"0877123334",
    "is_member":true,
    "discount_id":1
}
```

### Create Customer
- Request: POST
- Endpoint : `/customer/register`
- Body : 
```json
{
    "customer_name":"Sulaiman",
    "mobile_phone_no":"0877128875"
}
```

### Update Table
- Request: PUT
- Endpoint : `/table/update`
- Scheme : update description, update availability
- Body : 
```json
{
    "id":1,
    "tabledescription":"Ini Table 1",
    "isavailable":false
}
```

### Create Table
- Request: POST
- Endpoint : `/table/register`
- Body : 
```json
{
    "tabledescription":"meja 10"
}
```

### Update Discount
- Request: PUT
- Endpoint : `/discount/update`
- Scheme : update description, update discount value
- Body : 
```json
{
    "id":1,
    "description":"discount member baru",
    "pct":15
}
```

### Create Discount
- Request: POST
- Endpoint : `/discount/register`
- Body : 
```json
{
    "description":"discount member",
    "pct":10
}
```

### Get Customer - Get Table - Get Discount
- Request: GET
- Endpoint : `/customer/:id` `/table/:id` `/discount/:id`
- Variable : number

### Delete Customer - Delete Table - Delete Discount
- Request: DELETE
- Endpoint : `/customer/:id` `/table/:id` `/discount/:id`
- Variable : number

### Create Transaction
- Request: POST
- Endpoint : `/transaction/create`
- Body : 
```json
{
    "table_id":2,
    "transaction_type_id":"DI",
    "customer":{
        "customername":"Devi",
        "mobilephoneno":"0877745983"
    },
    "order_menu":[
        {
            "menu_id":1,
            "qty":1
        },
        {
            "menu_id":4,
            "qty":1
        }
    ]
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": "transaction created! id:5"
}
  ```

### Pay and Finish Bill
- Request: POST
- Endpoint : `/payment/pay`
- Scheme Payment Method : "cash" OR "lopei"
- Body : 
```json
{
    "BillId":3,
    "PaymentMethod":"lopei"
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": {
        "bill_id": 3,
        "transaction_date": "14 Jul 2022 20:54:18",
        "customer_name": "Devi",
        "transaction_type": "Dine In",
        "table_number": "3",
        "grand_total": 65000,
        "order_menu": [
            {
                "menu_name": "Nasi Putih",
                "menu_price": 5000,
                "qty": 10,
                "subtotal": 50000
            },
            {
                "menu_name": "Es Teh Tawar",
                "menu_price": 1500,
                "qty": 10,
                "subtotal": 15000
            }
        ]
    }
}
  ```

### Daily Revenue
- Request: GET
- Endpoint : `/transaction/revenue`
- Body : 
```json
{
    "transaction_date":"2022-07-06"
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": {
        "transaction_date": "2022-07-06",
        "total_revenue": 31070
    }
}
  ```

### Check Balance
- Request: GET
- Endpoint : `/transaction/payment/balance`
- Body : 
```json
{
    "MobilePhoneNo":"0877745983"
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": {
        "CustomerName": "Devi",
        "MobilePhoneNo": "0877745983",
        "Balance": 50000
    }
}
  ```