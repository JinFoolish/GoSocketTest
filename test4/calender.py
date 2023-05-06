#!/usr/bin/python

import cgi
import calendar

# Print HTTP header 
print("Content-Type: text/html")
print()

# Get current month and year 
form = cgi.FieldStorage()
if "month" in form and "year" in form:
    month = int(form["month"].value)
    year = int(form["year"].value)
else:
    month = int(cgi.escape(input("Month: ")))
    year = int(cgi.escape(input("Year: ")))

# Print HTML header
print("<html>")
print("<body>")
print("<h1>Calendar - %d/%d</h1>" % (month, year))

# Generate calendar 
c = calendar.HTMLCalendar()
calendar_html = c.formatmonth(year, month)

# Print HTML tail
print(calendar_html)  
print("</body>")
print("</html>")