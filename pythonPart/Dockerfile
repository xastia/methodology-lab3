

FROM python:3.10-alpine


RUN apk add --no-cache \
        build-base \
        lapack-dev \
        gcc \
        musl-dev \
        gfortran


COPY requirements.txt /app/requirements.txt
WORKDIR /app
RUN pip install --no-cache-dir -r requirements.txt


COPY . /app


EXPOSE 80
CMD ["uvicorn", "spaceship.main:app", "--host", "0.0.0.0", "--port", "80"]
