# import required module
import sys
 
# append the path of the
# parent directory
sys.path.append("..")

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from typing import Union

#from ai.bot import Chat, ask

csapi = FastAPI()


@csapi.get("/")
async def read_root():
    return {"Hello": "World"}


# @aiapi.post("/bot")
# async def ask_bot(request: Chat):
#     print(request)
#     response = ask(request)

#     print(response)
#     return response


@csapi.get("/items/{item_id}")
def read_item(item_id: int, q: Union[str, None] = None):
    return {"item_id": item_id, "q": q}