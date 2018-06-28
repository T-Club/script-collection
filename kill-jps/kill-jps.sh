#!/bin/bash
process_ids=`jps | sed '/Jps/d;/Main/d' |  sed 's/[a-zA-Z]//g'`
for id in $process_ids 
do 
  echo 'killing pid '$id
  `kill -9  $id`
done

