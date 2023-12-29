#! bin/bash

apt-get update
apt-get install ca-certificates
if [ ! -f ./docker/ciscoumbrellaroot.cer ]; then
  wget http://www.cisco.com/security/pki/certs/ciscoumbrellaroot.cer
  openssl x509 -inform DER -in ./docker/ciscoumbrellaroot.cer -out ./docker/ciscoumbrellaroot.crt
fi
cp ./docker/ciscoumbrellaroot.crt /usr/local/share/ca-certificates/ciscoumbrellaroot.crt
update-ca-certificates
