import calendar
import datetime

print()

# Get current month and year
now = datetime.datetime.now()
year = now.year
month = now.month

# Print HTML headers
print("<html>")
print("<head>")
print("<title>Current Month Calendar</title>")  
print("</head>")
print("<body>")

# Print month header 
print("<h2>{0} {1}</h2>".format(calendar.month_name[month], year))

# Get calendar 
cal = calendar.HTMLCalendar(calendar.MONDAY) 
table = cal.formatmonth(year, month) 

# Print calendar
print(table)  

# Finish HTML  
print("</body>")
print("</html>")