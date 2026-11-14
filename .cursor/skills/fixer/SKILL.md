---
name: fixer
description: >-
  Applies settled PR review findings for standards and spec breaches. Use when invoked as /fixer or when fixing a pull request from a factory instruction.
---

Apply the settled findings on the pull request named after this instruction. The ticket is the issue named after it. The spec is the issue in that ticket's "Parent" section. When there is none, the ticket is its own spec. Do not ask questions. Do not propose first.

Read these before you start and follow them:

- `CODING_STANDARDS.md` for all Go

Check out the pull request. Read its reviews and the checks on its head.

Apply a breach of a documented standard. Apply a requirement that is missing or partial, and one that is implemented wrong. Do not apply a baseline smell, or behaviour the spec did not ask for. List those in the pull request body under "Review findings not applied", each with its axis: Standards or Spec.

A red check is a failing test run. For each red check, take the failure the same way as a criterion: write only enough code to make that test pass, then run the package's tests. Leave a criterion's contract unchanged. When a fix needs a contract the spec does not have, stop and comment on the pull request with the criterion and the contract it needs.

Commit following the git-commit skill. Push to the pull request's branch. Update the pull request body: "Last test run", and "Review findings not applied" when any were left. Do not open a second pull request. Never mark it ready, never merge it, and never close the issue.

When the suite still fails, push what you have and leave the failure in "Last test run".
