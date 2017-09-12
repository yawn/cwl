# Cloudwatch Logs (`cwl`)

`cwl` is a commandline application for discovering and tailing AWS CloudWatch
log groups. It's multiregional in nature making it straighforward to e.g. monitor
replicated Lambda@Edge logs.

`cwl` can return all log groups in all regions with `groups` and optionally apply
a prefix filter using `--prefix`.

It can also tail groups named using the schema from groups (`group@region`) with
`tail` and optionally return JSON (instead of `\t` separated strings) using
`--json`.
