import email
import requests
from faker import Faker
from faker.providers import profile, phone_number

fake = Faker()
fake.add_provider(profile)
fake.add_provider(phone_number)

HOST = "http://127.0.0.1:8000"

def generate_user() -> dict:
    user = fake.profile()
    
    return {
        "email": user["mail"],
        "name": user["name"],
        "password": "qaswedfr",
        "phone": fake.phone_number(),
        "address": user['address']
    }

def sign_up(json) -> requests.Response:
    url = HOST + "/users/signup"
    response: requests.Response = requests.post(url=url, json=json)
    
    return response