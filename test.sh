curl http://0.0.0.0:9900
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/post  
curl -u test:test http://0.0.0.0:9900/json 
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/debit  
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/credit 
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/create
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/getCardDetails
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/block
curl -u test:test -d "param1=test&param2=test" http://0.0.0.0:9900/unblock  

