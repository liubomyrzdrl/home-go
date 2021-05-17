package config

// DSN You can set your database local settings
const DSN = "host=localhost user=macbook  dbname=?  port=5432"

// DSNREMOTE DSN remote database
const DSNREMOTE = `
 host=ec2-54-228-174-49.eu-west-1.compute.amazonaws.com 
 user=rcgqcnqqybbzxj 
 database=de9spueqlas61b
 password=f8a8e88ff4e1909aac2562e530f7ffa5357092896b2dede157633136f59edc5d
 port=5432
 sslmode=require
`

// PORT Set port
const PORT = "8000"
