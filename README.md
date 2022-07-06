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

## API Spec

### Get Menu
- Request: GET
- Endpoint : `/menu/`
- Body : 
```json
{
    "id":2
}
```
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
- Endpoint : `/menu/`
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
- Endpoint : `/menu/`
- Body : 
```json
{
    "id":2
}
```

### Create Menu
- Request: POST
- Endpoint : `/menu/`
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
- Endpoint : `/customer/`
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
- Endpoint : `/customer/`
- Body : 
```json
{
    "customer_name":"Sulaiman",
    "mobile_phone_no":"0877128875"
}
```

### Update Table
- Request: PUT
- Endpoint : `/table/`
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
- Endpoint : `/table/`
- Body : 
```json
{
    "tabledescription":"meja 10"
}
```

### Update Discount
- Request: PUT
- Endpoint : `/discount/`
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
- Endpoint : `/discount/`
- Body : 
```json
{
    "description":"discount member",
    "pct":10
}
```

### Get Customer - Get Table - Get Discount
- Request: GET
- Endpoint : `/customer/` `/table/` `/discount/`
- Body : 
```json
{
    "id":2
}
```

### Delete Customer - Delete Table - Delete Discount
- Request: DELETE
- Endpoint : `/customer/` `/table/` `/discount/`
- Body : 
```json
{
    "id":2
}
```

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

### Print Bill
- Request: GET
- Endpoint : `/transaction/print`
- Body : 
```json
{
    "id":5
}
```
- Response:
```json
{
    "response_code": "00",
    "response_message": "success",
    "data": {
        "bill_id": 5,
        "transaction_date": "6 Jul 2022 20:54:04",
        "customer_name": "Devi",
        "transaction_type": "Dine In",
        "table_number": "2",
        "grand_total": 6500,
        "order_menu": [
            {
                "menu_name": "Nasi Putih",
                "menu_price": 5000,
                "qty": 1,
                "subtotal": 5000
            },
            {
                "menu_name": "Es Teh Tawar",
                "menu_price": 1500,
                "qty": 1,
                "subtotal": 1500
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