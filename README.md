## DJT Webserver

A web server in Go that can get the Ethereum owner and transactions of an ERC-721 token by 
the ERC-721 contract address and the token ID

## How to test
* Clone the project and build
* The webserver utilizes caching through a backend DB (currently defaults to Postgres)
* In order to utilize it you will need to spin up a Postgres instance (or use an existing instance)
    * With a db called `djthash` 
    * And a table with the following schema
    ```postgresql
      CREATE TABLE token_transfer (
          contract_addr text,
          token_id bigint,
          block bigint,
          from_addr text,
          to_addr text,
          PRIMARY KEY(contract_addr, token_id)
      );
    ```
* Run the main package `server.go`

## Endpoints for testing
You can test it [using this contract - 0xcc62564d40c06e2be1f84287b0d8f6b734c856d3](https://ropsten.etherscan.io/address/0xcc62564d40c06e2be1f84287b0d8f6b734c856d3) which is already deployed on Ropsten
* Here are the endpoints with sample responses (assuming you're deploying on `localhost:8080`)
    * Token Owner: http://localhost:8080/contracts/0xcc62564D40C06e2Be1F84287b0d8F6B734c856D30xcc62564D40C06e2Be1F84287b0d8F6B734c856D3/1/owner   
  ```json
    {
      "owner": "0x9Fa06D81DFB8F6A713a0Bd1071A2C24f1f990629"
    }
    ```
    * Token Transfers: http://localhost:8080/contracts/0xcc62564D40C06e2Be1F84287b0d8F6B734c856D30xcc62564D40C06e2Be1F84287b0d8F6B734c856D3/2/transfers
    ```json
    {
      "transfers": [
        {
            "Block": 6216854,
            "From": "0x0000000000000000000000000000000000000000",
            "To": "0x9Fa06D81DFB8F6A713a0Bd1071A2C24f1f990629"
        }
      ]
    }
    ```

## Alternate DB as a cache server
* You can alternately choose to use a different DB (such as MySQL)
    * In which case create a corresponding client under the `platform/cache` package (check the pgclient for reference)
    * And update the db connection string in `server.go` `main()`
    
## TODO (Future Enhancements)
* Get the etherium network and db connection string programmatically through cmd options
* Use a websocket to connect to Ropsten and use a subscriber to automatically get notified on transaction events
* Create the DB schema automatically if it doesn't exist
* Unit Tests...
* Better error handling