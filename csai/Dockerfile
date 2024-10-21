FROM python:3.13.0

RUN mkdir /app 

COPY . /app
COPY pyproject.toml /app 
WORKDIR /app

RUN apt-get update

ENV PYTHONPATH=${PYTHONPATH}:${PWD} 

RUN pip3 install poetry
RUN poetry config virtualenvs.create false
RUN poetry install --no-root --only main

ENV PYTHONPATH "${PYTHONPATH}:/app/src"

CMD ["uvicorn", "main:csapi", "--host", "0.0.0.0", "--port", "7000" ,"--app-dir", "csai/server"]