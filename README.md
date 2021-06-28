# iitk-coin

repo for iitk-coin project 

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


