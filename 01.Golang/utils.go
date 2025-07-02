import (
	"encoding/json"
)


/**
 * 获取嵌套字段的字符串值
 * 从嵌套的 map[string]interface{} 中提取一个字符串字段，支持多种基础类型自动转换为字符串，适用于 JSON 反序列化后的数据结构。
 * @param m map[string]interface{}
 * @param keys ...string  // 用于指定字段路径的键序列
 * @return string  // 提取到的字符串值，若不存在或无法转换，则返回空字符串
 *
 * 示例：
 * data := map[string]interface{}{
 *   "user": map[string]interface{}{
 *     "profile": map[string]interface{}{
 *       "name": "小明",
 *       "age":  10,
 *       "vip":  true,
 *     },
 *   },
 * }
 * name := nestedStrField(data, "user", "profile", "name") // 输出: "小明"
 * age := nestedStrField(data, "user", "profile", "age")   // 输出: "10"
 * vip := nestedStrField(data, "user", "profile", "vip")   // 输出: "true"
 */
func nestedStrField(m map[string]interface{}, keys ...string) string {
	cur := m
	for i, key := range keys {
		val, ok := cur[key]
		if !ok {
			return ""
		}

		if i == len(keys)-1 {
			switch v := val.(type) {
			case string:
				return v
			case fmt.Stringer:
				return v.String()
			case json.Number:
				return v.String()
			case float64, float32, int, int64, int32, uint, uint64, bool:
				return fmt.Sprintf("%v", v)
			default:
				return ""
			}
		}

		// 中间层必须是 map[string]interface{}
		if next, ok := val.(map[string]interface{}); ok {
			cur = next
		} else {
			return ""
		}
	}
	return ""
}

/**
  * 获取嵌套字段的值
  * 从嵌套的 map[string]interface{} 中提取一个整数值，适用于 JSON 反序列化后的数据（数字默认解析为 float64）。
  * @param m map[string]interface{}
  * @param keys ...string
  * @return int64
  * 示例：
  * data := map[string]interface{}{
   	"user": map[string]interface{}{
   		"profile": map[string]interface{}{
   			"age": float64(28),
   		},
   	},
   }
   age := nestedIntField(data, "user", "profile", "age")
   fmt.Println(age) // 输出: 28
*/

func nestedIntField(m map[string]interface{}, keys ...string) int64 {
	cur := m
	for i, key := range keys {
		val, ok := cur[key]
		if !ok {
			return 0
		}

		if i == len(keys)-1 {
			switch v := val.(type) {
			case float64:
				return int64(v)
			case int:
				return int64(v)
			case int64:
				return v
			case json.Number:
				if i64, err := v.Int64(); err == nil {
					return i64
				}
			}
			return 0
		}

		// 中间层必须是 map[string]interface{}
		if next, ok := val.(map[string]interface{}); ok {
			cur = next
		} else {
			return 0
		}
	}
	return 0
}
