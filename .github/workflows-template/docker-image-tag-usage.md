## need `New repository secret`

- file `docker-image-tag.yml`
- `DOCKERHUB_TOKEN` from [hub.docker](https://hub.docker.com/settings/security)
    - if close push remote can pass `DOCKERHUB_TOKEN` setting 
- `DOCKERHUB_OWNER` for docker hub user
- `DOCKERHUB_REPO_NAME` for docker hub repo name

## usage at github action

```yml
jobs:
  version:
    name: version
    uses: ./.github/workflows/version.yml

  docker-image-tag:
    name: docker-image-tag
    needs:
      - version
      - golang-ci
    uses: ./.github/workflows/docker-image-tag.yml
    if: startsWith(github.ref, 'refs/tags/')
            # with:
    # push_remote_flag: false
    secrets:
      DOCKERHUB_OWNER: "${{ secrets.DOCKERHUB_OWNER }}"
      DOCKERHUB_REPO_NAME: "${{ secrets.DOCKERHUB_REPO_NAME }}"
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `true`