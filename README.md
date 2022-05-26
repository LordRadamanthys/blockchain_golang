
# Simple structure on how a Blockchain works

This very simple project is for study about blockchain and how that system works.


## How run it

First of all we need to check if dependencies have been downloaded so 
enter inside the folder project:

```bash
  go mod init
```
and
```bash
  go mod tidy
```
So after that the project should be ready to run with the command:

```bash
  go run main.go
```

Your terminal will show somenthing like that:

```bash
[GIN-debug] GET    /                         --> github.com/LordRadamanthys/blockchain_golang/src/application/ports.BlockchainController.GetBlockchain-fm (3 handlers)
[GIN-debug] POST   /                         --> github.com/LordRadamanthys/blockchain_golang/src/application/ports.BlockchainController.WriteBlockchain-fm (3 handlers)
[GIN-debug] POST   /book                     --> github.com/LordRadamanthys/blockchain_golang/src/application/ports.BlockchainController.NewBook-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
Prev. hash: 
Data: {
"book_id": "genesis",
"user": "genesis",
"checkout_date": "genesis",
"is_genesis": true
}
Hash: c4208a83e372d47ff683e910d7634e2f346d15d89a11ee372c2fee1b6fd310cc
size: 1
```
Alright, API is on now let's make some request to see what happens


## API Reference

#### Get blockchain

```http
  GET /blockchain
```
This will return all blocks inside the blockchain

```javascript
  {
	"Blocks": [
		{
			"Pos": 1,
			"Data": {
				"book_id": "genesis",
				"user": "genesis",
				"checkout_date": "genesis",
				"is_genesis": true
			},
			"TimeStamp": "2022-05-25 23:12:39.448564769 -0300 -03 m=+0.001742821",
			"Hash": "0177dc90b9e7cde51306dea64caf62a97de9bd908dacb6e07985ca59662a42a9",
			"PrevHash": ""
		},
		{
			"Pos": 2,
			"Data": {
				"book_id": "90a36e8c8d6c914166dc1b991bbcc6bb",
				"user": "Mary poppins",
				"checkout_date": "2022-05-26",
				"is_genesis": false
			},
			"TimeStamp": "2022-05-25 23:12:48.373040209 -0300 -03 m=+8.926218281",
			"Hash": "d2fbc7adf5958ab0c3aae6307c0487d27beb14865b4d16390452c1de95740845",
			"PrevHash": "0177dc90b9e7cde51306dea64caf62a97de9bd908dacb6e07985ca59662a42a9"
		}
	]
}
```

#### Get item

```http
  GET /api/items/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of item to fetch |

#### add(num1, num2)

Takes two numbers and returns the sum.

