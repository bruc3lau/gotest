import requests


def insert():
    url = "http://localhost:8080/movies"
    body = {
        "id": "1111",
        "name": "Liu",
        "director": {
            "name": "Andy"
        }
    }
    response = requests.post(url=url, json=body)
    print(response.content)
    pass


if __name__ == "__main__":
    insert()
    pass
