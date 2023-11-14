# aws-data-test-generator

With this application you can easily generate data to use in your tests


## Wich type of data can you create?

- uuid
- first-name
- last-name
- code (you need to inform the size of the code) `just numbers`
- currency (you need to inform an interval with the minimum and maximum value)
- state (just in brazil)
- city (just in brazil)


## What kind of output can be created?

- csv (for now is just possible to create a CSV based file)


## How can you use this app?

Create data it's really easy. You just need: 

```json
{
	"type": "csv",
	"size": 300,
    "headers": true,
	"data-structure": [
		{
			"column-id": 0,
			"header": "ID PESSOA CLIENTE",
			"type": "uuid"
		},
        {
			"column-id": 1,
			"header": "NOME CLIENTE",
			"type": "first-name"
		},
        {
			"column-id": 2,
			"header": "SOBRENOME CLIENTE",
			"type": "last-name"
		},
        {
			"column-id": 3,
			"header": "NUMERO CONTRATO",
			"type": "code",
			"size": 9
		}
    ]
}
```

1. Create your data structure in a JSON based file
2. Setup the app to read the JSON based file as input
3. Start the application
4. Get your CSV based data file as output


## Skills

[![Basic Skills](https://skillicons.dev/icons?i=vscode,go)](https://skillicons.dev)