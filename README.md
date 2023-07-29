# Multi-Blockchain Address Generator



The Multi-Blockchain Address Generator is a Golang-based tool that enathe Qoinpal team to generate multiple addresses for different blockchains. These addresses can be used by customers for various purposes. This repository provides a simple and efficient solution to manage blockchain addresses for our project.

## Features

- Supports multiple blockchains (Bitcoin, Ethereum, Litecoin, etc.).
- Generates a configurable number of addresses for each blockchain.
- Easy-to-use command-line interface.
- Lightweight and fast.

## Prerequisites

Make sure you have the following software installed on your system before running the project:

- Go programming language (version ^1.18)
- Make


## Installation

1. Clone this repository to your local machine:

```bash
git clone https://github.com/qoinpalhq/qoinpal_crypto.git 
```
## Usage

1. To start the development server, if make is installed on your machine
```bash
make run
```
else
```bash 
go run main.go
```
2. Once the development server is running on port 9002, you can make api calls e.g to generate Bitcoin address 
``` bash 
   curl GET http://localhost:9002/api/address/bitcoin
```
or use any other clients to make api requests
