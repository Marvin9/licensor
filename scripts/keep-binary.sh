if [ "${CI+1}" ]
then
    ls | grep -v licensor | xargs rm -rf
else
    echo "YOU ARE IN DEVELOPMENT ENVIRONMENT."
    echo "THIS COMMAND MUST NOT RUN OUTSIDE CI."
    exit 1
fi