
if [ "$MODE" = "DROP" ];
then
    rm mock/main.go
    rm -rf mock/foo
else
    echo "package main" > mock/main.go
    mkdir -p mock/foo
    echo "package foo" > mock/foo/bar.go
fi