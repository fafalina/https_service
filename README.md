# https_service
Use golang to create an https service connection

# Windows Power Shell command
## 生成 CA 私鑰
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 365 -out ca.crt -subj "/CN=MyCA"

## 生成伺服器私鑰
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
echo "subjectAltName = DNS:localhost" | Out-File -FilePath san_config.txt -Encoding ascii
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extfile san_config.txt
Remove-Item san_config.txt

## 生成客戶端私鑰
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
echo "subjectAltName = DNS:localhost" | Out-File -FilePath san_config.txt -Encoding ascii
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extfile san_config.txt
Remove-Item san_config.txt
