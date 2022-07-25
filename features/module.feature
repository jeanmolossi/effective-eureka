Feature: Modules endpoints
	As a student
	I want to be able to get a list of modules
	So that I can see what modules are available

	Scenario: Should get module info
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		Given there are "modules" with:
			| module_id								| module_title		| module_description			| module_thumb						| module_published	| course_id								|
			| 4aa77560-9c90-4128-b308-ad5c0515b5d7	| Effective Eureka	| A module for effective eureka	| http://www.example.com/thumb.jpg	| true				| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	|
		When I "GET" to "/module/4aa77560-9c90-4128-b308-ad5c0515b5d7"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
			"module_description": "A module for effective eureka",
			"module_id": "4aa77560-9c90-4128-b308-ad5c0515b5d7",
			"module_published": true,
			"module_thumbnail": "http://www.example.com/thumb.jpg",
			"module_title": "Effective Eureka"
		}
		"""

	Scenario: Should update specific module
		Given there are "courses" with:
			| course_id								| course_title		| course_description					| course_thumb							| course_published |
			| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	| Effective eureka	| This is a catalog video manager API.	| http://e2e_api_client:8080/img1.jpeg	| true				|
		Given there are "modules" with:
			| module_id								| module_title		| module_description			| module_thumb						| module_published	| course_id								|
			| 4aa77560-9c90-4128-b308-ad5c0515b5d7	| Effective Eureka	| A module for effective eureka	| http://www.example.com/thumb.jpg	| true				| b23d00ec-69f7-4fc9-b86c-3ba91c845bff	|
		When I "PUT" to "/module/4aa77560-9c90-4128-b308-ad5c0515b5d7" with:
		"""
		{
			"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
			"title":"Effective Eureka edited",
			"published":true
		}
		"""
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"course_id": "b23d00ec-69f7-4fc9-b86c-3ba91c845bff",
			"module_description": "A module for effective eureka",
			"module_id": "4aa77560-9c90-4128-b308-ad5c0515b5d7",
			"module_published": true,
			"module_thumbnail": "http://www.example.com/thumb.jpg",
			"module_title": "Effective Eureka edited"
		}
		"""
