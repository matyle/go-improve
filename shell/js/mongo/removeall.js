conn = new Mongo("mongodb://tadmin:ZssFX6KEdpqvCnqE@market-loan-dev.cluster-cq0fq8br8u7j.us-west-2.docdb.amazonaws.com:27017/?replicaSet=rs0");

conn.getDB("contractkline_indexprice").dropDatabase();
conn.getDB("contractkline_markprice").dropDatabase();
conn.getDB("contractkline_premiumprice").dropDatabase();

print("drop success");

