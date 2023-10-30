# Loan Application

## Instruction for Backend and Frontend
1) https://github.com/girishg4t/loan-application-system/blob/main/Backend/README.md
2) https://github.com/girishg4t/loan-application-system/blob/main/Frontend/README.md


### How to run the application
1) Create .env file in Backend folder with variable 
```sh
PORT=8080
API_KEY=super-secret
```
2) Run the docker compose command as
```sh
docker-compose up -d   
```

### Working demo
Balance sheet of Business ABC and provider Xero   
![Balance sheet ABC-Xero](./images/balance-sheet.png)
Balance sheet of Business XYZ and provider MYOB  
![Balance sheet XYZ-MYOB](./images/xyz-balance-sheet.png?raw=true)
When profit is positive and loan amount is less then asset value   
![ABC Loan Approved 100%](./images/full-approved-loan.png?raw=true)
When profit is positive and loan amount is greater then asset value (only 60 % is approved)
![ABC Loan Approved 60%](./images/appove-loand-60.png?raw=true)
When profit is negative so loan is not approved
![XYZ no Loan Approved](./images/loan-not-approved.png?raw=true)