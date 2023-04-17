#!/bin/python3

import sys
import json
from faker import Faker

fake = Faker(["zh_CN"])

count = int(sys.argv[1])

jsonData = []
for i in range(count):
    name = fake.name()
    address = fake.address()
    identifier = fake.ssn(min_age=20, max_age=55)
    phone = fake.phone_number()
    company = fake.company()
    email = fake.email()
    credit_card_number = fake.credit_card_number()
    jsonData.append({"name": name, "address": address, "identifier": identifier, "phone": phone,
                    "company": company, "email": email, "credit_card_number": credit_card_number})


print(json.dumps(jsonData, ensure_ascii=False))
