# iitk-coin

repo for iitk-coin project 

T3 done under "Task3"-- iitkcoin is executable


Server listens at 127.0.0.1:8000/


T3 Has 3 endpoints:

/create : Aceepts Json requests for coin/ user creation 
ex:  {"Roll": "200536", "Coins": "1000"} 


/check : Aceepts Json requests for checking Balance 
ex:  {"Roll": "200536"}


/transfer : Aceepts Json requests to transfer coins 
ex:  {"From": "200536", "To": "123456", "Coins": "120"} 
