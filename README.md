# NearPortal demo tasks

Small Go challenges for NearPortal bounty rehearsals. **Test funds only:** Base Sepolia USDC and Monerium sandbox EUR.

## Factorial — USDC bounty

Implement `Factorial` in `tasks/factorial/factorial.go`.

- Return the exact factorial for every input from 0 through 20, including `0! = 1`.
- For every input greater than 20, return 0 and a non-nil error.
- Preserve the function signature. Add no dependencies.
- Modify only `tasks/factorial/factorial.go`.

Required job: **`acceptance-factorial`**.

## Min/max — bank bounty

Implement `MinMax` in `tasks/minmax/minmax.go`.

- Return the smallest and largest values of a non-empty slice.
- Support negatives, duplicates and a single element.
- For empty or nil input, return 0, 0 and a non-nil error.
- Leave the input slice unchanged. Preserve the function signature. Add no dependencies.
- Modify only `tasks/minmax/minmax.go`.

Required job: **`acceptance-minmax`**.

## Contributing

1. Claim the corresponding funded bounty on NearPortal before working.
2. Follow its assignment instructions; one task per pull request targeting `main`.
3. Include the exact NearPortal assignment marker in the PR description.
4. Submit the PR through NearPortal before the assignment deadline.
5. The agreed job must pass on the submitted commit, then the creator reviews and merges the PR. Any GitHub merge method is accepted.

The initial implementations deliberately fail. The other task's failing job does not invalidate your solution: NearPortal must select only your task's job.

Tests and workflows must not be modified to make a submission pass. The creator must review the diff before merging. Public tests are acceptance examples; the full written criteria still apply.

## NearPortal configuration

- Workflow: `.github/workflows/acceptance.yml`
- Target branch: `main`
- Select exactly `acceptance-factorial` or `acceptance-minmax` for the corresponding bounty.
- Protect `.github/workflows/acceptance.yml`, `go.mod`, and both `*_test.go` files in NearPortal's agreed verification files.
- Do not protect the implementation file that contributors must change.
- Do not require both jobs to pass globally while another exercise remains unsolved.

Run independently:

```sh
go test -count=1 -timeout=20s -v ./tasks/factorial
go test -count=1 -timeout=20s -v ./tasks/minmax
```

CI pins Go 1.23.12 and uses no third-party Go packages, external APIs, randomness or timing benchmarks in the tests.
