package dao

import (
	"strings"
	"encoding/json"
	"database/sql"
	"github.com/nic-chen/nice"
	"../config"
	"strconv"
	_ "github.com/go-sql-driver/mysql"
)


type Tbl struct {
	Key string              //主键
	Name string             // 表名
	//Cols map[string]interface{}  // 表的所有列信息
}

func (d *Tbl) Fetch(value interface{}) (map[string]interface{}, error) {
	cache := nice.Instance(config.APP_NAME).Cache();
	if cache!=nil {
		data, _ := d.Fetch_cache(value);
		if data!=nil{
			nice.Instance(config.APP_NAME).Logger().Printf("data fetch from cache");
			return data, nil
		}
	}
	
	db  := nice.Instance(config.APP_NAME).Db();
	tbl := d.Name;
	col := d.Key;
	val := ConvertToString(value);
	sql := ` SELECT * FROM ` + tbl + ` WHERE ` + "`" + col + "`=?";
	tmp := make(map[string]interface{});
	res, err := db.Query(sql, val) // 执行语句并返回

	if err!=nil{
		return tmp, err;
	}

	if len(res)>0 {
		d.Store_cache(value, res[0]);
		return res[0], err
	}

	d.Store_cache(value, tmp);

	return tmp, err;
}

func (d *Tbl) Insert(data map[string]interface{}, replace bool) (sql.Result, error) {
	db := nice.Instance(config.APP_NAME).Db();
	tbl := d.Name;

	cmd := ` INSERT INTO `;
	if(replace){
		cmd =  ` REPLACE INTO `
	}

	kv := d.implode(data);

	sql := cmd+tbl+" SET "+kv;

	res, err := db.Exec(sql) // 执行语句并返回

	return res, err
}

func (d *Tbl) Update(value interface{}, data map[string]interface{}) (sql.Result, error) {
	db := nice.Instance(config.APP_NAME).Db();
	tbl := d.Name;
	cmd := ` UPDATE `;
	kv := d.implode(data);
	val := ConvertToString(value);
	sql := cmd+tbl+" SET "+kv+" WHERE `"+d.Key+"`=?";

	res, err := db.Exec(sql, val) // 执行语句并返回

	//删除缓存
	d.Delete_cache(value);

	return res, err
}

func (d *Tbl) Delete(value interface{}) (sql.Result, error) {
	db  := nice.Instance(config.APP_NAME).Db();
	tbl := d.Name;
	cmd := ` DELETE FROM `;
	val := ConvertToString(value);
	sql := cmd+tbl+" WHERE `"+d.Key+"`=?";
	res, err := db.Exec(sql, val) // 执行语句并返回

	//删除缓存
	d.Delete_cache(value);

	return res, err
}

func (d *Tbl) implode(data map[string]interface{}) string{
	sql := ``
	for key, value := range data { // 遍历要传入的数据
		// 拼接set后的字符串  a=1,b='2',c=11
		sql = sql + "`" + key + "`" + `='` + ConvertToString(value) + `',`;
	}
	sql = strings.TrimRight(sql, `,`) // 去掉最后的逗号

	return sql;
}

func (d *Tbl) Fetch_cache(value interface{}) (map[string]interface{}, error){
	cache := nice.Instance(config.APP_NAME).Cache();
	if cache!=nil {
		ckey := d.Name+"_"+ConvertToString(value);
		data,err := cache.Do("GET", ckey);
		if err!=nil || data==nil{
			return nil, nil
		}
		
		var m map[string]interface{}
		err = json.Unmarshal(data.([]byte), &m);
		if err!=nil{
			return nil, nil
		}
		return m, nil
	}
	return nil, nil
}

func (d *Tbl) Store_cache(value interface{}, data map[string]interface{}) bool{
	cache := nice.Instance(config.APP_NAME).Cache();
	if cache!=nil {
		ckey := d.Name+"_"+ConvertToString(value);
	    str, err := json.Marshal(data)

	    if err != nil {
	        return false;
	    }
		_, err = cache.Do("SET", ckey, str);
		return err==nil;
	}
	return false
}

func (d *Tbl) Delete_cache(value interface{}) bool{
	cache := nice.Instance(config.APP_NAME).Cache();
	if cache!=nil {
		ckey := d.Name+"_"+ConvertToString(value);
		_, err := cache.Do("DEL", ckey);
		return err==nil;
	}
	return false
}

// 把数据转换为字符串
func ConvertToString(m interface{}) string {
	switch m.(type) {
	case int64:
		return strconv.FormatInt(m.(int64),10)
	case int32:
		return strconv.Itoa(int(m.(int32)))
	case int:
		return strconv.Itoa(m.(int))
		break
	default:
		return m.(string)
	}
	return ""
}

// 把数据转换为字符串
func ConvertToInt32(m interface{}) int32 {
	switch m.(type) {
	case int64:
		return int32(m.(int64))
	case int:
		return int32(m.(int))
	case float64:
		return int32(m.(float64))
	case uint64:
		return int32(m.(uint64))
		
	default:
		return m.(int32)
	}
	return 0
}
