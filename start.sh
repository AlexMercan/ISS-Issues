currentDir=`pwd`
configFilePath="${currentDir}/nginx.config"
sudo nginx -s quit > /dev/null
sudo nginx -c ${configFilePath} > /dev/null
