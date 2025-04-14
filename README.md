# Linktok

A modern URL shortener that transforms lengthy links into memorable, shareable URLs.

![Linktok Preview](docs/hero.png)

Try it now: https://linktok-url-shortener.vercel.app/

## Overview

Linktok is a minimalist URL shortener that combines reliability with elegant design. Convert long, unwieldy URLs into short, memorable links perfect for social media, marketing campaigns, and personal use.

## Key Features

- 🔗 **Instant URL Shortening**: Create short links in milliseconds
- 🔍 **Duplicate Detection**: Automatic detection of existing URLs using MD5 checksums
- 🎨 **Animated UI**: Smooth animations and gradient backgrounds for enhanced user experience
- 📋 **One-Click Copy**: Instant copy-to-clipboard functionality
- 🔄 **Automatic Redirects**: Robust redirection system with URL validation
- 📱 **Responsive Design**: Works flawlessly across all device sizes
- 🛡 **Base64 Encoding**: Secure short code generation using base64 encoding

### How It Works

1. **Submit URL**: Enter any valid web address
2. **Checksum Verification**: System checks for existing URL using MD5 hash
3. **Short Code Generation**: Creates base64-encoded identifier
4. **Link Sharing**: Copy and share your shortened link
5. **Redirection**: Any access to short link instantly redirects to original URL

## Technology Stack

### Frontend

- Next.js 15 (App Router)
- TypeScript
- Tailwind CSS
- Framer Motion for animations
- Zustand state management

### Backend

- Go (Chi router)
- PostgreSQL database
- MD5 checksum validation
- Base64 encoding/decoding
- CORS-enabled API endpoints

## CI/CD

This project uses GitHub Actions for continuous integration and deployment:

- **Automated Testing**: All code changes are automatically tested
- **Docker Integration**: Docker images are built and pushed to Docker Hub
- **Deployment Pipeline**: Streamlined deployment process for production updates

To configure the CI/CD pipeline:

1. Add your Docker Hub credentials as repository secrets:
   - `DOCKER_USERNAME`: Your Docker Hub username
   - `DOCKER_PASSWORD`: Your Docker Hub access token or password

# License

[MIT License](LICENSE)
