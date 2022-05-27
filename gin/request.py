import requests
host="localhost:8080"
def get():
    url=f"{host}/weapon?type=ninjaStar"
    response=requests.get(url=url)
    print(response.text)

if __name__=="__main__":
    get()