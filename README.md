# Blog-gator 🐊

Blog-gator is a command-line blog aggregator that allows you to follow your favorite RSS/Atom feeds and browse their latest posts. It's built with Go and uses PostgreSQL for data persistence.

## Prerequisites

Before running Blog-gator, you need to have the following installed:

- **Go 1.24.3 or later** - [Download here](https://golang.org/dl/)
- **PostgreSQL** - [Download here](https://www.postgresql.org/download/)

## Installation

Install Blog-gator using Go's package manager:

```bash
go install github.com/VMT1312/blog-gator@latest
```

This will install the `blog-gator` binary to your `$GOPATH/bin` directory (make sure it's in your PATH).

## Configuration

### 1. Set up PostgreSQL Database

First, create a PostgreSQL database for Blog-gator:

```sql
CREATE DATABASE blog_gator;
```

### 2. Create Configuration File

Blog-gator uses a JSON configuration file that should be located at `~/.gatorconfig.json`. Create this file with the following structure:

```json
{
  "db_url": "postgres://username:password@localhost/blog_gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `username` and `password` with your PostgreSQL credentials.

### 3. Run Database Migrations

The application will automatically handle database schema creation when you first run it.

## Usage

Blog-gator provides several commands to manage users, feeds, and posts:

### User Management

**Register a new user:**
```bash
blog-gator register <username>
```

**Login as an existing user:**
```bash
blog-gator login <username>
```

**View all users:**
```bash
blog-gator users
```

**Reset the database (⚠️ deletes all data):**
```bash
blog-gator reset
```

### Feed Management

**Add a new RSS/Atom feed:**
```bash
blog-gator addfeed <feed_name> <feed_url>
```

**View all available feeds:**
```bash
blog-gator feeds
```

**Follow a feed:**
```bash
blog-gator follow <feed_url>
```

**View feeds you're following:**
```bash
blog-gator following
```

**Unfollow a feed:**
```bash
blog-gator unfollow <feed_url>
```

### Content Browsing

**Fetch latest posts from all followed feeds:**
```bash
blog-gator agg
```

**Browse recent posts:**
```bash
blog-gator browse [limit]
```

## Example Workflow

1. **Set up your user account:**
   ```bash
   blog-gator register john_doe
   ```

2. **Add some interesting feeds:**
   ```bash
   blog-gator addfeed "Go Blog" "https://go.dev/blog/feed.atom"
   blog-gator addfeed "GitHub Blog" "https://github.blog/feed/"
   ```

3. **Follow the feeds:**
   ```bash
   blog-gator follow "https://go.dev/blog/feed.atom"
   blog-gator follow "https://github.blog/feed/"
   ```

4. **Fetch the latest posts:**
   ```bash
   blog-gator agg
   ```

5. **Browse your feed:**
   ```bash
   blog-gator browse 10
   ```

## Development

To contribute to Blog-gator or run it from source:

1. Clone the repository:
   ```bash
   git clone https://github.com/VMT1312/blog-gator.git
   cd blog-gator
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the application:
   ```bash
   go run .
   ```

## License

This project is open source. Please check the repository for license details.