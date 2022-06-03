# Go Image Service

Built by [Aviciena Santoso](https://github.com/AviSantoso)

---

## Welcome

<b>Go Image Service</b> allows you to store 📦 and manage your images and albums in the Cloud ☁.

Get started by [logging in](/login) through the navigation bar using your Google account.

---

## Software Stack

- This site is hosted using [Docker](https://www.docker.com/), [Kubernetes](https://kubernetes.io/) and [Digital Ocean](https://www.digitalocean.com/).

- The frontend was built with [SolidJS](https://www.solidjs.com/) and deployed with [Nginx](https://www.nginx.com/).

- The backend is built with [Golang](https://go.dev/) using [Gin Gonic](https://gin-gonic.com/) .

- The images are stored on a [Redis](https://redis.io/) instance.

---

## Best Practices

This website uses 2022 best practices for building websites.

- Uses the [Jamstack](https://jamstack.wtf/) architecture (without CDN).
- 100% test coverage (untestable code is marked as ignored).
- CI / CD using [GitHub Actions](https://github.com/features/actions).
- Pages rendered using Markdown files through [Marked.js](https://marked.js.org/).

You can find the GitHub repository here: [GitHub](https://github.com/AviSantoso/go-image-service).
