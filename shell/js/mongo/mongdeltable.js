conn = new Mongo("mongodb://tadmin:ZssFX6KEdpqvCnqE@market-loan-dev.cluster-cq0fq8br8u7j.us-west-2.docdb.amazonaws.com:27017/?replicaSet=rs0",);

conn.getDB("contractkline_indexprice").drop();
conn.getDB("contractkline_markprice").getCollection("hot_data").drop();
conn.getDB("contractkline_premiumprice").getCollection("hot_data").drop();

print("drop success");
