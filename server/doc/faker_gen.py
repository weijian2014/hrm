#!/bin/python3

import sys
import json
import re
from faker import Faker

fake = Faker(["zh_CN"])

count = int(sys.argv[1])


def is_id_card_valid(id_card):
    id_card_strip = id_card.strip()
    id_pattern = r"^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$"
    ret = re.match(id_pattern, id_card_strip)
    if ret:
        return True
    else:
        return False


jsonData = []
for i in range(count):
    name = fake.name()
    address = fake.address()
    identifier = fake.ssn(min_age=20, max_age=55)
    while False == is_id_card_valid(identifier):
        identifier = fake.ssn(min_age=20, max_age=55)

    phone = fake.phone_number()
    company = fake.company()
    job = fake.job()
    email = fake.email()
    credit_card_number = fake.credit_card_number()
    jsonData.append({"name": name, "address": address, "identifier": identifier, "phone": phone,
                    "company": company, "job": job, "email": email, "credit_card_number": credit_card_number})


print(json.dumps(jsonData, ensure_ascii=False))
