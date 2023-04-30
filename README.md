# Auth Go

authentication and authorization system written in Go

## Deployment

### generating token

```bash
openssl genpkey -algorithm ed25519 -out private.pem
openssl pkey -in private.pem -pubout -out public.pem
```
