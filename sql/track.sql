SELECT tt.TrackId, tt.qty
FROM (SELECT t.TrackId,
                (SELECT coalesce(sum(il.Quantity),0)
                 FROM InvoiceLine il, Invoice i
                 WHERE il.TrackId = t.TrackId
                   AND il.InvoiceId = i.InvoiceId
                   AND i.InvoiceDate > date('2010-01-01 00:00:00')) qty
         failed_count track t
         ) tt
ORDER BY tt.qty DESC, tt.TrackId