
### Step 1: Install Vault

First, install Vault on your system:

```bash
# Download and install Vault
curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
sudo apt-get update && sudo apt-get install vault
```

### Step 2: Start Vault

Start Vault in development mode (for testing purposes):

```bash
vault server -dev
```

In another terminal, export the Vault address and token:

```bash
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='root'
```

### Step 3: Configure Vault

Enable the PKI secrets engine and configure it to act as a Certificate Authority (CA):

```bash
vault secrets enable pki

vault write pki/root/generate/internal common_name="surexsend.com" ttl=8760h

vault write pki/config/urls \
    issuing_certificates="$VAULT_ADDR/v1/pki/ca" \
    crl_distribution_points="$VAULT_ADDR/v1/pki/crl"

vault write pki/roles/example-dot-com \
    allowed_domains="surexsend.com" \
    allow_subdomains=true \
    max_ttl="72h"
```

Curl command to request a certificate with the generated token
```bash
curl -X POST -H "Authorization: Bearer YOUR_JWT_TOKEN" -H "Content-Type: application/json" -d '{"common_name": "test.surexsend.com"}' http://192.168.0.199:8080/request_cert
```
