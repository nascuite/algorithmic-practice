select t.country, t.total
from (select i.BillingCountry country, sum(i.Total) total
      from Invoice i
      group by i.BillingCountry) t
order by t.total