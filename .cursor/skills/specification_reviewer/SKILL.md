---
name: specification_reviewer
description: >-
  Reviews a pull request against its parent spec without changing code. Use when invoked as /specification_reviewer or when doing a specification review from a factory instruction.
---

Review the pull request named after this instruction against its spec. The fixed point is `origin/main`. Do not ask questions. Do not change the code.

The issue number and URL follow this instruction. Read that issue, following `docs/agents/issue-tracker.md`. The spec is the issue in its "Parent" section. When there is none, the ticket is its own spec. Read the spec the same way.

Check out the pull request. Mark new files with `git add -N <path>` so the diff sees them. The diff is `git diff $(git merge-base origin/main HEAD)`. The commits are `git log origin/main..HEAD --oneline`. Confirm `origin/main` resolves and the diff is non-empty before reviewing. A bad ref or an empty diff stops the review.

Report: (a) requirements the spec asked for that are missing or partial; (b) behaviour in the diff that wasn't asked for; (c) requirements that look implemented but where the implementation looks wrong. Quote the spec line for each finding. Under 400 words.

Submit that report as one review on the pull request. Request changes when a requirement is missing, partial, or implemented wrong. Approve when none of those are present. Behaviour the spec did not ask for is not a reason to request changes; list it in the review body on an approval.
