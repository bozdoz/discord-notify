FROM golang:1.21.3-bullseye

WORKDIR /app

RUN useradd --create-home gopher \
  && chown -R gopher:gopher /app

USER gopher

COPY --chown=gopher:gopher . .
