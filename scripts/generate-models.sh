swagger generate model -f $1 -m models
rm models/wrapped_*.go
sed -i ':a;N;$!ba;s/struct {\n\t\t\([a-zA-Z]*\)\n\t}/\1/g' models/*.go