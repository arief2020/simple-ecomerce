
# E-commerce API Project
![Status](https://img.shields.io/badge/status-Completed%20(with%20potential%20for%20future%20improvements)-blue?style=for-the-badge)

This project is the final assignment for the **Project-Based Virtual Intern - Backend Developer Evermos x Rakamin Academy**. The goal of the assignment is to develop a robust e-commerce API using **Golang**, the **Fiber framework**, and **MySQL database** while adhering to specific requirements and implementing best practices.

## Features and Requirements

1.  **Comprehensive Routing**  
    The API includes all required routes as specified in the [Postman collection](https://github.com/Fajar-Islami/go-example-cruid/blob/master/Rakamin%20Evermos%20Virtual%20Internship.postman_collection.json). Additional routes can be implemented, but none of the original routes may be removed.
    
2.  **Unique Constraints**  
    User **email** and **phone number** must be unique across the application.
    
3.  **Authentication & Authorization**
    
    -   JWT is used for secure user authentication.
    -   Users cannot access or manage other users' data, including addresses, store details, products, and transactions.
4.  **File Upload**  
    The API supports file uploads for features such as product images.
    
5.  **Automated Store Creation**  
    A store is automatically created when a user registers.
    
6.  **Address Management**  
    Address data is mandatory for product shipment purposes.
    
7.  **Admin-Only Features**  
    Only admin users can manage product categories.
    
8.  **Data Pagination and Filtering**  
    The API implements pagination and supports data filtering as demonstrated in the provided Postman collection.
    
9.  **Product Transaction Logging**
    
    -   A **Product Log Table** records details of products involved in each transaction.
    -   This ensures traceability and transactional accuracy.
10.  **Clean Architecture**  
    The project follows clean architecture principles to maintain scalability, testability, and maintainability.
    

## Resources

To assist in the development of this API, the following resources are provided:

1.  **Indonesian Regions Data**  
    Retrieve regional data for Indonesia using the API available at:  
    [https://www.emsifa.com/api-wilayah-indonesia/](https://www.emsifa.com/api-wilayah-indonesia/)
    
2.  **Database Design**  
    The database schema for this project can be viewed at the following link:  
    [https://www.emsifa.com/api-wilayah-indonesia/](https://www.emsifa.com/api-wilayah-indonesia/)
## Tech Stack  

This project is built using the following technologies:  

### Backend  
- **Language**: Golang  
- **Framework**: [Fiber](https://gofiber.io/) - A fast and lightweight web framework inspired by Express.js  
- **Database**: MySQL - A relational database management system  
- **Authentication**: JWT (JSON Web Tokens) for secure user authentication  
- **ORM**: GORM - An Object Relational Mapper for Golang  
- **Swagger**:  For API testing and documentation  

### Tools  
- **Postman**: For API testing and documentation  
- **Git**: Version control system  
- **Docker** : For containerization and development environment setup  
- **VS Code**: Preferred code editor  

### Deployment  
- **Platform**: (Specify platform if used, e.g., AWS, Heroku, or local server)  

## Running the Project  

To run this project locally, follow these steps:  

### Prerequisites  
Before starting, ensure you have the following installed:  
- **Go** (Golang) - The language used for the project. [Installation Guide](https://golang.org/doc/install)  
- **MySQL** - The database used by the application. [Installation Guide](https://dev.mysql.com/doc/refman/8.0/en/installing.html)  
- **Postman** (optional) - For API testing and exploring the endpoints.

### Setup Steps

1. **Clone the repository**  
   First, clone the project to your local machine:  

   ```bash
   git clone https://github.com/your-username/project-name.git
   cd project-name
   ```
2.  **Install Go dependencies**  
Navigate to the project directory and install all Go dependencies:
	```
	go mod tidy
	```
3. **Setup MySQL database**

-   Create a new MySQL database for the project.
-   Configure the database connection by updating the `config/db.go` file with your MySQL credentials.
	
	Example configuration in `db.go`:
	```
	const (
	    DB_HOST     = "localhost"
	    DB_PORT     = 3306
	    DB_USER     = "your_db_user"
	    DB_PASSWORD = "your_db_password"
	    DB_NAME     = "your_db_name"
	)
	```
4. **Run Migration and Start the API Server**  
Now, you can start the API server:
	```
	go run main.go
	```
	The API will start on `http://localhost:3000`.
	    
5. **Testing the API**
    -   Open **Postman** and import the Postman collection for API testing (provided in the documentation).
    -   Test the available endpoints according to the collection.
  6. ### Running in Docker (Optional)

		If you prefer to run the project inside a Docker container, follow these steps:

		1.  Build the Docker image:
		```
			docker build -t ecommerce-api .
		```
		2. Run the Docker container:
		```
			docker run -p 3000:3000 ecommerce-api
		```
		The API should now be accessible at `http://localhost:3000` within the Docker container.
