# rsgrants - Redshift Grants tool

AWS Redshift grants tool

## Built With

* go v1.16
* make
* [git-chglog](https://github.com/git-chglog/git-chglog)
* [goreleaser](https://goreleaser.com/install/)

## Deployment

Run `./release.sh $VERSION`

This will update docs, changelog, add the tag, push main and the tag to the repo. The `goreleaser` action will publish the binaries to the Github Release.

If you want to simulate the `goreleaser` process, run the following command:

```
$ curl -sL https://git.io/goreleaser | bash -s -- --rm-dist --skip-publish --snapshot
```

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull
requests to us.

1. Fork the [GoodwayGroup/rsgrants](https://github.com/GoodwayGroup/rsgrants) repo
1. Use `go >= 1.16`
1. Branch & Code
1. Run linters :broom: `golangci-lint run`
    - The project uses [golangci-lint](https://golangci-lint.run/usage/install/#local-installation)
1. Commit with a Conventional Commit
1. Open a PR

## Versioning

We employ [git-chglog](https://github.com/git-chglog/git-chglog) to manage the [CHANGELOG.md](CHANGELOG.md). For the
versions available, see the [tags on this repository](https://github.com/GoodwayGroup/rsgrants/tags).

## Authors

* **Derek Smith** - [@clok](https://github.com/clok)

See also the list of [contributors](https://github.com/GoodwayGroup/rsgrants/contributors) who participated in this
project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

## Sponsors

[![goodwaygroup][goodwaygroup]](https://goodwaygroup.com)

[goodwaygroup]: https://s3.amazonaws.com/gw-crs-assets/goodwaygroup/logos/ggLogo_sm.png "Goodway Group"
