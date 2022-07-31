Feature: Create course
	To create correctly a course we should validate
	build and answer properly

	Scenario: All information passing
		When I "POST" to "/course" with:
		"""
		{
			"title": "Effective eureka",
			"description": "This is a catalog video manager API."
		}
		"""
		Then the status code received should be 201


	Scenario: Fail if payload does not fill required rules
		When I "POST" to "/course" with:
		"""
		{
			"title": "",
			"thumbnail": "",
			"description": "0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789-0123456789"
		}
		"""
		Then the status code received should be 400
		Then the response received should match json:
		"""
		{
			"error": "Bad Request",
			"errors": [
				{
					"field": "title",
					"message": "title is required"
				},
				{
					"field": "description",
					"message": "description must be less than 255 characters"
				}
			]
		}
		"""


	Scenario: Should show the course
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		When I "GET" to "/course/b23d00ec-69f7-4fc9-b86c-3ba91c845bff"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
			"course_title": "Effective eureka",
			"course_description": "This is a catalog video manager API.",
			"course_thumbnail": "http://e2e_api_client:8080/img1.jpeg",
			"course_published": true
		}
		"""


	Scenario: If course not found should return 404
		When I "GET" to "/course/unknown-id"
		Then the status code received should be 404
		Then the response received should match json:
		"""
		{"error": "course not found"}
		"""


	Scenario: Should update the course
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		When I "PUT" to "/course/b23d00ec-69f7-4fc9-b86c-3ba91c845bff" with:
		"""
		{
			"title": "Effective eureka edited",
			"description": "This is a catalog video manager API."
		}
		"""
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
			"course_title": "Effective eureka edited",
			"course_description": "This is a catalog video manager API.",
			"course_thumbnail": "http://e2e_api_client:8080/img1.jpeg",
			"course_published": true
		}
		"""


	Scenario: We can create modules inside a course
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		When I "POST" to "/course/b23d00ec-69f7-4fc9-b86c-3ba91c845bff/module" with:
		"""
		{
			"description": "This is a catalog video manager API.",
			"published": true,
			"thumbnail": "https://effective-eureka.s3.amazonaws.com/courses/thumbnail/1.jpg",
			"title": "Effective Eureka Module"
		}
		"""
		Then the status code received should be 201


	Scenario: We can not create modules inside a course if send invalid payload
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		When I "POST" to "/course/b23d00ec-69f7-4fc9-b86c-3ba91c845bff/module" with:
		"""
		{
			"title": ""
		}
		"""
		Then the status code received should be 400
		Then the response received should match json:
		"""
		{
			"error": "Bad Request",
			"errors": [
				{
					"field": "title",
					"message": "title is required"
				}
			]
		}
		"""


	Scenario: Should get modules from specific course
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		Given there are "modules" with:
			| module_id								| module_title		| module_description			| module_thumb						| module_published	| course_id								|
			| 4aa77560-9c90-4128-b308-ad5c0515b5d7	| Effective Eureka	| A module for effective eureka	| http://www.example.com/thumb.jpg	| true				| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	|
		When I "GET" to "/course/b23d00ec-69f7-4fc9-b86c-3ba91c845bff/modules"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"data": [
				{
					"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
					"module_description": "A module for effective eureka",
					"module_id": "4aa77560-9c90-4128-b308-ad5c0515b5d7",
					"module_published": true,
					"module_thumbnail": "http://www.example.com/thumb.jpg",
					"module_title": "Effective Eureka"
				}
			],
			"meta": {
				"items_per_page": 10,
				"page": 1
			}
		}
		"""
