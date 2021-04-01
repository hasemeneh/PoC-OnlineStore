#!/bin/bash

mkdir -p /var/log/poc/

DIRDATABASE=$1
if [ -z "$DIRDATABASE" ] ;
then
  DIRDATABASE="./files/database"
fi

echo $DIRDATABASE

l=( $(ls $DIRDATABASE ) )

SCHEMA=$2
DBOPT=
if [ ! -z "$SCHEMA" ];
then
  DBOPT="${DBOPT} -D${SCHEMA}"
fi

if [ -z "$MYSQL_MASTER_HOST" ];
then
  MYSQL_MASTER_HOST="cart-db"
fi


echo "creating tables and inserting dumps"
for file in "${l[@]}"
do
 if [[ $file =~ \.sql ]];
 then
   echo "executing $file"
   mysql -uroot -h$MYSQL_MASTER_HOST $DBOPT <  $DIRDATABASE/$file
 fi
done
echo "Done"
