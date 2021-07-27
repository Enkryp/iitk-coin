# iitk-coin

## **Summer Project 2021**

## **SnT Project 2021, Programming Club**

This repository contains the backend code for the IITK Coin which is a reward based psuedo currency of IIT Kanpur.

### Relevant Links

- [Final Documentation](https://docs.google.com/document/d/1EeWv3Scq-kA00V1FJrz0wNHvRIWpM0l7oqCeJ1PEay8/edit?usp=sharing)
- [Poster](https://drive.google.com/file/d/1Iy7P4NNJWoIywhW9lNcm_IeZctU2mKc4/view?usp=sharing)
- [Midterm Evaluation presentation](https://docs.google.com/presentation/d/1kriN-7A3v1RlXUDL5NETX3roJKRMJInptkWofIxY8dg/edit?usp=sharing)
- [Midterm Documentation](https://docs.google.com/document/d/1bvOWH4k0U-l2pQ1jLWIDzOkJ2wbHNW4jJw7tMWkUV6o/edit?usp=sharing)

## Table Of Content

- [Development Environment](#development-environment)
- [Directory Structure](#directory-structure)
- [Usage](#usage)
- [Endpoints](#endpoints)
- [Models](#models)
- [Settings](#settings)

## Development Environment

```bash
- OS:           Fedora 34
- Kernel:       Linux 5.13 
- go version:   go1.16.6 linux/amd64    
- text editor:  VSCode    	                  # https://code.visualstudio.com/download
```

## Directory Structure

```


└── iitk-coin
    ├── README.md
    ├── Task 1
    │   ├── a.db
    │   ├── go.mod
    │   ├── go.sum
    │   ├── iitk-coin
    │   └── main.go
    ├── Task 2
    │   ├── a.db
    │   ├── go.mod
    │   ├── go.sum
    │   ├── iitk-coin2
    │   ├── index.html
    │   ├── login.go
    │   ├── login.html
    │   ├── main.go
    │   ├── ok.html
    │   ├── secret.go
    │   ├── signup.go
    │   ├── signup.html
    │   └── wrong.html
    ├── Task 3
    │   ├── a.db
    │   ├── check.go
    │   ├── create.go
    │   ├── go.mod
    │   ├── go.sum
    │   ├── iitkcoin
    │   ├── main.go
    │   └── transfer.go
    ├── Task 4
    │   ├── a.db
    │   ├── check.go
    │   ├── create.go
    │   ├── go.mod
    │   ├── go.sum
    │   ├── iitkcoin
    │   ├── login.go
    │   ├── main.go
    │   ├── signup.go
    │   └── transfer.go
    └── Task 5
        ├── a.db
        ├── add.go
        ├── approve.go
        ├── check.go
        ├── create.go
        ├── dockerfile
        ├── go.mod
        ├── go.sum
        ├── login.go
        ├── mail.go
        ├── main.go
        ├── redeem.go
        ├── signup.go
        └── transfer.go

6 directories, 52 files
```

## Endpoints

POST requests take place via `JSON` requests on localhost:8000/

 

### [UPD: OTP added to signup]

JSON req on /signup now requires OTP entry, put OTP = "NULL" if u want to generate new one.

OTPs expire in 5 mins. 

You get only 2 tries per OTP (to prevent bruteforce attack), failing both tries you need to generate a new one.

OTPs are 4 digits, may start with 0(s).




##example JSONs on endpoints 

/signup/ : New signups here...
```
'{"Roll":"200536", "Pass":"pass", "OTP":"NULL/xxxx"}'
```



/add/: Authorised person can create awards for GBM

```ex: {"Roll":"200536", "Award":"T-shirt", "Coins":"100", "JWT":"JWT"}
```



/redeem/: GBM can create redeem requests to Gensec, by selecting Id of award

```ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}
```



/approve/: Gensec can approve requests here by selecting txn id.

```ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}
```




/signup/: Accepts Requests for new user creation

```ex: {"Roll":"200536", "Pass":"pass"}
```



/login/: listens to login requests, returns JWT if credentials are found to be valid.

```ex: {"Roll":"200536", "Pass":"pass"}
```



/create/ : Aceepts Json requests for coin/user creation with JWT by a superuser(admimn)

```ex:  {"Roll": "200536", "Coins": "1000", "JWT":"Val" }
```



/check/ : Aceepts Json requests for checking Balance 

```ex:  {"Roll": "200536"}
```



/transfer/ : Aceepts Json requests to transfer coins  and JWT of the sender

```ex:  {"From": "200536", "To": "123456", "Coins": "120","JWT":"Val"} 
```



## Tables and their Schemas in a.db :

```
OTP :	CREATE TABLE Otp (id INTEGER PRIMARY KEY,pass TEXT, time TEXT, fail INTEGER)

Pending:	CREATE TABLE Pending (id INTEGER PRIMARY KEY,awardId INTEGER, coins INTEGER, Recipient INTEGER)

Reedem:		CREATE TABLE Redeem (id INTEGER PRIMARY KEY,Item TEXT , coins INTEGER, Recipient INTEGER)

Tx: 	CREATE TABLE Tx (id INTEGER PRIMARY KEY,Time TEXT , coins INTEGER, txfrom INTEGER, txto INTEGER)

User:	CREATE TABLE User (id INTEGER PRIMARY KEY, coins FLOAT, pass TEXT, adm INTEGER)



```




