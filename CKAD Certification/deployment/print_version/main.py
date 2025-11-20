from fastapi import FastAPI

app = FastAPI()

@app.get("/version")
def get_version():
    return {"version": "v2"}
