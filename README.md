# go-github-graphql

Get latest release from a github repository using [Github's GraphQL API](https://developer.github.com/v4/)

## How to build

    go build -v -o app .

## How to use

Export your personal Github access token in `GITHUB_ACCESS_TOKEN` environment variable
in order to use Github's API [docs](https://developer.github.com/v4/guides/forming-calls/#authenticating-with-graphql).
As per Github docs the following scopes are needed:

    user
    public_repo
    repo
    repo_deployment
    repo:status
    read:repo_hook
    read:org
    read:public_key
    read:gpg_key

Then you can use the application as follows:

    ./app -repo github.com/MediaBrowser/Emby.Releases
    4.4.0.9-beta: prerelease:true,
            https://github.com/MediaBrowser/Emby.Releases/releases/tag/4.4.0.9

    Looking for emby-server-deb_4.4.0.9_amd64.deb...
            emby-server-deb_4.4.0.9_amd64.deb: 69111666
            Download URL: https://github.com/MediaBrowser/Emby.Releases/releases/download/4.4.0.9/emby-server-deb_4.4.0.9_amd64.deb
