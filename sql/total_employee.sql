select tt.fio, tt.cnt_from_date
from (select t.fio, t.cnt_from_date
        from (select e.FirstName||' '||e.LastName fio,
                     count(i.Total) cnt_from_date,
                     e.EmployeeId
                   from Invoice i, Customer c, Employee e
                   where i.CustomerId = c.CustomerId
                     and c.SupportRepId = e.EmployeeId
                     and i.InvoiceDate > date('2010-01-01 00:00:00')
                   group by e.EmployeeId ) t
      order by cnt_from_date desc limit 3 ) tt
order by cnt_from_date desc