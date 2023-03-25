import requests
import os
import telegram
from flask import Flask, request, render_template
from flask_mail import Mail, Message
from config import Configs
from typing import List


# Set up the Telegram bot using the API token
bot = telegram.Bot(token=os.environ.get('TELEGRAM_API_TOKEN'))

# Init the flask app
app = Flask(__name__, emplate_folder="./templates")
app.config.from_object(Configs)
mail = Mail(app)

# backend endpoint
BACKEND_ENDPOINT = os.environ.get("BACKEND_ENDPOINT")


def get_users() -> List[str]:
    """The the Telegram handels of user that suscribed for jobs"""
    try:
        response = requests.get(BACKEND_ENDPOINT + "/subscribers")
        return response.json()
    except Exception as e:
        print(e)
        return []


def send_messages(message: str, handles: List[str]) ->None:
    """Send a message to a Telegram user"""
    for handle in handles:
        try:
            bot.send_message(chat_id=handle, text=message)
        except Exception as e:
            print(e)


def send_emails(message: str, html_message: str, emails: List[str]) ->None:
    """Send an email to a Telegram user"""

    msg = Message(
        "[Alert New JOB]",
        sender=os.environ.get("MAIL_USERNAME"),
        recipients=emails
    )
    msg.body = message
    msg.html = html_message

    try:
        mail.send(message)
    except Exception as e:
        print(e)


@app.route("/message-broadcast", methods=["POST"])
def send_message():
    users = get_users()
    emails = []
    telegram_handles = []

    for user in users:
        if user["email"]:
            emails.append(user["email"])

        if user ["telegram"]:
            telegram_handles.append(user["telegram"])

    text = render_template("mail.txt")
    html = render_template("mail.html")

    send_messages(text, telegram_handles)
    send_emails(text, html, emails)


if __name__ == "__main__":
    app.run(port=5000, host="0.0.0.0")