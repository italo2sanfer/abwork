# GCE1 Project

Automating the creation of a Google Cloud Engine (GCE) with a web server on Google Cloud Compute. Proven skills:

- IaC (Docker)
- IaC (Terraform)
- Google Compute Engine (GCE)

## Instructions

Up and down the infrastructure with docker-compose. You can enter in container and use the resources.

```Basic usage
$ docker-compose -f docker-compose-dev.yml up -d
$ docker exec -it gce1_gce_1 bash
$ docker-compose -f docker-compose-dev.yml down
```

Handle gcloud.

```
$ gloud --version
$ gcloud init
```

Follow the instructions and log in with your account using the token. Choose an existing project to handle or create a new project.

If you choose an existing project, it must have a terraform key created. If you want to create a new project, create a terraform key in it.

After choosing the project, fill in the project_id variable in variables.tf with the information. Proceed with the terraform instructions.

```
$ terraform init
$ terraform plan
$ terraform apply
$ terraform destroy
```