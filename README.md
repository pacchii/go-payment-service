# go-payment-service

## 💳 Payment API

### 🔹 Create Payment

**POST** `http://localhost:9090/payment`

#### Request Body
```json
{
  "amount": 1000.0
}
```

---

### 🔹 Get Payment by ID

**GET** `http://localhost:9090/payment/{id}`

#### Example
```
http://localhost:9090/payment/8be4219c-5233-40ba-a559-6fabbbeb21d3
```

#### Response
```json
{
  "id": "8be4219c-5233-40ba-a559-6fabbbeb21d3",
  "amount": 1000,
  "status": "SUCCESS"
}
```

---

### 🔹 Get All Payments

**GET** `http://localhost:9090/payments`

#### Response
```json
[
  {
    "id": "3a81e189-5e50-4d43-800d-dcd3f71c1588",
    "amount": 1000,
    "status": "FAILURE"
  },
  {
    "id": "b80045c6-8d27-4faa-a85c-32b43de1718c",
    "amount": 1000,
    "status": "SUCCESS"
  }
]
```
