FROM python:3.13

RUN mkdir /app

COPY . /app
COPY pyproject.toml /app
WORKDIR /app

RUN apt-get update

ENV PYTHONPATH=${PYTHONPATH}:${PWD}

# RUN pip3 install --upgrade packaging
RUN pip3 install --upgrade 'packaging>=24.2'

RUN pip3 install poetry
#RUN curl -sSL https://install.python-poetry.org | python3 - --uninstall
#RUN curl -sSL https://install.python-poetry.org | python3 -
#
RUN poetry config virtualenvs.create false
RUN poetry self add 'poetry-plugin-export>=1.9'
RUN poetry add fastapi
RUN poetry add uvicorn[standard]
RUN poetry install --no-root --only main

ENV PYTHONPATH "${PYTHONPATH}:/app/src"

CMD ["poetry", "run", "uvicorn", "main:csapi", "--host", "0.0.0.0", "--port", "7000" ,"--app-dir", "csai/server", "--log-level", "critical"]
