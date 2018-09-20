# Istio Gateway

Create certs secret for Istio ingress gateway
```
kubectl create -n istio-system secret tls istio-ingressgateway-certs --key privkey.pem --cert cert.pem
```
