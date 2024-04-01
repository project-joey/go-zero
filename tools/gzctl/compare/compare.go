package main

import "github.com/project-joey/go-zero/tools/gzctl/compare/cmd"

// EXPERIMENTAL: compare goctl generated code results between old and new, it will be removed in the feature.
// TODO: BEFORE RUNNING: export DSN=$datasource, the database must be gozero, and there has no limit for tables.
// TODO: AFTER RUNNING: diff --recursive old_fs new_fs

func main() {
	cmd.Execute()
}
