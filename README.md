<p align="center">
  <img src="./preview.jpeg" alt="preview" />
  <h2 align="center">blog0</h2>
<p>A modern, full-stack blog platform with AI-powered content processing, built with clean architecture principles and modern web technologies.</p>
</p>


[![Go Version](https://img.shields.io/badge/Go-1.24.5+-00ADD8?style=flat&logo=go)](https://golang.org) [![Next.js](https://img.shields.io/badge/Next.js-15-black?style=flat&logo=next.js)](https://nextjs.org) [![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE) [![TypeScript](https://img.shields.io/badge/TypeScript-5+-3178C6?style=flat&logo=typescript)](https://typescriptlang.org) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-13+-336791?style=flat&logo=postgresql)](https://postgresql.org)


## Architecture Overview

Blog0 is a comprehensive blogging platform consisting of three main components:

- **Backend** - Go REST API with clean architecture
- **UI** - Next.js 15 frontend application
- **Processor** - Background job processing with Trigger.dev

```mermaid
graph TB
    subgraph "Frontend Layer (Next.js 15)"
        UI[React Components]
        Pages[App Router Pages]
        AuthUI[Auth Components]
        Store[Zustand Store]
        Editor[Markdown Editor]
    end
    
    subgraph "Backend Layer (Go Clean Architecture)"
        subgraph "Presentation"
            Router[Gin Router]
            Handlers[HTTP Handlers]
            Middleware[Auth Middleware]
        end
        subgraph "Application"
            PostService[Post Service]
            UserService[User Service]
            AuthService[Auth Service]
        end
        subgraph "Domain"
            Entities[Domain Entities]
            Interfaces[DAO Interfaces]
        end
        subgraph "Infrastructure"
            PostgresDAO[PostgreSQL DAOs]
            JWTAuth[JWT Handler]
            OAuthClient[OAuth Client]
        end
    end
    
    subgraph "Background Processing (Trigger.dev)"
        ContentJobs[Content Processing Jobs]
        AudioJobs[Audio Generation Jobs]
        AIJobs[AI Enhancement Jobs]
        JobQueue[Job Queue]
    end
    
    subgraph "External Services"
        DB[(PostgreSQL Database)]
        GoogleOAuth[Google OAuth2]
        OpenAI[OpenAI GPT API]
        ElevenLabs[ElevenLabs TTS]
        Upload[File Storage]
    end
    
    %% Frontend to Backend
    UI -->|HTTP Requests| Router
    Pages -->|API Calls| Handlers
    AuthUI -->|OAuth Flow| Middleware
    Store -->|State Updates| UI
    Editor -->|Content Creation| Pages
    
    %% Backend Internal Flow
    Router --> Middleware
    Middleware --> Handlers
    Handlers --> PostService
    Handlers --> UserService
    Handlers --> AuthService
    PostService --> Interfaces
    UserService --> Interfaces
    AuthService --> JWTAuth
    Interfaces --> PostgresDAO
    JWTAuth --> GoogleOAuth
    OAuthClient --> GoogleOAuth
    
    %% Backend to Database
    PostgresDAO --> DB
    
    %% Backend to Background Jobs
    PostService -->|Trigger Jobs| JobQueue
    UserService -->|Trigger Jobs| JobQueue
    JobQueue --> ContentJobs
    JobQueue --> AudioJobs
    JobQueue --> AIJobs
    
    %% Background Jobs to External APIs
    ContentJobs --> OpenAI
    AudioJobs --> ElevenLabs
    AIJobs --> OpenAI
    ContentJobs --> Upload
    
    %% Background Jobs back to Database
    ContentJobs -->|Store Results| DB
    AudioJobs -->|Store Audio URLs| DB
    AIJobs -->|Store Enhanced Content| DB
    
    classDef frontend fill:#61dafb,stroke:#21759b,color:#000
    classDef backend fill:#00add8,stroke:#007d9c,color:#fff
    classDef processor fill:#ff6b6b,stroke:#cc5500,color:#fff
    classDef external fill:#f9f9f9,stroke:#999,color:#333
    classDef layer fill:#e8f4fd,stroke:#1976d2,color:#000
    
    class UI,Pages,AuthUI,Store,Editor frontend
    class Router,Handlers,Middleware,PostService,UserService,AuthService,Entities,Interfaces,PostgresDAO,JWTAuth,OAuthClient backend
    class ContentJobs,AudioJobs,AIJobs,JobQueue processor
    class DB,GoogleOAuth,OpenAI,ElevenLabs,Upload external
```

## Detailed Architecture Flow

```mermaid
sequenceDiagram
    participant User
    participant Frontend as Next.js Frontend
    participant Backend as Go Backend
    participant DB as PostgreSQL
    participant Jobs as Trigger.dev
    participant AI as OpenAI/ElevenLabs

    Note over User,AI: User Authentication Flow
    User->>Frontend: Login Request
    Frontend->>Backend: OAuth Redirect
    Backend->>DB: Google OAuth Flow
    DB-->>Backend: User Data
    Backend-->>Frontend: JWT Token
    Frontend-->>User: Authenticated Session

    Note over User,AI: Content Creation Flow
    User->>Frontend: Create Post
    Frontend->>Backend: POST /api/v1/me/posts
    Backend->>DB: Store Draft Post
    Backend->>Jobs: Trigger Content Enhancement
    Jobs->>AI: Process Content
    AI-->>Jobs: Enhanced Content
    Jobs->>DB: Update Post
    Backend-->>Frontend: Post Created
    Frontend-->>User: Success Message

    Note over User,AI: Background Processing Flow
    Jobs->>AI: Generate Audio
    AI-->>Jobs: Audio File
    Jobs->>DB: Store Audio URL
    Jobs->>Backend: Webhook Notification
    Backend->>Frontend: Real-time Update
    Frontend-->>User: Audio Available
```

## Features

### Core Functionality
- **User Authentication** - Google OAuth2 integration with JWT tokens
- **Content Management** - Create, edit, and publish blog posts with Markdown support
- **Social Features** - Comments, likes, bookmarks, and user following
- **Rich Media** - Image uploads and audio generation for posts

### AI-Powered Features
- **Content Enhancement** - AI-powered content suggestions and improvements
- **Audio Generation** - Text-to-speech conversion using ElevenLabs
- **Background Processing** - Asynchronous job processing for heavy tasks

### Technical Features
- **Clean Architecture** - Separation of concerns with domain-driven design
- **Type Safety** - Full TypeScript support across frontend and processing layers
- **API Documentation** - Swagger/OpenAPI documentation
- **Database Migrations** - Version-controlled schema changes
- **Real-time Updates** - Modern React patterns with server components

## Technology Stack

### Backend (Go)
- **Framework**: Gin web framework
- **Database**: PostgreSQL with database/sql
- **Authentication**: Google OAuth2 + JWT
- **Documentation**: Swagger/OpenAPI
- **Architecture**: Clean Architecture with DDD principles

### Frontend (Next.js)
- **Framework**: Next.js 15 with App Router
- **UI**: React 19 with Radix UI primitives
- **Styling**: Tailwind CSS v4
- **State**: Zustand for client state
- **Markdown**: Rich markdown editor with syntax highlighting

### Background Processing
- **Platform**: Trigger.dev for job orchestration
- **Runtime**: Bun for fast JavaScript execution
- **AI Integration**: OpenAI GPT models
- **Audio**: ElevenLabs text-to-speech
- **Database**: Drizzle ORM with PostgreSQL

## API Documentation

When the backend server is running, interactive API documentation is available at:
```
http://localhost:8080/api/swagger/index.html
```

## Architecture Principles

### Backend (Clean Architecture)
- **Domain Layer**: Entities, value objects, and business rules
- **Application Layer**: Use cases and application services
- **Infrastructure Layer**: Database, HTTP handlers, external services
- **Dependency Rule**: Dependencies point inward toward the domain

### Frontend (Component Architecture)
- **Pages**: App Router pages with server components
- **Components**: Reusable UI components with proper separation
- **State**: Client state with Zustand, server state with React Query patterns
- **Styling**: Utility-first CSS with Tailwind and design system components

### Processing (Job Architecture)
- **Jobs**: Discrete, idempotent background tasks
- **Triggers**: Event-driven job execution
- **Error Handling**: Retry logic and failure notifications
- **Monitoring**: Job status tracking and observability

## Security

- JWT token validation for API authentication
- CORS configuration for cross-origin requests
- SQL injection prevention with parameterized queries
- Input validation and sanitization
- OAuth2 secure authentication flow

## Performance

- Database connection pooling
- Efficient database queries with proper indexing
- Image optimization and lazy loading
- Code splitting and bundle optimization
- Background job processing for heavy operations

## Links

- [Backend Documentation](./backend/README.md)
- [Frontend Documentation](./ui/README.md)
- [Processor Documentation](./processor/README.md)