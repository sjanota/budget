package storage

type doc map[string]interface{}

const _id = "_id"
const expenses = "expenses"
const opPush = "$push"
const opPull = "$pull"
const opMatch = "$match"
const opUnwind = "$unwind"
const opReplaceRoot = "$replaceRoot"
const opElemMatch = "$elemMatch"
