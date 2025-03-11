# Online Auction Website

## Overview

This project is an online auction platform where users can register, list items for sale, bid on items, and manage their profiles. The platform supports real-time bidding, secure payment processing, and communication between buyers and sellers. It is built using **Go** for the backend, **PostgreSQL** for the database, and **React.js** for the frontend.

## Features

### User Management
- User registration and login with JWT-based authentication.
- Profile management (update profile, upload profile picture).
- Seller verification process.
- User ratings and feedback system.

### Listing Management
- Create, edit, and delete auction listings.
- Categorize listings with hierarchical categories and attributes.
- Upload multiple images for each listing.
- Set reserve prices and "Buy Now" options.

### Auction Engine
- Timed auctions with auto-extension to prevent sniping.
- Real-time bid updates using WebSocket connections.
- Bid validation and history tracking.
- Anti-sniping measures.

### Search and Discovery
- Full-text search for items.
- Browse items by category.
- Filter results by price range, attributes, and sorting options.

### Payment and Communication
- Secure payment processing using Stripe.
- Buyer-seller messaging system for pre- and post-sale communication.
- Transaction history and invoice generation.

### Notifications
- Outbid alerts.
- Auction ending reminders.
- Watchlist updates.
- Email notifications for key events.

## Project Roadmap

### Phase 1: Planning and Setup
- [x] Define detailed requirements and user stories.
- [x] Design database schema.
- [x] Create wireframes and UI/UX design.
- [x] Set up development environments.
- [x] Initialize Git repositories.
- [ ] Configure CI/CD pipelines.

### Phase 2: Core Backend Development
- Set up Go project structure with Gin/Echo.
- Implement database models and migrations.
- Create user authentication system.
- Develop RESTful API endpoints for user management, item listings, categories, and search functionality.
- Set up image upload and storage system.
- Implement basic testing.

### Phase 3: Auction Engine Development
- Develop bidding system logic.
- Create auction timing system with anti-sniping measures.
- Implement WebSocket connections for real-time updates.
- Add notification system for auction events.
- Set up background jobs for auction processing.
- Write comprehensive tests for auction logic.

### Phase 4: Frontend Development
- Set up React project with TypeScript.
- Implement component structure.
- Create user authentication UI.
- Develop main marketplace views (home page, category browsing, search results, item detail page).
- Build user dashboard (profile management, watchlist, active bids, selling items).
- Add responsive design for mobile users.

### Phase 5: Payment and Messaging
- Integrate payment processing system (Stripe).
- Implement buyer/seller messaging system.
- Create feedback and rating system.
- Add invoice generation.

### Phase 6: Testing and Optimization
- Perform load testing.
- Optimize database queries.
- Improve application performance.
- Fix bugs and issues.
- Implement security audits.
- Add analytics tracking.

### Phase 7: Deployment and Launch
- Set up production infrastructure.
- Deploy application.
- Monitor system performance.
- Handle initial user feedback.
- Fix critical issues.

### Phase 8: Post-Launch
- Collect and analyze user feedback.
- Implement new features.
- Continuous improvement and optimization.
- Regular security updates.

## Technologies Used

### Backend
- **Go**: Primary programming language.
- **Gin/Echo**: Web framework for building RESTful APIs.
- **PostgreSQL**: Relational database for storing user data, listings, and transactions.
- **GORM**: ORM for database interactions.
- **JWT**: For secure user authentication.

### Frontend
- **React.js**: JavaScript library for building the user interface.
- **TypeScript**: Adds static typing to JavaScript.
- **WebSocket**: For real-time updates in the auction engine.

### Infrastructure
- **Google Cloud**: Hosting and cloud services.
- **Cloud SQL**: Managed PostgreSQL database.
- **Cloud Storage**: For storing images and static files.
- **Stripe**: Payment processing.

## Setup Instructions

### Prerequisites
- Go 1.24+ installed.
- Node.js and npm/yarn installed.
- PostgreSQL database running locally or remotely.
- Google Cloud account (optional for deployment).

### Backend Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/jimsyyap/auctions-backend.git
   cd auctions-backend
   ```
2. Install dependencies:
   ```bash
   go mod tidy
   ```
3. Set up environment variables:
   - Create a `.env` file in the root directory with the following variables:
     ```env
     DB_HOST=localhost
     DB_USER=your_db_user
     DB_PASSWORD=your_db_password
     DB_NAME=your_db_name
     DB_PORT=5432
     JWT_SECRET=your_jwt_secret
     STRIPE_SECRET_KEY=your_stripe_secret_key
     ```
4. Run database migrations:
   ```bash
   go run cmd/migrate/main.go
   ```
5. Start the server:
   ```bash
   go run cmd/server/main.go
   ```

### Frontend Setup
1. Navigate to the frontend directory:
   ```bash
   cd auctions-frontend
   ```
2. Install dependencies:
   ```bash
   npm install
   ```
3. Start the development server:
   ```bash
   npm start
   ```

### Deployment
- Use Google Cloud Build or GitHub Actions for CI/CD.
- Deploy the backend to Google App Engine or Cloud Run.
- Deploy the frontend to Firebase Hosting or Google Cloud Storage.

## Contribution Guidelines

We welcome contributions from the community! To contribute:

1. Fork the repository.
2. Create a new branch for your feature or bugfix:
   ```bash
   git checkout -b feature-name
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add feature or fix"
   ```
4. Push your branch:
   ```bash
   git push origin feature-name
   ```
5. Open a pull request against the `main` branch.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contact

For questions or feedback, please contact:
- **Email**: your-email@example.com
- **GitHub**: [@jimsyyap](https://github.com/jimsyyap)

