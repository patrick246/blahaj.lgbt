# blahaj.lgbt
Front- and backend components of https://blahaj.lgbt, which displays stock levels of Blahaj (the big version) at all the Ikea stores world-wide.

## Components
The whole thing has the following components:

- The [Blahaj Stock Level Prometheus Exporter](https://github.com/patrick246/blahaj-exporter).
- A [Prometheus](https://github.com/prometheus/prometheus) instance that is configured to scrape the exporter every five minutes
- A server component that fetches the current stock level from Prometheus and serves it as a REST API. This component is located in `/server`.
- A front-end component written in React that fetches and renders stock levels. It uses [Leaflet](https://github.com/Leaflet/Leaflet) for rendering a map.

## Getting started

### Setting up the backend
The backend currently only supports Prometheus as a datasource for Stock data. It needs to be configured using the following environment variables:

- `API_PROMETHEUSDATASOURCE_ADDRESS`: The URL pointing to the Prometheus instance 
- `API_PROMETHEUSDATASOURCE_USERNAME`: Username for HTTP Basic Auth
- `API_PROMETHEUSDATASOURCE_PASSWORD`: Password for HTTP Basic Auth

Once these are set, the backend can be started with `make run`. To build a distribution version, use `make build`. The binary will be created at `out/bin/blahajserver`. A Docker Image can be created with `make docker`. You might need to adjust the docker registry at the top of the Makefile, though.

### Setting up the frontend
The frontend requires the backend to be running on `http://localhost:8080`. If it runs on any other host or port, you will need to modify the webpack dev server proxy config in webpack.config.js.

Dependencies can be installed with `npm install`. After this is done, the frontend can be started locally using `npm run serve`. To build a distribution version, use `npm run webpack`. A Docker Image can be created using `docker build --pull -t ghcr.io/patrick246/blahaj.lgbt/frontend:v1.0.3 .`.

