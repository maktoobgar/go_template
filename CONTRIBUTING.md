# Welcome to bookstore contributing guide

Thank you for investing your time in contributing to our project! Any contribution you make will be reflected on **nowhere** because I just started this but I promise I will set a domain and a workflow to see your changes and your efforts somewhere. :sparkles:. 
<!-- Edit Top Part -->
Read our [Code of Conduct](CODE_OF_CONDUCT.md) to keep our community approachable and respectable.

In this guide you will get an overview of the contribution workflow from opening an issue, creating a PR, reviewing, and merging the PR.

Use the table of contents icon <img src="./images/table-of-contents.png" width="25" height="25" /> on the top left corner of this document to get to a specific section of this guide quickly.

## New contributor guide

To get an overview of the project, read the [README](README.md).

### Issues

#### Create a new issue

If you spot a problem with the project or docs, search if an [issue already exists](https://github.com/maktoobgar/bookstore/issues). If a related issue doesn't exist, you can open a new issue using a relevant [issue form](https://github.com/maktoobgar/bookstore/issues/new/choose).
<!-- Don't Forget To Add A Form For Issue -->

#### Solve an issue

Scan through our [existing issues](https://github.com/maktoobgar/bookstore/issues) to find one that interests you. You can narrow down the search using `labels` as filters(that we still don't have). As a general rule, we don’t assign issues to anyone. If you find an issue to work on, you are welcome to open a PR with a fix.

### Make Changes

1. Fork the repository.
2. Git clone your forked version of project.
3. Make your changes locally.
4. Test and debug your changes.(If you want, you can add test cases too)
5. If your code is done, DON'T FORGET to add your changes to [**CHANGELOG**](./CHANGELOG.rst) under **UNRELEASED** section. (If you wanna know how to edit the CHANGELOG, have a look at [CHANGELOG](#changelog) section)
6. If you are sure of your code, commit your changes (there are rules for commits in [Commit Rules](#commit-rules) section, have a look)
7. Create a pull request.
8. Wait for the review:
   1.  If the reviewer asked you something to add or change to your code, do that and then comeback and inform us about the changes.
   2.  If the review rejected, we will inform you about the reason but I don't think we do that cause any improvement is improvement and acceptable but **IF** it is an improvement and that's a cheap **IF** cause I personally suck at programming. :)
   3.  If the review passed and pull request happened, thank you. we are happy to have you and we appreciate your work. :sparkles:. 

### Changelog

Under **UNRELEASED** section, write a general simple sentence of what you did.
For example:
- In `Errors` package when `New` function is called without a pointer but with a variable passed as parameters, it doesn't panic anymore.
- Interfaces for internal `Service` package added.

### Commit Rules

There is a template for how to commit:

[Sticker] [Verb] [Package], [Extra]

#### Stickers

There are some stickers to use in a commit which are:
1. 🐛 Fixed
   - Means you found and fixed a bug somewhere.
2. 🔧 Updated
    - Means you updated someone elses code with new code. (you made optimization or better behavior)
3. ✅ Added
    - Means you added something new. (a function, a variable, a whole new package...)
4. ❌ Deleted
    - Means you deleted something. (a file, a function, a variable, a whole package...)
5. ✏️ Renamed
    - Means you renamed a (folder-file)'s name to something else.
6. 📁 Moved
    - Means you moved a (folder-file) to somewhere else.
7. 🐞 #IssueID
   - Means you fixed an issue from [issues](https://github.com/maktoobgar/bookstore/issues) section of this repository.

Let's have some examples of each one of them:
1. 🐛 Fixed translator, bug in line 19
2. 🔧 Updated translator, `Translate` function code optimized
3. ✅ Added Config
4. ❌ Deleted translator, `translator.go` file, line 19
5. ✏️ Renamed logger, `warnning.go` to `warning.go`
6. 📁 Moved logger, `warning.go` to `./functions` folder
7. 🐞 #14 error, instead of panic, error returns

**Note**: If you changed more than one thing in your code, do a multiline commit or commit them separately.
**Note**: This is not really important to act exactly as I said in commits but at least use relevant **stickers** and say which **package** in your commits.
