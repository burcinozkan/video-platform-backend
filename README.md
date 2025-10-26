#  Video Platform Backend

A production-ready **video management backend** built with Go (Gin) and AWS S3.  
This project demonstrates how to upload, store, and manage video content with a scalable cloud-native architecture.

---

##  Tech Stack

| Layer | Technology |
|-------|-------------|
| Backend | Go (Gin Framework) |
| Database | PostgreSQL (via Docker Compose) |
| ORM | GORM |
| Storage | AWS S3 |
| Caching | Redis *(optional)* |
| Deployment | Docker |

---

##  Features

 Upload and store video files in AWS S3  
 Return public URLs after successful upload  
 CRUD operations for video tasks  
 PostgreSQL integration via GORM  
 Modular, production-oriented folder structure  
 Ready for Redis caching and Kubernetes scaling  

---

##  Project Structure
```
video-platform-backend/
├── cmd/ # App entry point
│ └── main.go
├── internal/
│ ├── db/ # Database connection
│ ├── handlers/ # API logic
│ ├── models/ # GORM models
│ ├── routes/ # API routes
│ └── storage/ # AWS S3 integration
├── docker-compose.yml # PostgreSQL container
├── go.mod / go.sum # Go module setup
├── .env.example # Environment variable template
└── README.md
```
---

##  Environment Setup

Create a `.env` file in the root directory:

```env
AWS_REGION=eu-central-1
AWS_ACCESS_KEY_ID=your-access-key
AWS_SECRET_ACCESS_KEY=your-secret-key
AWS_BUCKET_NAME=burcin-video-storage
POSTGRES_USER=admin
POSTGRES_PASSWORD=password
POSTGRES_DB=video_platform
```
## Docker Setup

Start PostgreSQL locally using Docker Compose:

```
docker-compose up -d
```
Verify:
```
docker ps
```
You should see video_db running on port 5432.

## Run Locally
```
go mod tidy
```
```
go run cmd/main.go
```

## Expected output:

Connected to PostgreSQL successfully

Connected to AWS S3 successfully

[GIN-debug] Listening and serving HTTP on :8080

 ## API Endpoints
 ## Video Upload
 
```
POST /videos/upload
```

Form-data:

Key	Type	Description
video	File	The .mp4 file to upload
Response
```
{
  "url": "https://burcin-video-storage.s3.eu-central-1.amazonaws.com/demo.mp4"
}
```

## Task Management
```
Method	Endpoint	Description
POST	/tasks	Create new task
GET	/tasks	List all tasks
GET	/tasks/:id	Get task by ID
PUT	/tasks/:id	Update task
DELETE	/tasks/:id	Delete task
```
## Example Workflow

Start the backend

Use Postman to send a POST /videos/upload request

Receive a public S3 video URL
