## need `New repository secret`

- file `docker-image-latest.yml`
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

  docker-image-latest:
    name: docker-image-latest
    needs:
      - version
      - golang-ci
      - go-build-check-main
    uses: ./.github/workflows/docker-image-latest.yml
    if: ${{ ( github.event_name == 'push' && github.ref == 'refs/heads/main' ) || ( github.base_ref == 'main' && github.event.pull_request.merged == true ) }}
            # with:
            # build_branch_name: 'main'
            # push_remote_flag: ${{ github.event.pull_request.merged == true }}
    # push_remote_flag: ${{ github.ref == 'refs/heads/main' }}
    secrets:
      DOCKERHUB_OWNER: "${{ secrets.DOCKERHUB_OWNER }}"
      DOCKERHUB_REPO_NAME: "${{ secrets.DOCKERHUB_REPO_NAME }}"
      DOCKERHUB_TOKEN: "${{ secrets.DOCKERHUB_TOKEN }}"
```

- `push_remote_flag` default is `false`