. /etc/os-release

HERE=`dirname "$0"`
cd "$HERE"

if [ "$UID" = 0 ] || [ `whoami` = root ]; then

    if [ "$ID" = alpine ]; then
        cp -i bin/paf-alpine /usr/local/bin/paf
    else
        cp -i bin/paf /usr/local/bin/paf
    fi

    echo Done.
else
    echo "You must be root to run this script"
    exit 1
fi