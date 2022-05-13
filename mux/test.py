import requests


def insert():
    url = "http://localhost:2022/movies"
    body = {
        "id": 222,
        "name": "Emperor",
        "director": {
            "name": "Fu"
        }
    }
    response = requests.post(url=url, json=body)
    print(response.content)


def delete():
    url = "http://localhost:2022/movies/1001"
    response = requests.delete(url=url)
    print(response.content)


if __name__ == "__main__":
    delete()
