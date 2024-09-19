package examples

const Json1 = `{
	"key1": [1, 2, 3]
}`
const Json2 = `{
	"key1": [1, 2, 4]
}`
const JsonStr_1 = `
{
"age": 30.0,
"name": "Alice",
"address": {
"city": "Wonderland",
"postalCode": 12345
},
"hobbies": ["reading", "chess"]
}`

const JsonStr_2 = `
{
"name": "Alice",
"age": 30,
"address": {
"city": "Wonderland"
},
"hobbies": ["reading", "chess", "swimming"],
"isActive": true
}`

const JsonStr1 = `{
	"organization": {
		"name": "Global Tech",
		"departments": {
			"hr": {
				"staff": [
					{
						"id": "hr1",
						"name": "John Doe",
						"role": "HR Manager",
						"details": {
							"age": 40,
							"address": {
								"street": "123 Elm Street",
								"city": "Metropolis",
								"postalCode": 10101
							}
						}
					}
				],
				"policies": {
					"leave": "20 days per year",
					"workFromHome": true
				}
			},
			"it": {
				"staff": [
					{
						"id": "it1",
						"name": "Jane Smith",
						"role": "Software Engineer",
						"details": {
							"age": 35,
							"address": {
								"street": "456 Maple Avenue",
								"city": "Tech City",
								"postalCode": 20202
							}
						}
					},
					{
						"id": "it2",
						"name": "Alice Brown",
						"role": "DevOps Engineer",
						"details": {
							"age": 28,
							"address": {
								"street": "789 Oak Lane",
								"city": "Tech City",
								"postalCode": 30303
							}
						}
					}
				],
				"projects": [
					{
						"title": "Project Alpha",
						"deadline": "2024-06-30",
						"status": "ongoing",
						"tasks": [
							{
								"task": "Develop Backend",
								"status": "in-progress"
							},
							{
								"task": "Setup CI/CD",
								"status": "completed"
							}
						]
					}
				]
			}
		}
	},
	"clients": [
		{
			"id": "c1",
			"name": "Acme Corp",
			"contracts": [
				{
					"contractId": "c1-1",
					"startDate": "2023-01-01",
					"endDate": "2025-01-01"
				}
			]
		}
	]
}`

const JsonStr2 = `{
	"organization": {
		"name": "Global Tech",
		"departments": {
			"hr": {
				"staff": [
					{
						"id": "hr1",
						"name": "John Doe",
						"role": "HR Director",
						"details": {
							"age": 41,
							"address": {
								"street": "123 Elm Street",
								"city": "Metropolis",
								"postalCode": 10101
							}
						}
					}
				],
				"policies": {
					"leave": "25 days per year",
					"workFromHome": true,
					"remoteWork": {
						"eligible": true,
						"maxDays": 3
					}
				}
			},
			"it": {
				"staff": [
					{
						"id": "it1",
						"name": "Jane Smith",
						"role": "Senior Software Engineer",
						"details": {
							"age": 36,
							"address": {
								"street": "456 Maple Avenue",
								"city": "Tech City",
								"postalCode": 20202
							}
						}
					},
					{
						"id": "it3",
						"name": "Charlie White",
						"role": "QA Engineer",
						"details": {
							"age": 30,
							"address": {
								"street": "101 Pine Street",
								"city": "Tech City",
								"postalCode": 40404
							}
						}
					}
				],
				"projects": [
					{
						"title": "Project Alpha",
						"deadline": "2024-06-30",
						"status": "completed",
						"tasks": [
							{
								"task": "Develop Backend",
								"status": "completed"
							},
							{
								"task": "Setup CI/CD",
								"status": "completed"
							},
							{
								"task": "Write Documentation",
								"status": "in-progress"
							}
						]
					}
				]
			}
		}
	},
	"clients": [
		{
			"id": "c1",
			"name": "Acme Corp",
			"contracts": [
				{
					"contractId": "c1-1",
					"startDate": "2023-01-01",
					"endDate": "2024-12-31"
				}
			]
		},
		{
			"id": "c2",
			"name": "Globex Inc",
			"contracts": [
				{
					"contractId": "c2-1",
					"startDate": "2022-07-01",
					"endDate": "2023-07-01"
				}
			]
		}
	]
}`
