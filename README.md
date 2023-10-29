# Loan Application

## Instruction for Backend and Frontend
1) https://github.com/girishg4t/loan-application-system/Backend/README.md
2) https://github.com/girishg4t/loan-application-system/Frontend/README.md
### How to run the application
1) Create .env file with variable 
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
![alt text](https://github.com/girishg4t/loan-application-system/balance-sheet.png?raw=true)
Balance sheet of Business XYZ and provider MYOB  
![alt text](https://github.com/girishg4t/loan-application-system/xyz-balance-sheet.png?raw=true)
When profit is positive and loan amount is less then asset value   
![alt text](https://github.com/girishg4t/loan-application-system/full-approved-loan.png?raw=true)
When profit is positive and loan amount is greater then asset value (only 60 % is approved)
![alt text](https://github.com/girishg4t/loan-application-system/appove-loand-60.png?raw=true)
When profit is negative so loan is not approved
![alt text](https://github.com/girishg4t/loan-application-system/loan-not-approved.png?raw=true)