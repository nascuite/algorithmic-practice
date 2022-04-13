select tt.TrackId, tt.qty
from (select t.TrackId,
                (select coalesce(sum(il.Quantity),0)
                 from InvoiceLine il, Invoice i
                 where il.TrackId = t.TrackId
                   and il.InvoiceId = i.InvoiceId
                   and i.InvoiceDate > date('2010-01-01 00:00:00')) qty
         from track t
         ) tt
order by tt.qty desc, tt.TrackId