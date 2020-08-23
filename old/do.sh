for (( i=1; i<=9; i=i+1 )); do
  mkdir 0$i
done

for (( i=10; i<=16; i=i+1 )); do
  mkdir $i
done

for i in `ls`;
do
    if [[ $i == g* ]];
    then
        mv $i ${i:2:2};
    fi
done

for i in `ls`;
do
    if [[ $i != "Thumbs.db" ]];
    then
        synoacltool -enforce-inherit $i
        synoacltool -add $i user:22$i:allow:rwxpdDaARWcCo:fd--
    fi
done

for i in `ls`;
do
    if [[ $i != "Thumbs.db" ]];
    then
        chown -R 22$i $i
    fi
done

for i in `ls`;
do
    if [[ $i != "Thumbs.db" ]];
    then
        pushd $i
        mkdir 语文 数学 英语 物理 化学 生物 政治 历史 地理 技术
        popd
    fi
done