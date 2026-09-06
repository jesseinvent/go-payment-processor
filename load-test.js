import http from 'k6/http';

/**
 * Simulate a load test with 100 virtual users for a duration of 1 minute
 * Result: No errors, all requests completed successfully with an average response time of 200ms and a maximum response time of 800ms, due to the SELECT FOR UPDATE row locking, subsequent requests had to wait for the lock to be released which led to increased response times.
 * Ideally, requests takes <15ms
 */

// export const options = {
//   vus: 100,
//   duration: '1m',
// };

// export default function () {

//     const payload = JSON.stringify({
//         "userId": 2,
//         "walletId": 2,
//         "amount": 100
//     });

//     http.post('http://localhost:5001/api/v1/payments/fund-wallet', payload, {
//     headers: {
//       'Content-Type': 'application/json',
//     },
//   });

// }

export const options = {
  vus: 1000,
  duration: '10s',
};

export default function () {

    const payload = JSON.stringify({
        "senderUserId": 4,
        "currencyId": 1,
        "amount": 500,
        "beneficiaryName": "John Doe",
        "beneficiaryAccountNumber": "1234568951",
        "beneficiaryBankCode": "133",
        "swiftCode": "11111",
        "sortCode": "0000"
    });

    http.post('http://localhost:5001/api/v1/payments/external-bank-transfer', payload, {
    headers: {
        'Content-Type': 'application/json',
        // 'Idempotency-Key': `test-${__VU}-${__ITER}` // Unique key for each request
        'Idempotency-Key': `test_12345WSWSEDDDRFDDD333DDD2` // Unique key for each request
    },
  });

}


