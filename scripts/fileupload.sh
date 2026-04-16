#!/bin/bash

# $1 - Directory Path
# $2 - File types to upload (default to PDF)

DIR=$1
APIKEY=$2
EXT=$3

if [ -z $DIR ]; then
  echo "No Directory provided"
  printf "\n"
  exit 1
fi 

if [ -z "$APIKEY" ]; then
  echo "APIKEY is empty and required."
  printf "\n"
  exit 1
fi

if [ -z $EXT ]; then
  echo "No File types specified, defaulting to PDF"
  printf "\n"
  EXT="pdf"
fi 

CNT=0
SUCCESS="_uploaded"
FAILED="_uploadfailed"
SUBDIR=""

for FILE in $DIR/*; do

  if [ "$CNT" -lt 50 ]; then

    echo -e "COUNT: $CNT\n"

    echo "checking if $FILE exists"
    printf "\n"

    if [ -f "$FILE" ]; then
      echo "found $FILE"
      printf "\n"
      
      FNAME=$(basename "$FILE")

      echo "uploading $FNAME"
      printf "\n"

      res=$(curl -H "x-api-key: $APIKEY" -F "file=@$FILE" -F "title=$FNAME" https://api.neosentience.org/files)

      if [ "$res" != "Success!" ]; then
        printf "Upload failed!\n"
        SUBDIR="$FAILED"
      else
        printf "Upload succeeded!\n"
        SUBDIR="$SUCCESS"
      fi

      printf "Moving $FILE to $DIR/$SUBDIR\n"
      mv "$FILE" "$DIR/$SUBDIR/"

      printf "\n"
    else
      echo "file $FILE does NOT exist..."
      printf "\n"
    fi
  fi

  let "CNT+=1"

done

