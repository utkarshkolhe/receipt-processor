# Fetch Receipt Processor
Containarized Go API for handling customer purchase reciepts.
## Installation

### Requirements
- Clone of the repository
- Docker

### Setting Up

1. Clone the repository.
2. Execute following commands:
```
docker build . -t go-containerized:latest
docker run -e PORT=9000 -p 9000:9000 go-containerized:latest
```
The API is up and running!!

<br>
<br>
Can provide any name in place of go-containerized and any port number in place of 9000.

```
docker build . -t CONTAINERNAME:latest
docker run -e PORT=PORTNUMBER -p PORTNUMBER:9000 CONTAINERNAME:latest
```



## Design

![Data_Model](/design_diagram.png)
The Diagram above shows high level view of how the API components interact.

### Input validation
Input Validation implemented in [Item_Model](/api/models/item_model.go) and [Receipt_Model](/api/models/receipt_model.go). <br>
Types of input validation done:
1. All fields (retailer, shortDescription, purchaseTime, purchaseDate, price, etc) follow their specified regex pattern as defined in [API_YML](/api.yml) . Can be updated in [Constants](/api/config/constants.go)
2. total of receipts equals sum of prices of items in Receipts
3. Receipt has at least MinItemsInReceipt(default=1) items. Can be updated in [Constants](/api/config/constants.go)

### Testing
Testing automatically done while building docker image.<br>
If required, testing in Go can be manually triggered by running:
```
go test ./...
```
#### Unit testing
Unit tests written for three modules in [ReceiptManagerTest](/api/db/receipt_manager_test.go), [IDGeneratorTest](/api/utils/id_generator_test.go) and [PointsCalculator](/api/utils/points_calculator_test.go)

#### End to End testing
End-To-End testing implemented in [MainTest](/main_test.go).
Evaluates on 4 valid test cases and 5 invalid test cases [TestCases](/test/).

#### Manual testing
You can send any other POST and GET requests on `localhost:9000/receipts/process` and `localhost:9000/receipts/__ID__/points` respectively.

### Logging
Logger is implemented in [Logger](/api/logger/logger.go)<br>
Logger performs logging on both console as well as writing to a file.<br>
Console or File Logging can be toggled by editing ConsoleLog and FileLog in [Constants](/api/config/constants.go) <br>
Example of a log file: [LogFile](/app.log)


### File Structure
```
fetch-receipt-processor/
├─ api/                             - Main source folder for the API
|    ├─ config/                     - Contains constants used across the API
|      ├─ constants.go              - Contains API config and other constants like Validation Patterns
|      ├─ msgs.go                   - Contains Error and Log Messages
|    ├─ db/                         - Contains logic for interactions with the database
|      ├─ reciept_manager.go        - Manages storage and retrieval of Receipts in the Database
|      ├─ reciept_manager_test.go   - Unit Testing for Receipt Manager
|    ├─ handelers/                  - Contains logic for all the handlers for the API endpoints
|      ├─ get_points.go             - Handles Endpoint: Get Points
|      ├─ process_receipts.go       - Handles Endpoint: Process Receipts
|    ├─ models/                     - Contains Definition for different Model Structures required by the API
│      ├─ id_model.go               - Structure for IDModel. used to send ID
│      ├─ points_model.go           - Structure for PointsModel, Used to send the points of Reciept
│      ├─ item_model.go             - Structure for Items and associated validation logic
│      ├─ receipt_model.go          - Structure for Receipts and associated validation logic
│    ├─ routes                      - Contains logic for assigning web routes to API
│      ├─ route.go                  - Assigns function to all route calls.
│    ├─ utils                       - Contains logic for common used functions across the API
│      ├─ id_generator.go           - Logic for generating unique ID.
│      ├─ id_generator_test.go      - Unit Testing for Id Generator.
│      ├─ points_calculator.go      - Logic for calculating points of a receipt.
│      ├─ id_generator_test.go      - Unit Testing for Point Calculator.
├─ test/                            - Contains test json data for End to End Testing
├─ .gitattributes                   - Attributes Specification for git
├─ app.log                          - Log File for the API
├─ dockerfile                       - Instructions to create docker container
├─ go.mod                           - Dependency Management for the Go Codebase
├─ go.sum                           - Checksum for Go. To avoid reinsatlling dependencies
├─ main.go                          - Entry point for the API Application
├─ main_test.go                     - Automated End-To-End API Testing
├─ api.yml                          - Formal Definition of the API
├─ design_diagram.png               - High Level Design of the API
├─ README.md                        - Documentation and general information about the project
```

## Summary of API Specification

### Endpoint: Process Receipts

* Path: `/receipts/process`
* Method: `POST`
* Payload: Receipt JSON
* Response: JSON containing an id for the receipt.

Description:

Takes in a JSON receipt (see example in the example directory) and returns a JSON object with an ID generated by your code.

The ID returned is the ID that should be passed into `/receipts/{id}/points` to get the number of points the receipt
was awarded.

How many points should be earned are defined by the rules below.

Reminder: Data does not need to survive an application restart. This is to allow you to use in-memory solutions to track any data generated by this endpoint.

Example Response:
```json
{ "id": "7fb1377b-b223-49d9-a31a-5a02701dd310" }
```

## Endpoint: Get Points

* Path: `/receipts/{id}/points`
* Method: `GET`
* Response: A JSON object containing the number of points awarded.

A simple Getter endpoint that looks up the receipt by the ID and returns an object specifying the points awarded.

Example Response:
```json
{ "points": 32 }
```

---

# Rules

These rules collectively define how many points should be awarded to a receipt.

* One point for every alphanumeric character in the retailer name.
* 50 points if the total is a round dollar amount with no cents.
* 25 points if the total is a multiple of `0.25`.
* 5 points for every two items on the receipt.
* If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
* 6 points if the day in the purchase date is odd.
* 10 points if the time of purchase is after 2:00pm and before 4:00pm.


## Examples

```json
{
  "retailer": "Target",
  "purchaseDate": "2022-01-01",
  "purchaseTime": "13:01",
  "items": [
    {
      "shortDescription": "Mountain Dew 12PK",
      "price": "6.49"
    },{
      "shortDescription": "Emils Cheese Pizza",
      "price": "12.25"
    },{
      "shortDescription": "Knorr Creamy Chicken",
      "price": "1.26"
    },{
      "shortDescription": "Doritos Nacho Cheese",
      "price": "3.35"
    },{
      "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
      "price": "12.00"
    }
  ],
  "total": "35.35"
}
```
```text
Total Points: 28
Breakdown:
     6 points - retailer name has 6 characters
    10 points - 4 items (2 pairs @ 5 points each)
     3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
     3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
     6 points - purchase day is odd
  + ---------
  = 28 points
```

----

```json
{
  "retailer": "M&M Corner Market",
  "purchaseDate": "2022-03-20",
  "purchaseTime": "14:33",
  "items": [
    {
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    },{
      "shortDescription": "Gatorade",
      "price": "2.25"
    }
  ],
  "total": "9.00"
}
```
```text
Total Points: 109
Breakdown:
    50 points - total is a round dollar amount
    25 points - total is a multiple of 0.25
    14 points - retailer name (M&M Corner Market) has 14 alphanumeric characters
                note: '&' is not alphanumeric
    10 points - 2:33pm is between 2:00pm and 4:00pm
    10 points - 4 items (2 pairs @ 5 points each)
  + ---------
  = 109 points
```

---
