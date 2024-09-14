# 🤔 Quiz App

Welcome to the Quiz App built using Golang! This simple console-based quiz application is designed to test your general knowledge. It reads questions from a CSV file and presents them to the user in a multiple-choice format. The user has a limited amount of time to answer each question, and at the end of the quiz, their score is displayed.

## Features

- Reads quiz questions from a CSV file.
- Multiple choice questions (options A, B, C, D).
- Timer functionality to limit the time per question.
- Keeps track of the user’s score.
- Gracefully handles timeouts if the user doesn't answer within the time limit.

## Requirements

- Golang 1.18 or higher

## Getting Started

1. Clone the Repository

   ```bash
    git clone https://github.com/your-username/quiz-app-go.git
    cd quiz-app-go
   ```

2. Build and Run the App

## How It Works

- The application reads the CSV file and loads the questions.
- It displays each question with four possible answers (A, B, C, D).
- The user is prompted to enter their answer (case-sensitive).
- The application tracks the user's correct answers.
- A timer runs for each question. If the user fails to answer within the specified time, the program moves on to the next question.
- At the end of the quiz, the user's total score is displayed.

## Project Structure

```bash
 quiz-app-go/
 │
 ├── main.go        # Main application logic
 ├── quiz.csv       # Sample CSV file with questions
 ├── README.md      # Project README file
 └── go.mod         # Go modules file
```

## Contributing

Feel free to fork this repository and submit pull requests. Contributions are always welcome!

- Fork the repository.
- Create a new branch (git checkout -b feature-branch).
- Make your changes.
- Commit your changes (git commit -m 'Add new feature').
- Push to the branch (git push origin feature-branch).
- Create a new Pull Request.

## License

This project is open-source and available under the [MIT License](./LICENSE).
