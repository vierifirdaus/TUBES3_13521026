# TUBES3_STIMA

> Penerapan String Matching dan Regular Expression dalam Pembuatan ChatGPT Sederhana

| Student ID | Name                  |
| ---------- | --------------------- |
| 13521026   | Kartini Copa          |
| 13521080   | Fajar Maulana Herawan |
| 13521099   | Vieri Fajar Firdaus   |

## Table of Contents

* [Desc](#deskripsi-umum)ription
* [Te](#teknologi)chnology
* [F](#fitur)eatures
* Folders and Description
* Setup

## Description

This web application, ChatGPT, is a simple chatbot that can assist users in finding answers to their questions. The application uses a straightforward QA approach by utilizing string matching algorithms such as Knuth-Morris-Pratt (KMP) and Boyer-Moore (BM) to find the most similar question to the one provided by the user. In addition, the application is equipped with various features such as text question, date feature, calculator feature, add question and answer to the database feature, and delete question feature. In conducting the search, the application uses the Levenshtein Distance method to calculate the similarity level between the user's question and the questions in the database. With its features, this application can help users quickly and easily find answers to their questions.

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


## Setup

### Installation

* [Node JS](https://nodejs.org/en/)
* [React](https://reactjs.org/)
* [Golang](https://go.dev/)

### Compilation

Frontend

1. Clone this repository in your own local directory
   `git clone https://github.com/vierifirdaus/TUBES3_STIMA.git`
2. Open the command line and change the directory to 'frontend' folder
   `cd Tubes3_13521026/src/frontend`
3. Run `npm install` on the command line
4. Run `npm run dev` on the command line

Backend

1. Navigate to `/src/backend` folder
2. Run `go get .` to make sure all dependencies installed
3. Run `go run server.go`
