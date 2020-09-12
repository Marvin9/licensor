if [ "${CI+1}" ]
then
    sh ./scripts/build-platform.sh
    ls | grep -v build | xargs rm -rf
else
    echo "YOU ARE IN DEVELOPMENT ENVIRONMENT."
    echo "THIS COMMAND MUST NOT RUN OUTSIDE CI."
    exit 1
fi