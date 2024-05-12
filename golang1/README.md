# GoLang Project

Competências comprovadas neste projeto:

- IaC (Docker);
- Golang; HTML; CSS;
- Google Cloud Run

## Demonstração

**Com docker**

> Instruções
$ cd ~
$ git clone https://github.com/italo2sanfer/abwork
$ cd ~/abwork/golang1/
$ docker build -f DockerfileDemo -t abwork-golang1:v1 .
$ docker run --rm -p 8080:8080 abwork-golang1:v1

Acesse http://localhost:8080/

**Pelo App no Google Cloud Run**

Colocar link

## Base

[GCP Samples](https://github.com/GoogleCloudPlatform/golang-samples/tree/main/run/helloworld)