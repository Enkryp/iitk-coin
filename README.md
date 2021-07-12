# iitk-coin

repo for iitk-coin project 



--------------------------------------------------------------------
[UPD:Task 5]:

Tables added to a.db :

Reedme (Contains Awards, Prices, AwardId)
Pending (Contains Pending Txns for Gensec to verify)


3 endpoints added:

/add/: Authorised person can create awards for GBM
ex: {"Roll":"200536", "Award":"T-shirt", "Coins":"100", "JWT":"JWT"}

/redeem/: GBM can create redeem requests to Gensec, by selecting Id of award
ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}

/approve/: Gensec can approve requests here by selecting txn id.
ex: {"Roll":"200536", "JWT":"JWT", "Id":"123"}




--------------------------------------------------------------------
T4 done under "Task 4"-- iitkcoin is executable


Server listens at 127.0.0.1:8000/

a.db has the database for users, transactions


T4 Has endpoints:


/signup/: Accepts Requests for new user creation
ex: {"Roll":"200536", "Pass":"pass"}


/login/: listens to login requests, returns JWT if credentials are found to be valid.
ex: {"Roll":"200536", "Pass":"pass"}


/create/ : Aceepts Json requests for coin/ user creation with JWT by a superuser(admimn)
ex:  {"Roll": "200536", "Coins": "1000", "JWT":"Val" }


/check/ : Aceepts Json requests for checking Balance 
ex:  {"Roll": "200536"}


/transfer/ : Aceepts Json requests to transfer coins  and JWT of the sender
ex:  {"From": "200536", "To": "123456", "Coins": "120","JWT":"Val"} 


