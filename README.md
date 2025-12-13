
### Twitter-like Backend Service (Go + MySQL)

This project is a **Twitter-like backend system** built using **Go (Golang)** and **MySQL**, focusing on clean architecture, data integrity, and scalable backend design.

The application supports **user management, tweet management, and hashtag/tag management**, with automated tag extraction and seamless association during tweet creation.


---
#### ✨ Key Features

* **User Management**

  * Secure user registration and authentication with **bcrypt password hashing**
  * Enforced **unique email and username constraints** at the database level
  * Clean separation of authentication and user domain logic

* **Tweet Management**

  * Create and retrieve tweets with optimized database queries
  * Automatic **hashtag extraction** from tweet content using text parsing
  * Designed for high-read scenarios similar to social media platforms

* **Tag / Hashtag Management**

  * Automatic creation of tags if they don’t already exist
  * Efficient **many-to-many relationship** between tweets and tags
  * Ensures tag reuse to avoid duplication and maintain consistency

* **Backend Architecture**

  * REST APIs following clean and modular Go project structure
  * Proper indexing and foreign key constraints in MySQL
  * Transaction-safe operations for tweet and tag creation

---
#### 🛠 Tech Stack

* **Language**: Go (Golang)
* **Database**: MySQL
* **Security**: bcrypt for password hashing
* **Architecture**: REST APIs, layered design (handlers, services, repositories)

---
