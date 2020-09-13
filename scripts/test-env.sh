
if [ "$MODE" = "DROP" ];
then
    rm mock/main.go
    rm -rf mock/foo
    rm -rf mock/integration
else
    echo "package main" > mock/main.go
    mkdir -p mock/foo
    echo "package foo" > mock/foo/bar.go

    # FOR INTEGRATION TESTS
    mkdir -p mock/integration
    echo "package integration" > mock/integration/integration.go
    printf '#Already commented\nprint("Hello")' > mock/integration/integration.py
fi