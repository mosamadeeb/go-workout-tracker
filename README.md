# Go Workout Tracker

This is an implementation of a workout tracker following the project idea on [roadmap.sh](https://roadmap.sh/projects/fitness-workout-tracker).

# Stack plan
- Authentication system with JWT
- RESTful API using net/http's `ServeMux`
- Documentation with OpenAPI
- Database migrations with [goose](https://github.com/pressly/goose)
- Query-to-code generation with [sqlc](https://github.com/sqlc-dev/sqlc)

# Milestones
- [ ] Basic server setup
    - [ ] Server state (dev/production modes)
    - [x] Database setup
- [ ] Sign up/login system
- [x] Database schema
- [x] Database seeding
- [ ] API endpoints for basic requirements
- [ ] User authorization
- [ ] JWT Authentication
- [ ] API Documentation (OpenAPI)
- [x] Unit tests for endpoints (preferrably through TDD)
- [ ] More requirements...
- [ ] Basic HTML ([htmx](https://htmx.org/)?) frontend for CRUD visualization
