# Blog UI

A Next.js 15 blog application with authentication, post creation, editing, and commenting features.

## Features

- User authentication with OAuth providers
- Create and edit blog posts with Markdown editor
- Post commenting system
- User follow functionality
- Personal post management
- Responsive design with Tailwind CSS

## Tech Stack

- **Framework**: Next.js 15 with App Router and Turbopack
- **UI Components**: Radix UI primitives with custom styling
- **Styling**: Tailwind CSS v4
- **State Management**: Zustand
- **Markdown**: @uiw/react-md-editor with syntax highlighting
- **Icons**: Lucide React
- **Language**: TypeScript

## Project Structure

```
src/
├── app/                    # Next.js App Router pages
│   ├── auth/              # Authentication callbacks
│   ├── create/            # Post creation page
│   ├── edit/[slug]/       # Post editing page
│   ├── my-posts/          # User's posts page
│   ├── post/[slug]/       # Individual post page
│   ├── layout.tsx         # Root layout
│   └── page.tsx           # Home page
├── components/
│   ├── auth/              # Authentication components
│   │   ├── AuthProvider.tsx
│   │   ├── LoginModal.tsx
│   │   └── UserMenu.tsx
│   ├── ui/                # Reusable UI components
│   │   ├── avatar.tsx
│   │   ├── badge.tsx
│   │   ├── button.tsx
│   │   ├── card.tsx
│   │   ├── separator.tsx
│   │   └── toast.tsx
│   ├── CommentForm.tsx
│   ├── CommentItem.tsx
│   ├── FollowButton.tsx
│   └── PostEditor.tsx
├── lib/
│   ├── api-client.ts      # API client utilities
│   └── utils.ts           # Utility functions
└── store/
    └── authStore.ts       # Authentication state management
```

## Getting Started

1. Install dependencies:
```bash
npm install
```

2. Run the development server:
```bash
npm run dev
```

3. Open [http://localhost:3000](http://localhost:3000) in your browser.

## Available Scripts

- `npm run dev` - Start development server with Turbopack
- `npm run build` - Build the application for production
- `npm start` - Start production server
- `npm run lint` - Run ESLint

## Development Tools

- **Linting**: ESLint with Next.js configuration
- **Formatting**: Biome for code formatting
- **Type Checking**: TypeScript strict mode
