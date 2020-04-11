#!/bin/bash

echo "Download Jmeter ..."
wget https://downloads.apache.org//jmeter/binaries/apache-jmeter-5.2.1.zip

echo "Unzip ..."
unzip apache-jmeter-5.2.1.zip

echo "Clear zip file ..."
rm apache-jmeter-5.2.1.zip

echo "Completed !!!"
