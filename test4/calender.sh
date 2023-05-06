#!/bin/sh

# Get current month and year 
month=$(date +%m)
year=$(date +%Y)

# Print HTML header
cat <<EOF 
<html>
<body>
<h1>Calendar - $month/$year</h1>
EOF

# Print calendar table 
echo "<table border=\"1\">"
echo "<tr><th>Sun</th><th>Mon</th><th>Tue</th><th>Wed</th><th>Thu</th><th>Fri</th><th>Sat</th></tr>"

# Get first day of month and total days 
firstday=$(date -d "$year-$month-01" +%A)
totaldays=$(date -d "$year-$month-01" +%t)

# Calculate empty cells before first day of month
case $firstday in
    Sun) empty=0 ;;
    Mon) empty=1 ;;
    Tue) empty=2 ;; 
    Wed) empty=3 ;;
    Thu) empty=4 ;; 
    Fri) empty=5 ;;
    Sat) empty=6 ;; 
esac

# Generate calendar table rows
day=1 
while [[ $day -le $totaldays ]] ; do
    echo "<tr>"  
    for i in $(seq 1 7); do
        if [ $i -le $empty ] ; then
            echo "<td>&nbsp;</td>"
        elif [ $day -le $totaldays ] ; then      
            echo "<td>$day</td>"  
            day=$((day + 1))
        else       
            echo "<td>&nbsp;</td>"
        fi
    done
    echo "</tr>"  
done

# Print HTML tail 
cat <<EOF
</table>
</body> 
</html>