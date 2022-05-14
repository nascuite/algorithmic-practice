SELECT tt.fio, tt.cnt_from_date
FROM (SELECT t.fio, t.cnt_from_date
        FROM (SELECT e.FirstName||' '||e.LastName fio,
                     count(i.Total) cnt_from_date,
                     e.EmployeeId
                   FROM Invoice i, Customer c, Employee e
                   WHERE i.CustomerId = c.CustomerId
                     AND c.SupportRepId = e.EmployeeId
                     AND i.InvoiceDate > date('2010-01-01 00:00:00')
                   GROUP BY e.EmployeeId ) t
      ORDER BY cnt_from_date DESC LIMIT 3 ) tt
ORDER BY cnt_from_date DESC