# Skill-Sharing-Platform

> Concept
Front-end: Typescript + React (Vercel)
Back-end: Go + Python (AWS)
Database: MongoDB (while in dev), PostgreSQL (when ready for prod and long-term scaling)


The User's Journey: A user will interact with the React/TypeScript frontend. When they perform an action, like booking a session, the frontend will make an API call.
Core Functionality: This API call will be handled by your GoLang backend. Go will process the request, interact with the MongoDB database to save the booking information, and handle any real-time notifications to the involved users.
Specialized Tasks: If the platform wants to recommend new skills to the user after the booking, the GoLang backend could send a message to a Python service. This Python service would then use its machine learning libraries to generate recommendations and store them in the database for the user to see later.
