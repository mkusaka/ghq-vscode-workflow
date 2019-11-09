#!/bin/bash
WORKFLOW_NAME=ghq-vscode-workflow

rm -f main workflow/main
go build main.go
mv main workflow/main
rm -f ${WORKFLOW_NAME}.alfredworkflow
cd workflow; zip -r ../${WORKFLOW_NAME}.zip *
cd ..
mv ${WORKFLOW_NAME}.zip ${WORKFLOW_NAME}.alfredworkflow
rm -f workflow/main
