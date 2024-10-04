package database

var Comparisons = []string{"$eq", "$ne", "$gt", "$gte", "$lt", "$lte", "$in", "$nin", "$regex", "$text"}
var LogicalOperators = []string{"$and", "$or"}
var UpdateOperators = []string{"$set", "$push", "$pull"} 