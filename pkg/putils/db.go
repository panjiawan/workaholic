package putils

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
)

func DbIsNotFound(err error) bool {
	if err == nil {
		return false
	}
	return err.Error() == gorm.ErrRecordNotFound.Error()
}

// 查询条件构建

func BuildCondition(where map[string]interface{}) (whereSql string,
	values []interface{}, err error) {
	for key, value := range where {
		conditionKey := strings.Split(key, " ")
		if len(conditionKey) > 2 {
			return "", nil, fmt.Errorf("" +
				"map构建的条件格式不对，类似于'age >'")
		}
		if whereSql != "" {
			whereSql += " AND "
		}
		switch len(conditionKey) {
		case 1:
			whereSql += fmt.Sprint(conditionKey[0], " = ?")
			values = append(values, value)
			break
		case 2:
			field := conditionKey[0]
			switch conditionKey[1] {
			case "=":
				whereSql += fmt.Sprint(field, " = ?")
				values = append(values, value)
				break
			case ">":
				whereSql += fmt.Sprint(field, " > ?")
				values = append(values, value)
				break
			case ">=":
				whereSql += fmt.Sprint(field, " >= ?")
				values = append(values, value)
				break
			case "<":
				whereSql += fmt.Sprint(field, " < ?")
				values = append(values, value)
				break
			case "<=":
				whereSql += fmt.Sprint(field, " <= ?")
				values = append(values, value)
				break
			case "in":
				whereSql += fmt.Sprint(field, " in (?)")
				values = append(values, value)
				break
			case "notin":
				whereSql += fmt.Sprint(field, " not in (?)")
				values = append(values, value)
				break
			case "like":
				whereSql += fmt.Sprint(field, " like ?")
				values = append(values, value)
				break
			case "<>":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			case "!=":
				whereSql += fmt.Sprint(field, " != ?")
				values = append(values, value)
				break
			}
			break
		}
	}
	return
}

/*
*
where查询
where:=map[string]map[string]string{
	"AND":{"status >":"1","role":"agent"},
	"OR":{"pass":"123456","pass <=":"123456"},
	"LIKE":{"name":Parament.Name,"id":"1"},
}
*/

func Where(wheres map[string]map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	var likeWhere, andWhereSql, orWhereSql string
	var andValues, orValues []interface{}
	for key, value := range wheres {
		//组装where
		if key == "AND" {
			for key2, value2 := range value {
				//andWhere[key2] = value2
				conditionKey := strings.Split(key2, " ")
				if len(conditionKey) == 1 {
					andWhereSql += fmt.Sprint(key2, " = ? AND ")
				} else {
					andWhereSql += fmt.Sprint(conditionKey[0], fmt.Sprintf(" %s ? AND ", conditionKey[1]))
				}
				andValues = append(andValues, value2)
			}
			if andValues != nil {
				andWhereSql = andWhereSql[:len(andWhereSql)-4]
			}
		}
		if key == "OR" {
			for key2, value2 := range value {
				conditionKey := strings.Split(key2, " ")
				if len(conditionKey) == 1 {
					orWhereSql += fmt.Sprint(key2, " = ? OR ")
				} else {
					orWhereSql += fmt.Sprint(conditionKey[0], fmt.Sprintf(" %s ? OR ", conditionKey[1]))
				}
				orValues = append(orValues, value2)
			}
			if orValues != nil {
				orWhereSql = orWhereSql[:len(orWhereSql)-3]
			}
		}
		if key == "LIKE" {
			for key2, value2 := range value {
				likeWhere += key2 + ` LIKE "%` + value2.(string) + `%" OR `
			}
			// 去除尾部OR
			likeWhere = likeWhere[:len(likeWhere)-3]
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(andWhereSql, andValues...).Where(likeWhere).Where(orWhereSql, orValues...)
	}
}
