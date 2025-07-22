# Distraction Free Reddit MVP
- Rewrite of my first spring boot project
  - single user, no login (simple to use)
- golang backend (gin framework)
  - serve as middleware to retrieve subreddit information
  - check for cached posts to reduce traffic
    - usage of redis for cache
  - health check endpoint to keep backend alive on vercel free service
    - creation of github action cron job to ping endpoint
- typescript frontend (angular framework)
  - localStorage to store user subreddits
  - progressive web app for native-like experience on mobile

