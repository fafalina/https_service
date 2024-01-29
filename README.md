# Simple Time.is
This is a simple example of Time.is. These project consists of three parts.
1. Use golang to create https service. (server)
2. The https service obtains data from the firebase real-time database. (SQL)
3. Build an Angular webui to get data from golang https service. (client)

# Github LFS Rule
Since the docker image is too large to be pushed to github, I need to use git lfs.
If you are unable to use "git clone" successfully, install git lfs and use "git lfs clone".
```bash
git lfs install
```
```bash
git lfs clone https://github.com/fafalina/https_service.git
```

# Docker Images
Please load docker images or use 'docker compose up'.
```bash
docker load -i https_service-golang-service.tar
docker load -i https_service-angular-service.tar

docker run -p 8080:8080 https_service-golang-service.tar
docker run -p 4200:4200 https_service-angular-service.tar
```

```bash
docker compose up
```

# Record Some Methods Here
## Generate root certificate and private key (windows power shell)
```shell
openssl genrsa -out ca.key 2048
openssl req -x509 -new -nodes -key ca.key -days 365 -out ca.crt -subj "/CN=MyCA"
```

## Generate server certificate and private key (windows power shell)
```shell
openssl genrsa -out server.key 2048
openssl req -new -key server.key -out server.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
echo "subjectAltName = DNS:localhost" | Out-File -FilePath san_config.txt -Encoding ascii
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -extfile san_config.txt
Remove-Item san_config.txt
```

## Generate client certificate and private key (windows power shell)
```shell
openssl genrsa -out client.key 2048
openssl req -new -key client.key -out client.csr -subj "/CN=localhost" -addext "subjectAltName=DNS:localhost"
echo "subjectAltName = DNS:localhost" | Out-File -FilePath san_config.txt -Encoding ascii
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -extfile san_config.txt
Remove-Item san_config.txt
```
