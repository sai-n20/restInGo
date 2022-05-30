# Golang REST server exercise

This repository contains code to scaffold a RESTful server in Go which has 2 endpoints:-

- An endpoint to persist application metadata (In memory). The API uses YAML as a valid payload format.
- An endpoint to search application metadata and retrieve a list that matches the query parameters.


This application was built by roughly keeping an MVC architecture in mind

- `Controller` - This folder houses the business logic being provided by the endpoints. The endpoints are also defined here
- `Services` - This folder houses the data store procedures. This seperates the business logic from the specific implementation of any service utilized in the app (cache, db, logging etc)
- `Utils` - Common utilities used throughout the application are stored here


- ## Endpoints

- `GET` /appData - Fetches any/all apps currently stored in the database. Supports query parameters to search a specific app using any of its metadata fields
- `POST` /appData - Endpoint to persist application metadata. Examples of valid/invalid payloads:-
   - Invalid payload
   ```
    title: App w/ Invalid maintainer email
	version: 1.0.1
	maintainers:
	- name: Firstname Lastname
  	email: apptwohotmail.com
	company: Upbound Inc.
	website: https://upbound.io
	source: https://github.com/upbound/repo
	license: Apache-2.0
	description: |
 	### blob of markdown
 	More markdown
   ```

   ```
    title: App w/ missing version
	maintainers:
	- name: first last
  	email: email@hotmail.com
	- name: first last
  	email: email@gmail.com
	company: Company Inc.
	website: https://website.com
	source: https://github.com/company/repo
	license: Apache-2.0
	description: |
 	### blob of markdown
 	More markdown
   ```

   - Valid payload
   ```
    title: Valid App 1
	version: 0.0.1
	maintainers:
	- name: firstmaintainer app1
  	email: firstmaintainer@hotmail.com
	- name: secondmaintainer app1
  	email: secondmaintainer@gmail.com
	company: Random Inc.
	website: https://website.com
	source: https://github.com/random/repo
	license: Apache-2.0
	description: |
 	### Interesting Title
 	Some application content, and description
   ```

   ```
	title: Valid App 2
	version: 1.0.1
	maintainers:
	- name: AppTwo Maintainer
  	email: apptwo@hotmail.com
	company: Upbound Inc.
	website: https://upbound.io
	source: https://github.com/upbound/repo
	license: Apache-2.0
	description: |
 	### Why app 2 is the best
 	Because it simply is...
	```