SELECT t.country, t.total
FROM (SELECT i.BillingCountry country, sum(i.Total) total
      FROM Invoice i
      GROUP BY i.BillingCountry) t
ORDER BY t.total