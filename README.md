# IFASSION Backend Service
## Website Development Competition - Technology Euphoria

### A Short Story...
IFASSION (InFormatics pASSION) is a website that can provide a recommendation of interest in informatics field. IFASSION works using expert system with forward chaining method.

IFASSION developed to tackle hesitations of newcomers in informatics fields. They may use IFASSION to get the recommendation of their area of interest.

### The Technology
Technologies used to develop IFASSION BE:
1. Go (progrmaming language)
2. MongoDB (database)
3. Gin (web framework)

### How to Run
#### 1. Install technologies used
- Install [Go Programming Language](https://go.dev/dl/)
- Install [MongoDB](https://www.mongodb.com/docs/manual/installation/) if you want to use local MongoDB server or create a database using [MongoDB Atlas](https://www.mongodb.com/atlas/database)  
Field template for collections:
```json
// indicators
{
    "_id": "",
    "code": "",
    "indicator": ""
}

// passions
{
    "_id": "",
    "code": "",
    "interest": ""
}

// rules
{
    "_id": "",
    "code": "",
    "if": [],
    "then": ""
}

// results
{
    "_id": "",
    "time": "",
    "database": {
        "true": [],
        "false": []
    },
    "rules": [],
    "status": true,
    "passion": ""
}
```
#### 2. Install dependencies
- Install project dependencies using this command
```bash
go install
```
#### 3. Set environment variables
- Create a .env file on root directory, use the template provided in [.env.example](https://github.com/tudemaha/ifassion-be/blob/main/.env.example)
- Fill all the required variables
#### 4. Run the project
- Run project this command:
```bash
go run main.go
```
