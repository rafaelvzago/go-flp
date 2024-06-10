# Insurance Claim Prediction Data API

## Description

This API provides sample data for the Insurance Claim Prediction project, specifically designed to support workshops and tutorials. The API offers the following endpoint:

- `/claim`: Returns a list of all claims in JSON format, providing sample claim data for testing and development purposes.

## Building the Image

To build the Docker image for the API, execute the following command:

```bash
podman build -t quay.io/USERNAME/insurance-claim-data .
```

### Parameters:
- `-t quay.io/USERNAME/insurance-claim-data`: Tags the image with the specified repository and image name. Replace `USERNAME` with your actual Quay.io username.

## Running the Image

To run the Docker image, use the following command:

```bash
podman run -d -p 8080:8080 -v $PWD/data:/app/data quay.io/USERNAME/insurance-claim-data 
```

### Parameters:
- `-d`: Runs the container in detached mode, meaning it runs in the background.
- `-p 8080:8080`: Maps port 8080 on the host to port 8080 on the container, making the API accessible via `http://localhost:8080`.
- `-v $PWD/data:/app/data`: Mounts the `data` directory from your current working directory (`$PWD/data`) to `/app/data` inside the container. This allows the container to access necessary data files from the host.

## Accessing the API

To access the API, navigate to `http://localhost:8080/claim` in your web browser or use a tool like `curl` to make a request:

```bash
curl http://localhost:8080/claim
```

## Context

This project is an adaptation of the "Streamline Insurance Claims with OpenShift AI" workshop powered by Red Hat. It leverages Red Hat Service Interconnect to provide data seamlessly for the workshop environment.

## License

This project is licensed under the Apache License 2.0. You can freely use, modify, and distribute this software under the terms of this license.

---
