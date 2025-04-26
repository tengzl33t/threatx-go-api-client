# Contribution guide

## Bug reports

A bug is a verifiable issue that is caused by the code in the repository or by the libraries used in this code.
Clear and detailed bug reports are extremely valuable.

### Guideline:

1. Check if the issue has already been reported
   in [GitHub issues list](https://github.com/tengzl33t/threatx-go-api-client/issues).
2. Check if the case is still actual and is reproducible using the latest main branch.
3. Isolate the problem and provide list of actions required to reproduce problematic behaviour.
4. Open the issue ticket with all the findings.

A good bug report should provide all the necessary information upfront,
so others don’t need to follow up for clarification.
Please try to include as many details as possible in your report.

#### Report example:

> Short and descriptive title
>
> A summary If possible, include the steps required to reproduce the issue.
>
> 1. First step
> 2. Second step
> 3. Etc.
>
> Any other information you want to share that is relevant to the issue being
> reported. This might include the lines of code that you have identified as
> causing the bug, and potential solutions (and your opinions on their
> merits).

## Feature requests

Feature requests are highly appreciated.
Please ensure your suggestions align with the project's goals and provide enough details
and context to support its value.

To open a feature request, create a new issue ticket with label 'enhancement'
in [GitHub issues](https://github.com/tengzl33t/threatx-go-api-client/issues).
Before opening a new FR, please if this has already been reported.

## Pull requests

A pull request is a proposal to merge a set of changes from one branch into another.
In a pull request, collaborators can review and discuss the proposed set of changes
before they integrate the changes into the main codebase.

All the commits pushed must follow [Conventional Commits](https://www.conventionalcommits.org) specification.

### Code style requirements:

These requirements are applied only to the Python code changes.

- Code must pass all the tests and formatting checks before being merged to the main branch.

### Guideline:

1. Clone the project: `git clone git@github.com:tengzl33t/threatx-go-api-client.git`
2. Create a new branch for the pull request: `git checkout -b branch_name`
3. Make required changes, commit and push.
4. Open a PR in GitHub with appropriate labels, title and description.

Please check [this guide](https://www.baeldung.com/ops/git-guide), if you are new to the git.

## Releases

This project follows [Semantic Versioning (SemVer)](https://semver.org/) for versioning its releases.
All releases are managed and pushed by the main maintainers,
ensuring adherence to these versioning guidelines for consistency and clarity of the repository.