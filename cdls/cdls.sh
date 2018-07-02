cdls(){
  if [ ! ${1} ];then
    cdls ~/;
  else
    cd "${1}";
    ls -lah;
  fi
}
alias cd='cdls'
