# Loan Application Backend


### Assumptions /Instructions
1) Currently data is added for only 2 businesses, ABC and XYZ, and only for 4 months, however application will work as per the rules specified for 12 months.
2) Loan application date will be considered as today.
3) The Average Asset value and profit and loss is calculated for last 12 months, however if the data is not present for some month then the average is calculated accordingly
3) Api's are secured based on X-API-KEY, so that only frontend with that key should able to call the backend.
4) Test cases are added for core logic which are present [here](https://github.com/girishg4t/loan-application-system/blob/main/pkg/rules/rule_test.go) 
5) Backend is secured with api key which is 'X-API-KEY: super-secret'


### Api's exposed are
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