rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=DK/ST=Copenhagen/L=Frederiksberg/O=ITU/OU=Education/CN=Mads Jensen/emailAddress=macj@itu.dk"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=DK/ST=Copenhagen/L=Hvidovre/O=Hospital/OU=Hospital/CN=Hospital/emailAddress=hospital@hospital.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 200 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

# 4. Generate client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout Alice-key.pem -out Alice-req.pem -subj "/C=DK/ST=Copenhagen/L=Frederiksberg/O=Patient/OU=Patient/CN=Alice/emailAddress=alice@patient.com"
openssl req -newkey rsa:4096 -nodes -keyout Bob-key.pem -out Bob-req.pem -subj "/C=DK/ST=Copenhagen/L=Amager/O=Patient/OU=Patient/CN=Bob/emailAddress=bob@patient.com"
openssl req -newkey rsa:4096 -nodes -keyout Charlie-key.pem -out Charlie-req.pem -subj "/C=DK/ST=Copenhagen/L=Valby/O=Patient/OU=Patient/CN=Charlie/emailAddress=charlie@patient.com"

# 5. Use CA's private key to sign client's CSR and get back the signed certificate
openssl x509 -req -in Alice-req.pem -days 200 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out Alice-cert.pem -extfile client-ext.cnf
openssl x509 -req -in Bob-req.pem -days 200 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out Bob-cert.pem -extfile client-ext.cnf
openssl x509 -req -in Charlie-req.pem -days 200 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out Charlie-cert.pem -extfile client-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text