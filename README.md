# terraform-provider-volcengine
Terraform Provider
==================
<svg width="313" height="88" viewBox="0 0 313 88" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M98.46 30.46H87.21v-6.72h30.121v6.72h-11.25v33.6h-7.62v-33.6z" fill="#000"></path><path d="M125.291 58.59a30.77 30.77 0 009-1.39l1.15 5.56a31.65 31.65 0 01-10.86 1.88c-9.25 0-12.46-4.3-12.46-11.38v-7.8c0-6.23 2.78-11.49 12.22-11.49 9.44 0 11.56 5.5 11.56 11.85v6.3h-16.32v1.51c0 3.57 1.23 4.96 5.71 4.96zm-5.71-12.4h9.38v-1.46c0-2.78-.85-4.71-4.48-4.71-3.63 0-4.9 1.93-4.9 4.71v1.46zm36.08-5.45a56.603 56.603 0 00-7.81 4.3v19.02h-7.38V34.57h6.23l.49 3.27a32.648 32.648 0 017.74-3.87l.73 6.77zm18.019 0a57.125 57.125 0 00-7.8 4.3v19.02h-7.38V34.57h6.23l.48 3.27a32.764 32.764 0 017.75-3.87l.72 6.77zm24.681 23.32h-6.05l-.54-2a16.153 16.153 0 01-8.77 2.6c-5.39 0-7.69-3.69-7.69-8.77 0-6 2.61-8.29 8.59-8.29h7.08v-3.11c0-3.26-.9-4.41-5.62-4.41-2.747.03-5.484.33-8.17.9l-.91-5.62a38.314 38.314 0 0110.1-1.39c9.26 0 12 3.26 12 10.64l-.02 19.45zm-7.38-11.16h-5.4c-2.42 0-3.09.67-3.09 2.91 0 2 .67 3 3 3 1.95-.03 3.861-.55 5.56-1.51l-.07-4.4zm30.069-25.04a21.154 21.154 0 00-4.24-.49c-2.9 0-3.32 1.27-3.32 3.51v3.69h7.5l-.41 5.87h-7.07v23.62h-7.38V40.44h-4.72v-5.87h4.72v-4.11c0-6.11 2.84-9.14 9.37-9.14a23.496 23.496 0 016.35.85l-.8 5.69zm14.39 36.78c-10.1 0-12.86-5.58-12.86-11.58v-7.48c0-6 2.72-11.61 12.82-11.61s12.83 5.56 12.83 11.61v7.48c.04 6-2.65 11.58-12.79 11.58zm0-24.38c-3.93 0-5.44 1.75-5.44 5.08v7.92c0 3.33 1.51 5.09 5.44 5.09 3.93 0 5.45-1.76 5.45-5.09v-7.92c0-3.28-1.51-5.08-5.45-5.08zm31.82.48a57.125 57.125 0 00-7.8 4.3v19.02h-7.38V34.57h6.23l.48 3.27a32.943 32.943 0 017.79-3.87l.68 6.77zm20.32 23.32v-20.6c0-1.57-.67-2.36-2.36-2.36s-5 1.09-7.68 2.49v20.47h-7.39V34.57h5.63l.73 2.48a29.593 29.593 0 0111.79-3.08c2.85 0 4.6 1.15 5.57 3.14a29.004 29.004 0 0111.86-3.14c4.9 0 6.65 3.44 6.65 8.71v21.38H305v-20.6c0-1.57-.67-2.36-2.36-2.36a19.425 19.425 0 00-7.68 2.49v20.47h-7.38z" fill="#000"></path><path fill-rule="evenodd" clip-rule="evenodd" d="M21.2 16.55l19.1 11.03v22.06L21.2 38.61V16.55zm21.19 11.03v22.06L61.5 38.61V16.55L42.39 27.58zM0 4.24V26.3l19.1 11.03V15.27L0 4.24zm21.2 58.85l19.1 11.03V52.06L21.2 41.03v22.06z" fill="#7B42BC"></path></svg>

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)


Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.13 (to build the provider plugin)


# Volcengine Provider

The Volcengine provider is used to interact with many resources supported by [Volcengine](https://www.volcengine.com/).
The provider needs to be configured with the proper credentials before it can be used.

Use the navigation on the left to read about the available resources.

-> **Note:** This guide requires an available Volcengine account or sub-account with project to create resources.

## Example Usage
```hcl
# Configure the Volcengine Provider
provider "volcengine" {
  access_key = "your ak"
  secret_key = "your sk"
  session_token = "sts token"
  region = "cn-beijing"
}

# Query Vpc
data "volcengine_vpcs" "default"{
  ids = ["vpc-mizl7m1kqccg5smt1bdpijuj"]
}

#Create vpc
resource "volcengine_vpc" "foo" {
  vpc_name = "tf-test-1"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8","114.114.114.114"]
}

```

## Authentication

The Volcengine provider offers a flexible means of providing credentials for
authentication. The following methods are supported, in this order, and
explained below:

- Static credentials
- Environment variables

### Static credentials

Static credentials can be provided by adding an `public_key` and `private_key` in-line in the
volcengine provider block:

Usage:

```hcl
provider "volcengine" {
   access_key = "your ak"
   secret_key = "your sk"
   region = "cn-beijing"
}
```

### Environment variables

You can provide your credentials via `VOLCENGINE_ACCESS_KEY` and `VOLCENGINE_SECRET_KEY`
environment variables, representing your volcengine public key and private key respectively.
`VOLCENGINE_REGION` is also used, if applicable:

```hcl
provider "volcengine" {
  
}
```

Usage:

```hcl
$ export VOLCENGINE_ACCESS_KEY="your_public_key"
$ export VOLCENGINE_SECRET_KEY="your_private_key"
$ export VOLCENGINE_REGION="cn-beijing"
$ terraform plan
```



