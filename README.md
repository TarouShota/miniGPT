# GPT-3 Telegram Bot

This is a Telegram bot that uses the OpenAI GPT-3.5 Turbo model to generate responses based on user messages. It integrates with the Telegram API to send and receive messages.

## Features

- **Start Command**: When the bot receives the `/start` command, it sends a greeting message to the user.

- **Message Handling**: The bot listens to incoming messages and sends them to the GPT-3.5 Turbo model for response generation. It then sends the generated response back to the user.

- **Environment Variables**: The bot reads sensitive information such as the OpenAI API key and Telegram bot token from a `.env` file.

## How to Use

1. Clone this repository to your local machine.

2. Create a `.env` file in the project root directory and add the following environment variables:
   - `AITOKEN`: Your OpenAI API key.
   - `TOKEN`: Your Telegram bot token.

3. Run the following command to install the required dependencies:

   ```
   go get -u github.com/joho/godotenv
   go get -u github.com/mymmrac/telego
   ```

4. Build and run the bot using the following command:

   ```
   go run main.go
   ```

5. Start a conversation with the bot on Telegram and send messages to generate responses.

## Dependencies

- [godotenv](https://github.com/joho/godotenv): A Go package for reading environment variables from a `.env` file.
- [telego](https://github.com/mymmrac/telego): A Go package that provides an API for interacting with the Telegram Bot API.
- [telegohandler](https://github.com/mymmrac/telego/tree/master/telegohandler): A telego extension that simplifies the process of handling incoming updates.

