# iitk-coin

repo for iitk-coin project 



--------------------------------------------------------------------
## [UPD:Task 5]:

### [UPD: OTP added to signup]

JSON req on /signup now requires OTP entry, put OTP = "NULL" if u want to generate new one.

OTP expire in 5 mins. 

You get only 2 tries per OTP (to prevent bruteforce attack), failing both tries you need to generate a new one.

OTPs are 4 digits, may start with 0(s).


## Tables added to a.db :



Redeem (Contains Awards, Prices, AwardId)

Pending (Contains Pending Txns for Gensec to verify)



## 3 endpoints added:



/add/: Authorised person can create awards for GBM

```ex: {"Roll":"200536", "Award":"T-shirt", "Coins":"100", "JWT":"JWT"}```



/redeem/: GBM can create redeem requests to Gensec, by selecting Id of award

```ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}```



/approve/: Gensec can approve requests here by selecting txn id.

```ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}```




--------------------------------------------------------------------
## T4 done under "Task 4"-- iitkcoin is executable


Server listens at 127.0.0.1:8000/

a.db has the tables for users, transactions


## T4 Has endpoints:


/signup/: Accepts Requests for new user creation

```ex: {"Roll":"200536", "Pass":"pass"}```



/login/: listens to login requests, returns JWT if credentials are found to be valid.

```ex: {"Roll":"200536", "Pass":"pass"}```



/create/ : Aceepts Json requests for coin/ user creation with JWT by a superuser(admimn)

```ex:  {"Roll": "200536", "Coins": "1000", "JWT":"Val" }```



/check/ : Aceepts Json requests for checking Balance 

```ex:  {"Roll": "200536"}```



/transfer/ : Aceepts Json requests to transfer coins  and JWT of the sender

```ex:  {"From": "200536", "To": "123456", "Coins": "120","JWT":"Val"} ```


