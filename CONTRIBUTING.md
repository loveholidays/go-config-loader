# Development

Thanks for contributing! We want to ensure that `go-config-loader` evolves and fulfills
its idea of extensibility and flexibility by seeing continuous improvements
and enhancements, no matter how small or big they might be.

## How to contribute?

We follow fairly standard but lenient rules around pull requests and issues.
Please pick a title that describes your change briefly, optionally in the imperative
mood if possible.

If you have an idea for a feature or want to fix a bug, consider opening an issue
first. We're also happy to discuss and help you open a PR and get your changes
in!

- If you have a question,
  try [creating a GitHub Discussions thread.](https://github.com/loveholidays/go-config-loader/discussions/new)
- If you think you've found a
  bug, [open a new issue.](https://github.com/loveholidays/go-config-loader/issues/new/choose)
- or, if you found a bug you'd like to fix, [open a PR.](https://github.com/loveholidays/go-config-loader/compare)
- If you'd like to propose a
  change [open a new issue.](https://github.com/loveholidays/go-config-loader/issues/new/choose)

### What are the issue conventions?

There are **no strict conventions**, but we do have two templates in place that will fit most
issues, since questions and other discussion start on GitHub Discussions. The bug template is fairly
standard and the rule of thumb is to try to explain **what you expected** and **what you got
instead.** Following this makes it very clear whether it's a known behavior, an unexpected issue,
or an undocumented quirk.

We do ask that issues _aren’t_ created for questions, or where a bug is likely to be either caused
by misusage or misconfiguration. In short, if you can’t provide a reproduction of the issue, then
it may be the case that you’ve got a question instead.

### How do I propose changes?

We follow **no strict process** when it comes to proposing changes.
Simply [raise an issue](https://github.com/loveholidays/go-config-loader/issues/new/choose)
in github. This allows us to track what's being worked on by who and keep our feature requests in a centralised place.

### What are the PR conventions?

This also comes with **no strict conventions**. We only ask you to follow the PR template we have
in place more strictly here than the templates for issues, since it asks you to list a summary
(maybe even with a short explanation) and a list of technical changes.

If you're **resolving** an issue please don't forget to add `Resolve #123` to the description so that
it's automatically linked, so that there's no ambiguity and which issue is being addressed (if any)

We also typically **name** our PRs with a slightly descriptive title, e.g. `(shortcode) - Title`,
where shortcode is either the name of a package, e.g. `(core)` and the title is an imperative mood
description, e.g. "Update X" or "Refactor Y."

## How do I set up the project?

Luckily it's not hard to get started. You can install dependencies
[using `make`](https://www.gnu.org/software/make/).

```sh
make build
```

Other useful make commands:

```sh
# Unit tests
make test

# Linting (golangci lint):
make check

```

You can find the main packages in `pkg/*`.

## How do I test my changes?

It's always good practice to run the tests when making changes. If you're unsure which packages
may be affected by your new tests or changes you may run `make test` in the root of
the repository.

## How do I lint my code?

We ensure consistency in `go-config-loader`'s codebase using `golangci lint`.
The lint can be run can be run using `make check`. It runs as part of the `make build` step of CI and will highlight any
errors when this step is run.
The rules for the lint can be found in `.golangci.yaml`.

## How do I release new versions of our packages?

We have a [GitHub Actions workflow](./.github/workflow/go-list.yml) which is triggered whenever a new
tag is created.
