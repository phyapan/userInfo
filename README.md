

# React + Vite & Firebase

This project now uses a React + Vite frontend and Firebase Firestore to save user info.

## Firebase Setup
1. Create a Firebase project at https://console.firebase.google.com/
2. Add a web app and copy your Firebase config.
3. Replace the config in `src/App.jsx`.
4. Install Firebase SDK:
	- `npm install firebase`

## Usage
The form will save user info to the `users` collection in Firestore.

## Push to GitHub
1. Initialize git if not done:
	- `git init`
2. Add remote:
	- `git remote add origin https://github.com/<your-username>/userInfo.git`
3. Add, commit, and push:
	- `git add .`
	- `git commit -m "Switch to Firebase for user info"`
	- `git push -u origin main`
