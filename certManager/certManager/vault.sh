#!/usr/bin/env bash
export VAULT_ADDR='http://127.0.0.1:8200'
export VAULT_TOKEN='root'


vault secrets enable pki

vault write pki/root/generate/internal common_name="surexsend.com" ttl=8760h

vault write pki/config/urls \
    issuing_certificates="$VAULT_ADDR/v1/pki/ca" \
    crl_distribution_points="$VAULT_ADDR/v1/pki/crl"

vault write pki/roles/surexsend-dot-com \
    allowed_domains="surexsend.com" \
    allow_subdomains=true \
    max_ttl="72h"
