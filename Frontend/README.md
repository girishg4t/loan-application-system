# Loan Application frontend


### Assumptions
1) The application will only work for 2 businesses ABC and XYZ, this is because data is only set for these companies.
2) Currently no validation is performed on the frontend site.
3) Backend api will only be called with X-API-KEY as 'super-secret'


### Backend api's
#### To get balance sheet
```sh
curl --location 'http://localhost:8080/api/v1/abc/balancesheet/xero' \
--header 'X-API-KEY: super-secret' \
--header 'Content-Type: application/json' \
--data '{}'

Response:
[
    {
        "year": 2022,
        "month": 12,
        "profitOrLoss": 250000,
        "assetsValue": 1234
    },
    {
        "year": 2022,
        "month": 11,
        "profitOrLoss": 1150,
        "assetsValue": 5789
    },
    {
        "year": 2022,
        "month": 10,
        "profitOrLoss": 2500,
        "assetsValue": 22345
    },
    {
        "year": 2023,
        "month": 1,
        "profitOrLoss": -187000,
        "assetsValue": 223452
    }
]

```

#### To submit application
```sh
curl --location 'http://localhost:8080/api/v1/ABC/submit' \
--header 'X-API-KEY: super-secret' \
--form 'business_name="ABC"' \
--form 'established_year="1999"' \
--form 'loan_amount="12300"' \
--form 'account_provider="XERO"'

Response:

{
    "decision": true,
    "approved_amount": 12300
}

```