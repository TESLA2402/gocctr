#!/bin/bash

green='\033[0;32m'
clear='\033[0m'

go build cctr.go
# Rename the executable if needed
mv cctr mytrtool
# Move the executable to a directory in the PATH
sudo mv mytrtool /usr/local/bin/
echo "Done!🚀"
printf "${green}Build completed. You can now run 'mytrtool' from the command line!${clear}🚀"