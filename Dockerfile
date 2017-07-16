FROM scratch

ADD build/vault-demo-injection-linux-amd64.tgz /app

CMD ["/app/vault-demo-injection"]