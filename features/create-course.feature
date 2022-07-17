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
