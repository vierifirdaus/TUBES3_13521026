# Tugas Besar 3 Strategi Algoritma

> Penerapan String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana

| Student ID | Name                  |
| ---------- | --------------------- |
| 13521026   | Kartini Copa          |
| 13521080   | Fajar Maulana Herawan |
| 13521099   | Vieri Fajar Firdaus   |

## Table of Contents

* [Description](#description)
* [Overview](#overview)
* [Technology](#technology)
* [Features](#features)
* [Folders and Description](#folders-and-description)
* [Setup](#setup)

## Description

This web application, ChatGPT, is a simple chatbot that can assist users in finding answers to their questions. The application uses a straightforward QA approach by utilizing string matching algorithms such as Knuth-Morris-Pratt (KMP) and Boyer-Moore (BM) to find the most similar question to the one provided by the user. In addition, the application is equipped with various features such as text question, date feature, calculator feature, add question and answer to the database feature, and delete question feature. In conducting the search, the application uses the Levenshtein Distance method to calculate the similarity level between the user's question and the questions in the database. With its features, this application can help users quickly and easily find answers to their questions.

## Overview
<p align="center">
<img width="513" alt="image" src="https://user-images.githubusercontent.com/102657926/236506079-5e19d1b5-9285-4523-bf5e-cfc85e80bb7e.png">
</p>

## Technology

- Golang
- JavaScript, HTML, CSS using React framework
- MySQL for the database tool

## Features

- Text questions
- Date queries
- Calculator operations
- Adding questions and answers to the database
- Deleting questions and answers from the database

## **Folders and Files Description**

```
.
├── .DS_Store
├── .github
│   └── workflows
│       └── go.yml
├── build
│   └── web.config
├── doc
├── README.md
└── src
    ├── client
    │   ├── .eslintrc.cjs
    │   ├── .gitignore
    │   ├── index.html
    │   ├── package-lock.json
    │   ├── package.json
    │   ├── postcss.config.js
    │   ├── public
    │   │   ├── logo.png
    │   │   ├── send.png
    │   │   ├── sendIcon.svg
    │   │   └── user.png
    │   ├── src
    │   │   ├── App.css
    │   │   ├── App.tsx
    │   │   ├── assets
    │   │   │   └── react.svg
    │   │   ├── components
    │   │   │   ├── chat.tsx
    │   │   │   ├── message.tsx
    │   │   │   ├── sendmessage.tsx
    │   │   │   ├── sidebar.tsx
    │   │   │   ├── sideButton.tsx
    │   │   │   └── typingMessage.tsx
    │   │   ├── index.css
    │   │   ├── interface
    │   │   │   └── index.tsx
    │   │   ├── main.tsx
    │   │   ├── pages
    │   │   │   └── Home.tsx
    │   │   ├── theme
    │   │   │   ├── button.tsx
    │   │   │   ├── Input.tsx
    │   │   │   ├── textarea.tsx
    │   │   │   └── theme.tsx
    │   │   └── vite-env.d.ts
    │   ├── tailwind.config.js
    │   ├── tsconfig.json
    │   ├── tsconfig.node.json
    │   └── vite.config.ts
    └── server
        ├── algorithm.go
        ├── calculator.go
        ├── calender.go
        ├── go.mod
        ├── go.sum
        └── server.go
```

## Setup

### Installation

* [Node JS](https://nodejs.org/en/)
* [React](https://reactjs.org/)
* [Golang](https://go.dev/)

### How to Run

Frontend

1. Clone this repository in your own local directory
   `git clone https://github.com/vierifirdaus/TUBES3_STIMA.git`
2. Open the command line and change the directory to 'client' folder
   `cd Tubes3_13521026/src/client`
3. Run `npm install` on the command line
4. Run `npm run dev` on the command line
5. Open in http://localhost:8080/

Backend

1. Navigate to `/src/server` folder
2. Run `go run server.go algorithm.go calculator.go calender.go`
