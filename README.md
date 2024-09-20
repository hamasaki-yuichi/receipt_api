# receipt_api
receipt api backend

## How to deploy api.

1. Build docker image.

    for Cloud Run '--platform: linux/x86_64'.

    ```
    docker build --platform linux/x86_64 -t {acount_name}/{project_name}:{tag} apps/api
    ```

1. Push image to dockerhub.

    ```
    docker push {acount_name}/{project_name}:{tag}
    ```

1. Deploy Cloud Run.

    ```
    image_url : docker.io/{acount_name}/{project_name}:{tag}
    port : 8080 (default)
    ```
