# Go Payment Processor

Simulates a secure implementation of internal wallets transfers and external payments operations in a Fintech system. It demonstrates the use atomic database transactions and idempotency keys to prevent race conditions, double processing, and handling concurrent requests properly.

- **Wallet Funding**: Simulates funding wallet from an external source e.g virtual account credit
- **Internal Wallet Transfer:** Simulates transfer between wallets, accounts for race conditions and idempotency.
- **External Bank Transfer:** Simulates USD/GRP/EUR transfers to external bank accounts.

## Setup with Docker

### 1) Clone repo
```bash
git clone https://github.com/jesseinvent/go-payment-processor
```

### 2) Build with Docker and run migrations
```bash
make build-docker
```

## Endpoints

### Create user - Simulates user registration
```bash
curl -X POST http://localhost:5001/api/v1/users \
-H "Content-Type: application/json" \
-d '{
  "name": "James",
  "email": "james@example.com",
  "phoneNumber": "09037148367"
}'
```

### Get user - Simulates fetching user details e.g during login
```bash
curl -X GET http://localhost:5001/api/v1/users/1
```

### Create Currency - Simulates adding a new currency to the system
```bash
curl -X POST http://localhost:5001/api/v1/currencies \
-H "Content-Type: application/json" \
-d '{
  "name": "USD",
  "symbol": "USD",
  "baseUnitFactor": 100,
  "iconUrl": "https://usd-icon.png"
}'
```

### Get Currencies - Simulates fetching all available currencies in the system
```bash
curl -X GET http://localhost:5001/api/v1/currencies
```

### Create wallet for user - Simulates wallet creation for a user in a specific currency
```bash
curl -X POST http://localhost:5001/api/v1/wallets \
-H "Content-Type: application/json" \
-d '{
  "userId": 1,
  "currencyId": 1
}'
```

### Get user wallet details - Simulates fetching a user's wallets details including balance and currency information
```bash
curl -X GET http://localhost:5001/api/v1/wallets/user/1
```

### Fund wallet - Simulates funding a user's wallet from an external source
```bash
curl -X POST http://localhost:5001/api/v1/payments/fund-wallet \
-H "Content-Type: application/json" \
-d '{
  "userId": 1,
  "walletId": 1,
  "amount": 1000
}'
```

### Internal wallet transfer - Simulates transferring payments in a specific currency between two wallets within the system
```bash
curl -X POST http://localhost:5001/api/v1/payments/internal-transfer \
-H "Content-Type: application/json" \
-H "Idempotency-Key: <unique-uuid-key>" \
-d '{
  "senderUserId": 2,
  "receiverUserId": 1,
  "currencyId": 1,
  "amount": 20.5
}'
```

### External bank transfer - Simulates transferring payments in a specific currency from a user's wallet to an external bank account
```bash
curl -X POST http://localhost:5001/api/v1/payments/external-bank-transfer \
-H "Content-Type: application/json" \
-H "Idempotency-Key: <unique-uuid-key>" \
-d '{
  "senderUserId": 1,
  "currencyId": 1,
  "amount": 10.5,
  "beneficiaryName": "John Doe",
  "beneficiaryAccountNumber": "1234568951",
  "beneficiaryBankCode": "133",
  "swiftCode": "11111",
  "sortCode": "0000"
}'
```

### Get transaction history - Simulates fetching a user's transaction history 
```bash
curl -X GET "http://localhost:5001/api/v1/transactions/user/1"
``` 
