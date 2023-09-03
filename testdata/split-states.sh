awk -F\; '{print $1}' municipios.csv | uniq | tail -n +2 | while read line; do grep "^$line;" municipios.csv > $line.csv ; done
