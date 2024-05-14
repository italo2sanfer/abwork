# GCE1 Project

Proven skills:

- IaC (Docker)
- IaC (Terraform)
- Google Compute Engine (GCE)

## Instructions

```Passo 1
$ gloud --version
```

```Passo 2
$ gcloud init
```

Follow the instructions and log in with your account using the token. Choose an existing project to handle or create a new project.

If you choose an existing project, it must have a terraform key created. If you want to create a new project, create a terraform key in it.

After choosing the project, fill in the project_id variable in varias.tf with the information. Proceed with the terraform instructions.

```Passo 3 - Terraform
$ terraform init
$ terraform plan
$ terraform apply
$ terraform destroy
```