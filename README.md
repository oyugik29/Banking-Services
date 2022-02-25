# BasicAuthentication
A simple basic authentication api with HTTPS

# Create security certificates
mkcert localhost

# Setup 
1. Clone Repo
2. Get libraries go mod init hubuc.com/basic-auth-example
3. Run  ( using temporary username and password ) AUTH_USERNAME=test AUTH_PASSWORD=FQfKWRCcGjZqz6g go run .
4. Test on the browser https://localhost:4000/secure
5. Should return the UserID = 1234456




